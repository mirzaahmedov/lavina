package tools

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"log"
)

type SignParams struct {
	Method string
	IsTLS  bool
	Host   string
	Path   string
	Body   string
	Secret string
}

func CompareSign(sign string, params *SignParams) bool {
	protocol := "http"

	if params.IsTLS {
		protocol += "s"
	}

	payload := params.Method + protocol + "://" + params.Host + params.Path + params.Body + params.Secret

	log.Println(payload)

	sum := make([]byte, 16)
	if _, err := hex.Decode(sum, []byte(sign)); err != nil {
		return false
	} else {
		result := md5.Sum([]byte(payload))
		return bytes.Equal(result[:], sum)
	}
}