package services

import (
	"context"
	"database/sql"

	"github.com/SyaibanAhmadRamadhan/api-pagination-golang/internal/helpers/pagination/models"
	"github.com/SyaibanAhmadRamadhan/api-pagination-golang/src/modules/profile/dto"
	satoriuuid "github.com/satori/go.uuid"
)

func (svc *ProfileServiceImpl) Insert(ctx context.Context, req *dto.CreateProfileRequest) (*dto.ProfileResponse, error) {
	err := svc.Transaction.RunTransaction(ctx, &sql.TxOptions{ReadOnly: false}, func(tx *sql.Tx) error {
		svc.ProfileCVT.ProfileCreateRequest = req
		uuid := satoriuuid.NewV4().String()

		profileCreate, err := svc.ProfileRepository.Insert(ctx, tx, *svc.ProfileCVT.CreateRequestToEntity(uuid))
		if err != nil {
			return err
		}
		svc.ProfileCVT.ProfileEntity = profileCreate

		return nil
	})
	if err != nil {
		return nil, err
	}

	return svc.ProfileCVT.EntityToRespon(), nil
}

func (svc *ProfileServiceImpl) Update(ctx context.Context, id string, req *dto.UpdateProfileRequest) (*dto.ProfileResponse, error) {
	err := svc.Transaction.RunTransaction(ctx, &sql.TxOptions{ReadOnly: false}, func(tx *sql.Tx) error {
		profile, err := svc.ProfileRepository.FindByID(ctx, svc.DB, id)
		if err != nil {
			return err
		}

		svc.ProfileCVT.ProfileUpdateRequest = req
		svc.ProfileCVT.ProfileEntity = profile

		profileUpdate, err := svc.ProfileRepository.Update(ctx, tx, *svc.ProfileCVT.UpdateRequestToEntity(profile.ID))
		if err != nil {
			return err
		}

		svc.ProfileCVT.ProfileEntity = profileUpdate

		return nil
	})
	if err != nil {
		return nil, err
	}
	return svc.ProfileCVT.EntityToRespon(), nil
}

func (svc *ProfileServiceImpl) Delete(ctx context.Context, id string) error {
	err := svc.Transaction.RunTransaction(ctx, &sql.TxOptions{ReadOnly: false}, func(tx *sql.Tx) error {
		_, err := svc.ProfileRepository.FindByID(ctx, svc.DB, id)
		if err != nil {
			return err
		}

		err = svc.ProfileRepository.Delete(ctx, tx, id)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (svc *ProfileServiceImpl) FindByID(ctx context.Context, id string) (*dto.ProfileResponse, error) {
	profile, err := svc.ProfileRepository.FindByID(ctx, svc.DB, id)
	if err != nil {
		return nil, err
	}

	svc.ProfileCVT.ProfileEntity = profile
	return svc.ProfileCVT.EntityToRespon(), nil
}

func (svc *ProfileServiceImpl) FindAll(ctx context.Context, page int, sort, filter string) (*[]dto.ProfileResponse, *models.Pagination, error) {
	limit := 2
	offset := limit * (page - 1)
	recordCount, err := svc.ProfileRepository.CountData(ctx, svc.DB)
	if err != nil {
		return nil, nil, err
	}
	pagination := svc.Pagination.CalculatePagination(limit, page, recordCount)

	profiles, err := svc.ProfileRepository.FindAll(ctx, svc.DB, limit, offset, sort, filter)
	if err != nil {
		return nil, nil, err
	}

	var profileResponses []dto.ProfileResponse
	for _, profile := range *profiles {
		svc.ProfileCVT.ProfileEntity = &profile
		profileResponse := svc.ProfileCVT.EntityToRespon()
		profileResponses = append(profileResponses, *profileResponse)
	}

	return &profileResponses, pagination, nil
}
