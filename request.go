package awslambda

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

// NewHTTPRequest returns an `*http.Request` given an `events.APIGatewayProxyRequest`.
func NewHTTPRequest(req events.APIGatewayProxyRequest) (*http.Request, error) {
	// Adaptation from https://stackoverflow.com/a/71727175/1908967
	httpReq, err := http.NewRequestWithContext(
		context.Background(),
		req.HTTPMethod,
		req.Path,
		strings.NewReader(req.Body))
	if err != nil {
		return nil, errors.New("failed to convert a lambda req into a http req")
	}

	// some headers may be important, let get all of them, just in case
	for name, value := range req.Headers {
		httpReq.Header.Add(name, value)
	}
	return httpReq, nil
}
