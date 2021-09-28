package todolist

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    int    `json:"deadline"`
}

type Tasks []*Task

// type tasksResponse struct {
// 	Msg  string `json:"msg"`
// 	Data []Task `json:"data"`
// }

// type taskResponse struct {
// 	Msg  string `json:"msg"`
// 	Data Task   `json:"data"`
// }
