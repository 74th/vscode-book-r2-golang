package usecase

import "github.com/74th/vscode-book-r2-golang/domain/repository"

type Interactor struct {
	Database repository.TaskDatabase
}
