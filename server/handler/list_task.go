package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListTask struct {
	Service ListTasksService
}

func (lt *ListTask) ListTask(ctx *gin.Context) {
	tasks, err := lt.Service.ListTask(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}
