package model

type (
	Search struct {
		Title  string
		Year   string
		ImdbID string
		Type   string
	}

	SearchResponse struct {
		Search       []Search
		Response     string
		Error        string
		totalResults int
	}
)
