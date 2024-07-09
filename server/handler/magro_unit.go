package handler

import (
	"fmt"
	"net/http"
	"strconv"

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
			"status":  entity.ER,
			"message": err.Error(),
		})
		return
	}

	// SPAなので諦める...?
	// // 0の場合はフロント側でやれ
	// if len(units) == 1 {
	// 	// Subunit取得して返す
	// 	ctx.JSON(http.StatusAccepted, gin.H{"message": "単一unitあとでsubunitを返す実装する"})
	// 	return
	// } else {
	// 	resp := &entity.RespBelongUnits{
	// 		Result:    entity.OK,
	// 		UnitCount: len(units),
	// 		Units:     units,
	// 	}
	// 	ctx.JSON(http.StatusAccepted, resp)
	// 	return
	// }
	// ctx.JSON(http.StatusNotImplemented, gin.H{
	// 	"units": units,
	// })
	// return

	resp := &entity.RespBelongUnits{
		Result:    entity.OK,
		UnitCount: len(units),
		Units:     units,
	}
	ctx.JSON(http.StatusAccepted, resp)
	return

}

func (lu *MaGROUnitList) ListUsersSubunit(ctx *gin.Context) {
	unit_id := ctx.Param("unit")

	int_unit_id, err := strconv.Atoi(unit_id)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  entity.ER,
			"message": err.Error(),
		})
		return
	}
	result, err := lu.Service.ListUsersSubunit(ctx, (*entity.UnitId)(&int_unit_id))
	if err != nil {
		// // TODO: ErrNoRowsのハンドリングはrepositoryでやるべきでは？
		// -> どうせ空かも
		// if err == sql.ErrNoRows {
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{
		// 		"status":  entity.ER,
		// 		"message": "empty",
		// 	})
		// 	return
		// } else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  entity.ER,
			"message": err.Error(),
		})
		return
	}
	fmt.Printf("from unit id: %d\ncurrent user: %#v\n", int_unit_id, result.CurrentUser)
	ctx.JSON(http.StatusOK, gin.H{
		"status": entity.OK,
		"data":   result,
	})
	return
}
