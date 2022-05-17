// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: post_service.proto

package post_service

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MakePost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Links    []string `protobuf:"bytes,1,rep,name=links,proto3" json:"links,omitempty"`
	Photos   []string `protobuf:"bytes,2,rep,name=photos,proto3" json:"photos,omitempty"`
	PostText string   `protobuf:"bytes,3,opt,name=postText,proto3" json:"postText,omitempty"`
}

func (x *MakePost) Reset() {
	*x = MakePost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakePost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakePost) ProtoMessage() {}

func (x *MakePost) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakePost.ProtoReflect.Descriptor instead.
func (*MakePost) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{0}
}

func (x *MakePost) GetLinks() []string {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *MakePost) GetPhotos() []string {
	if x != nil {
		return x.Photos
	}
	return nil
}

func (x *MakePost) GetPostText() string {
	if x != nil {
		return x.PostText
	}
	return ""
}

type MakePostPlusEmail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Links     []string `protobuf:"bytes,1,rep,name=links,proto3" json:"links,omitempty"`
	Photos    []string `protobuf:"bytes,2,rep,name=photos,proto3" json:"photos,omitempty"`
	PostText  string   `protobuf:"bytes,3,opt,name=postText,proto3" json:"postText,omitempty"`
	UserEmail string   `protobuf:"bytes,4,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
}

func (x *MakePostPlusEmail) Reset() {
	*x = MakePostPlusEmail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakePostPlusEmail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakePostPlusEmail) ProtoMessage() {}

func (x *MakePostPlusEmail) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakePostPlusEmail.ProtoReflect.Descriptor instead.
func (*MakePostPlusEmail) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{1}
}

func (x *MakePostPlusEmail) GetLinks() []string {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *MakePostPlusEmail) GetPhotos() []string {
	if x != nil {
		return x.Photos
	}
	return nil
}

func (x *MakePostPlusEmail) GetPostText() string {
	if x != nil {
		return x.PostText
	}
	return ""
}

func (x *MakePostPlusEmail) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostID       string      `protobuf:"bytes,1,opt,name=postID,proto3" json:"postID,omitempty"`
	UserEmail    string      `protobuf:"bytes,2,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	CreationTime string      `protobuf:"bytes,3,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
	MediaAssets  []*Media    `protobuf:"bytes,4,rep,name=mediaAssets,proto3" json:"mediaAssets,omitempty"`
	PostText     string      `protobuf:"bytes,5,opt,name=postText,proto3" json:"postText,omitempty"`
	Reactions    []*Reaction `protobuf:"bytes,6,rep,name=reactions,proto3" json:"reactions,omitempty"`
	Comments     []*Comment  `protobuf:"bytes,7,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{2}
}

func (x *Post) GetPostID() string {
	if x != nil {
		return x.PostID
	}
	return ""
}

func (x *Post) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *Post) GetCreationTime() string {
	if x != nil {
		return x.CreationTime
	}
	return ""
}

func (x *Post) GetMediaAssets() []*Media {
	if x != nil {
		return x.MediaAssets
	}
	return nil
}

func (x *Post) GetPostText() string {
	if x != nil {
		return x.PostText
	}
	return ""
}

func (x *Post) GetReactions() []*Reaction {
	if x != nil {
		return x.Reactions
	}
	return nil
}

func (x *Post) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type Media struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaID  string `protobuf:"bytes,1,opt,name=mediaID,proto3" json:"mediaID,omitempty"`
	Site     string `protobuf:"bytes,2,opt,name=site,proto3" json:"site,omitempty"`
	Filepath string `protobuf:"bytes,3,opt,name=filepath,proto3" json:"filepath,omitempty"`
}

func (x *Media) Reset() {
	*x = Media{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Media) ProtoMessage() {}

func (x *Media) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Media.ProtoReflect.Descriptor instead.
func (*Media) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{3}
}

func (x *Media) GetMediaID() string {
	if x != nil {
		return x.MediaID
	}
	return ""
}

func (x *Media) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *Media) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

type Reaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReactionID     string `protobuf:"bytes,1,opt,name=reactionID,proto3" json:"reactionID,omitempty"`
	PostID         string `protobuf:"bytes,2,opt,name=postID,proto3" json:"postID,omitempty"`
	UserEmail      string `protobuf:"bytes,3,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	CreationTime   string `protobuf:"bytes,4,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
	TypeOfReaction string `protobuf:"bytes,5,opt,name=typeOfReaction,proto3" json:"typeOfReaction,omitempty"`
}

func (x *Reaction) Reset() {
	*x = Reaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reaction) ProtoMessage() {}

func (x *Reaction) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reaction.ProtoReflect.Descriptor instead.
func (*Reaction) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{4}
}

func (x *Reaction) GetReactionID() string {
	if x != nil {
		return x.ReactionID
	}
	return ""
}

