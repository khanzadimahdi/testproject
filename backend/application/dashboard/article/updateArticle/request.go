package updatearticle

import (
	"time"
)

type validationErrors map[string]string

type Request struct {
	UUID        string    `json:"uuid"`
	Cover       string    `json:"cover"`
	Video       string    `json:"video"`
	Title       string    `json:"title"`
	Excerpt     string    `json:"excerpt"`
	Body        string    `json:"body"`
	PublishedAt time.Time `json:"published_at"`
	AuthorUUID  string    `json:"-"`
	Tags        []string  `json:"tags"`
}

func (r *Request) Validate() (bool, validationErrors) {
	return true, nil
}
