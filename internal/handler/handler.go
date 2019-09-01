package handler

import (
	"errors"
	"inmemdb/internal/models/column"
	"inmemdb/internal/models/database"
	"inmemdb/internal/models/record"
	"inmemdb/internal/models/table"
)

type DBInstanceHandler struct {
	dbMap   map[string]*database.Database
	dbInUse *database.Database
}

func New() *DBInstanceHandler {
	return &DBInstanceHandler{
		dbMap:   make(map[string]*database.Database),
		dbInUse: nil,
	}
}

func (dbih *DBInstanceHandler) CreateDB(name string) error {
	if _, ok := dbih.dbMap[name]; !ok {
		dbih.dbMap[name] = database.New(name)
		return nil
	}
	return errors.New("db already exists")
}

func (dbih *DBInstanceHandler) DropDB(name string) error {
	if _, ok := dbih.dbMap[name]; ok {
		delete(dbih.dbMap, name)
		return nil
	}
	return errors.New("db doesnt exists")
}

func (dbih *DBInstanceHandler) UseDB(name string) error {
	if _, ok := dbih.dbMap[name]; !ok {
		return errors.New("db doesn't exist")
	}
	dbih.dbInUse = dbih.dbMap[name]
	return nil
}

func (dbih *DBInstanceHandler) CreateTable(tab string, col []*column.Column) error {
	tbl := table.New("some.table", col)
	err := dbih.dbInUse.AddTable(tbl)
	if err != nil {
		return err
	}
	return nil
}

func (dbih *DBInstanceHandler) Insert(tab string, columns []string, values []interface{}) error {
	if len(columns) != len(values) {
		return errors.New("number of columns and values doesnt match")
	}
	return dbih.dbInUse.Insert(tab, columns, values)
}

func (dbih *DBInstanceHandler) Print(tab string) {
	dbih.dbInUse.PrintTable(tab)
}

func (dbih *DBInstanceHandler) DeleteTable(name string) error {
	return dbih.dbInUse.DeleteTableWithName(name)
}

func (dbih *DBInstanceHandler) GetRecords(tab string, col string, val interface{}) ([]*record.Record, error) {
	return dbih.dbInUse.GetRecords(tab, col, val)
}

func (dbih *DBInstanceHandler) AddIndexOnColumn(tab string, col string) error {
	return dbih.dbInUse.AddIndexOnColumn(tab, col)
}
