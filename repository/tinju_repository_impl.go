package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type tinjuRepositoryImpl struct {
	DB *sql.DB
}

func NewTinjuRepository(db *sql.DB) TinjuRepository {
	return &tinjuRepositoryImpl{DB: db}
}

func (repository *tinjuRepositoryImpl) Insert(ctx context.Context, tinju entity.Tinju) (entity.Tinju, error) {
	script := "INSERT INTO tinju(pemain,wasit) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, tinju.Pemain, tinju.Wasit)
	if err != nil {
		return tinju, err
	}
	id, err := result.LastInsertId()
	if err != nil {
	}
	tinju.Id = int32(id)
	return tinju, nil
}

func (repository *tinjuRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Tinju, error) {
	script := "SELECT id, pemain, wasit FROM tinju WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	tinju := entity.Tinju{}
	if err != nil {
	}
	defer rows.Close()
	if rows.Next() {
		//ada
		rows.Scan(&tinju.Id, &tinju.Pemain, &tinju.Wasit)
		return tinju, nil
	} else {
		//tidak ada
		return tinju, errors.New("Id " + strconv.Itoa(int(id)) + "Not Found")
	}
}

func (repository *tinjuRepositoryImpl) FindAll(ctx context.Context) ([]entity.Tinju, error) {
	script := "SELECT id, pemain, wasit FROM tinju"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tinjus []entity.Tinju
	for rows.Next() {
		tinju := entity.Tinju{}
		rows.Scan(&tinju.Id, &tinju.Pemain, &tinju.Wasit)
		tinjus = append(tinjus, tinju)
	}
	return tinjus, nil
}

func (repository *tinjuRepositoryImpl) Update(ctx context.Context, tinju *entity.Tinju) (*entity.Tinju, error) {
	script := "SELECT tinju wasit =?, WHERE Id = ?"
	rows, err := repository.DB.PrepareContext(ctx, script)
	if err != nil {
		return nil, err
	}
	_, err = rows.ExecContext(
		ctx,
		tinju.Pemain,
		tinju.Wasit,
		tinju.Id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return tinju, nil
}

func (repository *tinjuRepositoryImpl) Delete(ctx context.Context, Id int32) (bool, error) {
	script := "Delete tinju WHERE id=?"

	rows, err := repository.DB.PrepareContext(ctx, script)
	if err != nil {
		return false, err
	}
	_, err = rows.ExecContext(ctx, Id)
	if err != nil {
		return false, err
	}
	return true, nil

}
