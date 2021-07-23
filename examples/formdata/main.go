package main

import (
	"context"
	"encoding/json"
	"io"
	"path/filepath"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/grokify/go-awslambda"
)

type customStruct struct {
	Content       string
	FileName      string
	FileExtension string
}

func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	res := events.APIGatewayProxyResponse{}
	r, err := awslambda.NewReaderFormdata(req)
	if err != nil {
		return res, err
	}
	part, err := r.NextPart()
	if err != nil {
		return res, err
	}
	content, err := io.ReadAll(part)
	if err != nil {
		return res, err
	}
	custom := customStruct{
		Content:       string(content),
		FileName:      part.FileName(),
		FileExtension: filepath.Ext(part.FileName())}

	customBytes, err := json.Marshal(custom)
	if err != nil {
		return res, err
	}

	res = events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json"},
		Body: string(customBytes)}
	return res, nil
}

func main() {
	lambda.Start(handleRequest)
}
