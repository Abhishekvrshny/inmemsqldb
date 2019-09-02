package main

import (
	"fmt"
	"inmemdb/internal/services"
	"inmemdb/internal/models/column"
	"inmemdb/internal/models/constraint"
	"inmemdb/internal/models/field"
)

func main() {
	// get an instance of db service
	dbs := services.NewDBService()

	// create a database
	err := dbs.CreateDB("some.db")
	if err != nil {
		fmt.Println(err.Error())
	}

	// drop a database
	err = dbs.DropDB("some-non-existent.db")
	if err != nil {
		fmt.Println(err.Error())
	}

	// use a database
	err = dbs.UseDB("some.db")
	if err != nil {
		fmt.Println(err.Error())
	}

	// create a new column with constraints
	col := column.New("some.string.column", field.STRING, []constraint.Type{constraint.NOTNULL})

	// create a new table with the column
	err = dbs.CreateTable("some.table", []*column.Column{col})
	if err != nil {
		fmt.Println(err.Error())
	}

	// insert into table, NOTNULL constraint should break here
	err = dbs.Insert("some.table", []string{"some.string.column"}, []interface{}{nil})
	if err != nil {
		fmt.Println(err.Error())
	}

	// insert values in table
	err = dbs.Insert("some.table", []string{"some.string.column"}, []interface{}{"some.value"})
	if err != nil {
		fmt.Println(err.Error())
	}

	// print table
	dbs.Print("some.table")

	// add index on column
	dbs.AddIndexOnColumn("some.table", "some.string.column")

	// get records
	recs, err := dbs.GetRecords("some.table", "some.string.column", "some.value")
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, rec := range recs {
		fmt.Println(rec)
	}

	// delete table
	err = dbs.DeleteTable("some-non-existent.table")
	if err != nil {
		fmt.Println(err.Error())
	}
}
