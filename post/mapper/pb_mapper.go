package mapper

import (
	"XWS-Project/post/dto"
	"XWS-Project/post/model"
	"XWS-Project/proto/post_service"
	pb "XWS-Project/proto/post_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

func MapCommentDTO(dtoPb *post_service.NewComment) *dto.CommentDTO {
	id, _ := primitive.ObjectIDFromHex(dtoPb.PostID)
	rtrn := &dto.CommentDTO{
		CommentContent: dtoPb.CommentContent,
		PostID:         id,
	}
	return rtrn
}

func MapReactionDTO(dtoPb *post_service.NewReaction) *dto.ReactionDTO {
	id, _ := primitive.ObjectIDFromHex(dtoPb.PostID)
	rtrn := &dto.ReactionDTO{
		PostID:    id,
		React:     GetReactionType(dtoPb.React),
		UserEmail: dtoPb.UserEmail,
	}
	return rtrn
}

func GetReactionType(reaction string) model.ReactionType {
	if reaction == "LIKE" {
		return model.Like
	}
	if reaction == "DISLIKE" {
		return model.Dislike
	}
	return model.Like
}

func MapManyPosts(posts []*model.Post) *post_service.PostList {
	var retr *post_service.PostList
	for _, el := range posts {
		retr.Posts = append(retr.Posts, MapResponse(el))
	}
	return retr
}

func MapPhotoDTO(pbDto *post_service.PhotoMessage) *dto.PhotoDTO {
	rtrn := &dto.PhotoDTO{
		FileData: pbDto.FileData,
		FileName: pbDto.FileName,
	}
	return rtrn
}
