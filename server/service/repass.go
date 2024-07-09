package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"sort"
	"strings"
	"time"

	"math/rand/v2"

	"github.com/halllllll/MaGRO/auth"
	"github.com/halllllll/MaGRO/entity"
)

// uniqを使いたいのでslicesを使うためにsort済でないといけない
type Targets []*entity.UserPrimaryUniqID

func (t Targets) SortByUUID() {
	sort.SliceStable(t, func(i, j int) bool {
		return t[i].ID < t[j].ID
	})
}

type Repass struct {
	Repo MaGRORepasser
}


/*

👺 this is temporary implementation til asign entra app to user manager privilege 👹

*/
func (r *Repass) RepassUser(ctx context.Context, unitId *entity.UnitId, _target []*entity.UserPrimaryUniqID) (*entity.RespRepass, error) {
	fmt.Println("repass action:")

	// get access token
	token, err := auth.GetAccessTokenBySecret(ctx)
	if err != nil {
		fmt.Println("repass error: %w\n", err.Error())
		return nil, err
	}

	// contextからID
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user id not found")
	}
	// まずIDからUUIDを引く
	uuid, err := r.Repo.Me(ctx, &id)
	if err != nil {
		return nil, err
	}

	// uuidとunit idから所属してるメンバーを引く
	result, err := r.Repo.ListUsersSubunits(ctx, uuid, unitId)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}

	// uuid:Userのmapにする
	m := map[entity.UserUUID]entity.User{}
	for _, v := range result {
		m[entity.UserUUID(v.UserID)] = entity.User{
			UserID:      entity.UserUUID(v.UserID),
			UserName:    entity.UserID(v.AccountID),
			DisplayName: v.UserName,
			UserSortKey: *v.UserKananame,
			UserType:    entity.Role(v.Role),
			Status:      "", // TODO: なんか引き起こしそう
		}
	}

	// フロントから送られてきたデータと間違っていないか所属を確認　ついでに戻り値の型に合わせるためにUserとして取得しておく
	resultByUuid := map[entity.UserUUID]entity.RepassEveryResult{}

	// uniqする
	var target Targets
	target = _target
	target.SortByUUID()
	uniqTarget := slices.Compact(target)

	for _, v := range uniqTarget {
		x, ok := m[entity.UserUUID(v.ID)]
		if !ok {
			return nil, fmt.Errorf("include invalid id")
		}
		if x.UserName != v.Account {
			return nil, fmt.Errorf("include invalid id")
		}
		resultByUuid[entity.UserUUID(v.ID)] = entity.RepassEveryResult{
			User: x,
		}
	}

	alphabet := create2z()
	minn, maxn := 100000, 999999
	for _, t := range uniqTarget {
		// すでに保存しているuserだけ取り出してあとでほかのデータといっしょに入れ直す
		resultData, ok := resultByUuid[t.ID]
		if !ok {
			fmt.Printf("not found candidate ID: %s\n", t.ID)
			continue
		}
		user := resultData.User
		a := alphabet[rand.N(len(alphabet))]
		num := rand.N(maxn-minn+1) + minn
		newPassword := fmt.Sprintf("%s%s%d", strings.ToUpper(string(a)), strings.ToLower(string(a)), num)

		result := entity.RepassEveryResult{User: user}


		// テスト: とりあえず半分の確率で失敗するようにする
		// if rand.N(4) < 1 {
		// 	result.Result = entity.ER
		// 	result.Message = "エラー: desu"
		// } else {
		// 	result.Result = entity.OK
		// 	result.Message = "成功"
		// 	result.Issue = newPassword
		// }

		// exponentialにすべきだが面倒なのでとりあえず
		time.Sleep(time.Duration(100 * rand.N(5)) * time.Millisecond)

		if err := repass(ctx, token, newPassword, entity.UserID(user.UserName)); err != nil{
			result.Result = entity.ER
			result.Message = err.Error()
		}else{
			result.Result = entity.OK
			result.Message = "update password"
			result.Issue = newPassword
		}

		// 入れ直す
		resultByUuid[t.ID] = result
	}
	results := []entity.RepassEveryResult{}
	for _, v := range resultByUuid {
		results = append(results, v)
	}

	return &entity.RespRepass{Result: results}, nil
}

func create2z() string {
	alphabet := make([]rune, 0)
	for cur := []rune("a"); ; cur[0] += 1 {
		alphabet = append(alphabet, cur[0])
		if string(cur[0]) == "z" {
			break
		}
	}
	return string(alphabet)
}


func repass(ctx context.Context, token string, password string, id entity.UserID)(error){
	// graph api
	// https://learn.microsoft.com/en-us/graph/api/user-update?view=graph-rest-1.0&tabs=http#example-3-update-the-passwordprofile-of-a-user-and-reset-their-password
	type passwordProfile struct{
		ForceChangePassordNextSignIn bool `json:"forceChangePasswordNextSignIn"`
		Password string `json:"password"`
	}
	type requestBody struct{
		PasswordProfile passwordProfile `json:"passwordProfile"`
	}

	body := &requestBody{
		PasswordProfile: passwordProfile{
			ForceChangePassordNextSignIn: true,
			Password: password,
		},
	}

	jsonBody, err := json.Marshal(&body)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://graph.microsoft.com/v1.0/users/%s", id)

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	
	defer resp.Body.Close()

	//　リファレンスによると成功した場合は
	// `HTTP/1.1 204 No Content`
	fmt.Printf("resp:\n%#v\n", resp)

	if resp.StatusCode != http.StatusNoContent{
		return fmt.Errorf("%s", resp.Status)
	}

	fmt.Println("repass done;")
	return nil
}
