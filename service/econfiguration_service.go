package service

import (
	"github.com/jeffcail/cgncode/models"
	"github.com/jeffcail/cgncode/dto"
)

type EConfigurationService struct{}

func NewEConfigurationService() *EConfigurationService {
    return &EConfigurationService{}
}

// CreateEConfiguration handles the creation logic for EConfiguration
func (this *EConfigurationService) CreateEConfiguration (input *models.EConfiguration) error {
    return dto.NewEConfigurationDto().CreateEConfiguration(input)
}

// GetEConfiguration handles the retrieval logic for EConfiguration
func (this *EConfigurationService) GetEConfiguration(id string) (*models.EConfiguration, error) {
    return dto.NewEConfigurationDto().GetEConfiguration(id)
}

// UpdateEConfiguration handles the update logic for EConfiguration
func (this *EConfigurationService) UpdateEConfiguration(id string, input *models.EConfiguration) error {
    return dto.NewEConfigurationDto().UpdateEConfiguration(id, input)
}

// DeleteEConfiguration handles the deletion logic for EConfiguration
func (this *EConfigurationService) DeleteEConfiguration (id string) error {
    return dto.NewEConfigurationDto().DeleteEConfiguration(id)
}
