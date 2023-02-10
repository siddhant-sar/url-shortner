package models


type UrlObject struct {
	LongUrl   *string `json:"long_url,omitempty"`
	ShortUrl  *string `json:"short_url,omitempty"`
	TotalHits *int64  `json:"hits,omitempty"`
	Ttl       *int64  `json:"ttl,omitempty"`
}
