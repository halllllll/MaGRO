package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/halllllll/MaGRO/entity"
)

type MaGRORepass struct {
	Service MaGRORepassService
}

func (rp *MaGRORepass) Repass(ctx *gin.Context) {

	unit_id := ctx.Param("unit")

	var req *entity.RepassRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  entity.ER,
			"message": err.Error(),
		})
		return
	}

	// paramにあるunit idを取得
	int_unit_id, err := strconv.Atoi(unit_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  entity.ER,
			"message": err.Error(),
		})
		return
	}


	var reqData []*entity.UserPrimaryUniqID
	for _, v := range req.TargetUsers {
		reqData = append(reqData, &entity.UserPrimaryUniqID{
			ID:      v.ID,
			Account: v.Account,
		})
	}

	result, err := rp.Service.RepassUser(ctx, (*entity.UnitId)(&int_unit_id), reqData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  entity.ER,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  entity.OK,
		"body": result.Result,
	})
}
