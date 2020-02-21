package auth

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// Authenticator ...
type Authenticator struct {
	Username  string
	AccessKey string
	URL       string
	HttpClient *http.Client
}

// DoAuth ...
func (auth *Authenticator) DoAuth() (bool, string, string) {
	body, err := json.Marshal(map[string]string{
		"username":   auth.Username,
		"access_key": auth.AccessKey,
	})
	if err != nil {
		// TODO: handle can't create body
	}
	req, err := http.NewRequest("POST", "http://localhost:8072/tests/auth", bytes.NewBuffer(body))
	if err != nil {
		//
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := auth.HttpClient.Do(req)
	if err != nil {
		logrus.Error(err)
		return false, "", ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		errMsg := "failed to authenticate, error code:  " + strconv.Itoa(resp.StatusCode)
		err  = NewAuthError(errMsg)
		logrus.Error(err)
		return false, "", ""
	}
	authResp := authResponse{}
	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {
		logrus.Errorf("unexpected resp body: %v", err)
	}
	if !authResp.IsAllowed {
		return false, "not allowed", ""
	}
	logrus.Infof("Authenticated for %s", auth.Username)
	return authResp.IsAllowed, authResp.RemainLimit, authResp.UserID

}

type authResponse struct {
	IsAllowed   bool   `json:"is_allowed"`
	RemainLimit string `json:"remain_limit"`
	UserID      string `json:"user_id"`
}
