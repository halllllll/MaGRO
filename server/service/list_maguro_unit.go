package service

import (
	"context"
	"fmt"

	"github.com/halllllll/MaGRO/auth"
	"github.com/halllllll/MaGRO/entity"
)

type ListUnit struct {
	Repo MaGROLister
}

func (l *ListUnit) ListUnit(ctx context.Context) ([]*entity.Unit, error) {

	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user id not found")
	}

	result, err := l.Repo.ListUnits(ctx, &id)

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
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user id not found")
	}

	// まずidからlookupする必要があるよね
	uuid, err := l.Repo.Me(ctx, &id)
	if err != nil {
		return nil, err
	}
	result, err := l.Repo.ListUsersSubunits(ctx, uuid, unitid)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return &entity.RespUsersSubunit{}, nil
	}
	// TODO: 並列処理で取得・加工して返すAPI
	// 現状かなり無理矢理やっとる

	var subunits []*entity.SubUnit
	subunitIdMap := make(map[int]*entity.SubUnit)
	var usersWithSubgroupIds []*entity.UserWithSubgroups

	// belongs subgroups by each user
	belongsSubunitByUser := make(map[entity.UserUUID][]entity.SubunitId)

	userIdMap := make(map[string]bool)
	subunitOperators := make(map[int][]*entity.UserID)

	ret := &entity.RespUsersSubunit{
		Unit: &entity.Unit{
			UnitID: *unitid,
			Name:   result[0].Unit,
		},
		CurrentUser: &entity.User{
			UserID:   entity.UserUUID(*uuid),
			UserName: entity.UserID(id),
			// TODO: rest of current user data
		},
		UserGroups:    []*entity.UserWithSubgroups{},
		SubunitGroups: []*entity.OperatorSubunits{},
	}

	for _, v := range result {
		fmt.Printf("%s, role %s\n", v.AccountID, v.RoleAlias)
		if _, ok := subunitIdMap[int(v.SubunitID)]; !ok {
			newSubunit := &entity.SubUnit{
				SubunitID: entity.SubunitId(v.SubunitID),
				Name:      v.Subunit,
				IsPublic:  v.Public,
				CreatdAt:  v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			}
			subunitIdMap[int(v.SubunitID)] = newSubunit
			subunits = append(subunits, newSubunit)
		}
		if v.Role != string(entity.RoleGuest) {
			if _, okok := subunitOperators[int(v.SubunitID)]; !okok {
				subunitOperators[int(v.SubunitID)] = []*entity.UserID{}
			}
			operators := subunitOperators[int(v.SubunitID)]
			// uuidではなくuserのaccount idを示したい
			operators = append(operators, (*entity.UserID)(&v.AccountID))
			subunitOperators[int(v.SubunitID)] = operators
		}
		if _, ok := userIdMap[v.UserID]; !ok {
			userIdMap[v.UserID] = true
			usersWithSubgroupIds = append(usersWithSubgroupIds, &entity.UserWithSubgroups{
				User: entity.User{
					UserID:      entity.UserUUID(v.UserID),
					UserName:    entity.UserID(v.AccountID),
					DisplayName: v.UserName,
					UserSortKey: *v.UserKananame,
					UserType:    entity.Role(v.Role),
				},
				BelongsSubunit: []entity.SubunitId{},
			})
			belongsSubunitByUser[(entity.UserUUID)(v.UserID)] = []entity.SubunitId{}
		}

		// belogns subunit
		belongsSubunitByUser[(entity.UserUUID)(v.UserID)] = append(belongsSubunitByUser[(entity.UserUUID)(v.UserID)], (entity.SubunitId)(v.SubunitID))
	}

	// operator subunit
	for _, v := range subunits {
		rs := &entity.OperatorSubunits{
			Subunit:   v,
			Operators: subunitOperators[int(v.SubunitID)],
		}
		ret.SubunitGroups = append(ret.SubunitGroups, rs)
	}

	// subunit
	for _, v := range usersWithSubgroupIds {
		vv, ok := belongsSubunitByUser[v.UserID]
		// 判定する必要ないと思うが一応
		if ok {
			v.BelongsSubunit = vv
		}
	}

	ret.UserGroups = usersWithSubgroupIds
	ret.UserCount = len(usersWithSubgroupIds)
	ret.SubunitCount = len(subunits)

	return ret, nil
}
