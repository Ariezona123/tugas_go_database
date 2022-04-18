package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type bukuRepositoryImpl struct {
	DB *sql.DB
}

func NewBukuRepository(db *sql.DB) BukuRepository {
	return &bukuRepositoryImpl{DB: db}
}

func (repository *bukuRepositoryImpl) Insert(ctx context.Context, buku entity.Buku) (entity.Buku, error) {
	script := "INSERT INTO buku(judul,halaman) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, buku.Judul, buku.Halaman)
	if err != nil {
		return buku, err
	}
	id, err := result.LastInsertId()
	if err != nil {
	}
	buku.Id = int32(id)
	return buku, nil
}

func (repository *bukuRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Buku, error) {
	script := "SELECT id, judul, halaman FROM buku WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	buku := entity.Buku{}
	if err != nil {
	}
	defer rows.Close()
	if rows.Next() {
		//ada
		rows.Scan(&buku.Id, &buku.Judul, &buku.Halaman)
		return buku, nil
	} else {
		//tidak ada
		return buku, errors.New("Id " + strconv.Itoa(int(id)) + "Not Found")
	}
}

func (repository *bukuRepositoryImpl) FindAll(ctx context.Context) ([]entity.Buku, error) {
	script := "SELECT id, judul, halaman FROM buku"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var bukus []entity.Buku
	for rows.Next() {
		buku := entity.Buku{}
		rows.Scan(&buku.Id, &buku.Judul, &buku.Halaman)
		bukus = append(bukus, buku)
	}
	return bukus, nil
}

func (repository *bukuRepositoryImpl) Update(ctx context.Context, buku *entity.Buku) (*entity.Buku, error) {
	script := "SELECT buku halaman =?, WHERE Id = ?"
	rows, err := repository.DB.PrepareContext(ctx, script)
	if err != nil {
		return nil, err
	}
	_, err = rows.ExecContext(
		ctx,
		buku.Judul,
		buku.Halaman,
		buku.Id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return buku, nil
}

func (repository *bukuRepositoryImpl) Delete(ctx context.Context, Id int32) (bool, error) {
	script := "Delete buku WHERE id=?"

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

