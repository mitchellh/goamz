package iam

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/mitchellh/goamz/aws"
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
		payloadHash = hex.EncodeToString(Sha256(payload))

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
	stringToSign := params["X-Amz-Algorithm"] + "\n" + params["X-Amz-Date"] + "\n" + credentialScope + "\n" + hex.EncodeToString(Sha256(canonicalReq))
	signKey := signatureKey(iam.Auth.SecretKey, timestamp, iam.Region, "iam")

	params["X-Amz-Signature"] = hex.EncodeToString(HMac(signKey, stringToSign))
}

func signPostV4(iam *IAM, method, cannocialUri, payload string, headers map[string]string, utcnow time.Time) {
	var (
		timestamp         = utcnow.Format(ISO8601BasicFormatShort)
		amzTimeStamp      = utcnow.Format(ISO8601BasicFormat)
		cannocialQueryStr = ""

		payloadHash = hex.EncodeToString(Sha256(payload))
	)

	headers["X-Amz-Date"] = amzTimeStamp
	cannocialHeaders, signedHeaders := formatHeader(headers)

	canonicalReq := method + "\n" + cannocialUri + "\n" + cannocialQueryStr + "\n" + cannocialHeaders + "\n" + signedHeaders + "\n" + payloadHash
	credentialScope := timestamp + "/" + iam.Region.Name + "/iam/aws4_request"
	stringToSign := Algorithm + "\n" + amzTimeStamp + "\n" + credentialScope + "\n" + hex.EncodeToString(Sha256(canonicalReq))
	signKey := signatureKey(iam.Auth.SecretKey, timestamp, iam.Region, "iam")

	headers["Authorization"] = Algorithm + " " + "Credential=" + iam.Auth.AccessKey + "/" + credentialScope + "," + "SignedHeaders=" + signedHeaders + "," + "Signature=" + hex.EncodeToString(HMac(signKey, stringToSign))
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

func HMac(key, message string) []byte {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

func Sha256(message string) []byte {
	hash := sha256.New()
	hash.Write([]byte(message))
	return hash.Sum(nil)
}

func signatureKey(securityKey string, datestamp string, region aws.Region, serviceName string) string {
	signedDate := string(HMac("AWS4"+securityKey, datestamp))
	signRegion := string(HMac(signedDate, region.Name))
	signSrv := string(HMac(signRegion, serviceName))
	return string(HMac(signSrv, "aws4_request"))
}
