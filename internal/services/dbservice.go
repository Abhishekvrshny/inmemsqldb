package services

import (
	"errors"
	"inmemdb/internal/models/column"
	"inmemdb/internal/models/database"
	"inmemdb/internal/models/record"
	"inmemdb/internal/models/table"
)

type DBService struct {
	dbMap   map[string]*database.Database
	dbInUse *database.Database
}

func NewDBService() *DBService {
	return &DBService{
		dbMap:   make(map[string]*database.Database),
		dbInUse: nil,
	}
}

func (dbs *DBService) CreateDB(name string) error {
	if _, ok := dbs.dbMap[name]; !ok {
		dbs.dbMap[name] = database.New(name)
		return nil
	}
	return errors.New("db already exists")
}

func (dbs *DBService) DropDB(name string) error {
	if _, ok := dbs.dbMap[name]; ok {
		delete(dbs.dbMap, name)
		return nil
	}
	return errors.New("db doesnt exists")
}

func (dbs *DBService) UseDB(name string) error {
	if _, ok := dbs.dbMap[name]; !ok {
		return errors.New("db doesn't exist")
	}
	dbs.dbInUse = dbs.dbMap[name]
	return nil
}

func (dbs *DBService) CreateTable(tab string, col []*column.Column) error {
	tbl := table.New("some.table", col)
	err := dbs.dbInUse.AddTable(tbl)
	if err != nil {
		return err
	}
	return nil
}

func (dbs *DBService) Insert(tab string, columns []string, values []interface{}) error {
	if len(columns) != len(values) {
		return errors.New("number of columns and values doesnt match")
	}
	return dbs.dbInUse.Insert(tab, columns, values)
}

func (dbs *DBService) Print(tab string) {
	dbs.dbInUse.PrintTable(tab)
}

func (dbs *DBService) DeleteTable(name string) error {
	return dbs.dbInUse.DeleteTableWithName(name)
}

func (dbs *DBService) GetRecords(tab string, col string, val interface{}) ([]*record.Record, error) {
	return dbs.dbInUse.GetRecords(tab, col, val)
}

func (dbs *DBService) AddIndexOnColumn(tab string, col string) error {
	return dbs.dbInUse.AddIndexOnColumn(tab, col)
}
