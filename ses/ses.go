package ses

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/mitchellh/goamz/aws"
)

type Server struct {
	aws.Auth
	aws.Region
}

func (s *Server) SendEmail(from, to, cc, subject, body string) (string, error) {
	data := make(url.Values)
	data.Add("Action", "SendEmail")
	data.Add("Source", from)
	data.Add("Destination.ToAddresses.member.1", to)
	data.Add("Message.Subject.Data", subject)
	//data.Add("Message.Body.Text.Data", body)
	data.Add("Message.Body.Html.Data", body)
	data.Add("AWSAccessKeyId", s.Auth.AccessKey)

	return s.sesGet(data)
}

func (s *Server) authorizationHeader(date string) []string {
	h := hmac.New(sha256.New, []uint8(s.Auth.SecretKey))
	h.Write([]uint8(date))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	auth := fmt.Sprintf("AWS3-HTTPS AWSAccessKeyId=%s, Algorithm=HmacSHA256, Signature=%s", s.Auth.AccessKey, signature)
	return []string{auth}
}

func (s *Server) sesGet(data url.Values) (string, error) {
	headers := http.Header{}

	now := time.Now().UTC()
	// date format: "Tue, 25 May 2010 21:20:27 +0000"
	date := now.Format("Mon, 02 Jan 2006 15:04:05 -0700")
	headers.Set("Date", date)

	h := hmac.New(sha256.New, []uint8(s.Auth.SecretKey))
	h.Write([]uint8(date))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	auth := fmt.Sprintf("AWS3-HTTPS AWSAccessKeyId=%s, Algorithm=HmacSHA256, Signature=%s", s.Auth.AccessKey, signature)
	headers.Set("X-Amzn-Authorization", auth)

	headers.Set("Content-Type", "application/x-www-form-urlencoded")

	body := strings.NewReader(data.Encode())
	req, err := http.NewRequest("POST", s.Region.SESEndpoint, body)
	if err != nil {
		return "", err
	}

	if s.Auth.Token != "" {
		headers.Set("X-Amz-Security-Token", s.Auth.Token)
		//fmt.Printf("Ali: SecToken = %s \n", s.Auth.Token)
	}

	req.Header = headers

	//c.Debugf("%+v", req)

	//client := urlfetch.Client(c)
	client := &http.Client{}

	r, err := client.Do(req)
	if err != nil {
		return "", err
	}

	resultbody, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	if r.StatusCode != 200 {
		return "", fmt.Errorf("error, status = %d; response = %s", r.StatusCode, resultbody)
	}

	return string(resultbody), nil
}
