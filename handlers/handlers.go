package handlers

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/jsonandrx/twitterGo/models"
)

func Majeadores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {
	fmt.Println("Voy a procesar "+ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))
	var r models.RespApi
	r.Status = 400

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
			
		}

	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
			
		}
	case "DELETe":
		switch ctx.Value(models.Key("path")).(string) {
			
		}
	}
	r.Message = "Method Invalid"
	return r
}