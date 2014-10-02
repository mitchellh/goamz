package document_test

import (
	"github.com/mitchellh/goamz/cloudsearch/document"
	. "github.com/motain/gocheck"
	"testing"
	"encoding/xml"
)

func Test(t *testing.T) {
	TestingT(t)
}

type S struct {
	document *document.CloudsearchDocument
}

var _ = Suite(&S{})

func (s *S) TestBatchSerialization(c *C) {
	batch := &document.Batch{}
	add := batch.Add("test")
	add.AddField("Foo", "Bar")
	batch.Delete("gone")

	c.Assert(*batch, DeepEquals, document.Batch{Adds: []*document.BatchAdd{&document.BatchAdd{Id: "test", Fields: []document.BatchAddField{document.BatchAddField{"Foo", "Bar"}}}}, Deletes: []document.BatchDelete{document.BatchDelete{"gone"}}})

	x, _ := xml.MarshalIndent(batch, " ", " ")
	println(string(x))
	//	c.Assert(req.Form["Action"], DeepEquals, []string{"CreateDBSecurityGroup"})
	//	c.Assert(req.Form["DBSecurityGroupName"], DeepEquals, []string{"foobarbaz"})
	//	c.Assert(req.Form["DBSecurityGroupDescription"], DeepEquals, []string{"test description"})
	//	c.Assert(err, IsNil)
	//	c.Assert(resp.RequestId, Equals, "e68ef6fa-afc1-11c3-845a-476777009d19")
}
