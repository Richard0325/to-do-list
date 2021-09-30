package todolist

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    int    `json:"deadline"`
}

type Tasks []*Task
