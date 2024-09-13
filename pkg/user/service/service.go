package service

import (
	"bytes"

	"io"
	"mime/multipart"
	"s3service/pkg/client"
)

type serviceInt interface {
	Upload(file multipart.File, head *multipart.FileHeader) error
	Retrive(file string) (string, error)
}
type Service struct {
	S3 client.S3pathway
}

func (s *Service) Retrive(file string) (string, error) {
	url, err := s.S3.RetriveFile(file)
	if err != nil {
		return "", err

	}
	return url, nil
}

func (s *Service) Upload(file multipart.File, head *multipart.FileHeader) error {

	filebyte, err := convertToByte(file)
	if err != nil {
		return err
	}

	err = s.S3.UploadFile(filebyte, head.Filename)
	if err != nil {

		return err
	}
	return nil
}

func convertToByte(file multipart.File) ([]byte, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, file)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

}
