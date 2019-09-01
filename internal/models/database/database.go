package database

import (
	"inmemdb/internal/models/record"
	"inmemdb/internal/models/table"

	"github.com/pkg/errors"
)

type Database struct {
	name   string
	tables map[string]*table.Table
}

func New(name string) *Database {
	return &Database{name, make(map[string]*table.Table)}
}

func (db *Database) AddTable(table *table.Table) error {
	if _, ok := db.tables[table.GetName()]; !ok {
		db.tables[table.GetName()] = table
		return nil
	}
	return errors.New("table already exists")
}

func (db *Database) DeleteTable(tab *table.Table) error {
	return db.DeleteTableWithName(tab.GetName())
}

func (db *Database) DeleteTableWithName(name string) error {
	if _, ok := db.tables[name]; ok {
		delete(db.tables, name)
		return nil
	} else {
		return errors.New("table doesnt exist")
	}
}

func (db *Database) GetRecords(tab string, col string, val interface{}) ([]*record.Record, error) {
	if _, ok := db.tables[tab]; !ok {
		return nil, errors.New("table doesnt exist")
	} else {
		return db.tables[tab].GetRecords(col, val)
	}
}

func (db *Database) AddIndexOnColumn(tab string, col string) error {
	if _, ok := db.tables[tab]; !ok {
		return errors.New("table doesnt exist")
	} else {
		return db.tables[tab].AddIndexOnColumn(col)
	}
}

func (db *Database) Insert(tab string, cols []string, values []interface{}) error {
	if _, ok := db.tables[tab]; !ok {
		return errors.New("table doesnt exist")
	} else {
		return db.tables[tab].Insert(cols, values)
	}
}

func (db *Database) PrintTable(tab string) error {
	if _, ok := db.tables[tab]; !ok {
		return errors.New("table doesnt exist")
	} else {
		db.tables[tab].Print()
	}
	return nil
}
