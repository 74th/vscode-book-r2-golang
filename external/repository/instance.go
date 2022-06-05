package repository

import (
	"fmt"

	"github.com/74th/vscode-book-r2-golang/domain/entity"
	"github.com/74th/vscode-book-r2-golang/domain/usecase"
)

// repository タスク	リポジトリの実装
type instance struct {
	tasks []entity.Task
}

// New リポジトリの作成
func New() usecase.Repository {
	s := new(instance)
	s.tasks = make([]entity.Task, 2, 20)
	s.tasks[0] = entity.Task{
		ID:   1,
		Text: "task1",
		Done: false,
	}
	s.tasks[1] = entity.Task{
		ID:   2,
		Text: "task2",
		Done: false,
	}
	return s
}

// Add タスクの追加
func (s *instance) Add(task *entity.Task) (int, error) {
	task.ID = len(s.tasks) + 1
	s.tasks = append(s.tasks, *task)
	return task.ID, nil
}

// List 未完了のタスクの一覧
func (s *instance) List() ([]*entity.Task, error) {
	result := []*entity.Task{}
	for _, task := range s.tasks {
		if !task.Done {
			result = append(result, task.Clone())
		}
	}
	return result, nil
}

// Update タスクを更新にする
func (s *instance) Update(task *entity.Task) error {
	for i, currentTask := range s.tasks {
		if currentTask.ID == task.ID {
			s.tasks[i] = *task
			return nil
		}
	}
	return fmt.Errorf("Not found id:%d", task.ID)
}

// Get タスクを取得する
func (s *instance) Get(id int) (*entity.Task, error) {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks[i].Done = true
			return task.Clone(), nil
		}
	}
	return nil, nil
}
