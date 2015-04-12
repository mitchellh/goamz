package aws_test

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/goamz/aws"
	. "github.com/motain/gocheck"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
	"testing"
	"time"
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
	c.Assert(auth.AccessKey, Equals, "access")
	c.Assert(auth.SecretKey, Equals, "secret")
}

func (s *S) TestEnvAuthWithToken(c *C) {
	os.Clearenv()
	os.Setenv("AWS_SECRET_ACCESS_KEY", "secret")
	os.Setenv("AWS_ACCESS_KEY_ID", "access")
	os.Setenv("AWS_SECURITY_TOKEN", "token")
	auth, err := aws.EnvAuth()
	c.Assert(err, IsNil)
	c.Assert(auth.AccessKey, Equals, "access")
	c.Assert(auth.SecretKey, Equals, "secret")
	c.Assert(auth.Token(), Equals, "token")
}

func (s *S) TestEnvAuthAlt(c *C) {
	os.Clearenv()
	os.Setenv("AWS_SECRET_KEY", "secret")
	os.Setenv("AWS_ACCESS_KEY", "access")
	auth, err := aws.EnvAuth()
	c.Assert(err, IsNil)
	c.Assert(auth.AccessKey, Equals, "access")
	c.Assert(auth.SecretKey, Equals, "secret")
}

func (s *S) TestGetAuthStatic(c *C) {
	auth, err := aws.GetAuth("access", "secret")
	c.Assert(err, IsNil)
	c.Assert(auth.AccessKey, Equals, "access")
	c.Assert(auth.SecretKey, Equals, "secret")
}

func (s *S) TestGetAuthEnv(c *C) {
	os.Clearenv()
	os.Setenv("AWS_SECRET_ACCESS_KEY", "secret")
	os.Setenv("AWS_ACCESS_KEY_ID", "access")
	auth, err := aws.GetAuth("", "")
	c.Assert(err, IsNil)
	c.Assert(auth.AccessKey, Equals, "access")
	c.Assert(auth.SecretKey, Equals, "secret")
}

type fakeIamServer struct {
	listener  net.Listener
	nRequests uint32
}

func (s *S) newFakeIamServer(c *C, expirationDelay time.Duration) *fakeIamServer {
	listenAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	c.Assert(err, IsNil)

	listener, err := net.ListenTCP("tcp", listenAddr)
	c.Assert(err, IsNil)

	iamServer := &fakeIamServer{
		listener: listener,
	}

	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		creds := struct {
			AccessKeyId     string
			SecretAccessKey string
			Token           string
			Expiration      string
		}{
			"access",
			"secret",
			"token",
			time.Now().UTC().Add(expirationDelay).Format("2006-01-02T15:04:05Z"),
		}

		atomic.AddUint32(&iamServer.nRequests, 1)

		json.NewEncoder(w).Encode(creds)
	})

	server := &http.Server{
		Handler: handler,
	}

	go server.Serve(listener)

	return iamServer
}

func (r *fakeIamServer) Stop() {
	r.listener.Close()
}

func (r *fakeIamServer) NRequests() uint32 {
	return atomic.LoadUint32(&r.nRequests)
}

func (s *S) TestRoleAuth(c *C) {
	server := s.newFakeIamServer(c, 5*time.Hour)
	defer server.Stop()

	os.Setenv("GOAWS_INSTANCE_METADATA_URL", fmt.Sprintf("http://%s/", server.listener.Addr().String()))

	auth, err := aws.RoleAuth()
	c.Assert(err, IsNil)
	c.Assert(auth.AccessKey, Equals, "access")
	c.Assert(auth.SecretKey, Equals, "secret")
	c.Assert(auth.Token(), Equals, "token")
	// 2 since one request for fetching the role, one for the token
	c.Assert(server.NRequests(), Equals, uint32(2))

	for i := 0; i < 10; i++ {
		auth.Token()
	}

	// Many calls to Token should not hit IAM if the token is not expired
	c.Assert(server.NRequests(), Equals, uint32(2))
}

func (s *S) TestRoleAuthExpiration(c *C) {
	server := s.newFakeIamServer(c, 0*time.Hour)
	defer server.Stop()

	os.Setenv("GOAWS_INSTANCE_METADATA_URL", fmt.Sprintf("http://%s/", server.listener.Addr().String()))

	auth, err := aws.RoleAuth()
	c.Assert(err, IsNil)

	for i := 0; i < 10; i++ {
		auth.Token()
	}

	c.Assert(server.NRequests(), Equals, uint32(22)) // (Initial + 10 expiries) x 2 requests
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
