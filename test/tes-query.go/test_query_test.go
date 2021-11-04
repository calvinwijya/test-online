package tesquerygo

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateTable1(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "CREATE TABLE Transaction(id int not null auto_increment,tanggal_order TIMESTAMP not null,status_pelunasan ENUM(`pending`,`lunas`),tanggal pembayaran TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP,primary key(id),unique_key user_unique (id_user);"

	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}
	fmt.Print("success create table")
}

func TestCreateTable2(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "CREATE TABLE Transaction(id_user varchar(10) not null ,id int not null auto_increment,harga int64,jumlah_barang int not null,total int not null,primary key(id),unique_key user_unique (id_user)"

	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}
	fmt.Print("success create table")
}

func TestInsert(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "insert into Transaction(status,jumlah_barang,total) values('lunas','20000','2','20000')"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}
	fmt.Print("success create table")
}
