package iam

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/mitchellh/goamz/aws"
	"hash"
	"strings"
	"time"
)

const (
	ISO8601BasicFormat      = "20060102T150405Z"
	ISO8601BasicFormatShort = "20060102"

	Algorithm = "AWS4-HMAC-SHA256"
)

// ----------------------------------------------------------------------------
// Version 4 signing (http://goo.gl/HyL72W), which will support China (Beijing) and EU (Frankfurt)

func signGetV4(iam *IAM, method, cannocialUri, payload string, params, headers map[string]string, utcnow time.Time) {
	var (
		payloadHash = hashAsHex("", payload)

		timestamp    = utcnow.Format(ISO8601BasicFormatShort)
		amzTimeStamp = utcnow.Format(ISO8601BasicFormat)
	)

	cannocialHeaders, signedHeaders := formatHeader(headers)
	credentialScope := timestamp + "/" + iam.Region.Name + "/iam/aws4_request"

	params["X-Amz-Algorithm"] = Algorithm
	params["X-Amz-Date"] = amzTimeStamp
	params["X-Amz-Expires"] = "30"
	params["X-Amz-Credential"] = iam.Auth.AccessKey + "/" + credentialScope
	params["X-Amz-SignedHeaders"] = signedHeaders

	canonicalReq := method + "\n" + cannocialUri + "\n" + multimap(params).Encode() + "\n" + cannocialHeaders + "\n" + signedHeaders + "\n" + payloadHash
	stringToSign := params["X-Amz-Algorithm"] + "\n" + params["X-Amz-Date"] + "\n" + credentialScope + "\n" + hashAsHex("", canonicalReq)
	signKey := getSignatureKey(iam.Auth.SecretKey, timestamp, iam.Region, "iam")

	params["X-Amz-Signature"] = hashAsHex(signKey, stringToSign)
}

func signPostV4(iam *IAM, method, cannocialUri, payload string, headers map[string]string, utcnow time.Time) {
	var (
		timestamp         = utcnow.Format(ISO8601BasicFormatShort)
		amzTimeStamp      = utcnow.Format(ISO8601BasicFormat)
		cannocialQueryStr = ""

		payloadHash = hashAsHex("", payload)
	)

	headers["X-Amz-Date"] = amzTimeStamp
	cannocialHeaders, signedHeaders := formatHeader(headers)

	canonicalReq := method + "\n" + cannocialUri + "\n" + cannocialQueryStr + "\n" + cannocialHeaders + "\n" + signedHeaders + "\n" + payloadHash
	credentialScope := timestamp + "/" + iam.Region.Name + "/iam/aws4_request"
	stringToSign := Algorithm + "\n" + amzTimeStamp + "\n" + credentialScope + "\n" + hashAsHex("", canonicalReq)
	signKey := getSignatureKey(iam.Auth.SecretKey, timestamp, iam.Region, "iam")

	headers["Authorization"] = Algorithm + " " + "Credential=" + iam.Auth.AccessKey + "/" + credentialScope + "," + "SignedHeaders=" + signedHeaders + "," + "Signature=" + hashAsHex(signKey, stringToSign)
}

func formatHeader(headers map[string]string) (string, string) {
	var (
		cHeaders []string
		sHeaders []string
	)

	for k, v := range headers {
		lh := strings.ToLower(k)
		cHeaders = append(cHeaders, lh+":"+v)
		sHeaders = append(sHeaders, lh)
	}

	return strings.Join(cHeaders, "\n") + "\n", strings.Join(sHeaders, ";")
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
