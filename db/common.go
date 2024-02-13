package db

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"

	"github.com/19sebastian95/Gambit/gambituser/models"
	"github.com/19sebastian95/Gambit/gambituser/secretmanager"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretmanager.GetSecretManager(os.Getenv("SecretName"))
	return err
}

func DbConnectRDSAWS() error {
	u := &url.URL{
		Scheme: "",
		User:   url.UserPassword(SecretModel.Username, SecretModel.Password),
		Host:   fmt.Sprintf("%s:%d", SecretModel.Host, SecretModel.Port),
	}

	Db, err = sql.Open("sqlserver", u.String())

	if err != nil {
		fmt.Println(err.Error())
	}

	err = Db.Ping()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Conexión exitosa")
	return nil
}
