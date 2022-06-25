package usecase

import (
	"fmt"

	"github.com/74th/vscode-book-r2-golang/domain/entity"
)

var RepositoryError = fmt.Errorf("RepositoryError")
var TaskNotFoundError = fmt.Errorf("TaskNotFoundError")

func (it *Interactor) ShowTasks() ([]*entity.Task, error) {
	return it.Repository.List()
}

func (it *Interactor) CreateTask(task *entity.Task) (*entity.Task, error) {
	newID, err := it.Repository.Add(task)
	if err != nil {
		return nil, RepositoryError
	}
	task.ID = newID
	return task, nil
}

func (it *Interactor) DoneTask(id int) (*entity.Task, error) {
	task, err := it.Repository.Get(id)
	if err != nil {
		return nil, RepositoryError
	}
	if task == nil {
		return nil, TaskNotFoundError
	}

	task.Done = true

	err = it.Repository.Update(task)
	if err != nil {
		return nil, RepositoryError
	}
	return task, nil
}
