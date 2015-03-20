package ecs

import (
	"strings"

	. "github.com/motain/gocheck"
)

// See http://goo.gl/wtteTH for test values
func (s *S) TestBasicSignature(c *C) {
	kSecret := "wJalrXUtnFEMI/K7MDENG+bPxRfiCYEXAMPLEKEY"
	dateStamp := "20110909"
	regionName := "us-east-1"
	serviceName := "iam"
	kSigning := getSignatureKey(kSecret, dateStamp, regionName, serviceName)

	expected := []byte{
		152, 241, 216, 137, 254, 196, 244, 66, 26, 220, 82, 43, 171, 12, 225, 248, 46, 105, 41, 194, 98, 237, 21, 229, 169, 76, 144, 239, 209, 227, 176, 231,
	}

	c.Assert(kSigning, DeepEquals, expected)
}

func (s *S) TestHashCanonicalRequest(c *C) {
	canonicalString := strings.Join([]string{
		"GET",
		"/",
		"Action=DescribeClusters&Version=2014-11-13&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAJTKX7CK2RJVD5MSA%2F20150223%2Fus-east-1%2Fecs%2Faws4_request&X-Amz-Date=20150223T140640Z&X-Amz-SignedHeaders=host",
		"host:ecs.us-east-1.amazonaws.com",
		"",
		"host",
		"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
	}, "\n")

	c.Assert(hexEncode(canonicalString), Equals, "57c130e03223a31650a9a5c762a1024384330c1c132a35029b1c5c468da875f0")
}
