package profileconverter

import (
	"time"

	"github.com/SyaibanAhmadRamadhan/api-pagination-golang/src/modules/profile/entities"
)

func (cvt *ProfileConverterImpl) CreateRequestToEntity(uuid string) *entities.Profile {
	return &entities.Profile{
		ID:        uuid,
		Name:      cvt.ProfileCreateRequest.Name,
		Gender:    cvt.ProfileCreateRequest.Gender,
		Phone:     cvt.ProfileCreateRequest.Phone,
		Email:     cvt.ProfileCreateRequest.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (cvt *ProfileConverterImpl) UpdateRequestToEntity(id string) *entities.Profile {
	return &entities.Profile{
		ID:        id,
		Name:      cvt.ProfileUpdateRequest.Name,
		Gender:    cvt.ProfileUpdateRequest.Gender,
		Phone:     cvt.ProfileUpdateRequest.Phone,
		Email:     cvt.ProfileUpdateRequest.Email,
		CreatedAt: cvt.ProfileEntity.CreatedAt,
		UpdatedAt: time.Now(),
	}
}
