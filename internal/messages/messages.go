package messages

import (
	"fmt"
	"github.com/hidalgopl/sailor/internal/status"
	"strings"
	"time"
)

// StartTestSuitePub ...
type StartTestSuitePub struct {
	TestSuiteID string    `json:"test_suite_id"`
	URL         string    `json:"url"`
	Tests       []string  `json:"tests"`
	Timestamp   time.Time `json:"timestamp"`
	UserID      string    `json:"user_id"`
}

// Print ...
func (msg *StartTestSuitePub) Print() string {
	return fmt.Sprintf("[%s] URL: %s {%s}", msg.Timestamp, msg.URL, strings.Join(msg.Tests[:], ","))
}

// TestFinishedPub ...
type TestFinishedPub struct {
	TestSuiteID string            `json:"test_suite_id"`
	Result      status.TestStatus `json:"result"`
	TestCode    string            `json:"test_code"`
	Timestamp   time.Time         `json:"timestamp"`
}

// TestSuiteFinishedPub ...
type TestSuiteFinishedPub struct {
	TestSuiteID string            `json:"test_suite_id"`
	URL         string            `json:"url"`
	Tests       []TestFinishedPub `json:"tests"`
	Timestamp   time.Time         `json:"timestamp"`
	UserID      string            `json:"user_id"`
}
