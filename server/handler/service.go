package handler

import (
	"context"

	"github.com/halllllll/MaGRO/entity"
)

type ListTasksService interface {
	ListTask(ctx context.Context) (*entity.Tasks, error)
}
type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}

type RegisterUserService interface {
	RegisterUser(ctx context.Context, name, password, role string) (*entity.User_T, error)
}

type AdminMutateService interface {
	// TODO: いったん返さずrepositoryで結果を確認する
	UpdateRole(ctx context.Context, req *entity.ReqNewRoleAlias) error
}

type SystemService interface {
	GetSystemInfo(ctx context.Context) (*entity.System, error)
}

type MaGROUnitService interface {
	ListUnit(ctx context.Context) ([]*entity.Unit, error)

	ListUsersSubunit(ctx context.Context, unitId *entity.UnitId) (*entity.RespUsersSubunit, error)
}

// TODO: とりあえず受け取るだけの仮実装
type MaGRORepassService interface{
	RepassUser(ctx context.Context)
}