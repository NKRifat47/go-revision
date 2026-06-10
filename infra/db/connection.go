package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func GetConnectionString() string{
	return "user=postgres password=admin123 host=localhost port=5432 dbname=ecommerce"
}

func NewConnection() (*sqlx.DB, error){
	dbSource := GetConnectionString()
	dbcon, err:= sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dbcon, nil
}