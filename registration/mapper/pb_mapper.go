package mapper

import pb "XWS-Project/proto/registration_service"
import "XWS-Project/registration/dto"

func MapPBToDTO(pbDto *pb.RegisterMessage) *dto.UserDTO {
	rtrn := &dto.UserDTO{
		Name:       pbDto.Name,
		Surname:    pbDto.Surname,
		BDateDay:   int(pbDto.BDateDay),
		BDateMonth: int(pbDto.BDateMonth),
		BDateYear:  int(pbDto.BDateYear),
		Bio:        pbDto.Bio,
		Username:   pbDto.Username,
		Password:   pbDto.Password,
		Numtel:     pbDto.Numtel,
		Email:      pbDto.Email,
		IsPrivate:  pbDto.IsPrivate,
	}

	for _, el := range pbDto.Skills {
		rtrn.Skills = append(rtrn.Skills, el)
	}
	for _, el := range pbDto.Experience {
		rtrn.Experience = append(rtrn.Experience, el)
	}
	for _, el := range pbDto.Education {
		rtrn.Education = append(rtrn.Education, el)
	}
	for _, el := range pbDto.Interests {
		rtrn.Interests = append(rtrn.Interests, el)
	}
	return rtrn
}
