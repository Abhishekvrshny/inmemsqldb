package table

import (
	"fmt"
	"inmemdb/internal/models/column"
	"inmemdb/internal/models/field"
	"inmemdb/internal/models/index"
	"inmemdb/internal/models/record"

	"github.com/pkg/errors"
)

type Table struct {
	name    string
	cols    []*column.Column
	records []*record.Record
	indexes map[string]*index.Index
}

func New(name string, cols []*column.Column) *Table {
	return &Table{name, cols, nil, make(map[string]*index.Index)}
}

func (t *Table) GetName() string {
	return t.name
}

func (t *Table) addRecord(rec *record.Record) {
	t.records = append(t.records, rec)
	t.updateIndex(rec)
}

func (t *Table) AddIndexOnColumn(col string) error {
	ind := index.New()
	t.indexes[col] = ind
	for _, rec := range t.records {
		t.updateIndex(rec)
	}
	return nil
}

func (t *Table) Print() {
	for _, rec := range t.records {
		fmt.Printf("%+v\n", rec)
	}
}

func (t *Table) updateIndex(rec *record.Record) {
	for col, idx := range t.indexes {
		idx.AddRecord(rec.GetColumnValue(col), rec)
	}
}

func (t *Table) IndexedGet(column string, value interface{}) ([]*record.Record, error) {
	if _, ok := t.indexes[column]; !ok {
		return nil, errors.New("column not indexed")
	} else {
		return t.indexes[column].GetRecords(value), nil
	}
}

func (t *Table) GetRecords(col string, val interface{}) ([]*record.Record, error) {
	if _, ok := t.indexes[col]; !ok {
		return t.scanGet(col, val)
	} else {
		return t.IndexedGet(col, val)
	}
}

func (t *Table) scanGet(column string, value interface{}) ([]*record.Record, error) {
	//TODO: to be implemented
	return nil, nil
}

func (t *Table) Insert(cols []string, values []interface{}) error {
	cols, values = t.getRecordTuple(cols, values)
	for i := 0; i < len(cols); i++ {
		col, err := t.checkAndGetColumn(cols[i])
		if err != nil {
			return err
		}
		err = field.Validate(col.GetType(), values[i])
		if err != nil {
			return err
		}
		err = col.Validate(values[i])
		if err != nil {
			return fmt.Errorf("contraint error on %s : %s", col.GetName(), err.Error())
		}
	}
	rec := record.New()
	for i := 0; i < len(cols); i++ {
		col, _ := t.checkAndGetColumn(cols[i])
		rec.Add(col, values[i])
	}
	t.addRecord(rec)
	return nil
}

func (t *Table) checkAndGetColumn(col string) (*column.Column, error) {
	for _, c := range t.cols {
		if c.GetName() == col {
			return c, nil
		}
	}
	return nil, errors.New("column not found")
}

func (t *Table) getRecordTuple(cols []string, values []interface{}) ([]string, []interface{}) {
	var c []string
	var v []interface{}
	for _, cc := range t.cols {
		flag := 0
		for i := 0; i < len(cols); i++ {
			if cc.GetName() == cols[i] {
				c = append(c, cols[i])
				v = append(v, values[i])
				flag = 1
				break
			}
		}
		if flag == 0 {
			c = append(c, cc.GetName())
			v = append(v, nil)
		}
	}
	return c, v
}
