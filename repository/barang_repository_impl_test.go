package repository

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestBarangInsert(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())

	ctx := context.Background()
	barang := entity.Barang{
		Name: "Meja",
		Stok: 15,
	}
	result, err := barangRepository.Insert(ctx, barang)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestBarangFindById(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())

	barang, err := barangRepository.FindById(context.Background(), 25)
	if err != nil {
		panic(err)
	}

	fmt.Println(barang)
}

func TestBarangFindAll(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())

	barangs, err := barangRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, barang := range barangs {
		fmt.Println(barang)
	}
}

func TestUpdate(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())

	ctx := context.Background()
	barang := entity.Barang{
		Name: "meja",
		Stok: 15,
	}

	result, err := barangRepository.Update(ctx, 1, barang)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())

	ctx := context.Background()

	result, err := barangRepository.Delete(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
