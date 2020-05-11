package requests

type Create struct {
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body"`
	Status string `json:"status" binding:"required"`
}