func (x *Reaction) GetPostID() string {
	if x != nil {
		return x.PostID
	}
	return ""
}

func (x *Reaction) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *Reaction) GetCreationTime() string {
	if x != nil {
		return x.CreationTime
	}
	return ""
}

func (x *Reaction) GetTypeOfReaction() string {
	if x != nil {
		return x.TypeOfReaction
	}
	return ""
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentID    string `protobuf:"bytes,1,opt,name=commentID,proto3" json:"commentID,omitempty"`
	PostID       string `protobuf:"bytes,2,opt,name=postID,proto3" json:"postID,omitempty"`
	UserEmail    string `protobuf:"bytes,3,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	CreationTime string `protobuf:"bytes,4,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
	CommentData  string `protobuf:"bytes,5,opt,name=commentData,proto3" json:"commentData,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{5}
}

func (x *Comment) GetCommentID() string {
	if x != nil {
		return x.CommentID
	}
	return ""
}

func (x *Comment) GetPostID() string {
	if x != nil {
		return x.PostID
	}
	return ""
}

func (x *Comment) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *Comment) GetCreationTime() string {
	if x != nil {
		return x.CreationTime
	}
	return ""
}

func (x *Comment) GetCommentData() string {
	if x != nil {
		return x.CommentData
	}
	return ""
}

type PhotoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	FileData string `protobuf:"bytes,2,opt,name=fileData,proto3" json:"fileData,omitempty"`
}

func (x *PhotoMessage) Reset() {
	*x = PhotoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhotoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhotoMessage) ProtoMessage() {}

func (x *PhotoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhotoMessage.ProtoReflect.Descriptor instead.
func (*PhotoMessage) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{6}
}

func (x *PhotoMessage) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *PhotoMessage) GetFileData() string {
	if x != nil {
		return x.FileData
	}
	return ""
}

type NewComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentContent string `protobuf:"bytes,1,opt,name=commentContent,proto3" json:"commentContent,omitempty"`
	PostID         string `protobuf:"bytes,2,opt,name=postID,proto3" json:"postID,omitempty"`
}

func (x *NewComment) Reset() {
	*x = NewComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewComment) ProtoMessage() {}

func (x *NewComment) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewComment.ProtoReflect.Descriptor instead.
func (*NewComment) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{7}
}

func (x *NewComment) GetCommentContent() string {
	if x != nil {
		return x.CommentContent
	}
	return ""
}

func (x *NewComment) GetPostID() string {
	if x != nil {
		return x.PostID
	}
	return ""
}

type NewReaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	React     string `protobuf:"bytes,1,opt,name=react,proto3" json:"react,omitempty"`
	UserEmail string `protobuf:"bytes,2,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	PostID    string `protobuf:"bytes,3,opt,name=postID,proto3" json:"postID,omitempty"`
}

func (x *NewReaction) Reset() {
	*x = NewReaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewReaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewReaction) ProtoMessage() {}

func (x *NewReaction) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewReaction.ProtoReflect.Descriptor instead.
func (*NewReaction) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{8}
}

func (x *NewReaction) GetReact() string {
	if x != nil {
		return x.React
	}
	return ""
}

func (x *NewReaction) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *NewReaction) GetPostID() string {
	if x != nil {
		return x.PostID
	}
	return ""
}

type GetListOfFollowing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListOfFollowing []string `protobuf:"bytes,1,rep,name=listOfFollowing,proto3" json:"listOfFollowing,omitempty"`
}

func (x *GetListOfFollowing) Reset() {
	*x = GetListOfFollowing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListOfFollowing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListOfFollowing) ProtoMessage() {}

func (x *GetListOfFollowing) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListOfFollowing.ProtoReflect.Descriptor instead.
func (*GetListOfFollowing) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetListOfFollowing) GetListOfFollowing() []string {
	if x != nil {
		return x.ListOfFollowing
	}
	return nil
}

type PostList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *PostList) Reset() {
	*x = PostList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostList) ProtoMessage() {}

func (x *PostList) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostList.ProtoReflect.Descriptor instead.
func (*PostList) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{10}
}

func (x *PostList) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type EmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessage) Reset() {
	*x = EmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessage) ProtoMessage() {}

func (x *EmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessage.ProtoReflect.Descriptor instead.
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{11}
}

var File_post_service_proto protoreflect.FileDescriptor

