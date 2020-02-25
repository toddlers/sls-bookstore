package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// logs any error to os.Stderr and returns 500 server error
// that AWS APU Gateway unterstands

func serverError(err error) (events.APIGatewayProxyResponse, error) {

	errorLogger.Println(err.Error())
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
		Body: fmt.Sprintf("%v", err.Error()),
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}
