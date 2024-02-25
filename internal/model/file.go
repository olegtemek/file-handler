package model

import "time"

type File struct {
	Id        int       `json:"id"`
	FilePath  string    `json:"filepath"`
	Tag       string    `json:"tag"`
	Timestamp time.Time `json:"timestamp"`
}
