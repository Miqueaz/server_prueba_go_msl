package orm_sql

import (
	"database/sql"
	"fmt"
)

func InitPostgres(connection Connection) (*sql.DB, error) {
	var err error
	// connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
	// 	connection.Host,
	// 	connection.Port,
	// 	connection.User,
	// 	connection.Password,
	// 	connection.Database,
	// 	connection.SSLMode,
	// )

	connStr := "host=172.18.7.66 port=5432 user=postgres password=DarthMonkus117 dbname=droply sslmode=disable"

	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = DB.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("✅ Conectado a PostgreSQL correctamente.")

	return DB, nil
}
