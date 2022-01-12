package models

import (
	"net/http"

	db "github.com/xvbnm48/echo-restfull/db"
)

type Pegawai struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama_pegawai"`
	Alamat  string `json:"alamat"`
	Telepon string `json:"telepon"`
}

func FetchPegawai() (Response, error) {
	var obj Pegawai
	var arrobj []Pegawai
	var res Response

	con := db.CreateCon()

	sqlStatement := `SELECT * FROM pegawai`

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		res.Status = 500
		res.Message = "Internal server error"
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telepon)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj
	return res, nil
}
