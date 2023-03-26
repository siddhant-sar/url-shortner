package models

type User struct {
	Active      bool   `'json:"is_active"`
	CustomAlias string `'json:"custom_alias,omitempty"`
	FirstName   string `'json:"first_name"`
	LastName    string `'json:"last_name"`
	Password    string `'json:"password"`
}

type UrlObjectInput struct {
	CustomAlias string `json:"custom_alias,omitempty"`
	LongUrl     string `json:"long_url"`
	Ttl         string `json:"ttl"`
}

type UrlObject struct {
	LongUrl     string `json:"long_url"`
	ShortUrlKey string `json:"short_url"`
	TotalHits   int64  `json:"hits"`
	Ttl         int64  `json:"ttl"`
	User        *User  `json:"user,omitempty"`
}
