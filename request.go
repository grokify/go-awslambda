package awslambda

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

// NewHTTPRequest returns an `*http.Request` given an `events.APIGatewayProxyRequest`.
func NewHTTPRequest(ctx context.Context, req events.APIGatewayProxyRequest) (*http.Request, error) {
	// Adaptation from https://stackoverflow.com/a/71727175/1908967
	if httpReq, err := http.NewRequestWithContext(
		ctx,
		req.HTTPMethod,
		req.Path,
		strings.NewReader(req.Body)); err != nil {
		return nil, wrapError("failed to convert lambda request to http request", err)
	} else {
		// some headers may be important, let get all of them, just in case
		for name, value := range req.Headers {
			httpReq.Header.Add(name, value)
		}
		return httpReq, nil
	}
}

func wrapError(msg string, err error) error {
	if err == nil || msg == "" {
		return err
	} else {
		return fmt.Errorf("%s: [%w]", msg, err)
	}
}
