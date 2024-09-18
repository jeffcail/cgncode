package service

import (
	"github.com/jeffcail/cgncode/dto"
	"github.com/jeffcail/cgncode/models"
)

type InfoService struct{}

func NewInfoService() *InfoService {
	return &InfoService{}
}

// CreateInfo handles the creation logic for Info
func (this *InfoService) CreateInfo(input *models.Info) error {
	return dto.NewInfoDto().CreateInfo(input)
}

// GetInfo handles the retrieval logic for Info
func (this *InfoService) GetInfo(id string) (*models.Info, error) {
	return dto.NewInfoDto().GetInfo(id)
}

// UpdateInfo handles the update logic for Info
func (this *InfoService) UpdateInfo(id string, input *models.Info) error {
	return dto.NewInfoDto().UpdateInfo(id, input)
}

// DeleteInfo handles the deletion logic for Info
func (this *InfoService) DeleteInfo(id string) error {
	return dto.NewInfoDto().DeleteInfo(id)
}
