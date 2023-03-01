package todos

type Todo struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
