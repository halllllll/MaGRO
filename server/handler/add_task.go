package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddTask struct {
	// Validator *validator.Validate
	Service AddTaskService
}

func (at *AddTask) AddTask(ctx *gin.Context) {
	var b struct {
		Title string `json:"title" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}
	task, err := at.Service.AddTask(ctx, b.Title)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	rsp := struct {
		ID int `json:"id"`
	}{
		ID: int(task.ID),
	}
	ctx.JSON(http.StatusOK, rsp)
	return
}
