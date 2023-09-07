package model

type NewItem struct {
	Title  string `json:"title"`
	Owner  string `json:"owner"`
	Rating int    `json:"rating"`
}

type UpdateItem struct {
	Title  *string `json:"title"`
	Owner  *string `json:"owner"`
	Rating *int    `json:"rating"`
}
