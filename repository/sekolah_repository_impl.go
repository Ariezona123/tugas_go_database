package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type sekolahRepositoryImpl struct {
	DB *sql.DB
}

func NewSekolahRepository(db *sql.DB) SekolahRepository {
	return &sekolahRepositoryImpl{DB: db}
}

func (repository *sekolahRepositoryImpl) Insert(ctx context.Context, sekolah entity.Sekolah) (entity.Sekolah, error) {
	script := "INSERT INTO sekolah(guru,mapel) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, sekolah.Guru, sekolah.Mapel)
	if err != nil {
		return sekolah, err
	}
	id, err := result.LastInsertId()
	if err != nil {
	}
	sekolah.Id = int32(id)
	return sekolah, nil
}

func (repository *sekolahRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Sekolah, error) {
	script := "SELECT id, guru, mapel FROM barang WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	sekolah := entity.Sekolah{}
	if err != nil {
	}
	defer rows.Close()
	if rows.Next() {
		//ada
		rows.Scan(&sekolah.Id, &sekolah.Guru, &sekolah.Mapel)
		return sekolah, nil
	} else {
		//tidak ada
		return sekolah, errors.New("Id " + strconv.Itoa(int(id)) + "Not Found")
	}
}

func (repository *sekolahRepositoryImpl) FindAll(ctx context.Context) ([]entity.Sekolah, error) {
	script := "SELECT id, guru, mapel FROM sekolah"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sekolahan []entity.Sekolah
	for rows.Next() {
		sekolah := entity.Sekolah{}
		rows.Scan(&sekolah.Id, &sekolah.Guru, &sekolah.Mapel)
		sekolahan = append(sekolahan, sekolah)
	}
	return sekolahan, nil
}

func (repository *sekolahRepositoryImpl) Update(ctx context.Context, sekolah *entity.Sekolah) (*entity.Sekolah, error) {
	script := "SELECT sekolah mapel =?, WHERE Id = ?"
	rows, err := repository.DB.PrepareContext(ctx, script)
	if err != nil {
		return nil, err
	}
	_, err = rows.ExecContext(
		ctx,
		sekolah.Guru,
		sekolah.Mapel,
		sekolah.Id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return sekolah, nil
}

func (repository *sekolahRepositoryImpl) Delete(ctx context.Context, Id int32) (bool, error) {
	script := "Delete sekolah WHERE id=?"

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
