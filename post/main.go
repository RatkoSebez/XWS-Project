package main

import (
	"XWS-Project/post/handlers"
	"XWS-Project/post/repository"
	"XWS-Project/post/service"
	"XWS-Project/utilities"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func initDB() *mongo.Client {
	//mongoUri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI("mongodb://" + "localhost:27017" + "/?connect=direct")

	// Connect to MongoDB
	//ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return client
}

func initRepo(postDB *mongo.Client) *repository.PostRepository {
	return &repository.PostRepository{Client: postDB}
}

func initService(repo *repository.PostRepository) *service.PostService {
	return &service.PostService{PostRepo: repo}
}

func initHandler(service *service.PostService) *handlers.PostHandler {
	return &handlers.PostHandler{PostService: service}
}

func handlerFunc(handler *handlers.PostHandler) {
	fmt.Println("Post server started...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/make-new-post", handler.Preflight).Methods("OPTIONS")
	router.HandleFunc("/make-new-post", handler.MakePost).Methods("POST")
	router.HandleFunc("/save-photo", handler.Preflight).Methods("OPTIONS")
	router.HandleFunc("/save-photo", handler.SavePhoto).Methods("POST")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	utilities.TracerInit("post")
	dbClient := initDB()
	postRepo := initRepo(dbClient)
	postService := initService(postRepo)
	postHandler := initHandler(postService)
	handlerFunc(postHandler)
}
