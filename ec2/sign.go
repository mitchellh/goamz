package ec2

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"
	"time"

	"github.com/mitchellh/goamz/aws"
)

// ----------------------------------------------------------------------------
// EC2 signing (http://goo.gl/fQmAN)

func sha256hex(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}

func signs(key, msg string) string {
	x := hmac.New(sha256.New, []byte(key))
	x.Write([]byte(msg))
	return string(x.Sum(nil))
}

func getSignatureKey(key, dateStamp, regionName, serviceName string) string {
	kDate := signs("AWS4"+key, dateStamp)
	kRegion := signs(kDate, regionName)
	kService := signs(kRegion, serviceName)
	kSigning := signs(kService, "aws4_request")
	return kSigning
}

func makeSortedRawQuery(data map[string]string) string {
	keys := []string{}
	sarray := []string{}
	for k, _ := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		sarray = append(sarray, aws.Encode(k)+"="+aws.Encode(data[k]))
	}
	joined := strings.Join(sarray, "&")
	return joined
}

func sign(auth aws.Auth, method, canonical_uri string, params map[string]string, host string) {
	hp := strings.Split(host, ".")
	var service = hp[0]
	var region = hp[1]

	var access_key = auth.AccessKey
	var secret_key = auth.SecretKey

	var amz_date = time.Now().UTC().Format("20060102T150405Z")
	var datestamp = time.Now().UTC().Format("20060102")

	credential_scope := datestamp + "/" + region + "/" + service + "/" + "aws4_request"
	signed_headers := "host"

	algorithm := "AWS4-HMAC-SHA256"

	params["X-Amz-Algorithm"] = algorithm
	params["X-Amz-Credential"] = access_key + "/" + credential_scope
	params["X-Amz-Date"] = amz_date
	params["X-Amz-Expires"] = "30"
	params["X-Amz-SignedHeaders"] = signed_headers
	canonical_querystring := makeSortedRawQuery(params)
	canonical_headers := "host:" + host + "\n"

	payload_hash := sha256hex([]byte{}) // empty in case of get

	canonical_request := method + "\n" + canonical_uri + "\n" + canonical_querystring + "\n" + canonical_headers + "\n" + signed_headers + "\n" + payload_hash

	string_to_sign := algorithm + "\n" + amz_date + "\n" + credential_scope + "\n" + sha256hex([]byte(canonical_request))

	signing_key := getSignatureKey(secret_key, datestamp, region, service)

	signature := hex.EncodeToString([]byte(signs(signing_key, string_to_sign)))

	params["X-Amz-Signature"] = signature
}
