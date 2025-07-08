package postgres

import (
	"database/sql"
	"fmt"
	"log"
	ORM "main/pkg/sql"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error
	port, err := strconv.Atoi(os.Getenv("PORT_DB_POSTGRES"))
	connection := ORM.Connection{
		Host:     os.Getenv("HOST_DB_POSTGRES"),
		Port:     port,
		User:     os.Getenv("USER_DB_POSTGRES"),
		Password: os.Getenv("PASSWORD_DB_POSTGRES"),
		Database: os.Getenv("DATABASE_POSTGRES"),
		SSLMode:  os.Getenv("SSLMODE_DB_POSTGRES"),
	}
	DB, err = ORM.InitPostgres(connection)
	if DB == nil {
		log.Fatal("Error: La conexión a la base de datos PostgreSQL no se ha inicializado." + err.Error())
	}
	fmt.Println("PostgreSQL está listo para usar.")
}
