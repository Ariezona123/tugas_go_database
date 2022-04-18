package repository

import (
	"context"
	"go_database/entity"
)

type TinjuRepository interface {
	Insert(ctx context.Context, tinju entity.Tinju) (entity.Tinju, error)
	FindById(ctx context.Context, id int32) (entity.Tinju, error)
	FindAll(ctx context.Context) ([]entity.Tinju, error)
	Update(ctx context.Context, tinju *entity.Tinju) (*entity.Tinju, error)
	Delete(ctx context.Context, Id int32) (bool, error)
}
