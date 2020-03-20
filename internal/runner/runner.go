package runner

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"net/http"
	url2 "net/url"
	"time"

	"github.com/hidalgopl/sailor/internal/config"
	"github.com/hidalgopl/sailor/internal/messages"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
)

type NatsSubjects struct {
	subWildcard   string
	suiteStart    string
	suiteComplete string
}

func generateSubjects(testSuiteID string) *NatsSubjects {
	return &NatsSubjects{
		subWildcard:   fmt.Sprintf("test_suite.%s.>", testSuiteID),
		suiteStart:    fmt.Sprintf("test_suite.%s.created", testSuiteID),
		suiteComplete: fmt.Sprintf("test_suite.%s.completed", testSuiteID),
	}
}

func queryTestUrl(testUrl string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, err := client.Get(testUrl)
	if err != nil {
		return &http.Response{}, err
	}
	defer r.Body.Close()
	return r, nil
}


// Run ...
func Run(conf *config.Config, userID string) error {
	err := checkUrl(conf.URL)
	if err != nil {
		return err
	}
	testSuiteID := xid.New().String()
	subjects := generateSubjects(testSuiteID)
	// Connect to NATS
	// Connect Options.
	opts := []nats.Option{nats.Name("sailor")}
	opts = setupConnOptions(opts)
	nc, err := nats.Connect(conf.NatsURL, opts...)
	defer nc.Close()
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer ec.Close()
	if err != nil {
		return err
	}
	r, err := queryTestUrl(conf.URL)
	if err != nil {
		return err
	}
	pubMsg := messages.StartTestSuitePub{
		TestSuiteID: testSuiteID,
		URL:         conf.URL,
		Tests:       messages.TestNames,
		Timestamp:   time.Now(),
		UserID:      userID,
		Headers:     r.Header,
	}
	ec.Publish(subjects.suiteStart, pubMsg)

	sub, err := nc.SubscribeSync(subjects.subWildcard)
	if err != nil {
		return err
	}
	finalMsg := ""
	// Wait for a message
	for i := 1; i <= (len(messages.TestNames) + 1); i++ {
		msg, err := sub.NextMsg(30 * time.Second)
		if err != nil {
			return err
		}
		switch msg.Subject {
		case subjects.suiteComplete:
			link := "http://secureapi.com/tests/" + conf.Username + "/" + testSuiteID
			finalMsg = "all tasks executed successfully. Link to your test suite: " + link
		default:
			var decodedMsg messages.TestFinishedPub
			json.Unmarshal(msg.Data, &decodedMsg)
			logrus.Infof("[%s] -> %s : result: %v", decodedMsg.TestSuiteID, decodedMsg.TestCode, decodedMsg.Result)
		}

	}
	logrus.Info(finalMsg)
	return nil
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second
	//opts = append(opts, nats.DefaultTimeout())
	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to: %s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}

func checkUrl(testUrl string) error {
	u, err := url2.ParseRequestURI(testUrl)
	if err != nil {
		errMsg := "Provided test url: " + testUrl + " is not valid URI"
		return errors.New(errMsg)
	}
	u, err = url2.Parse(testUrl)
	if err != nil || u.Scheme == "" || u.Host == "" {
		errMsg := "Provided test url: " + testUrl + " is not valid URI"
		return errors.New(errMsg)
	}
	return nil
}
