package startup

import (
	"XWS-Project/api_gateway/infrastructure/api"
	cfg "XWS-Project/api_gateway/startup/config"
	"XWS-Project/proto/follow_service"
	"XWS-Project/proto/job_offer_service"
	"XWS-Project/proto/login_service"
	"XWS-Project/proto/post_service"
	"XWS-Project/proto/profile_service"
	"XWS-Project/proto/registration_service"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	server.initCustomHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := login_service.RegisterAuthenticationHandlerFromEndpoint(context.TODO(), server.mux, server.config.AuthenticationHost+":"+server.config.AuthenticationPort, opts)
	if err != nil {
		panic(err)
	}
	err = post_service.RegisterPostServiceHandlerFromEndpoint(context.TODO(), server.mux, server.config.PostPort, opts)
	if err != nil {
		panic(err)
	}
	err = registration_service.RegisterRegistrationServiceHandlerFromEndpoint(context.TODO(), server.mux, server.config.RegistrationPort, opts)
	if err != nil {
		panic(err)
	}
	err = follow_service.RegisterFollowServiceHandlerFromEndpoint(context.TODO(), server.mux, server.config.FollowPort, opts)
	if err != nil {
		panic(err)
	}
	err = profile_service.RegisterProfileServiceHandlerFromEndpoint(context.TODO(), server.mux, server.config.ProfilePort, opts)
	if err != nil {
		panic(err)
	}
	err = job_offer_service.RegisterJobOfferServiceHandlerFromEndpoint(context.TODO(), server.mux, server.config.JobOfferPort, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) initCustomHandlers() {
	postEndp := server.config.PostHost + ":" + server.config.PostPort
	authEndp := server.config.AuthenticationHost + ":" + server.config.AuthenticationPort
	regisEndp := server.config.RegistrationHost + ":" + server.config.RegistrationPort
	followEndp := server.config.FollowHost + ":" + server.config.FollowPort
	profileEndp := server.config.ProfileHost + ":" + server.config.ProfilePort
	jobOfferEndp := server.config.JobOfferHost + ":" + server.config.JobOfferPort

	authHandler := api.NewCustomHandler(postEndp, authEndp, regisEndp, followEndp, profileEndp, jobOfferEndp)
	authHandler.Init(server.mux)
}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
}
