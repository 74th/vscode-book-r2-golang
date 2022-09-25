package usecase

import (
	"fmt"
	"log"

	"errors"

	"github.com/74th/vscode-book-r2-golang/domain/entity"
)

var TaskNotFoundError = errors.New("TaskNotFoundError")

// 未完了タスクの一覧
func (it *Interactor) ShowTasks() ([]*entity.Task, error) {
	tasks, err := it.Database.SearchUnfinished()
	if err != nil {
		log.Printf("Database error occurred: %s", err)
		return nil, fmt.Errorf("Database error occurred: %w", err)
	}
	return tasks, nil
}

// タスクの生成
func (it *Interactor) CreateTask(task *entity.Task) (*entity.Task, error) {
	newID, err := it.Database.Add(task)
	if err != nil {
		log.Printf("Database error occurred: %s", err)
		return nil, fmt.Errorf("Database error occurred: %w", err)
	}
	task.ID = newID
	return task, nil
}

// タスクの完了
func (it *Interactor) DoneTask(id int) (*entity.Task, error) {
	task, err := it.Database.Get(id)
	if err != nil {
		log.Printf("Database error occurred: %s", err)
		return nil, fmt.Errorf("Database error occurred: %w", err)
	}
	if task == nil {
		return nil, fmt.Errorf("%w", TaskNotFoundError)
	}

	task.Done = true

	err = it.Database.Update(task)
	if err != nil {
		log.Printf("Database error occurred: %s", err)
		return nil, fmt.Errorf("Database error occurred: %w", err)
	}
	return task, nil
}
