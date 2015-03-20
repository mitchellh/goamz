package ecs

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/mitchellh/goamz/aws"
)

const (
	AmzDateTimeFormat = "20060102T150405Z"
	AmzDateFormat     = "20060102"
	AmzAlgorithm      = "AWS4-HMAC-SHA256"
)

const (
	service = "ecs"
)

// See http://goo.gl/xZjQRs for Version 4 signing details
func sign(auth aws.Auth, region aws.Region, host string, params parameters) {
	now := time.Now().UTC()
	amzDateTime := now.Format(AmzDateTimeFormat)
	amzDate := now.Format(AmzDateFormat)

	// ---- add parameters ---------------------------------

	credentialScope := fmt.Sprintf("%s/%s/%s/aws4_request", amzDate, region.Name, service)
	signedHeaders := "host" // only signing host header

	params["X-Amz-Credential"] = auth.AccessKey + "/" + credentialScope
	params["X-Amz-Algorithm"] = AmzAlgorithm
	params["X-Amz-Date"] = amzDateTime
	params["X-Amz-SignedHeaders"] = signedHeaders

	// ---- generate the canonical request -----------------

	method := "GET"
	path := "/"
	canonicalQueryString := params.encoded()
	canonicalHeaders := fmt.Sprintf("host:%s", host)
	requestPayload := hexEncode("")

	canonicalRequest := strings.Join([]string{
		method,
		path,
		canonicalQueryString,
		canonicalHeaders,
		"",
		signedHeaders,
		requestPayload,
	}, "\n")

	// ---- create the string to sign ----------------------

	requestDate := amzDateTime

	stringToSign := strings.Join([]string{
		AmzAlgorithm,
		requestDate,
		credentialScope,
		hexEncode(canonicalRequest),
	}, "\n")

	// ---- calculate the signature ------------------------

	kSigning := getSignatureKey(auth.SecretKey, amzDate, region.Name, service)
	kSignature := hmacSHA256(kSigning, stringToSign)
	signature := hex.EncodeToString(kSignature)

	// ---- add to request ---------------------------------

	params["X-Amz-Signature"] = signature
}

func getSignatureKey(secretKey, dateStamp, regionName, serviceName string) []byte {
	kDate := hmacSHA256([]byte("AWS4"+secretKey), dateStamp)
	kRegion := hmacSHA256(kDate, regionName)
	kService := hmacSHA256(kRegion, serviceName)
	return hmacSHA256(kService, "aws4_request")
}

func hmacSHA256(base []byte, plus string) []byte {
	hash := hmac.New(sha256.New, base)
	hash.Write([]byte(plus))
	return hash.Sum(nil)
}

func hexEncode(text string) string {
	hasher := sha256.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
