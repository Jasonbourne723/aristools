package dto

type AddTodoDto struct {
	Name string `json:"name"`
	DoAt string `json:"do_at"`
}

type TodoDto struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	DoAt   string `json:"do_at"`
	DoneAt string `json:"done_at"`
}
