package iam

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/mitchellh/goamz/aws"
	"hash"
	"time"
)

const (
	ISO8601BasicFormat      = "20060102T150405Z"
	ISO8601BasicFormatShort = "20060102"
)

// ----------------------------------------------------------------------------
// Version 4 signing (http://goo.gl/HyL72W), which will support China (Beijing) and EU (Frankfurt)

func signV4(iam *IAM, method, path string, params map[string]string, host string) {
	var (
		now               = time.Now().In(time.UTC)
		cannocial_uri     = path
		cannocial_headers = "host:" + host + "\n"
		signed_headers    = "host"
		credential_scope  = now.Format(ISO8601BasicFormatShort) + "/" + iam.Region.Name + "/iam/aws4_request"
		payloadHash       = hashAsHex("", "")
	)

	params["Version"] = "2010-05-08"
	params["X-Amz-Algorithm"] = "AWS4-HMAC-SHA256"
	params["X-Amz-Date"] = now.Format(ISO8601BasicFormat)
	params["X-Amz-Expires"] = "30"
	params["X-Amz-Credential"] = iam.Auth.AccessKey + "/" + credential_scope
	params["X-Amz-SignedHeaders"] = signed_headers

	canonical_req := method + "\n" +
		cannocial_uri + "\n" +
		multimap(params).Encode() + "\n" +
		cannocial_headers + "\n" +
		signed_headers + "\n" + payloadHash

	stringToSign := params["X-Amz-Algorithm"] + "\n" +
		params["X-Amz-Date"] + "\n" +
		credential_scope + "\n" +
		hashAsHex("", canonical_req)

	signKey := getSignatureKey(iam.Auth.SecretKey, now.Format(ISO8601BasicFormatShort), iam.Region, "iam")

	params["X-Amz-Signature"] = hashAsHex(signKey, stringToSign)
}

func HashAsStr(key string, target string) string {
	var hash hash.Hash

	if key == "" {
		hash = sha256.New()
	} else {
		hash = hmac.New(sha256.New, []byte(key))
	}
	hash.Write([]byte(target))
	return string(hash.Sum(nil))
}

func hashAsHex(key string, target string) string {
	var hash hash.Hash

	if key == "" {
		hash = sha256.New()
	} else {
		hash = hmac.New(sha256.New, []byte(key))
	}
	hash.Write([]byte(target))
	return hex.EncodeToString(hash.Sum(nil))
}

func getSignatureKey(securityKey string, datestamp string, region aws.Region, serviceName string) string {
	signedDate := HashAsStr("AWS4"+securityKey, datestamp)
	signRegion := HashAsStr(signedDate, region.Name)
	signSrv := HashAsStr(signRegion, serviceName)
	return HashAsStr(signSrv, "aws4_request")
}
