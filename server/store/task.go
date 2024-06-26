package store

import (
	"context"

	"github.com/halllllll/MaGRO/entity"
	db "github.com/halllllll/MaGRO/gen/sqlc"
	"github.com/halllllll/MaGRO/store/transaction"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) ListTasks(ctx context.Context) (*entity.Tasks, error) {
	tasks := entity.Tasks{}
	// トランザクションを使わない
	result, err := r.query.AllTask(ctx)
	if err != nil {
		return nil, err
	}

	for _, t := range result {
		tasks = append(tasks, &entity.Task{
			ID:       entity.TaskID(t.ID),
			Title:    t.Title,
			Status:   entity.TaskStatus(t.Status),
			Created:  t.Created,
			Modified: t.Modified,
		})
	}

	return &tasks, nil
}

func (r *Repository) AddTask(ctx context.Context, t *entity.Task) (*entity.Task, error) {
	// トランザクションを使う
	tr := transaction.NewTransaction(r.pool)

	// DBでDEFAULT CURRENT_TIMEを使っているのと、sqlcで生成したコードを使っているから今回は不要だが、CreatedAtなどをつけるときはここでつける(こんな感じ↓)
	t.Created = r.Clocker.Now()

	args := &db.AddTaskParams{
		Title:  t.Title,
		Status: string(t.Status),
		// Created もしDBでDEFAULT CURRENT_TIMEがなければ（追加するSQLを書いてsqlcで生成したAddParamsで）追加する
	}
	var task *entity.Task
	// トランザクション実行
	err := tr.WithTransaction(ctx, func(tx pgx.Tx) error {
		txQuery := r.query.WithTx(tx)
		newTask, err := txQuery.AddTask(ctx, *args)
		if err != nil {
			return err
		}
		task = &entity.Task{
			ID:       entity.TaskID(newTask.ID),
			Status:   entity.TaskStatus(newTask.Status),
			Created:  newTask.Created,
			Modified: newTask.Modified,
		}
		return nil
	})
	return task, err
}
