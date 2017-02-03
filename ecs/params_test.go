package ecs

import (
	"testing"

	. "github.com/motain/gocheck"
)

func Test(t *testing.T) {
	TestingT(t)
}

var _ = Suite(&S{})

type S struct {
	ecs *ECS
}

func (s *S) TestMakeParams(c *C) {
	params := makeParams("Sample")
	value, found := params["Action"]
	c.Assert(value, Equals, "Sample")
	c.Assert(found, Equals, true)
}

type Mock struct {
	S1  string   `form:"string1"`
	S2  string   `form:"string1"`
	B1  bool     `form:"bool1"`
	B2  bool     `form:"bool2"`
	I1  int      `form:"int1"`
	I2  int      `form:"int2"`
	SA1 []string `form:"array1"`
	SA2 []string `form:"array2"`
}

func (s *S) TestSetParameters(c *C) {
	params := makeParams("blah")
	params.Set(Mock{
		S1:  "hello",
		B1:  true,
		I1:  1,
		SA1: []string{"a", "b", "c"},
	})

	c.Assert(params["string1"], Equals, "hello")
	c.Assert(params["bool1"], Equals, "true")
	c.Assert(params["bool2"], Equals, "false")
	c.Assert(params["int1"], Equals, "1")

	c.Assert(params["array1.1"], Equals, "a")
	c.Assert(params["array1.2"], Equals, "b")
	c.Assert(params["array1.3"], Equals, "c")

	var found bool
	_, found = params["string2"]
	c.Assert(found, Equals, false)

	_, found = params["int2"]
	c.Assert(found, Equals, false)

	_, found = params["array2.1"]
	c.Assert(found, Equals, false)
}

func (s *S) TestHexEncode(c *C) {
	params := makeParams("ListUsers")
	params["Version"] = "2010-05-08"
	c.Assert(hexEncode(params.encoded()), Equals, "b6359072c78d70ebee1e81adcbab4f01bf2c23245fa365ef83fe8f1f955085e2")
}
