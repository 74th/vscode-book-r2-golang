package repository

import "github.com/74th/vscode-book-r2-golang/domain/entity"

// タスクデータベース
type TaskDatabase interface {
	Add(*entity.Task) (int, error)
	SearchUnfinished() ([]*entity.Task, error)
	Update(*entity.Task) error
	Get(id int) (*entity.Task, error)
}
