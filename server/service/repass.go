package service

import (
	"context"
	"fmt"
	"slices"
	"sort"

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

func (r *Repass) RepassUser(ctx context.Context, unitId *entity.UnitId, _target []*entity.UserPrimaryUniqID) (*entity.RespRepass, error) {
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

	fmt.Printf("target? %#v\n", uniqTarget)
	// TODO: ここで処理をするが、とりあえず仮に、結果をシミュレートしてフロントにデータだけ返す

	fmt.Println("わ〜〜い")

	return nil, nil
}
