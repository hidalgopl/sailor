package checks

import (
	"github.com/hidalgopl/sailor/internal/messages"
	"github.com/hidalgopl/sailor/internal/status"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/textproto"
	"testing"
)

func TestXContentTypeOptionsNoSniff(t *testing.T) {
	tt := []struct {
		testName    string
		headers     map[string][]string
		expectedErr bool
		expectedRes status.TestStatus
	}{
		{
			testName: "happy path",
			headers: http.Header{
				textproto.CanonicalMIMEHeaderKey("X-Content-Type-Options"): {"nosniff"},
			},
			expectedRes: status.Passed,
		},
	}
	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			resultChan := make(chan messages.TestFinishedPub, 1)
			err := XContentTypeOptionsNoSniff(tc.headers, resultChan)
			assert.NoError(t, err)
			res := <-resultChan
			assert.Equal(t, tc.expectedRes, res.Result)
			close(resultChan)
		})
	}
}

func TestXFrameOptionsDeny(t *testing.T) {
	tt := []struct {
		testName    string
		headers     map[string][]string
		expectedErr bool
		expectedRes status.TestStatus
	}{
		{
			testName: "happy path",
			headers: http.Header{
				textproto.CanonicalMIMEHeaderKey("X-Frame-Options"): {"deny"},
			},
			expectedRes: status.Passed,
		},
		{
			testName: "happy path#2",
			headers: http.Header{
				textproto.CanonicalMIMEHeaderKey("X-Frame-Options"): {"DENY"},
			},
			expectedRes: status.Passed,
		},
	}
	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			resultChan := make(chan messages.TestFinishedPub, 1)
			err := XFrameOptionsDeny(tc.headers, resultChan)
			assert.NoError(t, err)
			res := <- resultChan
			assert.Equal(t, tc.expectedRes, res.Result)
			close(resultChan)
		})
	}
}

func TestXXSSProtection(t *testing.T) {
	tt := []struct {
		testName    string
		headers     map[string][]string
		expectedErr bool
		expectedRes status.TestStatus
	}{
		{
			testName: "happy path",
			headers: http.Header{
				textproto.CanonicalMIMEHeaderKey("X-XSS-Protection"): {"1; mode=block"},
			},
			expectedRes: status.Passed,
		},
	}
	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			resultChan := make(chan messages.TestFinishedPub, 1)
			err := XXSSProtection(tc.headers, resultChan)
			assert.NoError(t, err)
			res := <- resultChan
			assert.Equal(t, tc.expectedRes, res.Result)
			close(resultChan)
		})
	}
}

func TestContentSecurityPolicy(t *testing.T) {
	tt := []struct {
		testName    string
		headers     map[string][]string
		expectedErr bool
		expectedRes status.TestStatus
	}{
		{
			testName: "happy path",
			headers: http.Header{
				textproto.CanonicalMIMEHeaderKey("Content-Security-Policy"): {"default-src 'none'"},
			},
			expectedRes: status.Passed,
		},
	}
	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			resultChan := make(chan messages.TestFinishedPub, 1)
			err := ContentSecurityPolicy(tc.headers, resultChan)
			assert.NoError(t, err)
			res := <- resultChan
			assert.Equal(t, tc.expectedRes, res.Result)
			close(resultChan)
		})
	}
}

func TestDetectFingerprintHeaders(t *testing.T) {
	tt := []struct {
		testName    string
		headers     map[string][]string
		expectedErr bool
		expectedRes status.TestStatus
	}{
		{
			testName: "happy path",
			headers: http.Header{
				"x-powered-by": {"flask"},
			},
			expectedRes: status.Failed,
		},
	}
	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			resultChan := make(chan messages.TestFinishedPub, 1)
			err := DetectFingerprintHeaders(tc.headers, resultChan)
			assert.NoError(t, err)
			res := <- resultChan
			assert.Equal(t, tc.expectedRes, res.Result)
			close(resultChan)
		})
	}
}

func TestCORSconfigured(t *testing.T) {
	tt := []struct {
		testName    string
		headers     map[string][]string
		expectedErr bool
		expectedRes status.TestStatus
	}{
		{
			testName: "happy path",
			headers: http.Header{
				"Access-Control-Allow-Origin": {"*"},
			},
			expectedRes: status.Failed,
		},
	}
	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			resultChan := make(chan messages.TestFinishedPub, 1)
			err := CORSconfigured(tc.headers, resultChan)
			assert.NoError(t, err)
			res := <- resultChan
			assert.Equal(t, tc.expectedRes, res.Result)
			close(resultChan)
		})
	}
}

func TestStrictTransportSecurity(t *testing.T) {
	tt := []struct {
		testName    string
		headers     map[string][]string
		expectedErr bool
		expectedRes status.TestStatus
	}{
		{
			testName: "happy path",
			headers: http.Header{
				"Strict-Transport-Security": {"max-age=3600; includeSubDomains"},
			},
			expectedRes: status.Passed,
		},
	}
	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			resultChan := make(chan messages.TestFinishedPub, 1)
			err := StrictTransportSecurity(tc.headers, resultChan)
			assert.NoError(t, err)
			res := <- resultChan
			assert.Equal(t, tc.expectedRes, res.Result)
			close(resultChan)
		})
	}
}

func TestSetCookieSecureHttpOnly(t *testing.T) {
	tt := []struct {
		testName    string
		headers     map[string][]string
		expectedErr bool
		expectedRes status.TestStatus
	}{
		{
			testName: "happy path",
			headers: http.Header{
				"Set-Cookie": {"cookie-without-secureandhttponly"},
			},
			expectedRes: status.Failed,
		},
		{
			testName: "happy path#2",
			headers: http.Header{},
			expectedRes: status.Passed,
		},
	}
	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			resultChan := make(chan messages.TestFinishedPub, 1)
			err := SetCookieSecureHttpOnly(tc.headers, resultChan)
			assert.NoError(t, err)
			res := <- resultChan
			assert.Equal(t, tc.expectedRes, res.Result)
			close(resultChan)
		})
	}
}

func TestCacheControlOrExpires(t *testing.T) {
	tt := []struct {
		testName    string
		headers     map[string][]string
		expectedErr bool
		expectedRes status.TestStatus
	}{
		{
			testName: "happy path",
			headers: http.Header{},
			expectedRes: status.Failed,
		},
		{
			testName: "only Expires",
			headers: http.Header{
				"Expires": {"0"},
			},
			expectedRes: status.Passed,
		},
		{
			testName: "only Cache-Control",
			headers: http.Header{
				"Cache-Control": {"max-age=480"},
			},
			expectedRes: status.Passed,
		},
		{
			testName: "empty Cache-Control and Expires",
			headers: http.Header{
				"Cache-Control": {""},
				"Expires": {""},
			},
			expectedRes: status.Failed,
		},
	}
	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			resultChan := make(chan messages.TestFinishedPub, 1)
			err := CacheControlOrExpires(tc.headers, resultChan)
			assert.NoError(t, err)
			res := <- resultChan
			assert.Equal(t, tc.expectedRes, res.Result)
			close(resultChan)
		})
	}
}
