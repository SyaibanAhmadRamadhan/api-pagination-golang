package profileconverter

import (
	"time"

	"github.com/SyaibanAhmadRamadhan/technical-test-pt-zahir-international/src/modules/profile/dto"
)

func (cvt *ProfileConverterImpl) EntityToRespon() *dto.ProfileResponse {
	return &dto.ProfileResponse{
		ID:        cvt.ProfileEntity.ID,
		Name:      cvt.ProfileEntity.Name,
		Gender:    cvt.ProfileEntity.Gender,
		Phone:     cvt.ProfileEntity.Phone,
		Email:     cvt.ProfileEntity.Email,
		CreatedAt: cvt.ProfileEntity.CreatedAt.Local().Format(time.RFC3339),
		UpdatedAt: cvt.ProfileEntity.UpdatedAt.Local().Format(time.RFC3339),
	}
}
