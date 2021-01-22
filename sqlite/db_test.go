package sqlite

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/admpub/mysql-schema-sync/internal"
	"github.com/webx-top/com"

	_ "github.com/mattn/go-sqlite3"
)

func TestDb(t *testing.T) {
	return
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

func TestSync(t *testing.T) {
	return
	cfg := &internal.Config{}
	cfg.Drop = true
	cfg.Sync = true
	cfg.SetComparer(NewCompare())
	destDB := New(`/Users/hank/go/src/github.com/admpub/nging/dist/localtest/nging_darwin_amd64/nging.db`, `dest`)
	content, err := ioutil.ReadFile(`/Users/hank/Downloads/nging.sqlite.sql`)
	if err != nil {
		panic(err)
	}
	sourceDB := NewSchemaData(string(content), `source`)
	st := internal.CheckSchemaDiff(cfg, sourceDB, destDB)
	ioutil.WriteFile(`./testSync.html`, []byte(st.String()), os.ModePerm)
	//panic(``)
}
