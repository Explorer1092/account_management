package database

import (
	."common/logs"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"accout_management/conf"
	_ "github.com/go-sql-driver/mysql"
)

var pbDb *sql.DB

func init() {

	err, db := OpenDB("mysql", conf.MyDbUser, conf.MyDbPwd, conf.MyDbHost, conf.MyDbPort, conf.MyDbName)
	if err != nil {
		fmt.Println(err)
	}
	pbDb = db
}

func OpenDB(driverName, usrName, psswd, hostName, port, dbName string) (error, *sql.DB) {
	db_pbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Asia%%2FShanghai&timeout=30s&readTimeout=35s&writeTimeout=35s", usrName, psswd, hostName, port, dbName)

	db, err := sql.Open(driverName, db_pbLink)
	if err != nil {
		Error("sql open failed %s", err.Error())
		return err, nil
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(30)
	db.SetConnMaxLifetime(120 * time.Second)

	return err, db
}

func CloseDB() error {
	if pbDb == nil {
		Error("CloseDB: db is null")
		return errors.New("db is null")
	}

	return pbDb.Close()
}

func GetDb() *sql.DB {
	return pbDb
}

func dbExec(cmd string, columnList []string, keyName string, args ...interface{}) (result sql.Result, err error) {
	if pbDb == nil {
		Error("dbExec pbDb is null ")
		return nil, errors.New("not open db")
	}

	jobCmd := cmd + " SET "
	for idx := 0; idx < len(columnList); idx++ {
		if idx == 0 {
			jobCmd = fmt.Sprintf("%s %s=?", jobCmd, columnList[idx])
		} else {
			jobCmd = fmt.Sprintf("%s,%s=?", jobCmd, columnList[idx])
		}
	}

	if keyName != "" {
		jobCmd = fmt.Sprintf("%s WHERE %s=?", jobCmd, keyName)
	}

	stmt, err := pbDb.Prepare(jobCmd)
	if err != nil {
		Error("dbExec sql Prepare error: ", err)
		return nil, err
	}
	defer stmt.Close()

	result, err = stmt.Exec(args...)
	if err != nil {
		Error("dbExec sql stmt.Exec() failed: ", err)
		return result, err
	}

	return result, nil
}

func Insert(dbtName string, colNameList []string, args ...interface{}) (result sql.Result, err error) {
	Info(dbtName, colNameList, args)
	if pbDb == nil {
		Error("Insert pbDb is null ")
		return nil, errors.New("not open db")
	}

	result, err = dbExec("INSERT "+dbtName, colNameList, "", args...)
	if err != nil {
		Error("Insert error:  ", err)
		return
	}

	return
}

func Update(dbtName string, colNameList []string, keyName string, args ...interface{}) error {

	if pbDb == nil {
		Error("Update pbDb is null ")
		return errors.New("not open db")
	}

	_, err := dbExec("UPDATE "+dbtName, colNameList, keyName, args...)
	if err != nil {
		Error("Update error: ", err)
		return err
	}

	return nil
}

func QueryCount(dbtName string, whereKey []string, whereVal []interface{}) (count int, err error) {

	if pbDb == nil {
		Error("Query pbDb is null ")
		return 0, errors.New("not open db")
	}

	queryCmd := fmt.Sprintf("SELECT count(*) FROM %s ", dbtName)
	for index, _ := range whereKey {
		if index == 0 {
			if _, ok := whereVal[index].(string); ok {
				queryCmd = fmt.Sprintf("%s WHERE %s='%s'", queryCmd, whereKey[index], whereVal[index])
			} else {
				queryCmd = fmt.Sprintf("%s WHERE %s=%v", queryCmd, whereKey[index], whereVal[index])
			}
			continue
		}
		if _, ok := whereVal[index].(string); ok {
			queryCmd = fmt.Sprintf("%s and %s='%s'", queryCmd, whereKey[index], whereVal[index])
		} else {
			queryCmd = fmt.Sprintf("%s and %s=%v", queryCmd, whereKey[index], whereVal[index])
		}
	}
	queryCmd = fmt.Sprintf("%s", queryCmd)
	Info("Query qureyCmd=[", queryCmd, "].")
	err = pbDb.QueryRow(queryCmd).Scan(&count)
	Info("Query success.")
	if err != nil {
		Error("Query failed: ", err.Error())
	}
	return
}

func Query(queryKeys []string, dbtName string, whereKey []string, whereVal []interface{}, orderBySentence string, limitSentence string) (rows *sql.Rows, err error) {

	if pbDb == nil {
		Error("Query pbDb is null ")
		return nil, errors.New("not open db")
	}
	queryFieldStr := "*"

	for idx, key := range queryKeys {
		if idx == 0 {
			queryFieldStr = key
		} else {
			queryFieldStr = fmt.Sprintf("%s,%s", queryFieldStr, key)
		}
	}

	queryCmd := fmt.Sprintf("SELECT %s FROM %s ", queryFieldStr, dbtName)
	for index, _ := range whereKey {
		if index == 0 {
			if _, ok := whereVal[index].(string); ok {
				queryCmd = fmt.Sprintf("%s WHERE %s='%s'", queryCmd, whereKey[index], whereVal[index])
			} else {
				queryCmd = fmt.Sprintf("%s WHERE %s=%v", queryCmd, whereKey[index], whereVal[index])
			}
			continue
		}
		if _, ok := whereVal[index].(string); ok {
			queryCmd = fmt.Sprintf("%s and %s='%s'", queryCmd, whereKey[index], whereVal[index])
		} else {
			queryCmd = fmt.Sprintf("%s and %s=%v", queryCmd, whereKey[index], whereVal[index])
		}
	}
	queryCmd = fmt.Sprintf("%s %s %s", queryCmd, orderBySentence, limitSentence)
	Info("Query qureyCmd=[", queryCmd, "].")
	rows, err = pbDb.Query(queryCmd)
	Info("Query success.")
	if err != nil {
		Error("Query failed: ", err.Error())
	}
	return rows, err

}
