package document

import (
	"encoding/xml"
	"github.com/mitchellh/goamz/aws"
	"net/http"
	"net/url"
	"bytes"
	"fmt"
	awsauth "github.com/smartystreets/go-aws-auth"
	"strings"
)

type CloudsearchDocument struct {
	aws.Auth
	aws.Region
	endpoint   string
	httpClient *http.Client
}

const apiVersion = "2013-01-01"

func New(auth aws.Auth, region aws.Region, endpoint string) *CloudsearchDocument {
	return NewWithClient(auth, region, endpoint, aws.RetryingClient)
}

func NewWithClient(auth aws.Auth, region aws.Region, endpoint string, httpClient *http.Client) *CloudsearchDocument {
	return &CloudsearchDocument{auth, region, endpoint, httpClient}
}

// Requests
type Batch struct {
	XMLName xml.Name      `xml:"batch"`
	Adds    []*BatchAdd   `xml:"add"`
	Deletes []BatchDelete `xml:"delete"`
}

func (b *Batch) Add(id string) *BatchAdd {
	add := &BatchAdd{id, []BatchAddField{}}
	b.Adds = append(b.Adds, add)
	return add
}
func (b *Batch) Delete(id string) {
	b.Deletes = append(b.Deletes, BatchDelete{id})
}

type BatchDelete struct {
	Id string `xml:"id,attr"`
}
type BatchAdd struct {
	Id     string          `xml:"id,attr"`
	Fields []BatchAddField `xml:"field"`
}

func (ba *BatchAdd) AddField(name string, value string) {
	field := &BatchAddField{name, value}
	ba.Fields = append(ba.Fields, *field)
}

// Yes, there can be multiple fields with the same name
func (ba *BatchAdd) GetFields(name string) []BatchAddField {
	res := []BatchAddField{}
	for _, v := range ba.Fields {
		if v.Name == name {
			res = append(res, v)
		}
	}
	return res
}

type BatchAddField struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",innerxml"`
}

type BatchResult struct {
	Status   string   `xml:"status,attr"`
	Adds     int      `xml:"adds,attr"`
	Deletes  int      `xml:"deletes,attr"`
	Errors   []string `xml:"errors>error"`
	Warnings []string `xml:"warnings>warning"`
}

func (b *BatchResult) Error() string {
	return strings.Join(b.Errors, "\n")
}

func (r *CloudsearchDocument) SubmitBatch(batch Batch) (*BatchResult, error) {
	res := &BatchResult{}
	err := r.query("POST", fmt.Sprintf("/%v/documents/batch", apiVersion), batch, res)
	if err != nil {
		return nil, err
	}
	if res.Status == "error" {
		return nil, res
	}
	return res, nil
}

func (r *CloudsearchDocument) query(method string, path string, request interface{}, resp interface{}) error {
	endpoint, err := url.Parse(r.endpoint)
	endpoint.Path = path
	xmlBytes, err := xml.Marshal(request)
	if err != nil {
		return err
	}
	xmlReader := bytes.NewReader(xmlBytes)
	hReq, err := http.NewRequest(method, endpoint.String(), xmlReader)
	hReq.Header.Set("Content-Type", "application/xml")
	hReq.Close = true
	if err != nil {
		return err
	}

	awsauth.Sign(hReq, awsauth.Credentials{
		AccessKeyID:     r.Auth.AccessKey,
		SecretAccessKey: r.Auth.SecretKey,
	})

	re, err := r.httpClient.Do(hReq)
	if err != nil {
		return err
	}
	defer re.Body.Close()

	if re.StatusCode > 200 {
		return fmt.Errorf("Cloudsearch returned unexpected http %v", re.Status)
	}
	return xml.NewDecoder(re.Body).Decode(resp)
}
