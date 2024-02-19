package dto

import "mime/multipart"

type FileCreateDto struct {
	File *multipart.File `form:"file" validate:"required"`
	Tag  string          `form:"tag" validate:"required"`
}
