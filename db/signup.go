package db

import (
	"context"
	"fmt"
	"log"
)

// func SignUp(sig models.SignUp) error {
func SignUp() error {
	fmt.Println("Comienza Registro")

	err := DbConnectSlqServer()
	if err != nil {
		return err
	}

	//err = SelectUsert()

	return nil
}

func SelectUsert() error {
	ctx := context.Background()

	err := db.QueryRowContext(ctx, "SELECT * FROM Users").Scan()

	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
