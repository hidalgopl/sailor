package checks

import (
	"github.com/hidalgopl/sailor/internal/messages"
	"github.com/hidalgopl/sailor/internal/status"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

var (
	FingerPrintHeaders = []string{
		"x-powered-by",
		"x-generator",
		"server",
		"x-aspnet-version",
		"x-aspnetmvc-version",
	}
)

type TestChan struct {
	Result   status.TestStatus `json:"result"`
	TestCode string            `json:"test_code"`
}

func NotifyCheckFinished(testCode string, status status.TestStatus, resultChan chan<- messages.TestFinishedPub) error {
	msg := &messages.TestFinishedPub{
		Result:      status,
		TestCode:    testCode,
		Timestamp:   time.Now(),
	}
	logrus.Debugf("%s finished with %v, publishing results...", testCode, status)
	resultChan <- *msg
	logrus.Debug("Pushed info to channel")
	return nil
}

func XContentTypeOptionsNoSniff(headers http.Header, resultChan chan<- messages.TestFinishedPub) error {
	var Status status.TestStatus
	testCode := "SEC0001"
	header := headers.Get("X-Content-Type-Options")
	if header == "nosniff" {
		Status = status.Passed
	} else {
		Status = status.Failed
	}
	err := NotifyCheckFinished(testCode, Status, resultChan)
	if err != nil {
		return err
	}
	return nil
}

func XFrameOptionsDeny(headers http.Header, resultChan chan<- messages.TestFinishedPub) error {
	var Status status.TestStatus
	testCode := "SEC0002"
	header := headers.Get("X-Frame-Options")
	header = strings.ToLower(header)
	if header == "deny" || header == "sameorigin"{
		Status = status.Passed
	} else {
		Status = status.Failed
	}
	err := NotifyCheckFinished(testCode, Status, resultChan)
	if err != nil {
		return err
	}
	return nil
}

func XXSSProtection(headers http.Header, resultChan chan<- messages.TestFinishedPub) error {
	var Status status.TestStatus
	testCode := "SEC0003"
	header := headers.Get("X-XSS-Protection")
	if header == "1" || header == "1; mode=block" {
		Status = status.Passed
	} else {
		Status = status.Failed
	}
	err := NotifyCheckFinished(testCode, Status, resultChan)
	if err != nil {
		return err
	}
	return nil
}

func ContentSecurityPolicy(headers http.Header, resultChan chan<- messages.TestFinishedPub) error {
	var Status status.TestStatus
	testCode := "SEC0004"
	header := headers.Get("Content-Security-Policy")
	if header != "" {
		Status = status.Passed
	} else {
		Status = status.Failed
	}
	err := NotifyCheckFinished(testCode, Status, resultChan)
	if err != nil {
		return err
	}
	return nil
}

func DetectFingerprintHeaders(headers http.Header, resultChan chan<- messages.TestFinishedPub) error {
	var Status status.TestStatus
	testCode := "SEC0005"
	Status = status.Passed
	for _, key := range FingerPrintHeaders {
		if _, ok := headers[key]; ok {
			Status = status.Failed
		}
	}
	err := NotifyCheckFinished(testCode, Status, resultChan)
	if err != nil {
		return err
	}
	return nil

}

func CORSconfigured(headers http.Header, resultChan chan<- messages.TestFinishedPub) error {
	var Status status.TestStatus
	testCode := "SEC0006"
	Status = status.Passed
	header := headers.Get("Access-Control-Allow-Origin")
	if header == "*" {
		Status = status.Failed
	}
	err := NotifyCheckFinished(testCode, Status, resultChan)
	if err != nil {
		return err
	}
	return nil
}

func StrictTransportSecurity(headers http.Header, resultChan chan<- messages.TestFinishedPub) error {
	var Status status.TestStatus
	testCode := "SEC0007"
	Status = status.Passed
	header := headers.Get("Strict-Transport-Security")
	properlySet := strings.Contains(header, "max-age=")
	if !properlySet {
		Status = status.Failed
	}
	err := NotifyCheckFinished(testCode, Status, resultChan)
	if err != nil {
		return err
	}
	return nil
}

func SetCookieSecureHttpOnly(headers http.Header, resultChan chan<- messages.TestFinishedPub) error {
	var Status status.TestStatus
	testCode := "SEC0008"
	Status = status.Passed
	header := headers.Get("Set-Cookie")
	if !strings.Contains(header, "Secure") {
		Status = status.Failed
	}
	if !strings.Contains(header, "HttpOnly") {
		Status = status.Failed
	}
	if header == "" {
		Status = status.Passed
	}
	err := NotifyCheckFinished(testCode, Status, resultChan)
	if err != nil {
		return err
	}
	return nil
}

func CacheControlOrExpires(headers http.Header, resultChan chan<- messages.TestFinishedPub) error {
	var Status status.TestStatus
	testCode := "SEC0009"
	Status = status.Passed
	headerCacheControl := headers.Get("Cache-Control")
	headerExpires := headers.Get("Expires")
	if headerCacheControl+headerExpires == "" {
		Status = status.Failed
	}
	err := NotifyCheckFinished(testCode, Status, resultChan)
	if err != nil {
		return err
	}
	return nil
}

var (
	TestCodes = map[string]func(http.Header, chan<- messages.TestFinishedPub) error{
		"SEC0001": XContentTypeOptionsNoSniff,
		"SEC0002": XFrameOptionsDeny,
		"SEC0003": XXSSProtection,
		"SEC0004": ContentSecurityPolicy,
		"SEC0005": DetectFingerprintHeaders,
		"SEC0006": CORSconfigured,
		"SEC0007": StrictTransportSecurity,
		"SEC0008": SetCookieSecureHttpOnly,
		"SEC0009": CacheControlOrExpires,
	}
)
