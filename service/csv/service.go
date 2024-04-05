package csv

import "io"

type Service interface {
	Import(file io.Reader) error
	Export(fileType string) ([]byte, error)
}

type ServiceImpl struct{}

func NewService() *ServiceImpl {
	return &ServiceImpl{}
}

func (s *ServiceImpl) Import(file io.Reader) error {
	return nil
}

func (s *ServiceImpl) Export(fileType string) ([]byte, error) {
	return []byte{}, nil
}
