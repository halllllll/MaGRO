package service

import (
	"context"
	"fmt"

	"github.com/halllllll/MaGRO/entity"
)

type ListTask struct {
	Repo TaskLister
}

func (l *ListTask) ListTask(ctx context.Context) (*entity.Tasks, error) {
	tasks, err := l.Repo.ListTasks(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return tasks, nil
}
