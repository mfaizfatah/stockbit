package model

type OMDBResponse struct {
	Search       []Search `json:"Search,omitempty"`
	TotalResults string   `json:"totalResults"`
	Response     string   `json:"Response"`
	Error        string   `json:"Error,omitempty"`
}

type Search struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}
