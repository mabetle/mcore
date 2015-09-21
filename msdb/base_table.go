package msdb

import (
	"github.com/mabetle/mcore"
)

type BaseTable struct {
	*Cusor
	StringGetter
	Header []string
}

func (t *BaseTable) Next() bool {
	return t.Cusor.Next()
}

func (t *BaseTable) GetRowIndex() int {
	return t.Cusor.RowIndex
}

func (t *BaseTable) GetRows() int {
	return t.Cusor.MaxIndex
}

func (t *BaseTable) GetCols() int {
	return len(t.Header)
}

func (t *BaseTable) GetColIndex(colName string) (colIndex int) {
	for i := 0; i < t.GetCols(); i++ {
		if colName == t.Header[i] {
			colIndex = i
		}
	}
	return
}

func (t *BaseTable) GetStringByColName(colName string) (value string) {
	col := t.GetColIndex(colName)
	return t.GetString(col)
}

func (t *BaseTable) GetColNames() []string {
	return t.Header
}

func (t *BaseTable) IsHasColumn(columnName string) bool {
	return mcore.String(columnName).IsInArrayIgnoreCase(t.GetColNames())
}

func (t *BaseTable) Demo() {
	DemoSimpleTable(t)
}

// Random Access
// TODO not implement yet.
func (t *BaseTable) GetRowColString(row int, col int) (result string) {

	return
}
