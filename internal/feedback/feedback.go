package feedback

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type FeedbackProcessor struct {
	Username   string
	AccessKey  string
	HttpClient *http.Client
	APIUrl string
}

func NewFeedbackProcessor(username string, accessKey string, apiUrl string) *FeedbackProcessor {
	feedProc := &FeedbackProcessor{
		Username:   username,
		AccessKey:  accessKey,
		HttpClient: &http.Client{},
		APIUrl: apiUrl,
	}
	return feedProc
}

func (f *FeedbackProcessor) collectAnswers() (error, int, int, int, float64, string) {
	var easyToUse int
	fmt.Println(FirstQuestion)
	_, err := fmt.Scan(&easyToUse)
	if err != nil {
		return err, 0, 0, 0, 0, ""
	}

	var likelinessOfRecommend int
	fmt.Println(SecondQuestion)
	_, err = fmt.Scan(&likelinessOfRecommend)
	if err != nil {
		return err, 0, 0, 0, 0, ""
	}

	var overallScore int
	fmt.Println(ThirdQuestion)
	_, err = fmt.Scan(&overallScore)
	if err != nil {
		return err, 0, 0, 0, 0, ""
	}

	var proposedPrice float64
	fmt.Println(FourthQuestion)
	_, err = fmt.Scan(&proposedPrice)
	if err != nil {
		return err, 0, 0, 0, 0, ""
	}

	//reading a string
	var openFeedback string
	fmt.Println(FifthQuestion)
	r := bufio.NewReader(os.Stdin)
	openFeedback, err = r.ReadString('\n')
	if err != nil {
		return err, 0, 0, 0, 0, ""
	}
	fmt.Println(ThanksCopy)
	return nil, easyToUse, likelinessOfRecommend, overallScore, proposedPrice, openFeedback
}

type feedbackRequest struct {
	EaseOfUse             int     `json:"ease_of_use"`
	LikelinessOfRecommend int     `json:"likeliness_of_recommend"`
	OverallScore          int     `json:"overall_score"`
	ProposedPrice         float64 `json:"proposed_price"`
	OpenFeedback          string  `json:"open_feedback"`
}

func (f *FeedbackProcessor) buildURL() string {
	return f.APIUrl + "/feedback/"
}

func (f *FeedbackProcessor) sendFeedback(fReq *feedbackRequest) error {
	body, err := json.Marshal(fReq)
	if err != nil {
		return err
	}
	uri := f.buildURL()
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(body))
	header := f.Username + ":" + f.AccessKey
	req.Header.Set("X-CLI-CREDS", header)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	resp, err := f.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return err
	}
	return nil
}

func (f *FeedbackProcessor) Process() error {
	err, EaseOfUse, LikelinessOfRecommend, OverallScore, ProposedPrice, OpenFeedback := f.collectAnswers()
	if err != nil {
		return err
	}
	fReq := &feedbackRequest{
		EaseOfUse:             EaseOfUse,
		LikelinessOfRecommend: LikelinessOfRecommend,
		OverallScore:          OverallScore,
		ProposedPrice:         ProposedPrice,
		OpenFeedback:          OpenFeedback,
	}
	err = f.sendFeedback(fReq)
	if err != nil {
		return err
	}
	fmt.Println("Your feedback is sent.")
	return nil
}
