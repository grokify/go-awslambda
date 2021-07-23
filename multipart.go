package awslambda

import (
	"bytes"
	"encoding/base64"
	"errors"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

var (
	ContentTypeHeaderMissingError         = errors.New("content type header missing")
	ContentTypeHeaderNotMultipartError    = errors.New("content type header not multipart error")
	ContentTypeHeaderMissingBoundaryError = errors.New("content type header missing boundary error")
)

func StandardHeader(header map[string]string) http.Header {
	h := http.Header{}
	for k, v := range header {
		h.Add(strings.TrimSpace(k), v)
	}
	return h
}

// NewReaderMultipart returns a `*multipart.Reader` given an API Gateway Proxy
// Request.
func NewReaderMultipart(req events.APIGatewayProxyRequest) (*multipart.Reader, error) {
	headers := StandardHeader(req.Headers)
	ct := headers.Get("content-type")
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

	paramsInsensitiveKeys := StandardHeader(params)
	boundary := paramsInsensitiveKeys.Get("boundary")
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
