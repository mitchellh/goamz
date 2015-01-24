package aws_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/mitchellh/goamz/aws"
	. "github.com/motain/gocheck"
)

func Test(t *testing.T) {
	TestingT(t)
}

var _ = Suite(&S{})

type S struct {
	environ []string
}

func (s *S) SetUpSuite(c *C) {
	s.environ = os.Environ()
}

func (s *S) TearDownTest(c *C) {
	os.Clearenv()
	for _, kv := range s.environ {
		l := strings.SplitN(kv, "=", 2)
		os.Setenv(l[0], l[1])
	}
}

func (s *S) TestSharedAuthNoHome(c *C) {
	os.Clearenv()
	os.Setenv("AWS_PROFILE", "foo")
	_, err := aws.SharedAuth()
	c.Assert(err, ErrorMatches, "Could not get HOME")
}

func (s *S) TestSharedAuthNoCredentialsFile(c *C) {
	os.Clearenv()
	os.Setenv("AWS_PROFILE", "foo")
	os.Setenv("HOME", "/tmp")
	_, err := aws.SharedAuth()
	c.Assert(err, ErrorMatches, "Couldn't parse AWS credentials file")
}

func (s *S) TestSharedAuthNoProfileInFile(c *C) {
	os.Clearenv()
	os.Setenv("AWS_PROFILE", "foo")

	d, err := ioutil.TempDir("", "")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(d)

	err = os.Mkdir(d+"/.aws", 0755)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(d+"/.aws/credentials", []byte("[bar]\n"), 0644)
	os.Setenv("HOME", d)

	_, err = aws.SharedAuth()
	c.Assert(err, ErrorMatches, "Couldn't find profile in AWS credentials file")
}

func (s *S) TestSharedAuthNoKeysInProfile(c *C) {
	os.Clearenv()
	os.Setenv("AWS_PROFILE", "bar")

	d, err := ioutil.TempDir("", "")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(d)

	err = os.Mkdir(d+"/.aws", 0755)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(d+"/.aws/credentials", []byte("[bar]\nawsaccesskeyid = AK.."), 0644)
	os.Setenv("HOME", d)

	_, err = aws.SharedAuth()
	c.Assert(err, ErrorMatches, "AWS_SECRET_ACCESS_KEY not found in credentials file")
}

func (s *S) TestSharedAuthDefaultCredentials(c *C) {
	os.Clearenv()

	d, err := ioutil.TempDir("", "")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(d)

	err = os.Mkdir(d+"/.aws", 0755)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(d+"/.aws/credentials", []byte("[default]\naws_access_key_id = access\naws_secret_access_key = secret\n"), 0644)
	os.Setenv("HOME", d)

	auth, err := aws.SharedAuth()
	c.Assert(err, IsNil)
	c.Assert(auth, Equals, aws.Auth{SecretKey: "secret", AccessKey: "access"})
}

func (s *S) TestSharedAuth(c *C) {
	os.Clearenv()
	os.Setenv("AWS_PROFILE", "bar")

	d, err := ioutil.TempDir("", "")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(d)

	err = os.Mkdir(d+"/.aws", 0755)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(d+"/.aws/credentials", []byte("[bar]\naws_access_key_id = access\naws_secret_access_key = secret\n"), 0644)
	os.Setenv("HOME", d)

	auth, err := aws.SharedAuth()
	c.Assert(err, IsNil)
	c.Assert(auth, Equals, aws.Auth{SecretKey: "secret", AccessKey: "access"})
}

func (s *S) TestEnvAuthNoSecret(c *C) {
	os.Clearenv()
	_, err := aws.EnvAuth()
	c.Assert(err, ErrorMatches, "AWS_SECRET_ACCESS_KEY or AWS_SECRET_KEY not found in environment")
}

func (s *S) TestEnvAuthNoAccess(c *C) {
	os.Clearenv()
	os.Setenv("AWS_SECRET_ACCESS_KEY", "foo")
	_, err := aws.EnvAuth()
	c.Assert(err, ErrorMatches, "AWS_ACCESS_KEY_ID or AWS_ACCESS_KEY not found in environment")
}

