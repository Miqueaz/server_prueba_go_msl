package main

import (
	"fmt"
	"log"
	"net"
	controller "server/controller"
	"server/core/connection/db"
	_ "server/modules/models"
	"server/routes/proto"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Cargar configuración del entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	// Inicializar la conexión a MongoDB
	db.InitMongoDB()
	defer db.CloseConnection() // Cerrar conexión al finalizar

	// Configurar el servidor gRPC
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al iniciar el listener: %v", err)
	}

	// Crear el servidor gRPC
	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(16*1024*1024), // Tamaño máximo del mensaje recibido (16 MB)
		grpc.MaxSendMsgSize(16*1024*1024), // Tamaño máximo del mensaje enviado (16 MB)
	)

	// Registrar el servicio del controlador
	proto.RegisterBaseServer(grpcServer, &controller.Server{})

	// Registrar la reflexión (opcional)
	reflection.Register(grpcServer)

	// Iniciar el servidor
	fmt.Println("Servidor escuchando en el puerto 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
