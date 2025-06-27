package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

// InitMongoDB establece la conexión con la base de datos MongoDB
func InitMongoDB() {
	uri := os.Getenv("CLUSTER_CONNECTION")
	// Configuración de la conexión con MongoDB
	clientOptions := options.Client().ApplyURI(uri)
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error al conectar a MongoDB: %v", err)
	}

	// Verificar la conexión
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error al hacer ping a MongoDB: %v", err)
	}

	log.Println("Conexión a MongoDB establecida con éxito")
}

// InsertDocument inserta un documento en la colección
func InsertDocument(document interface{}) (*mongo.InsertOneResult, error) {
	return collection.InsertOne(context.Background(), document)
}

// FindDocuments encuentra múltiples documentos en la colección
func FindDocuments(filter interface{}, collectionName string, page int64, pageSize int64) (*mongo.Cursor, error) {
	// Obtener el nombre de la base de datos desde el archivo .env
	dbName := os.Getenv("DATABASE")

	// Obtener la colección de la base de datos
	collection := client.Database(dbName).Collection(collectionName)

	// Calcular el número de documentos a omitir (skip) y el límite de la página
	skip := (page - 1) * pageSize
	limit := pageSize

	// Configuración de las opciones de consulta
	opts := options.Find().
		SetSkip(skip).
		SetLimit(limit)

	// Contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Realizar la consulta con paginación
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	return cursor, nil
}

// CloseConnection cierra la conexión a MongoDB
func CloseConnection() {
	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatalf("Error al cerrar la conexión con MongoDB: %v", err)
	}
	log.Println("Conexión a MongoDB cerrada")
}
