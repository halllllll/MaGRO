package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/halllllll/MaGRO/entity"
)

type MaGROUnitList struct {
	Service MaGROUnitService
}

func (lu *MaGROUnitList) ListUnit(ctx *gin.Context) {
	units, err := lu.Service.ListUnit(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusNotImplemented, gin.H{
		"units": units,
	})
	return
}

// TODO: GETなのでrequest bodyはない....
// - あとで実装するときにちゃんとやる
func (lu *MaGROUnitList) ListUsersSubunit(ctx *gin.Context) {
	// とりあえずDBで呼ぶだけ
	var i struct {
		UnitID entity.UnitId `json:"unit_id" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&i); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	result, err := lu.Service.ListUsersSubunit(ctx, &i.UnitID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"body": result,
	})
	return
}