func (s *S) TestEnvAuth(c *C) {
	os.Clearenv()
	os.Setenv("AWS_SECRET_ACCESS_KEY", "secret")
	os.Setenv("AWS_ACCESS_KEY_ID", "access")
	auth, err := aws.EnvAuth()
	c.Assert(err, IsNil)
	c.Assert(auth, Equals, aws.Auth{SecretKey: "secret", AccessKey: "access"})
}

func (s *S) TestEnvAuthWithToken(c *C) {
	os.Clearenv()
	os.Setenv("AWS_SECRET_ACCESS_KEY", "secret")
	os.Setenv("AWS_ACCESS_KEY_ID", "access")
	os.Setenv("AWS_SECURITY_TOKEN", "token")
	auth, err := aws.EnvAuth()
	c.Assert(err, IsNil)
	c.Assert(auth, Equals, aws.Auth{SecretKey: "secret", AccessKey: "access", Token: "token"})
}

func (s *S) TestEnvAuthAlt(c *C) {
	os.Clearenv()
	os.Setenv("AWS_SECRET_KEY", "secret")
	os.Setenv("AWS_ACCESS_KEY", "access")
	auth, err := aws.EnvAuth()
	c.Assert(err, IsNil)
	c.Assert(auth, Equals, aws.Auth{SecretKey: "secret", AccessKey: "access"})
}

func (s *S) TestGetAuthStatic(c *C) {
	auth, err := aws.GetAuth("access", "secret")
	c.Assert(err, IsNil)
	c.Assert(auth, Equals, aws.Auth{SecretKey: "secret", AccessKey: "access"})
}

func (s *S) TestGetAuthEnv(c *C) {
	os.Clearenv()
	os.Setenv("AWS_SECRET_ACCESS_KEY", "secret")
	os.Setenv("AWS_ACCESS_KEY_ID", "access")
	auth, err := aws.GetAuth("", "")
	c.Assert(err, IsNil)
	c.Assert(auth, Equals, aws.Auth{SecretKey: "secret", AccessKey: "access"})
}

func (s *S) TestEncode(c *C) {
	c.Assert(aws.Encode("foo"), Equals, "foo")
	c.Assert(aws.Encode("/"), Equals, "%2F")
}

func (s *S) TestRegionsAreNamed(c *C) {
	for n, r := range aws.Regions {
		c.Assert(n, Equals, r.Name)
	}
}

func (s *S) TestEnvRegionNoRegion(c *C) {
	_, err := aws.EnvRegion()
	c.Assert(err, ErrorMatches, "AWS_REGION or aws_region not found in environment")
}

func (s *S) TestEnvRegion(c *C) {
	os.Clearenv()
	os.Setenv("AWS_REGION", "eu-west-1")
	region, err := aws.EnvRegion()
	c.Assert(err, IsNil)
	c.Assert(region.Name, Equals, "eu-west-1")
}

func (s *S) TestEnvRegionAlt(c *C) {
	os.Clearenv()
	os.Setenv("aws_region", "eu-west-1")
	region, err := aws.EnvRegion()
	c.Assert(err, IsNil)
	c.Assert(region.Name, Equals, "eu-west-1")
}

func (s *S) TestEnvRegionInvalid(c *C) {
	os.Clearenv()
	os.Setenv("AWS_REGION", "eu-west-never")
	_, err := aws.EnvRegion()
	errorString := fmt.Sprintf("%v region not found", os.Getenv("AWS_REGION"))
	c.Assert(err, ErrorMatches, errorString)
}

func (s *S) TestGetRegionStatic(c *C) {
	region, err := aws.GetRegion("eu-west-1")
	c.Assert(err, IsNil)
	c.Assert(region.Name, Equals, "eu-west-1")
}

func (s *S) TestGetRegionEnv(c *C) {
	os.Clearenv()
	os.Setenv("AWS_REGION", "eu-west-1")
	region, err := aws.GetRegion("")
	c.Assert(err, IsNil)
	c.Assert(region.Name, Equals, "eu-west-1")
}
