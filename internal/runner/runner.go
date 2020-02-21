package runner

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/hidalgopl/sailor/internal/config"
	"github.com/hidalgopl/sailor/internal/messages"
	"github.com/nats-io/nats.go"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"time"
)

// Run ...
func Run(conf *config.Config, userID string) error {
	testSuiteID := xid.New().String()
	startTestSuiteSubject := fmt.Sprintf("test_suite.%s.created", testSuiteID)
	subscribeWildcard := fmt.Sprintf("test_suite.%s.>", testSuiteID)

	testSuiteCompletedSubject := fmt.Sprintf("test_suite.%s.completed", testSuiteID)
	// Connect to NATS
	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Sample Queue Subscriber")}
	opts = setupConnOptions(opts)
	nc, err := nats.Connect(conf.NatsURL, opts...)
	defer nc.Close()
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer ec.Close()
	if err != nil {
		log.Fatal(err)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, err := client.Get(conf.URL)
	defer r.Body.Close()
	if err != nil {
		panic(err)
		// TODO
	}
	pubMsg := messages.StartTestSuitePub{
		TestSuiteID: testSuiteID,
		URL:         conf.URL,
		Tests:       messages.TestNames,
		Timestamp:   time.Now(),
		UserID:      userID,
	}
	ec.Publish(startTestSuiteSubject, pubMsg)

	sub, err := nc.SubscribeSync(subscribeWildcard)
	if err != nil {
		log.Fatal(err)
	}
	finalMsg := ""
	// Wait for a message
	for i := 1;  i<= (len(messages.TestNames) + 1); i++ {
		msg, err := sub.NextMsg(30 * time.Second)
		if err != nil {
			log.Fatal(err)
		}
		switch msg.Subject {
		case testSuiteCompletedSubject:
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
