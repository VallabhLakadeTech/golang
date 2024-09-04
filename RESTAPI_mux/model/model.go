package model

type Books struct {
	ID              string `json:"id"`
	BookName        string `json:"bookname"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publicationyear"`
}
