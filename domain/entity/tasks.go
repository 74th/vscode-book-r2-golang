package entity

// Task タスク
type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func (t Task) Clone() *Task {
	return &t
}
