package main

import (
	"XWS-Project/profile/handlers"
	"XWS-Project/profile/repository"
	"XWS-Project/profile/service"
	pb "XWS-Project/proto/profile_service"
	"XWS-Project/utilities"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initDatabase() *mongo.Client {
	//mongoUri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI("mongodb://" + "localhost:27017" + "/?connect=direct")

	// Connect to MongoDB
	//ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//testData(client) //treba mi samo prvi put
	fmt.Println("Connected to MongoDB!")

	return client
}

func initRepo(clientDB *mongo.Client) *repository.ProfileRepository {
	return &repository.ProfileRepository{Client: clientDB}
}

func initService(repo *repository.ProfileRepository) *service.ProfileService {
	return &service.ProfileService{ProfileRepository: repo}
}

func initHandler(service *service.ProfileService) *handlers.ProfileHandler {
	return &handlers.ProfileHandler{ProfileService: service}
}

func handlerFunc(handler *handlers.ProfileHandler) {
	/*fmt.Println("Profile server started...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{email}", handler.Edit).Methods("PUT")
	router.HandleFunc("/{email}", handler.GetProfileByMail).Methods("GET")
	router.HandleFunc("/{email}", handler.Preflight).Methods("OPTIONS")

	log.Fatal(http.ListenAndServe(":8081", router))*/
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", "8084"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterProfileServiceServer(grpcServer, handler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
func main() {
	utilities.TracerInit("profile")
	dbClient := initDatabase()
	regRepo := initRepo(dbClient)
	regService := initService(regRepo)
	regHandler := initHandler(regService)
	handlerFunc(regHandler)
}
