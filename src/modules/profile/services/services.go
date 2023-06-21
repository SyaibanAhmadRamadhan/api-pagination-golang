package services

import (
	"context"
	"database/sql"

	"github.com/SyaibanAhmadRamadhan/api-pagination-golang/infrastructure/db/transaction"
	"github.com/SyaibanAhmadRamadhan/api-pagination-golang/internal/helpers/pagination"
	"github.com/SyaibanAhmadRamadhan/api-pagination-golang/internal/helpers/pagination/models"
	"github.com/SyaibanAhmadRamadhan/api-pagination-golang/internal/helpers/profileconverter"
	"github.com/SyaibanAhmadRamadhan/api-pagination-golang/src/modules/profile/dto"
	profilerepo "github.com/SyaibanAhmadRamadhan/api-pagination-golang/src/modules/profile/repository"
)

type ProfileService interface {
	Insert(ctx context.Context, req *dto.CreateProfileRequest) (*dto.ProfileResponse, error)
	Update(ctx context.Context, id string, req *dto.UpdateProfileRequest) (*dto.ProfileResponse, error)
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*dto.ProfileResponse, error)
	FindAll(ctx context.Context, page int, sort, filter string) (*[]dto.ProfileResponse, *models.Pagination, error)
}

type ProfileServiceImpl struct {
	DB                *sql.DB
	Transaction       *transaction.TransactionImpl
	ProfileRepository profilerepo.ProfileRepository
	ProfileCVT        *profileconverter.ProfileConverterImpl
	Pagination        *pagination.PaginationImpl
}

func NewProfileServiceImpl(
	db *sql.DB,
	transaction *transaction.TransactionImpl,
	profileRepository profilerepo.ProfileRepository,
	profileCVT *profileconverter.ProfileConverterImpl,
	pagination *pagination.PaginationImpl,
) ProfileService {
	return &ProfileServiceImpl{
		DB:                db,
		Transaction:       transaction,
		ProfileRepository: profileRepository,
		ProfileCVT:        profileCVT,
		Pagination:        pagination,
	}
}
