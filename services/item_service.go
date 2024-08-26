package services

import (
	"github.com/Takahito-Uchino/gin-fleamarket/models"
	"github.com/Takahito-Uchino/gin-fleamarket/repositories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
}

type ItemService struct {
	repository repositories.IItemRepository
}

func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}
