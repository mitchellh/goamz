package document_test

import (
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/cloudsearch/document"
	"github.com/mitchellh/goamz/testutil"
	. "github.com/motain/gocheck"
	"io/ioutil"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type S struct {
	document *document.CloudsearchDocument
}

var _ = Suite(&S{})

var testServer = testutil.NewHTTPServer()

func (s *S) SetUpSuite(c *C) {
	testServer.Start()
	auth := aws.Auth{"abc", "123", ""}

	s.document = document.NewWithClient(auth, aws.USEast, testServer.URL, testutil.DefaultClient)
}

func (s *S) TearDownTest(c *C) {
	testServer.Flush()
}

func (s *S) TestSubmitBatch(c *C) {
	testServer.Response(200, nil, SubmitBatchSuccessResponse)
	batch := document.Batch{}
	add := batch.Add("test")
	add.AddField("Foo", "Bar")
	batch.Delete("gone")

	batchResult, err := s.document.SubmitBatch(batch)
	req := testServer.WaitRequest()
	c.Assert(err, IsNil)
	c.Assert(batchResult, NotNil)
	c.Assert(batchResult.Adds, Equals, 1)
	c.Assert(batchResult.Deletes, Equals, 1)

	c.Assert(req.Method, Equals, "POST")
	c.Assert(req.URL.Path, Equals, "/2013-01-01/documents/batch")

	data, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	c.Assert(err, IsNil)
	c.Assert(string(data), Equals, `<batch><add id="test"><field name="Foo">Bar</field></add><delete id="gone"></delete></batch>`)

}

func (s *S) TestSubmitBatchXmlEncoding(c *C) {
	testServer.Response(200, nil, SubmitBatchSuccessResponse)
	batch := document.Batch{}
	add := batch.Add("test")
	add.AddField("Foo", "Bar & Bar")
	batch.Delete("gone")

	batchResult, err := s.document.SubmitBatch(batch)
	req := testServer.WaitRequest()
	c.Assert(err, IsNil)
	c.Assert(batchResult, NotNil)
	c.Assert(batchResult.Adds, Equals, 1)
	c.Assert(batchResult.Deletes, Equals, 1)

	c.Assert(req.Method, Equals, "POST")
	c.Assert(req.URL.Path, Equals, "/2013-01-01/documents/batch")

	data, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	c.Assert(err, IsNil)
	c.Assert(string(data), Equals, `<batch><add id="test"><field name="Foo">Bar &amp; Bar</field></add><delete id="gone"></delete></batch>`)
}

func (s *S) TestSubmitBatchReturnsError(c *C) {
	testServer.Response(200, nil, SubmitBatchErrorResponse)
	batch := document.Batch{}
	add := batch.Add("test")
	add.AddField("Foo", "Bar")
	batch.Delete("gone")

	batchResult, err := s.document.SubmitBatch(batch)
	testServer.WaitRequest()
	c.Assert(err, NotNil)
	c.Assert(batchResult, IsNil)
	c.Assert(err.Error(), Equals, "Something went wrong")

}

func (s *S) TestBatchApi(c *C) {
	batch := document.Batch{}
	add := batch.Add("test")
	add.AddField("Foo", "Bar")
	batch.Delete("gone")

	c.Assert(batch, DeepEquals, document.Batch{Adds: []*document.BatchAdd{&document.BatchAdd{Id: "test", Fields: []document.BatchAddField{document.BatchAddField{"Foo", "Bar"}}}}, Deletes: []document.BatchDelete{document.BatchDelete{"gone"}}})
}
