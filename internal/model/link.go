package model

type LinkRequest struct {
	Url string `json:"url"`
}

// TODO
type LinkResponse struct {
	ID        string `json:"id"`
	LongLink  string `json:"long_lin"`
	ShortLink string `json:"short_link"`
}
