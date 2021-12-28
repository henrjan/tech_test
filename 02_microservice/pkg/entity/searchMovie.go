package entity

type SearchMovie struct {
	SearchResult []Movie `json:"Search,omitempty"`
	TotalResult  string  `json:"totalResults,omitempty"`
	Response     string  `json:"Response,omitempty"`
	Error        string  `json:"Error,omitempty"`
}
