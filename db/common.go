package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/19sebastian95/Gambit/gambituser/models"
	"github.com/19sebastian95/Gambit/gambituser/secretmanager"
	_ "github.com/denisenkom/go-mssqldb"
)

var SecretModel models.SecretRDSJson
var err error
var db *sql.DB

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

	db, err = sql.Open("sqlserver", u.String())

	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.Ping()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Conexión exitosa")
	return nil
}

func DbConnectSlqServer() error {
	var server = "SEBASTIAN"
	var port = 1433
	var user = "sa"
	var password = "123456$"
	var database = "Gambit"

	var connectionString = fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s;",
		server, port, user, password, database)

	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err.Error())
	}

	return err
}
