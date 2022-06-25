package usecase

import (
	"github.com/74th/vscode-book-r2-golang/domain/entity"
)

type Interactor struct {
	Repository Repository
}

// Repository タスクリポジトリ
type Repository interface {
	Add(*entity.Task) (int, error)
	List() ([]*entity.Task, error)
	Update(*entity.Task) error
	Get(id int) (*entity.Task, error)
}
