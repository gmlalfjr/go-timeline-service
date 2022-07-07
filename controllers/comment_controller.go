package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gmlalfjr/timeline-service/services"
)

type ICommentController interface {
	GetTimeline(ctx *gin.Context)
}

type CommentController struct {
	service services.ICommentService
}

func NewCommentController(service services.ICommentService) ICommentController {
	return &CommentController{service: service}
}

func (t CommentController) GetTimeline(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "success",
	})
	return
}
