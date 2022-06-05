package repository

import (
	"testing"

	"github.com/74th/vscode-book-r2-golang/domain/entity"
)

func TestNew(t *testing.T) {
	rep := New()
	if len(rep.(*instance).tasks) != 2 {
		t.Errorf("初期化時点で2つのタスクが格納されていること %d", len(rep.(*instance).tasks))
	}
}
func TestAdd(t *testing.T) {
	rep := New()
	newID, err := rep.Add(&entity.Task{
		Text: "new task",
	})
	if err != nil {
		t.Errorf("エラーを返さないこと")
	}

	if len(rep.(*instance).tasks) != 3 {
		t.Errorf("タスクが追加されていること %d", len(rep.(*instance).tasks))
	} else {
		addedTask := rep.(*instance).tasks[2]

		if addedTask.Text != "new task" {
			t.Errorf("追加したタスクが末尾に追加されていること %s", addedTask.Text)
		}

		if addedTask.ID <= 2 {
			t.Errorf("タスクに新しいIDが振られること %d", addedTask.ID)
		}
		if addedTask.ID != newID {
			// TODO: メッセージ
			t.Errorf("かえされたID")
		}

		for i, task := range rep.(*instance).tasks {
			if i != 2 {
				if addedTask.ID == task.ID {
					t.Errorf("既存のタスクとは異なるIDが振られていること %d == %d", addedTask.ID, task.ID)
				}
				if addedTask.Text == task.Text {
					t.Errorf("既存のタスクとを上書きしていないこと %s == %s", addedTask.Text, task.Text)
				}
			}
		}

	}
}

func TestGet(t *testing.T) {
	// TODO:
}

func TestUpdate(t *testing.T) {
	// TODO:
}
