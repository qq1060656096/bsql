# bsql
business sql

```go
package bsql_test

import (
  "fmt"
  "github.com/qq1060656096/bsql"
)

sb := bsql.NewSQLBuilder()
sb.WriteString("select * from users where 1 = 1")
sb.WriteIf(true, " and uid = ?", "v1")
sb.Write(" and name = ?", "v2")
sb.AddValues("v3", "v4")
sql := sb.SQL()
values := sb.Values()

fmt.Println(sql)
fmt.Println(values)
// Output: select * from users where 1 = 1 and uid = ? and name = ?
// [v1 v2 v3 v4]
```