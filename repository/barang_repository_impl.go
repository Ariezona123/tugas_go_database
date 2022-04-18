package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type barangRepositoryImpl struct {
	DB *sql.DB
}

func NewBarangRepository(db *sql.DB) BarangRepository {
	return &barangRepositoryImpl{DB: db}
}

func (repository *barangRepositoryImpl) Insert(ctx context.Context, barang entity.Barang) (entity.Barang, error) {
	script := "INSERT INTO barang(nama_barang, stok_barang) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, barang.Name, barang.Stok)
	if err != nil {
		return barang, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return barang, err
	}
	barang.Id = int32(id)
	return barang, nil
	//panic("implement me")
}

func (repository *barangRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Barang, error) {
	script := "SELECT id, name, stok FROM barang WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	barang := entity.Barang{}
	if err != nil {
	}
	defer rows.Close()
	if rows.Next() {
		//ada
		rows.Scan(&barang.Id, &barang.Name, &barang.Stok)
		return barang, nil
	} else {
		//tidak ada
		return barang, errors.New("Id " + strconv.Itoa(int(id)) + "Not Found")
	}
}

func (repository *barangRepositoryImpl) FindAll(ctx context.Context) ([]entity.Barang, error) {
	script := "SELECT id, nama, stok FROM barang"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var barangs []entity.Barang
	for rows.Next() {
		barang := entity.Barang{}
		rows.Scan(&barang.Id, &barang.Name, &barang.Stok)
		barangs = append(barangs, barang)
	}
	return barangs, nil
}

func (repository *barangRepositoryImpl) Update(ctx context.Context, id int32, barang entity.Barang) (entity.Barang, error) {
	//TODO implement me
	script := "SELECT id, name, stok FROM barang WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return barang, err
	}
	if rows.Next() {
		// yes
		script := "UPDATE barang SET name = ?, stok = ? WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, barang.Name, barang.Stok, id)
		if err != nil {
			return barang, err
		}
		barang.Id = id
		return barang, nil
	} else {
		// no
		return barang, errors.New(("Id " + strconv.Itoa(int(id)) + " Not Found"))
	}
}

func (repository *barangRepositoryImpl) Delete(ctx context.Context, id int32) (string, error) {
	script := "SELECT id, name, stok FROM barang WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return "Gagal", err
	}
	if rows.Next() {

		script := "DELETE FROM barang WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, id)
		if err != nil {
			return "Id" + strconv.Itoa(int(id)) + "Gagal", err
		}
		return "Id" + strconv.Itoa(int(id)) + "Sukses", nil
	} else {

		return "Gagal", errors.New(("Id" + strconv.Itoa(int(id)) + "tidak ada"))
	}
}
