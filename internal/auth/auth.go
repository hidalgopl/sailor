package auth

import (
	"bytes"
	"encoding/json"
	"github.com/hidalgopl/sailor/internal/config"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)


type Auth interface {
	DoAuth() (bool, string, error) // returns is_allowed, error
}

// Authenticator ...
type Authenticator struct {
	Username   string
	AccessKey  string
	URL        string
	HttpClient *http.Client
}

func (auth *Authenticator) BuildAuthURL() string {
	if config.APIURL == "" {
		config.APIURL = "http://localhost:8072"
	}
	return config.APIURL + "/tests/auth"
}

// DoAuth ...
func (auth *Authenticator) DoAuth() (bool, string, error) {
	body, err := json.Marshal(map[string]string{
		"username":   auth.Username,
		"access_key": auth.AccessKey,
	})
	if err != nil {
		return false, "", err
	}
	authURL := auth.BuildAuthURL()
	req, err := http.NewRequest("POST", authURL, bytes.NewBuffer(body))
	if err != nil {
		return false, "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := auth.HttpClient.Do(req)
	if err != nil {
		logrus.Error(err)
		return false, "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		errMsg := "failed to authenticate, error code:  " + strconv.Itoa(resp.StatusCode)
		err = NewAuthError(errMsg)
		logrus.Error(err)
		return false, "", err
	}
	authResp := authResponse{}
	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {
		logrus.Errorf("unexpected resp body: %v", err)
	}
	if !authResp.IsAllowed {
		return false, "not allowed", err
	}
	logrus.Infof("Authenticated for %s", auth.Username)
	return authResp.IsAllowed, authResp.UserID, nil

}

type authResponse struct {
	IsAllowed   bool   `json:"is_allowed"`
	RemainLimit string `json:"remain_limit"`
	UserID      string `json:"user_id"`
}
