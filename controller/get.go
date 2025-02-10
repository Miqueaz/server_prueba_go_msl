// controlador/server.go
package controller

import (
	"context"
	"fmt"
	"log"
	"server/core/base"
	"server/helpers"
	"server/routes/proto"

	"go.mongodb.org/mongo-driver/bson"
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
	for key, value := range filters {
		filter[key] = value
		fmt.Printf("filter[%s] = %v\n", key, value)
	}

	controller := base.GetController(collection)
	data, err := controller.Read(filter)
	cursor := data
	if err != nil {
		log.Printf("Error al buscar documentos en MongoDB: %v", err)
		return &proto.BaseResponse{
			Status:  404,
			Message: "No se encontraron documentos en la base de datos.",
		}, nil
	}

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

	var dataList []*proto.Data
	for _, doc := range documents {
		convertedDoc := helpers.ConvertBsonToMap(doc) // Convertir BSON a map compatible
		convertedToStringDoc, err := helpers.ConvertMapToStrings(convertedDoc)
		if err != nil {
			log.Printf("Error al convertir el documento a string: %v", err)
			return nil, err
		}
		protoDoc := &proto.Data{Data: convertedToStringDoc}
		dataList = append(dataList, protoDoc)
	}

	return &proto.BaseResponse{
		Status:  200,
		Message: "Documentos encontrados en la base de datos.",
		Data:    dataList,
	}, nil
}
