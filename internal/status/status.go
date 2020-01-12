package status

// TestStatus ...
type TestStatus string

const (
	// Passed ...
	Passed = TestStatus("passed")
	// Failed ...
	Failed = TestStatus("failed")
	// Error ...
	Error  = TestStatus("error")
)
