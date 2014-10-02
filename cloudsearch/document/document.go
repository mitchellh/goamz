package document

import (
	"net/http"
	"github.com/mitchellh/goamz/aws"
	"net/url"
	"encoding/xml"
	"strconv"
	awsauth "github.com/smartystreets/go-aws-auth"
	"bytes"
	"fmt"
)

type CloudsearchDocument struct {
	aws.Auth
	aws.Region
	endpoint string
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
	XMLName           xml.Name    `xml:"batch"`
	Adds              []*BatchAdd    `xml:"add"`
	Deletes           []BatchDelete `xml:"delete"`
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
	Id             string  `xml:"id,attr"`
}
type BatchAdd struct {
	Id               string           `xml:"id,attr"`
	Fields           []BatchAddField  `xml:"field"`
}

func (ba *BatchAdd) AddField(name string, value string) {
	field := &BatchAddField{name, value}
	ba.Fields = append(ba.Fields, *field)
}

type BatchAddField struct {
	Name           string    `xml:"name,attr"`
	Value          string    `xml:",innerxml"`
}

type BatchResult struct {
	Status            string `xml:"status,attr"`
	Adds              int    `xml:"adds,attr"`
	Deletes           int    `xml:"deletes,attr"`
}

type Error struct {
	// HTTP status code of the error.
	StatusCode int

	// AWS code of the error.
	Code string

	// Message explaining the error.
	Message string
}

func (e *Error) Error() string {
	var prefix string
	if e.Code != "" {
		prefix = e.Code+": "
	}
	if prefix == "" && e.StatusCode > 0 {
		prefix = strconv.Itoa(e.StatusCode)+": "
	}
	return prefix + e.Message
}

type xmlErrors struct {
	Errors []Error `xml:"Error"`
}

func (r *CloudsearchDocument) SubmitBatch(batch Batch) (*BatchResult, error) {

	res := &BatchResult{}
	err := r.query("POST", fmt.Sprintf("/%v/documents/batch", apiVersion), batch, res)
	switch {
	case err != nil: return nil, err
	default: return res, nil
	}

}

func (r *CloudsearchDocument) query(method string, path string, request interface{}, resp interface{}) error {
	endpoint, err := url.Parse(r.endpoint)
	endpoint.Path = path
	xmlBytes, _ := xml.Marshal(request)
	xmlReader := bytes.NewReader(xmlBytes)
	hReq, err := http.NewRequest(method, endpoint.String(), xmlReader)
	hReq.Header.Set("Content-Type", "application/xml")
	if err != nil {
		return err
	}

	awsauth.Sign(hReq, awsauth.Credentials{
		AccessKeyID: r.Auth.AccessKey,
		SecretAccessKey: r.Auth.SecretKey,
	})

	re, err := r.httpClient.Do(hReq)
	if err != nil {
		return err
	}
	defer re.Body.Close()

	if re.StatusCode > 200 {
		return buildError(re)
	}
	decoder := xml.NewDecoder(re.Body)
	decodedBody := decoder.Decode(resp)
	return decodedBody
}

func buildError(r *http.Response) error {
	var (
		err    Error
		errors xmlErrors
	)
	xml.NewDecoder(r.Body).Decode(&errors)
	if len(errors.Errors) > 0 {
		err = errors.Errors[0]
	}
	err.StatusCode = r.StatusCode
	if err.Message == "" {
		err.Message = r.Status
	}
	return &err
}
