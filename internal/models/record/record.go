package record

import "inmemdb/internal/models/column"

type Record struct {
	entry map[string]interface{}
}

func New() *Record {
	return &Record{make(map[string]interface{})}
}

func (r *Record) Add(column *column.Column, value interface{}) {
	r.entry[column.GetName()] = value
}

func (r *Record) GetColumnValue(column string) interface{} {
	return r.entry[column]
}
