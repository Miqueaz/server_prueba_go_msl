package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitPostgres() {
	var err error
	connStr := "host=172.17.4.128 port=5432 user=secureuser password=DarthMonkus117 dbname=droply sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error al conectar a PostgreSQL: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("No se pudo hacer ping a PostgreSQL: %v", err)
	}

	fmt.Println("✅ Conectado a PostgreSQL correctamente.")
}

func init() {
	InitPostgres()
	if DB == nil {
		log.Fatal("Error: La conexión a la base de datos PostgreSQL no se ha inicializado.")
	}
	fmt.Println("PostgreSQL está listo para usar.")
}
