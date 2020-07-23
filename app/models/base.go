package models

import (
	"database/sql"
	"fmt"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/snjssk/whimsy-painter-backend/app/config"
	"log"
)

var DbConnection *sql.DB

func init() {
	var err error
	dns := fmt.Sprintf("host=%s dbname=%s user=%s password=% sslmode=disable", config.Config.Host, config.Config.Name, config.Config.User, config.Config.Password)
	DbConnection, err = sql.Open("cloudsqlpostgres", dns)
	// defer DbConnection.Close()

	//log.Println(DbConnection)
	log.Println(err)
}
