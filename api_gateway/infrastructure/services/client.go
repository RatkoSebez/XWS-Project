package services

import (
	"XWS-Project/proto/login_service"
	"XWS-Project/proto/post_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewAuthenticationClient(address string) login_service.AuthenticationClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Authentication service: %v", err)
	}
	return login_service.NewAuthenticationClient(conn)
}

func NewPostClient(address string) post_service.PostServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Post service: %v", err)
	}
	return post_service.NewPostServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
