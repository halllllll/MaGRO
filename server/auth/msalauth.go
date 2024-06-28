package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/halllllll/MaGRO/entity"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jws"
	"github.com/lestrrat-go/jwx/v2/jwt"

	"github.com/gin-gonic/gin"
)

type userIDKey struct{}
type roleKey struct{}

func GetUserID(ctx context.Context) (entity.UserID, bool) {
	id, ok := ctx.Value(userIDKey{}).(entity.UserID)
	return id, ok
}

func SetUserID(ctx context.Context, uid entity.UserID) context.Context {
	return context.WithValue(ctx, userIDKey{}, uid)
}

func SetRole(ctx context.Context, role entity.Role) context.Context {
	return context.WithValue(ctx, roleKey{}, role)
}

func GetRole(ctx context.Context) (string, bool) {
	role, ok := ctx.Value(roleKey{}).(string)
	return role, ok
}

func IsAdmin(ctx context.Context) bool {
	role, ok := GetRole(ctx)
	if !ok {
		return false
	}
	return role == string(entity.RoleAdmin)
}

func MsalAuthMiddleware(clientId string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		reqTokens := ctx.Request.Header["Authorization"]
		if len(reqTokens) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid token",
			})
			ctx.Abort()
			return
		}
		reqToken := reqTokens[0]
		if reqToken == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			ctx.Abort()
			return
		}
		idToken := strings.Replace(reqToken, "Bearer ", "", 1)

		fmt.Printf("IdToken %s\n", idToken)

		// TODO: 2回くらいまではチャレンジしてもいいかも
		// check ms oidc
		req, err := http.NewRequestWithContext(ctx, "GET", "https://login.microsoftonline.com/common/v2.0/.well-known/openid-configuration", nil)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()

			return
		}

		client := &http.Client{
			Timeout:   10 * time.Second,
			Transport: http.DefaultTransport,
		}

		oidcResp, err := client.Do(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()

			return
		}

		defer oidcResp.Body.Close()

		if oidcResp.StatusCode != http.StatusOK {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"error": fmt.Errorf("falied to request openid-configuration"),
			})
			ctx.Abort()

			return
		}

		// retrieve jwkset uri from oidc response
		var generic map[string]interface{}
		if err = json.NewDecoder(oidcResp.Body).Decode(&generic); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()

			return
		}

		jwksUri, ok := generic["jwks_uri"].(string)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Errorf("conversion failed"),
			})
			ctx.Abort()

			return
		}

		keySet, err := jwk.Fetch(ctx, jwksUri)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Errorf("failed to fetching JWK sets"),
			})
			ctx.Abort()

			return

		}

		tok, err := jwt.ParseString(idToken, jwt.WithKeySet(keySet, jws.WithInferAlgorithmFromKey(true)), jwt.WithValidate(true), jwt.WithAudience(clientId), jwt.WithContext(ctx))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Sprint("conversion failed"),
			})
			ctx.Abort()

			return
		}

		fmt.Println("Valid token")
		fmt.Printf("iss: %s\n", tok.Issuer())
		fmt.Printf("aud: %v\n", tok.Audience())
		fmt.Printf("exp: %s\n", tok.Expiration())
		fmt.Printf("sub: %s\n", tok.Subject())
		fmt.Printf("jti: %s\n", tok.JwtID())

		// for azure entra id user account
		preferred_username, ok := tok.Get("preferred_username")

		if !ok {
			fmt.Println("nothing field `preferred_username`")
		} else {
			str_username := fmt.Sprintf("%v", preferred_username)
			id, ok := GetUserID(ctx)
			if !ok {
				// 登録する
				// ginのContextとginのRequest.Contextがある + SetUserIDは素のnet/httpを扱うことにしているのでややこしい。基本的に素のContext単体の世界で考えるようにする
				_tmpCtx := SetUserID(ctx.Request.Context(), entity.UserID(str_username))
				// WARN: RequestとContextの置き換えで意図せぬ上書きや抜け落ちがあるかもしれない
				ctx.Request = ctx.Request.WithContext(_tmpCtx)

			} else if id != entity.UserID(str_username) {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"error": fmt.Errorf("not match user id"),
				})
				ctx.Abort()
				return
			}

		}
		ctx.Next()
		// return
	}
}
