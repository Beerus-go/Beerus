package commons

import "mime/multipart"

type BeeFile struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}
