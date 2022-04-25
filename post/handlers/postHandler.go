package handlers

import (
	"XWS-Project/post/dto"
	"XWS-Project/post/model"
	"XWS-Project/post/service"
	"XWS-Project/utilities"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"strings"
)

type PostHandler struct {
	PostService *service.PostService
}

func (handler *PostHandler) MakePost(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

	var request dto.NewPostDTO
	span := utilities.Tracer.StartSpanFromRequest("Post-handler", r)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	var post *model.Post
	ctx := utilities.Tracer.ContextWithSpan(context.Background(), span)
	userID := utilities.GetLoggedUserIDFromToken(r)
	post, err = handler.PostService.MakePost(ctx, request, userID)
	if err != nil {
		log.Fatal(err)
		return
	}
	respJson, err := json.Marshal(post)
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

func (handler *PostHandler) Preflight(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
	rw.WriteHeader(http.StatusOK)
}

func (handler *PostHandler) SavePhoto(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

	var request dto.PhotoDTO
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	var fileData = request.FileData
	idx := strings.Index(fileData, ";base64,")
	ImageType := fileData[11:idx]
	log.Println(ImageType)
	unbased, err := base64.RawStdEncoding.DecodeString(fileData[idx+8:])
	if err != nil {
		log.Fatal(err)
	}
	byteReader := bytes.NewReader(unbased)

	switch ImageType {
	case "png":
		im, err := png.Decode(byteReader)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.OpenFile("../photos/"+request.FileName, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
		}

		png.Encode(f, im)
	case "jpeg":
		im, err := jpeg.Decode(byteReader)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.OpenFile("../photos/"+request.FileName, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
		}

		jpeg.Encode(f, im, nil)
	case "gif":
		im, err := gif.Decode(byteReader)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.OpenFile("../photos/"+request.FileName, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
		}

		gif.Encode(f, im, nil)
	}

}
