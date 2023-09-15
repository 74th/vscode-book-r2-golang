package usecase

import (
	"github.com/74th/vscode-book-r2-golang/domain/entity"
)

// タスクデータベース
type TaskDatabase interface {
	// タスクの追加
	Add(*entity.Task) (int, error)
	// 未完了のタスク一覧
	SearchUnfinished() ([]*entity.Task, error)
	// タスクの更新
	Update(*entity.Task) error
	// タスクの取得
	Get(id int) (*entity.Task, error)
}
