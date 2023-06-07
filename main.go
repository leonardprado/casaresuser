package casaresuser

import (
	"context"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"

)
func main() {
	lambda.Start(EjecutoLambda)

}
func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error){

}