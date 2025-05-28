package controller

import (
	"context"
	"fmt"
	"log"
	"server/helpers"
	"server/modules/handlers"
	"server/routes/proto"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

// Definir el servicio
type Server struct {
	proto.UnimplementedBaseServer
}

// Implementación de `Get`
func (s *Server) Get(ctx context.Context, req *proto.DynamicRequest) (*proto.BaseResponse, error) {
	filter := bson.M{}
	filters := req.GetFilters()
	collection := req.GetCollection()
	config := map[string]int{
		"page":     int(req.GetConfig().Page),
		"pageSize": int(req.GetConfig().PageSize),
	}

	// Construir el filtro a partir de los parámetros
	for key, value := range filters {
		filter[key] = value
		fmt.Printf("filter[%s] = %v\n", key, value)
	}

	// Obtener el controlador para la colección y los documentos
	_, ok := handlers.GetController(collection)
	if !ok {
		log.Println("ERROR: No se encontró el controlador para 'Comentarios'")
		return nil, status.Error(codes.NotFound, "Controlador no encontrado")
	}

	documents, err := handlers.CommentController.Read(filter, config)
	if err != nil || len(documents) == 0 {
		log.Printf("Error al buscar documentos en MongoDB: %v", err)
		return &proto.BaseResponse{
			Status:  404,
			Message: "No se encontraron documentos en la base de datos.",
		}, nil
	}

	// Convertir los documentos BSON a google.protobuf.Struct
	var dataList []*proto.Data
	for _, doc := range documents {
		// Convertir el BSON a un mapa de any
		convertedDoc := helpers.ConvertBsonToMap(doc)

		// Convertir el mapa a google.protobuf.Struct
		structDoc, err := structpb.NewStruct(convertedDoc)
		if err != nil {
			log.Printf("Error al convertir BSON a Struct: %v", err)
			return nil, err
		}

		// Crear un objeto proto.Data con el struct y agregarlo a la lista
		protoDoc := &proto.Data{Data: structDoc}
		dataList = append(dataList, protoDoc)
	}

	// Devolver la respuesta con los documentos encontrados
	return &proto.BaseResponse{
		Status:  200,
		Message: "Documentos encontrados en la base de datos.",
		Data:    dataList,
	}, nil
}
