package mapper

import "XWS-Project/post/dto"
import "XWS-Project/post/model"
import pb "XWS-Project/proto/post_service"

func MapPBtoDTORequest(pbDTO *pb.MakePostPlusEmail) *dto.NewPostDTO {
	rtrn := &dto.NewPostDTO{
		Links:    pbDTO.Links,
		Photos:   pbDTO.Photos,
		PostText: pbDTO.PostText,
	}
	return rtrn
}

func MapResponse(respPost *model.Post) *pb.Post {
	rtrn := &pb.Post{
		CreationTime: respPost.CreationTime.String(),
		PostID:       respPost.PostID.String(),
		PostText:     respPost.PostText,
		UserEmail:    respPost.UserEmail,
	}

	for _, el := range respPost.Comments {
		rtrn.Comments = append(rtrn.Comments, &pb.Comment{
			CommentData:  el.CommentData,
			CommentID:    el.CommentID.String(),
			CreationTime: el.CreationTime.String(),
			PostID:       el.PostID.String(),
			UserEmail:    el.UserEmail,
		})
	}
	for _, el := range respPost.Reactions {
		rtrn.Reactions = append(rtrn.Reactions, &pb.Reaction{
			CreationTime:   el.CreationTime.String(),
			PostID:         el.PostID.String(),
			ReactionID:     el.ReactionID.String(),
			TypeOfReaction: el.TypeOfReaction.String(),
			UserEmail:      el.UserEmail,
		})
	}
	for _, el := range respPost.MediaAssets {
		rtrn.MediaAssets = append(rtrn.MediaAssets, &pb.Media{
			Filepath: el.Filepath,
			MediaID:  el.MediaID.String(),
			Site:     el.Site,
		})
	}

	return rtrn
}
