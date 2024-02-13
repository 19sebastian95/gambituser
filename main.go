package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/19sebastian95/Gambit/gambituser/awsgo"
	"github.com/19sebastian95/Gambit/gambituser/db"
	"github.com/19sebastian95/Gambit/gambituser/models"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutorLambda)
}

func EjecutorLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InicializoAWS()

	if !ValidoParametros() {
		fmt.Println("Error en los parámetros, debe enviar 'SecretName'")
		err := errors.New("Error en los parámetros, debe enviar 'SecretName'")

		return event, err
	}

	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
		case "sub":
			datos.UserUUID = att
		}
	}

	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error al Leer el Secret" + err.Error())
		return event, err
	}

	err = db.SignUp(datos)

	return event, err
}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
