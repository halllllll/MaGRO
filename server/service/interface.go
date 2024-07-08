package service

import (
	"context"

	"github.com/halllllll/MaGRO/entity"
	db "github.com/halllllll/MaGRO/gen/sqlc"
)

// とりあえず本書からコピペした（moq使ってない）
// 戻り値は(sqlcで生成したコードを使う都合上)本書と違ってmodelもかえってくる
type TaskAdder interface {
	AddTask(ctx context.Context, t *entity.Task) (*entity.Task, error)
}

type TaskLister interface {
	ListTasks(ctx context.Context) (*entity.Tasks, error)
}

type UserRegister interface {
	RegisterUser(ctx context.Context, u *entity.User_T) (*entity.User_T, error)
}

type Infoer interface {
	GetSystemInfo(ctx context.Context) (*entity.System, error)
}

type ListUsersSubunits interface{
	ListUsersSubunits(ctx context.Context, userUuid *entity.UserUUID, unitId *entity.UnitId) ([]db.GetSubunitsByUserUuIDAndUnitIdRow, error) // 同じ
}

type Me interface{
	Me(ctx context.Context, userId *entity.UserID) (*entity.UserUUID, error)
}

type MaGROLister interface {
	ListUnits(ctx context.Context, unitid *entity.UserID) ([]db.Unit, error)
	// Me(ctx context.Context, userId *entity.UserID) (*entity.UserUUID, error)
	// ListUsersSubunits(ctx context.Context, userUuid *entity.UserUUID, unitId *entity.UnitId) ([]db.GetSubunitsByUserUuIDAndUnitIdRow, error) // 同じ
	Me
	ListUsersSubunits
}

type MaGROMutater interface {
	UpdateRole(ctx context.Context, roles *entity.ReqNewRoleAlias) error
}

type MaGRORepasser interface {
	Repass(ctx context.Context)
	// Me(ctx context.Context, userId *entity.UserID) (*entity.UserUUID, error)
	// ListUsersSubunits(ctx context.Context, userUuid *entity.UserUUID, unitId *entity.UnitId) ([]db.GetSubunitsByUserUuIDAndUnitIdRow, error) // 同じ
	Me
	ListUsersSubunits

}
