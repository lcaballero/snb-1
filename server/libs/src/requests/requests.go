package requests

import (
  "fmt"
	_ "github.com/bmizerany/pq"
	"database/sql"
	"bytes"
)

type DataPoint struct {
	ColumnName string
	FieldNumber int
	Value interface{}
}

type DataResult struct {
	FieldCount int
	Columns []string
	Rows [][]DataPoint
}

type ScanBuffer struct {
	rows *sql.Rows
	count int
	columns []string
	values []interface{}
	pointers []interface{}
}

func ToMapping(rows *sql.Rows) []map[string]interface{} {
	sb, err := ToScanBuffer(rows)
	if err != nil {
		fmt.Println("ToMapping(rows) : ", err)
		return make([]map[string]interface{}, 0)
	}
	output := sb.ToDataResult()
	mapping := output.ToMapping()
	return mapping
}

func (r *DataResult) ToMapping() []map[string]interface{} {

	m := make([]map[string]interface{}, 0)

	for _, row := range r.Rows {

		fields := make(map[string]interface{})
		m = append(m, fields)

		for _, dp := range row {
			fields[dp.ColumnName] = dp.Value
		}
	}

	return m
}



func (r *DataResult) String() string {

	if (r == nil) {
		return "<nil>" // Special case, useful for debugging
	}

	var buf bytes.Buffer

	for i, dps := range r.Rows {
		for _, dp := range dps {
			s := fmt.Sprintf("Row[%v].%v = %v\n", i, dp.ColumnName, dp.Value)
			buf.Write([]byte(s))
		}	
	}

	return buf.String()
}


func ToScanBuffer(rows *sql.Rows) (sb *ScanBuffer, err error) {
	
	sb = &ScanBuffer{}

	sb.rows = rows
	sb.columns, err = rows.Columns();
	sb.count = len(sb.columns)
	sb.values = make([]interface{}, sb.count)
	sb.pointers = make([]interface{}, sb.count)

	for i := 0; i < sb.count; i++ {
		sb.pointers[i] = &sb.values[i]
	}

	return sb, err
}

func (sb *ScanBuffer) ToDataResult() (r *DataResult) {

	r = &DataResult{ sb.count, sb.columns, make([][]DataPoint, 0, 20) }

	for sb.rows.Next() {

		sb.rows.Scan(sb.pointers...)

		dps := sb.toDataPoints()

		r.Rows = append(r.Rows, dps)
	}

	return r
}


func (sb *ScanBuffer) toDataPoints() (dps []DataPoint) {
	
	dps = make([]DataPoint, sb.count)

	for i, col := range sb.columns {

		var v interface{}

		val := sb.values[i]

		b, ok := val.([]byte)

		if (ok) {
			v = string(b)
		} else {
			v = val
		}

		dp := DataPoint{ ColumnName: col, FieldNumber: i, Value: v }

		dps[i] =  dp
	}

	return dps
}
