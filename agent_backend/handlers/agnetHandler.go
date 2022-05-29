package handlers

import (
	"XWS-Project/agent_backend/dto"
	"XWS-Project/agent_backend/model"
	"XWS-Project/agent_backend/service"
	"XWS-Project/utilities"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type AgentHandler struct {
	Service *service.AgentService
}

func (handler *AgentHandler) Login(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
	var request dto.LoginDTO
	span := utilities.Tracer.StartSpanFromRequest("Login-handler", r)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	var user *model.User
	ctx := utilities.Tracer.ContextWithSpan(context.Background(), span)
	user, err = handler.Service.Login(ctx, request)
	if err != nil {
		utilities.Tracer.LogError(span, err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	token, errr := utilities.CreateToken(user.Email, "authentication_service")
	if errr != nil {
		utilities.Tracer.LogError(span, errr)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	responseDTO := dto.LoginResponseDTO{
		Token:    token,
		Email:    user.Email,
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
	}

	respJson, err := json.Marshal(responseDTO)
	if err != nil {
		utilities.Tracer.LogError(span, err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusOK)
	_, _ = rw.Write(respJson)
	rw.Header().Set("Content-Type", "application/json")
	utilities.Tracer.FinishSpan(span)
}

func (handler *AgentHandler) Register(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

	span := utilities.Tracer.StartSpanFromRequest("Register-handler", r)
	ctx := utilities.Tracer.ContextWithSpan(context.Background(), span)

	var request dto.RegisterUserDTO
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var user = handler.Service.FindUserByEmail(ctx, request.Email)
	if user != nil {
		fmt.Println("Email already taken: " + request.Email)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.RegisterUser(ctx, request)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
	//_, _ = rw.Write(respJson)
	rw.Header().Set("Content-Type", "application/json")
	utilities.Tracer.FinishSpan(span)
}

func (handler *AgentHandler) Preflight(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
	rw.WriteHeader(http.StatusOK)
}

func (handler *AgentHandler) Test(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
	respJson, err := json.Marshal("vrati bilo sta")
	rw.Header().Add("content-type", "application/json")
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	_, _ = rw.Write(respJson)
}

func (handler *AgentHandler) CreateCompany(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

	span := utilities.Tracer.StartSpanFromRequest("CreateCompany company handler", r)
	ctx := utilities.Tracer.ContextWithSpan(context.Background(), span)

	var company model.Company
	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// company has unique name
	var user = handler.Service.FindCompanyByName(ctx, company.Name)
	if user != nil {
		fmt.Println("Name already taken: " + company.Name)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.CreateCompany(ctx, company)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	utilities.Tracer.FinishSpan(span)
}

func (handler *AgentHandler) ApproveCompany(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

	span := utilities.Tracer.StartSpanFromRequest("Approve company handler", r)
	ctx := utilities.Tracer.ContextWithSpan(context.Background(), span)

	var request dto.ApproveCompanyDTO
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	handler.Service.ApproveCompany(ctx, &request)

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	utilities.Tracer.FinishSpan(span)
}

func (handler *AgentHandler) GetAllCompanies(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

	span := utilities.Tracer.StartSpanFromRequest("Get all not approved companies handler", r)
	ctx := utilities.Tracer.ContextWithSpan(context.Background(), span)

	companies := handler.Service.GetAllCompanies(ctx)

	respJson, _ := json.Marshal(companies)
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	_, _ = rw.Write(respJson)
	utilities.Tracer.FinishSpan(span)
}

func (handler *AgentHandler) EditCompany(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

	span := utilities.Tracer.StartSpanFromRequest("Edit company handler", r)
	ctx := utilities.Tracer.ContextWithSpan(context.Background(), span)

	var company model.Company
	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	handler.Service.EditCompany(ctx, &company)

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	utilities.Tracer.FinishSpan(span)
}

//func (handler *AgentHandler) GetAllUsersCompanies(rw http.ResponseWriter, r *http.Request) {
//	rw.Header().Set("Access-Control-Allow-Origin", "*")
//	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
//	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
//
//	span := utilities.Tracer.StartSpanFromRequest("Get all not approved companies handler", r)
//	ctx := utilities.Tracer.ContextWithSpan(context.Background(), span)
//
//	userEmail, _ := mux.Vars(r)["email"]
//
//	companies := handler.Service.GetAllUsersCompanies(ctx, userEmail)
//
//	respJson, _ := json.Marshal(companies)
//	rw.WriteHeader(http.StatusOK)
//	rw.Header().Set("Content-Type", "application/json")
//	_, _ = rw.Write(respJson)
//	utilities.Tracer.FinishSpan(span)
//}
