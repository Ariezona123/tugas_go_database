package repository

import (
	"context"
	"go_database/entity"
)

type BukuRepository interface {
	Insert(ctx context.Context, buku entity.Buku) (entity.Buku, error)
	FindById(ctx context.Context, id int32) (entity.Buku, error)
	FindAll(ctx context.Context) ([]entity.Buku, error)
	Update(ctx context.Context, buku *entity.Buku) (*entity.Buku, error)
	Delete(ctx context.Context, Id int32) (bool, error)
}
