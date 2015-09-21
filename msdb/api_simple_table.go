package msdb

// SimpleTable
//
type SimpleTable interface {
	GetColNames() []string
	GetRows() int
	GetRowIndex() int
	GetCols() int
	Next() bool
	StringGetter
	GetStringByColName(colName string) string
	IsHasColumn(columName string) bool
}

type StringGetter interface {
	GetString(col int) string //default by col index
}
