package service

import "github.com/bhavdipd/demo/domain"

type AlbumService interface {
	GetAllAlbum() ([]domain.Album, error)
}

type DefaultAlbumService struct {
	repo domain.AlbumRepository
}

func (s DefaultAlbumService) GetAllAlbum() ([]domain.Album, error) {
	return s.repo.FindAll()
}

func NewAlbumService(repository domain.AlbumRepository) DefaultAlbumService {
	return DefaultAlbumService{repository}
}
