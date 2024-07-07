package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/halllllll/MaGRO/entity"
)

type MaGRORepass struct{
	Service MaGRORepassService
}

// TODO: 受け取って中身を見るだけの仮実装
func (rp *MaGRORepass)Repass(ctx *gin.Context){
	unit_id := ctx.Param("unit")
	
	var req *entity.RepassRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  entity.ER,
			"message": err.Error(),
		})
		return
	}

	fmt.Println("中身をみてみる")
	fmt.Printf("%#v\n", req)


	// paramにあるunit idを取得
	int_unit_id, err := strconv.Atoi(unit_id)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  entity.ER,
			"message": err.Error(),
		})
		return
	}
	fmt.Printf("from unit: %d\n", int_unit_id)

	// TODO: unit idとidtokenからの取得とか認証はServiceで使うが、とりあえず中身を見るだけの仮実装
	// rp.Service.RepassUser(ctx)
}