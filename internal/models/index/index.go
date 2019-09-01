package index

import (
	"inmemdb/internal/models/record"
)

type Index struct {
	entries map[interface{}]*record.Record
}

func New() *Index {
	return &Index{make(map[interface{}]*record.Record)}
}

func (i *Index) AddRecord(value interface{}, rec *record.Record) {
	i.entries[value] = rec
}

func (i *Index) GetRecords(value interface{}) []*record.Record {
	var recs []*record.Record
	recs = append(recs, i.entries[value])
	return recs
}

func (i *Index) RemoveRecord(value interface{}, rec *record.Record) {
	delete(i.entries, value)
}
