package main

import (
	"fmt"
	"inmemdb/internal/handler"
	"inmemdb/internal/models/column"
	"inmemdb/internal/models/constraint"
	"inmemdb/internal/models/field"
)

func main() {
	// get a db instance handler
	dbHandler := handler.New()

	// create a database
	err := dbHandler.CreateDB("some.db")
	if err != nil {
		fmt.Println(err.Error())
	}

	// drop a database
	err = dbHandler.DropDB("some-non-existent.db")
	if err != nil {
		fmt.Println(err.Error())
	}

	// use a database
	err = dbHandler.UseDB("some.db")
	if err != nil {
		fmt.Println(err.Error())
	}

	// create a new column with constraints
	col := column.New("some.string.column", field.STRING, []constraint.ConstraintType{constraint.NOTNULL})

	// create a new table with the column
	err = dbHandler.CreateTable("some.table", []*column.Column{col})
	if err != nil {
		fmt.Println(err.Error())
	}

	// insert into table, NOTNULL constraint should break here
	err = dbHandler.Insert("some.table", []string{"some.string.column"}, []interface{}{nil})
	if err != nil {
		fmt.Println(err.Error())
	}

	// insert values in table
	err = dbHandler.Insert("some.table", []string{"some.string.column"}, []interface{}{"some.value"})
	if err != nil {
		fmt.Println(err.Error())
	}

	// print table
	dbHandler.Print("some.table")

	// add index on column
	dbHandler.AddIndexOnColumn("some.table", "some.string.column")

	// get records
	recs, err := dbHandler.GetRecords("some.table", "some.string.column", "some.value")
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, rec := range recs {
		fmt.Println(rec)
	}

	// delete table
	err = dbHandler.DeleteTable("some-non-existent.table")
	if err != nil {
		fmt.Println(err.Error())
	}
}