var file_post_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x08, 0x4d, 0x61, 0x6b, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x54, 0x65, 0x78, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x54, 0x65, 0x78, 0x74, 0x22, 0x7b,
	0x0a, 0x11, 0x4d, 0x61, 0x6b, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x6c, 0x75, 0x73, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x68, 0x6f,
	0x74, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x54, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x84, 0x02, 0x0a, 0x04,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2d,
	0x0a, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x52, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6f, 0x73, 0x74, 0x54, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6f, 0x73, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x2c, 0x0a, 0x09, 0x72, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x72, 0x65,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x29, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x51, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x70, 0x61, 0x74, 0x68, 0x22, 0xac, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x74, 0x79, 0x70, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x79, 0x70, 0x65, 0x4f, 0x66, 0x52, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa3, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x46, 0x0a, 0x0c, 0x50, 0x68,
	0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x4c, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44,
	0x22, 0x59, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x72, 0x65, 0x61, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x22, 0x3e, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e,
	0x67, 0x12, 0x28, 0x0a, 0x0f, 0x6c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x69, 0x73, 0x74,
	0x4f, 0x66, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x0a, 0x08, 0x50,
	0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd6, 0x03, 0x0a, 0x0b, 0x50, 0x6f,
	0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x4d,
	0x61, 0x6b, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x6c, 0x75, 0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x1a, 0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x19, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x6d, 0x61, 0x6b, 0x65, 0x2d, 0x6e, 0x65, 0x77, 0x2d,
	0x70, 0x6f, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x4b, 0x0a, 0x09, 0x53, 0x61, 0x76, 0x65, 0x50,
	0x68, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x68, 0x6f, 0x74,
	0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x16, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x2d, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x3a, 0x01, 0x2a, 0x12, 0x3c, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x10, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x3a,
	0x01, 0x2a, 0x12, 0x48, 0x0a, 0x0c, 0x4d, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x11, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x6d, 0x61, 0x6b, 0x65,
	0x2d, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x57, 0x0a, 0x0e,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x18,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x1a, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x22, 0x10, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x6f, 0x73,
	0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x4b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x11, 0x12, 0x0f, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x70, 0x6f, 0x73,
	0x74, 0x73, 0x42, 0x20, 0x5a, 0x1e, 0x58, 0x57, 0x53, 0x2d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_post_service_proto_rawDescOnce sync.Once
	file_post_service_proto_rawDescData = file_post_service_proto_rawDesc
)

func file_post_service_proto_rawDescGZIP() []byte {
	file_post_service_proto_rawDescOnce.Do(func() {
		file_post_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_post_service_proto_rawDescData)
	})
	return file_post_service_proto_rawDescData
}

var file_post_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_post_service_proto_goTypes = []interface{}{
	(*MakePost)(nil),           // 0: post.MakePost
	(*MakePostPlusEmail)(nil),  // 1: post.MakePostPlusEmail
	(*Post)(nil),               // 2: post.Post
	(*Media)(nil),              // 3: post.Media
	(*Reaction)(nil),           // 4: post.Reaction
	(*Comment)(nil),            // 5: post.Comment
	(*PhotoMessage)(nil),       // 6: post.PhotoMessage
	(*NewComment)(nil),         // 7: post.NewComment
	(*NewReaction)(nil),        // 8: post.NewReaction
	(*GetListOfFollowing)(nil), // 9: post.GetListOfFollowing
	(*PostList)(nil),           // 10: post.PostList
	(*EmptyMessage)(nil),       // 11: post.EmptyMessage
}
var file_post_service_proto_depIdxs = []int32{
	3,  // 0: post.Post.mediaAssets:type_name -> post.Media
	4,  // 1: post.Post.reactions:type_name -> post.Reaction
	5,  // 2: post.Post.comments:type_name -> post.Comment
	2,  // 3: post.PostList.posts:type_name -> post.Post
	1,  // 4: post.PostService.CreatePost:input_type -> post.MakePostPlusEmail
	6,  // 5: post.PostService.SavePhoto:input_type -> post.PhotoMessage
	7,  // 6: post.PostService.Comment:input_type -> post.NewComment
	8,  // 7: post.PostService.MakeReaction:input_type -> post.NewReaction
	9,  // 8: post.PostService.FollowingPosts:input_type -> post.GetListOfFollowing
	11, // 9: post.PostService.GetUserPosts:input_type -> post.EmptyMessage
	2,  // 10: post.PostService.CreatePost:output_type -> post.Post
	11, // 11: post.PostService.SavePhoto:output_type -> post.EmptyMessage
	2,  // 12: post.PostService.Comment:output_type -> post.Post
	2,  // 13: post.PostService.MakeReaction:output_type -> post.Post
	10, // 14: post.PostService.FollowingPosts:output_type -> post.PostList
	10, // 15: post.PostService.GetUserPosts:output_type -> post.PostList
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_post_service_proto_init() }
func file_post_service_proto_init() {
	if File_post_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_post_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakePost); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakePostPlusEmail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Media); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhotoMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewComment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewReaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListOfFollowing); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_post_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_post_service_proto_goTypes,
		DependencyIndexes: file_post_service_proto_depIdxs,
		MessageInfos:      file_post_service_proto_msgTypes,
	}.Build()
	File_post_service_proto = out.File
	file_post_service_proto_rawDesc = nil
	file_post_service_proto_goTypes = nil
	file_post_service_proto_depIdxs = nil
}
