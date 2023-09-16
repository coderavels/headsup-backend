package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	"github.com/coderavels/headsup-backend/handlers"
)

var ginLambda *ginadapter.GinLambda
var router *gin.Engine

func init() {
	router = gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/decks", handlers.GetDecks)
	router.GET("/decks/:id/cards", handlers.GetDeckCards)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ginLambda = ginadapter.New(router)
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambdaRoot := os.Getenv("LAMBDA_TASK_ROOT")
	fmt.Printf("lambdaroot is: %s\n", lambdaRoot)
	if lambdaRoot != "" {
		// executing within aws lambda
		lambda.Start(Handler)
	} else {
		// run locally
		router.Run("localhost:8080")
	}
}
