package usecase_test

import (
	"testing"

	"github.com/74th/vscode-book-r2-golang/domain/entity"
	"github.com/74th/vscode-book-r2-golang/domain/usecase"
	"github.com/74th/vscode-book-r2-golang/repository"
)

func newInteractor() usecase.Interactor {
	return usecase.Interactor{
		Repository: repository.New(),
	}
}

func TestTaskWork(t *testing.T) {
	it := newInteractor()

	tasks, err := it.ShowTasks()
	if err != nil {
		t.Error("エラーが返らないこと")
		return
	}
	if len(tasks) > 0 {
		t.Error("初期状態のリポジトリからはからのタスクが引けること")
		return
	}

	newTask := &entity.Task{
		Text: "task1",
	}

	it.CreateTask(newTask)
}
