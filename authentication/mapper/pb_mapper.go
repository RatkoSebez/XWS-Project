package mapper

import "XWS-Project/authentication/dto"
import pb "XWS-Project/proto/login_service"

func MapRequest(dto *dto.LoginDTO) *pb.LoginDTO {
	lDTO := &pb.LoginDTO{
		Email:    dto.Email,
		Password: dto.Password,
	}
	return lDTO
}

func MapResponse(dto *dto.LoginResponseDTO) *pb.LoginResponseDTO {
	rDTO := &pb.LoginResponseDTO{
		Email:    dto.Email,
		Token:    dto.Token,
		Username: dto.Username,
	}
	return rDTO
}

func MapRequestObrnuto(req *pb.LoginDTO) *dto.LoginDTO {
	lDTO := &dto.LoginDTO{
		Email:    req.Email,
		Password: req.Password,
	}
	return lDTO
}
