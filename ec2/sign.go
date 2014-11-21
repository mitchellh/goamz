package ec2

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"net/url"
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

func sorted(data map[string]string) string {
	keys := []string{}
	sarray := []string{}
	for k, _ := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		//sarray = append(sarray, aws.Encode(k)+"="+aws.Encode(data[k]))
		sarray = append(sarray, k+"="+data[k])
	}
	joined := strings.Join(sarray, "&")
	return joined
}

func sign(auth aws.Auth, method, path string, params map[string]string, host string) {
	delete(params, "Timestamp")
	hp := strings.Split(host, ".")
	var service = hp[0]
	var region = hp[1]

	var access_key = auth.AccessKey
	var secret_key = auth.SecretKey

	var amz_date = time.Now().UTC().Format("20060102T150405Z")
	var datestamp = time.Now().UTC().Format("20060102")

	credential_scope := datestamp + "/" + region + "/" + service + "/" + "aws4_request"
	signed_headers := "host" //;x-amz-date"

	canonical_uri := path
	params["X-Amz-Algorithm"] = "AWS4-HMAC-SHA256"
	params["X-Amz-Credential"] = url.QueryEscape(access_key + "/" + credential_scope)
	params["X-Amz-Date"] = amz_date
	params["X-Amz-Expires"] = "30"
	params["X-Amz-SignedHeaders"] = signed_headers
	canonical_querystring := sorted(params)
	canonical_headers := "host:" + host + "\n" // + "x-amz-date:" + amzdate + "\n"

	payload_hash := sha256hex([]byte{}) // empty in case of get

	algorithm := "AWS4-HMAC-SHA256"

	canonical_request := method + "\n" + canonical_uri + "\n" + canonical_querystring + "\n" + canonical_headers + "\n" + signed_headers + "\n" + payload_hash

	string_to_sign := algorithm + "\n" + amz_date + "\n" + credential_scope + "\n" + sha256hex([]byte(canonical_request))

	signing_key := getSignatureKey(secret_key, datestamp, region, service)

	signature := hex.EncodeToString([]byte(signs(signing_key, string_to_sign)))

	params["X-Amz-Signature"] = signature

	params["X-Amz-Credential"] = access_key + "/" + credential_scope
}

var b64 = base64.StdEncoding

func signV2(auth aws.Auth, method, path string, params map[string]string, host string) {
	params["AWSAccessKeyId"] = auth.AccessKey
	params["SignatureVersion"] = "2"
	params["SignatureMethod"] = "HmacSHA256"
	if auth.Token != "" {
		params["SecurityToken"] = auth.Token
	}

	// AWS specifies that the parameters in a signed request must
	// be provided in the natural order of the keys. This is distinct
	// from the natural order of the encoded value of key=value.
	// Percent and equals affect the sorting order.
	var keys, sarray []string
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		sarray = append(sarray, aws.Encode(k)+"="+aws.Encode(params[k]))
	}
	joined := strings.Join(sarray, "&")
	payload := method + "\n" + host + "\n" + path + "\n" + joined
	hash := hmac.New(sha256.New, []byte(auth.SecretKey))
	hash.Write([]byte(payload))
	signature := make([]byte, b64.EncodedLen(hash.Size()))
	b64.Encode(signature, hash.Sum(nil))

	params["Signature"] = string(signature)
}
