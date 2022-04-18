package repository

import (
	"context"
	"go_database/entity"
)

type SekolahRepository interface {
	Insert(ctx context.Context, sekolah entity.Sekolah) (entity.Sekolah, error)
	FindById(ctx context.Context, id int32) (entity.Sekolah, error)
	FindAll(ctx context.Context) ([]entity.Sekolah, error)
	Update(ctx context.Context, sekolah *entity.Sekolah) (*entity.Sekolah, error)
	Delete(ctx context.Context, Id int32) (bool, error)
}
