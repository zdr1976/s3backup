package main

import "time"

type ObjectAttributes struct {
	ContentLength int64     `json:"content_length"`
	ContentType   string    `json:"content_type"`
	ETag          string    `json:"etag"`
	LastModified  time.Time `json:"last_modified"`
}
