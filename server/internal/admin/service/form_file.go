// @author AlphaSnow

package service

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
)

type FormFileService struct {
	file *multipart.FileHeader
}

func NewFormFileService(f *multipart.FileHeader) *FormFileService {
	return &FormFileService{file: f}
}

func (s *FormFileService) GetHashName() (string, error) {
	src, err := s.file.Open()
	defer src.Close()
	if err != nil {
		return "", err
	}

	hash := md5.New()
	if _, err := io.Copy(hash, src); err != nil {
		return "", err
	}

	hashName := fmt.Sprintf("%x", hash.Sum(nil))

	ext := filepath.Ext(s.file.Filename)
	fileExt := ext[1:]

	return hashName + "." + fileExt, nil
}
