package repository

import (
	"bioskop/structs"
	"database/sql"
	"fmt"
)

func GetAllBioskop(db *sql.DB) (result []structs.Bioskop, err error) {
    sql := "SELECT id, nama, lokasi, rating FROM bioskop"

    rows, err := db.Query(sql)
    if err != nil {
       return
    }

    defer rows.Close()
    for rows.Next() {
       var bioskop structs.Bioskop

       err = rows.Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating)
       if err != nil {
          return
       }

       result = append(result, bioskop)
    }

    return
}

func GetBioskop(db *sql.DB, bioskop structs.Bioskop) (result structs.Bioskop, err error) {
    sql := "SELECT id, nama, lokasi, rating  FROM bioskop WHERE id = $1"

    rows, err := db.Query(sql,  bioskop.ID)
    if err != nil {
        return
    }
    defer rows.Close()

    if rows.Next() {
       err = rows.Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating)
        if err != nil {
            return
        }
    } else {
        return result, fmt.Errorf("bioskop dengan ID %d tidak ditemukan", bioskop.ID)
    }

    return result, nil
}

func InsertBioskop(db *sql.DB, bioskop structs.Bioskop) (err error) {
    sql := "INSERT INTO bioskop(nama, lokasi, rating) VALUES ($1, $2, $3)"

    errs := db.QueryRow(sql, bioskop.Nama, bioskop.Lokasi, bioskop.Rating)

    return errs.Err()
}

func UpdateBioskop(db *sql.DB, bioskop structs.Bioskop) (err error) {
    sql := "UPDATE bioskop SET nama = $1, lokasi = $2, rating = $4 WHERE id = $3"

    errs := db.QueryRow(sql, bioskop.Nama, bioskop.Lokasi, bioskop.Rating, bioskop.ID)

    return errs.Err()
}

func DeleteBioskop(db *sql.DB, bioskop structs.Bioskop) (err error) {
    sql := "DELETE FROM bioskop WHERE id = $1"

    errs := db.QueryRow(sql, bioskop.ID)
    return errs.Err()
}