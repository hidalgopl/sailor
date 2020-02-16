package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
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
		fmt.Printf("oops: %v", err)
		return false, "", ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("oops")
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
	logrus.Info("auth went well")
	return authResp.IsAllowed, authResp.RemainLimit, authResp.UserID

}

type authResponse struct {
	IsAllowed   bool   `json:"is_allowed"`
	RemainLimit string `json:"remain_limit"`
	UserID      string `json:"user_id"`
}
