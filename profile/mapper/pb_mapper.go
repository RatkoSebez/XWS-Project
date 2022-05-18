package mapper

import (
	"XWS-Project/profile/dto"
	"XWS-Project/profile/model"
	pb "XWS-Project/proto/profile_service"
)

func MapProfileDTO(pbDto *pb.ProfileDTO) *dto.ProfileDTO {
	rtrn := &dto.ProfileDTO{
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

func MapResponseProfile(info *model.ProfileInfo) *pb.ProfileInfo {
	rtrn := &pb.ProfileInfo{
		Name:       info.Name,
		Surname:    info.Surname,
		BDateDay:   int32(info.BDateDay),
		BDateMonth: int32(info.BDateMonth),
		BDateYear:  int32(info.BDateYear),
		Bio:        info.Bio,
		Username:   info.Username,
		Password:   info.Password,
		Numtel:     info.Numtel,
		Email:      info.Email,
		IsPrivate:  info.IsPrivate,
	}

	for _, el := range info.Skills {
		rtrn.Skills = append(rtrn.Skills, el)
	}
	for _, el := range info.Experience {
		rtrn.Experience = append(rtrn.Experience, el)
	}
	for _, el := range info.Education {
		rtrn.Education = append(rtrn.Education, el)
	}
	for _, el := range info.Interests {
		rtrn.Interests = append(rtrn.Interests, el)
	}
	return rtrn
}
