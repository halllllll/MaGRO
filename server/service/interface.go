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

// TODO:
type MaGROLister interface {
	ListUnits(ctx context.Context, unitid *entity.UserID) ([]db.Unit, error)                        // いったん返さずDBで出力して確認する
	ListUsersSubunits(ctx context.Context, userid *entity.UnitId) ([]db.GetUsersSubunitsRow, error) // 同じ
}

type MaGROMutater interface {
	UpdateRole(ctx context.Context, roles *entity.ReqNewRoleAlias) error
}
