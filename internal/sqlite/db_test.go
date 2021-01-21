package sqlite

import (
	"fmt"
	"testing"

	"github.com/webx-top/com"

	_ "github.com/mattn/go-sqlite3"
)

func TestDb(t *testing.T) {
	c := New(`/Users/hank/Downloads/nging_hk.db`, ``)
	tables := c.GetTableNames()
	com.Dump(tables)
	//return
	for _, table := range tables {
		schema := c.GetTableSchema(table)
		fmt.Println(schema)
		sc := ParseSchema(schema)
		com.Dump(sc)
		break
	}
	//panic(``)
}
