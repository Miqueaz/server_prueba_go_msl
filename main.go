package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"server/core/connection/db"
	"server/core/connection/proto"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Definir el servicio
type server struct {
	proto.UnimplementedBaseServer
}

// Función para convertir BSON a `map[string]interface{}` y manejar `primitive.ObjectID`
func convertBsonToMap(doc bson.M) map[string]interface{} {
	result := make(map[string]interface{})
	for key, value := range doc {
		result[key] = fmt.Sprintf("%v", value)
	}
	return result
}

func convertMapToStrings(input map[string]interface{}) (map[string]string, error) {
	result := make(map[string]string)
	for key, value := range input {
		strValue, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("el valor para la clave %s no es un string", key)
		}
		result[key] = strValue
	}
	return result, nil
}

// Implementación de `Get`
func (s *server) Get(ctx context.Context, req *proto.DynamicRequest) (*proto.BaseResponse, error) {
	filter := bson.M{}
	for key, value := range req.GetFilters() {
		filter[key] = value
		fmt.Printf("filter[%s] = %v\n", key, value)
	}

	cursor, err := db.FindDocuments(filter, req.GetCollection())
	if err != nil {
		log.Printf("Error al buscar documentos en MongoDB: %v", err)
		return &proto.BaseResponse{
			Status:  404,
			Message: "No se encontraron documentos en la base de datos.",
		}, nil
	}
	defer cursor.Close(ctx)

	var documents []bson.M
	if err := cursor.All(ctx, &documents); err != nil {
		log.Printf("Error al decodificar documentos: %v", err)
		return nil, err
	}

	if len(documents) == 0 {
		return &proto.BaseResponse{
			Status:  404,
			Message: "No se encontraron documentos en la base de datos.",
		}, nil
	}

	var DataList []*proto.Data
	for _, doc := range documents {
		convertedDoc := convertBsonToMap(doc) // Convertir BSON a map compatible
		converteToStingDoc, err := convertMapToStrings(convertedDoc)
		if err != nil {
			log.Printf("Error al convertir el documento a string: %v", err)
			return nil, err
		}
		protoDoc := &proto.Data{Data: converteToStingDoc}

		DataList = append(DataList, protoDoc)
	}

	// Retornar respuesta con la lista de estructuras
	return &proto.BaseResponse{
		Status:  200,
		Message: "Documentos encontrados en la base de datos.",
		Data:    DataList, // Lista de *structpb.Struct
	}, nil
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	// Inicializar la conexión a MongoDB
	db.InitMongoDB()

	// Configurar el servidor gRPC
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Crear el servidor gRPC
	s := grpc.NewServer(
		grpc.MaxRecvMsgSize(16*1024*1024), // Tamaño máximo del mensaje recibido (16 MB)
		grpc.MaxSendMsgSize(16*1024*1024), // Tamaño máximo del mensaje enviado (16 MB)
	)
	proto.RegisterBaseServer(s, &server{})

	// Registrar la reflexión (opcional)
	reflection.Register(s)

	// Iniciar el servidor
	fmt.Println("Servidor escuchando en puerto 50051...")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	// Cerrar la conexión con MongoDB al finalizar
	defer db.CloseConnection()
}
