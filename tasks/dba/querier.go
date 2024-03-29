// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package taskmodel

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateTask(ctx context.Context, arg CreateTaskParams) (sql.Result, error)
	DeleteTask(ctx context.Context, id int64) error
	GetTask(ctx context.Context, id int64) (*Task, error)
	ListCalcTasks(ctx context.Context) ([]*Task, error)
	ListCalcTasksByJob(ctx context.Context, jobID int64) ([]*Task, error)
	ListTasks(ctx context.Context) ([]*Task, error)
	ListTasksByJob(ctx context.Context, jobID int64) ([]*Task, error)
	UpdateTask(ctx context.Context, arg UpdateTaskParams) (sql.Result, error)
	UpdateTaskStatus(ctx context.Context, arg UpdateTaskStatusParams) (sql.Result, error)
}

var _ Querier = (*Queries)(nil)
