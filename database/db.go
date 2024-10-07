package database

import "database/sql"

type DBDetails struct{
	DBName string
	DBUrl string
}

func NewDbDetails(name string , url string) *DBDetails {
	return &DBDetails{
		DBName: name,
		DBUrl: url,
	}
}

func(d *DBDetails) ConnectToDatabase() (*sql.DB , error){
	db , err := sql.Open(d.DBName , d.DBUrl)
	if err != nil {
		return nil , err
	}
	return db , nil
}