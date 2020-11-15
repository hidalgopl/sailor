package runner

import (
	"crypto/tls"
	"errors"
	"github.com/hidalgopl/sailor/internal/checks"
	"github.com/hidalgopl/sailor/internal/sectests"
	"github.com/hidalgopl/sailor/internal/status"
	"net/http"
	url2 "net/url"
	"sync"
	"time"

	"github.com/hidalgopl/sailor/internal/config"
	"github.com/hidalgopl/sailor/internal/messages"
	"github.com/sirupsen/logrus"
)


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
func Run(conf *config.Config) error {
	err := checkUrl(conf.URL)
	if err != nil {
		return err
	}
	r, err := queryTestUrl(conf.URL)
	if err != nil {
		return err
	}
	resultChan := make(chan messages.TestFinishedPub)
	pubMsg := messages.StartTestSuitePub{
		URL:         conf.URL,
		Timestamp:   time.Now(),
		Headers:     r.Header,
	}
	wg := &sync.WaitGroup{}
	for _, testCode := range messages.TestNames {
		wg.Add(1)
		go func(testCode string, wg *sync.WaitGroup) {
			_ = checks.TestCodes[testCode](pubMsg.Headers, resultChan)
			wg.Done()
		}(testCode, wg)
	}
	finalMsg := ""
	var failedCodes []string

	// Wait for a message
	for range messages.TestNames {
		wg.Add(1)
		go func(msg messages.TestFinishedPub, wg *sync.WaitGroup) {
			if msg.Result == status.Failed {
				failedCodes = append(failedCodes, msg.TestCode)
			}
			logrus.Infof("%s : result: %v", msg.TestCode, msg.Result)
			wg.Done()
		}(<-resultChan, wg)
	}
	wg.Wait()
	logrus.Info(finalMsg)
	if len(failedCodes) > 0 {
		sectests.PrintExplanation(failedCodes)
	}
	sectests.PrintSummary(len(failedCodes), len(messages.TestNames))
	return nil
}

func checkUrl(testUrl string) error {
	_, err := url2.ParseRequestURI(testUrl)
	if err != nil {
		errMsg := "Provided test url: " + testUrl + " is not valid URI"
		return errors.New(errMsg)
	}
	u, err := url2.Parse(testUrl)
	if err != nil || u.Scheme == "" || u.Host == "" {
		errMsg := "Provided test url: " + testUrl + " is not valid URI"
		return errors.New(errMsg)
	}
	return nil
}
