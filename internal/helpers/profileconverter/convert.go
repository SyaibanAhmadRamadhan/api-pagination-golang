package profileconverter

import (
	"github.com/SyaibanAhmadRamadhan/api-pagination-golang/src/modules/profile/dto"
	"github.com/SyaibanAhmadRamadhan/api-pagination-golang/src/modules/profile/entities"
)

type ProfileConverterImpl struct {
	ProfileEntity        *entities.Profile
	ProfileCreateRequest *dto.CreateProfileRequest
	ProfileUpdateRequest *dto.UpdateProfileRequest
}

func NewProfileConverterImpl() *ProfileConverterImpl {
	return &ProfileConverterImpl{}
}
