package main

import (
	"github.com/andersnormal/franz/internal/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	h := handler.New()

	lambda.Start(h.Handler)
}
