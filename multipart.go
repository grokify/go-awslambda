package awslambda

import (
	"bytes"
	"encoding/base64"
	"errors"
	"mime"
	"mime/multipart"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

var (
	ContentTypeHeaderMissingError         = errors.New("content type header missing")
	ContentTypeHeaderNotMultipartError    = errors.New("content type header not multipart error")
	ContentTypeHeaderMissingBoundaryError = errors.New("content type header missing boundary error")
)

// Headers represents the AWS Lambda Headers definition which is a
// `map[string]string`.
type Headers map[string]string

// MustGet returns a header value or empty string. It matches header names in
// a case insensitive fasion as described in
// https://github.com/aws/aws-lambda-go/issues/117.
func (h Headers) MustGet(key string) string {
	key = strings.ToLower(strings.TrimSpace(key))
	for k, v := range h {
		if strings.ToLower(strings.TrimSpace(k)) == key {
			return v
		}
	}
	return ""
}

// NewReaderMultipart returns a `*multipart.Reader` given an
// API Gateway Proxy Request.
func NewReaderMultipart(req events.APIGatewayProxyRequest) (*multipart.Reader, error) {
	headers := Headers(req.Headers)
	ct := headers.MustGet("content-type")
	if len(ct) == 0 {
		return nil, ContentTypeHeaderMissingError
	}

	mediatype, params, err := mime.ParseMediaType(ct)
	if err != nil {
		return nil, err
	}

	if strings.Index(strings.ToLower(strings.TrimSpace(mediatype)), "multipart/") != 0 {
		return nil, ContentTypeHeaderNotMultipartError
	}

	paramsInsensitiveKeys := Headers(params)
	boundary := strings.TrimSpace(paramsInsensitiveKeys.MustGet("boundary"))
	if len(boundary) == 0 {
		return nil, ContentTypeHeaderMissingBoundaryError
	}

	if req.IsBase64Encoded {
		decoded, err := base64.StdEncoding.DecodeString(req.Body)
		if err != nil {
			return nil, err
		}
		return multipart.NewReader(bytes.NewReader(decoded), boundary), nil
	}

	return multipart.NewReader(strings.NewReader(req.Body), boundary), nil
}
