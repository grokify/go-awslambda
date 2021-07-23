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

type Headers map[string]string

func (h Headers) MustGet(headerName string) string {
	headerName = strings.ToLower(strings.TrimSpace(headerName))
	for key, val := range h {
		if strings.ToLower(strings.TrimSpace(key)) == headerName {
			return val
		}
	}
	return ""
}

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

	boundary, ok := params["boundary"]
	if !ok {
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
