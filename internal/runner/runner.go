package runner

import (
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

func Run(conf *config.Config) error {
	client := &http.Client{}
	fmt.Println(conf.PrettyPrint())
	fmt.Println("I'm running high as fuck!")
	url := "http://localhost:8072/tests/run"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("err: %v", err)
		return err
	}
	subResp := runnerResp{}
	err = json.NewDecoder(resp.Body).Decode(&subResp)
	if err != nil {
		logrus.Errorf("unexpected resp body: %v", err)
	}
	testSuiteID := xid.New().String()
	startTestSuiteSubject := fmt.Sprintf("test_suite.%s.created", testSuiteID)
	subscribeWildcard := fmt.Sprintf("test_suite.%s.>", testSuiteID)

	testSuiteCompletedSubject := fmt.Sprintf("test_suite.%s.completed", testSuiteID)
	// Connect to NATS
	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Sample Queue Subscriber")}
	opts = setupConnOptions(opts)
	fmt.Println("trying to connect")
	nc, err := nats.Connect(conf.NatsURL, opts...)
	defer nc.Close()
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer ec.Close()
	fmt.Println("connected")
	if err != nil {
		log.Fatal(err)
	}
	pubMsg := messages.StartTestSuitePub{
		TestSuiteID: testSuiteID,
		Url:         conf.Url,
		Tests:       []string{"SEC#0001", "SEC#0002", "SEC#0003", "SEC#0004"},
		Timestamp:   time.Now(),
	}
	ec.Publish(startTestSuiteSubject, pubMsg)

	sub, err := nc.SubscribeSync(subscribeWildcard)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for a message
	for {
		msg, err := sub.NextMsg(30 * time.Second)
		if err != nil {
			log.Fatal(err)
		}
		switch msg.Subject {
		case testSuiteCompletedSubject:
			link := "http://secureapi.com/tests/username/" + testSuiteID
			log.Printf("all tasks executed successfully. Link to your test suite: %s", link)
			return nil
		default:
			var decodedMsg messages.TestFinishedPub
			_ = json.Unmarshal(msg.Data, &decodedMsg)
			log.Printf("[%s] -> %s : result: %v", decodedMsg.TestSuiteID, decodedMsg.TestCode, decodedMsg.Result)
		}

	}

	return nil
}

type runnerResp struct {
	Subject string `json:"subject"`
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
