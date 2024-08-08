package bsql_test

import (
	"github.com/qq1060656096/bsql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSqlBuilder(t *testing.T) {
	sb := bsql.NewSQLBuilder()
	sb.WriteString("select * from users where 1 = 1")
	sb.WriteIf(true, " and uid = ?", 20)
	sb.WriteIf(true, " and is_deleted = ?", 0)
	sb.WriteIf(false, " and name like ?", "%user%")
	sb.WriteIf(true, " and account_name = ?", "test")
	sb.WriteIf(true, " and (test1 = ? or test2 = ?)", "ifMultiV1", "ifMultiV2")
	sb.Write(" and (write = ? or write = ?)", "writeMulti1", "writeMulti2")
	sb.AddValues("rawValue1", "rawValue2")
	sql := `select * from users where 1 = 1 and uid = ? and is_deleted = ? and account_name = ? and (test1 = ? or test2 = ?) and (write = ? or write = ?)`
	values := []interface{}{
		20,
		0,
		"test",
		"ifMultiV1",
		"ifMultiV2",
		"writeMulti1",
		"writeMulti2",
		"rawValue1",
		"rawValue2",
	}
	assert.Equal(t, sql, sb.SQL())
	assert.Equal(t, values, sb.Values())

	sb.Reset()
	assert.Equal(t, "", sb.SQL())
	assert.Equal(t, 0, len(sb.Values()))

	sb.WriteString("select * from users where 1 = 1")
	sb.WriteIf(true, " and uid = ?", 20)
	sb.WriteIf(true, " and is_deleted = ?", 0)
	sb.WriteIf(false, " and name like ?", "%user%")
	sb.WriteIf(true, " and account_name = ?", "test")
	sb.WriteIf(true, " and (test1 = ? or test2 = ?)", "ifMultiV1", "ifMultiV2")
	sb.Write(" and (write = ? or write = ?)", "writeMulti1", "writeMulti2")
	sb.AddValues("rawValue1", "rawValue2")
	assert.Equal(t, sql, sb.SQL())
	assert.Equal(t, values, sb.Values())
}
