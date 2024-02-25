package response

import "github.com/olegtemek/file-handler/internal/model"

type FileCreate struct {
	*model.File `json:"file"`
	Status      int `json:"status"`
}

type FileGetAll struct {
	Files  []*model.File `json:"files"`
	Status int           `json:"status"`
}
type FileGetAllTags struct {
	Tags   []*string `json:"tags"`
	Status int       `json:"status"`
}

type FileDelete struct {
	*model.File `json:"file"`
	Status      int `json:"status"`
}

type FileGetOne struct {
	*model.File `json:"file"`
	Status      int `json:"status"`
}
