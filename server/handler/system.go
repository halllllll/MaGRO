package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/halllllll/MaGRO/entity"
)

/*
汎用的に使えるやつ
*/
type MagroSystem struct {
	Service SystemService
}

func (m *MagroSystem) GetSystemInfoHandler(ctx *gin.Context) {
	info, err := m.Service.GetSystemInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, info)
}

/*
for Admin
*/
type MaGROAdmin struct {
	MutateService AdminMutateService
}

func (m *MaGROAdmin) UpdateRoleNameHandler(ctx *gin.Context) {
	var req *entity.ReqNewRoleAlias
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := m.MutateService.UpdateRole(ctx, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "renew",
	})
	return
}
