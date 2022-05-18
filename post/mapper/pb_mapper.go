package mapper

import (
	"XWS-Project/post/dto"
	"XWS-Project/post/model"
	pb "XWS-Project/proto/post_service"
	//"fmt"
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

func MapCommentDTO(dtoPb *pb.NewComment) *dto.CommentDTO {
	id, _ := primitive.ObjectIDFromHex(dtoPb.PostID)
	rtrn := &dto.CommentDTO{
		CommentContent: dtoPb.CommentContent,
		PostID:         id,
	}
	return rtrn
}

func MapReactionDTO(dtoPb *pb.NewReaction) *dto.ReactionDTO {
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

func MapManyPosts(posts []*model.Post) *pb.PostList {
	retr := &pb.PostList{
		Posts: nil,
	}

	for _, el := range posts {

		postPb := &pb.Post{
			CreationTime: el.CreationTime.String(),
			PostID:       el.PostID.String(),
			PostText:     el.PostText,
			UserEmail:    el.UserEmail,
		}

		for _, elm := range el.Comments {
			postPb.Comments = append(postPb.Comments, &pb.Comment{
				CommentData:  elm.CommentData,
				CommentID:    elm.CommentID.String(),
				CreationTime: elm.CreationTime.String(),
				PostID:       elm.PostID.String(),
				UserEmail:    elm.UserEmail,
			})
		}
		for _, elm := range el.Reactions {
			postPb.Reactions = append(postPb.Reactions, &pb.Reaction{
				CreationTime:   elm.CreationTime.String(),
				PostID:         elm.PostID.String(),
				ReactionID:     elm.ReactionID.String(),
				TypeOfReaction: elm.TypeOfReaction.String(),
				UserEmail:      elm.UserEmail,
			})
		}
		for _, elm := range el.MediaAssets {
			postPb.MediaAssets = append(postPb.MediaAssets, &pb.Media{
				Filepath: elm.Filepath,
				MediaID:  elm.MediaID.String(),
				Site:     elm.Site,
			})
		}
		retr.Posts = append(retr.Posts, postPb)
	}
	return retr
}

func MapPhotoDTO(pbDto *pb.PhotoMessage) *dto.PhotoDTO {
	rtrn := &dto.PhotoDTO{
		FileData: pbDto.FileData,
		FileName: pbDto.FileName,
	}
	return rtrn
}
