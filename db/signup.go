package db

import (
	"fmt"

	"github.com/19sebastian95/Gambit/gambituser/models"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza Registro")

	err := DbConnectRDSAWS()
	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "SELECT * FROM Users"

	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
