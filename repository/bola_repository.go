package repository

import (
	"context"
	"go_database/entity"
)

type BolaRepository interface {
	Insert(ctx context.Context, bola entity.Bola) (entity.Bola, error)
	FindById(ctx context.Context, Id int32) (entity.Bola, error)
	FindAll(ctx context.Context) ([]entity.Bola, error)
	Update(ctx context.Context, Id int32, bola entity.Bola) (entity.Bola, error)
	Delete(ctx context.Context, Id int32) (string, error)
}
