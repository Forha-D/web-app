package  driver

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

type  MySQLConfig struct {
	Host        string
	User        string
	Password    string
	Port        string
	Db          string
}



func ConnectToMySQL(config  MySQLConfig) (*sql.DB,error) {
	connectionString := fmt.Sprintf("\"%v:%v@tcp(%v:%v)/%v", config.User,config.Password,config.Host,config.Port,config.Db)
	db, err:= sql.Open("mysql" , connectionString)
	if err!= nil{
		return nil, err
	}
	return db, nil

}


