package auth

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Authenticator struct {
	Username  string
	AccessKey string
	Url       string
}

func (auth *Authenticator) redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.SetBasicAuth(auth.Username, auth.AccessKey)
	return nil
}

func (auth *Authenticator) DoAuth() (bool, string) {
	client := http.Client{
		CheckRedirect: auth.redirectPolicyFunc,
	}
	req, err := http.NewRequest("GET", "http://localhost:8072/tests/auth", nil)
	if err != nil {
		//
	}
	req.SetBasicAuth(auth.Username, auth.AccessKey)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("oops: %v", err)
		return false, ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("oops")
		return false, ""
	}
	authResp := authResponse{}
	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {
		logrus.Errorf("unexpected resp body: %v", err)
	}
	if !authResp.IsAllowed {
		return false, "not allowed"
	}
	fmt.Printf("auth went well \n")
	return authResp.IsAllowed, authResp.RemainLimit

}

type authResponse struct {
	IsAllowed   bool   `json:"is_allowed"`
	RemainLimit string `json:"remain_limit"`
}
