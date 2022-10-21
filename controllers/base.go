package controllers

import (
	"finalproject/views"

	"github.com/gin-gonic/gin"
)

func WriteJsonResponse(ctx *gin.Context, payload *views.Response) {
	ctx.JSON(payload.Status, payload)
}

func WriteJsonResponse_Failed(ctx *gin.Context, data *views.Failed) {
	ctx.JSON(data.Status, data)
}

func WriteJsonResponse_Succes(ctx *gin.Context, data *views.Resp_Register_Success) {
	ctx.JSON(data.Status, data)
}

func WriteJsonResponse_Login(ctx *gin.Context, data *views.Resp_Login) {
	ctx.JSON(data.Status, data)
}

func WriteJsonResponse_Put(ctx *gin.Context, data *views.Resp_Put) {
	ctx.JSON(data.Status, data)
}

func WriteJsonResponse_Delete(ctx *gin.Context, data *views.Resp_Delete) {
	ctx.JSON(data.Status, data)
}

// //// Foto JSON RESPONSE //////
func WriteJsonResponse_PostPhoto(ctx *gin.Context, data *views.Resp_Post_Photo) {
	ctx.JSON(data.Status, data)
}

func WriteJsonResponse_GetPhoto(ctx *gin.Context, data *views.Get_Photos) {
	ctx.JSON(data.Status, data)
}

func WriteJsonResponse_PutPhoto(ctx *gin.Context, data *views.Put_Photos) {
	ctx.JSON(data.Status, data)
}

// ///// Comend JSON RESPONSE /////////
func WriteJsonResponse_PostComments(ctx *gin.Context, data *views.Comments_Post) {
	ctx.JSON(data.Status, data)
}

func WriteJsonResponse_GetComment(ctx *gin.Context, data *views.Get_Comment) {
	ctx.JSON(data.Status, data)
}

func WriteJsonResponse_PutComment(ctx *gin.Context, data *views.Put_Comment) {
	ctx.JSON(data.Status, data)
}

// ///// Social Media JSON RESPONSE /////////
func WriteJsonResponse_PostSocialMedia(ctx *gin.Context, data *views.Post_Social_Media) {
	ctx.JSON(data.Status, data)
}

func WriteJsonResponse_GetSocialMedia(ctx *gin.Context, data *views.Post_Social_Media_Get) {
	ctx.JSON(data.Status, data)
}

func WriteJsonResponse_PutSocialMedia(ctx *gin.Context, data *views.Put_Social_Media_Get) {
	ctx.JSON(data.Status, data)
}
