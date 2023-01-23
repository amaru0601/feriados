package models

type ApiResponse struct {
	Status string `json:"status"`
	Data   []Data `json:"data"`
}

type Data struct {
	Date        string `json:"date"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Inalienable bool   `json:"inalienable"`
	Extra       string `json:"extra"`
}
