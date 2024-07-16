package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/halllllll/MaGRO/entity"
	"github.com/jackc/pgx/v5"
)

type roleKey struct{}

type regularRoler interface {
	GetRole(ctx context.Context, id *entity.UserID) (entity.Role, error)
}

// serviceを介さず直接DB使っちゃう（というかauthをserviceに移動してもいいか
type ensureRegularAccount struct {
	db regularRoler
}

func NewEnsureRegularAccountMiddleware(db regularRoler) *ensureRegularAccount {
	return &ensureRegularAccount{db: db}
}

// guest以外
func (er *ensureRegularAccount) EnsureRegularAccountMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userid, ok := GetUserID(ctx)
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  entity.ER,
				"message": "empty user id",
			})
			ctx.Abort()
			return
		}
		role, err := er.db.GetRole(ctx, &userid)
		// 直接dbから帰ってきてるerr
		if err != nil {
			fmt.Printf("DB Error: %s -  %s\n", userid, err.Error())
			if err == pgx.ErrNoRows {
				fmt.Println("error no rows")
				ctx.JSON(http.StatusOK, gin.H{
					"status":  entity.ER,
					"message": "no data",
				})
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  entity.ER,
				"message": fmt.Errorf("DB Error"),
			})
			ctx.Abort()
			return
		}

		// guest以外
		if role == entity.RoleGuest {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  entity.OK,
				"message": "not allowed",
			})
			ctx.Abort()
			return
		}

		// ginのContextとginのRequest.Contextがある + SetUserIDは素のnet/httpを扱うことにしているのでややこしい。基本的に素のContext単体の世界で考えるようにする
		_tmpCtx := SetRole(ctx.Request.Context(), role)
		// WARN: RequestとContextの置き換えで意図せぬ上書きや抜け落ちがあるかもしれない
		ctx.Request = ctx.Request.WithContext(_tmpCtx)

		ctx.Next()
	}
}

func SetRole(ctx context.Context, role entity.Role) context.Context {
	return context.WithValue(ctx, roleKey{}, role)
}

func GetRole(ctx context.Context) (entity.Role, bool) {
	role, ok := ctx.Value(roleKey{}).(entity.Role)
	return role, ok
}

func IsAdmin(ctx context.Context) bool {
	role, ok := GetRole(ctx)
	if !ok {
		return false
	}
	return role == entity.RoleAdmin
}

func IsGuest(ctx context.Context) bool {
	role, ok := GetRole(ctx)
	if !ok {
		return false
	}
	return role == entity.RoleGuest
}
