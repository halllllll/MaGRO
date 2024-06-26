package service

import (
	"context"

	"github.com/halllllll/MaGRO/entity"
)

type AddTask struct {
	Repo TaskAdder
}

func (a *AddTask) AddTask(ctx context.Context, title string) (*entity.Task, error) {
	t := &entity.Task{
		Title:  title,
		Status: entity.TaskStatusTodo,
	}
	// 本書だと引数に渡したtにDBから格納されたデータが入るのでそれを使うのだがsqlcの生成したコードだと(生成されたmodelの)Taskが返ってくる
	task, err := a.Repo.AddTask(ctx, t)

	if err != nil {
		return nil, err
	}

	return &entity.Task{
		ID:       entity.TaskID(task.ID),
		Title:    task.Title,
		Status:   entity.TaskStatus(task.Status),
		Created:  task.Created,
		Modified: task.Modified,
	}, nil
}
