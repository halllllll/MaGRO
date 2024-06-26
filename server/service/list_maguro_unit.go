package service

import (
	"context"

	"github.com/halllllll/MaGRO/entity"
)

type ListUnit struct {
	Repo MaGROLister
}

// TODO: いったん返さず確認する
func (l *ListUnit) ListUnit(ctx context.Context, userid *entity.UserID) ([]*entity.Unit, error) {
	result, err := l.Repo.ListUnits(ctx, userid)
	if err != nil {
		return nil, err
	}
	units := []*entity.Unit{}
	for _, v := range result {
		units = append(units, &entity.Unit{
			Name:   v.Name,
			UnitID: entity.UnitId(v.ID),
		})
	}
	return units, nil
}

// TODO: 中身を加工しよう
func (l *ListUnit) ListUsersSubunit(ctx context.Context, unitid *entity.UnitId) (*entity.RespUsersSubunit, error) {
	result, err := l.Repo.ListUsersSubunits(ctx, unitid)
	if err != nil {
		return nil, err
	}
	// どういう情報になったらいい？
	// ロイロをみてみよう
	// 全体情報
	// - currentUser
	//		- id, name(学校), username,isAdmin
	//		- school
	//			- id, name（学校）,timezone, isXXXX系
	//
	// - school
	//		- ↑と同じ内容
	// - district
	//	- 自治体とか,全体管理から設定されるデータっぽい(sso)
	//
	// ページごとの個別情報
	// 生徒一覧
	// 	- そのページのデータ
	//    - 総ページ数・pagenation
	//  - user_groups(array)
	//  	- user data
	//			- id, username, sort_key, displayname, schoolId
	//			- user_groups(在籍してるやつ)（そのschoolIdのやつね）
	//				- id, name,grade_string(学年), date情報
	//  		↑がたくさん
	//	- usertype
	// 	- totalUsersCount
	// TODO: 並列処理で取得・加工して返すAPI

	var subunits []*entity.SubUnit
	subunitIdMap := make(map[int]*entity.SubUnit)
	var users []*entity.User
	userIdMap := make(map[string]bool)
	subunitOperators := make(map[int][]*entity.User)

	ret := &entity.RespUsersSubunit{
		Unit: &entity.Unit{
			UnitID: *unitid,
			Name:   result[0].Unit,
		},
		// TODO: current user
		UserGroups:    []*entity.User{},
		SubunitGroups: []*entity.OperatorSubunits{},
	}

	for _, v := range result {
		if _, ok := subunitIdMap[int(v.SubunitID)]; !ok {
			newSubunit := &entity.SubUnit{
				SubunitID: int(v.SubunitID),
				Name:      v.Subunit,
				IsPublic:  v.Public,
				CreatdAt:  v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			}
			subunitIdMap[int(v.SubunitID)] = newSubunit
			subunits = append(subunits, newSubunit)
			if v.Role != string(entity.RoleGuest) {
				if _, okok := subunitOperators[int(v.SubunitID)]; !okok {
					subunitOperators[int(v.SubunitID)] = []*entity.User{}
				}
				operators := subunitOperators[int(v.SubunitID)]
				operators = append(operators, &entity.User{
					UserID:      entity.UserID(v.UserID),
					UserName:    v.AccountID,
					DisplayName: v.UserName,
					UserSortKey: *v.UserKananame,
					UserType:    entity.Role(v.Role),
				})
			}
		}
		if _, ok := userIdMap[v.UserID]; !ok {
			userIdMap[v.UserID] = true
			users = append(users, &entity.User{
				UserID:      entity.UserID(v.UserID),
				UserName:    v.AccountID,
				DisplayName: v.UserName,
				UserSortKey: *v.UserKananame,
				UserType:    entity.Role(v.Role),
			})
		}
	}

	// operator subunit
	for _, v := range subunits {
		rs := &entity.OperatorSubunits{
			Subunit:   v,
			Operators: subunitOperators[v.SubunitID],
		}
		ret.SubunitGroups = append(ret.SubunitGroups, rs)
	}

	ret.UserGroups = users
	ret.UserCount = len(users)
	ret.SubunitCount = len(subunits)

	return ret, nil
}
