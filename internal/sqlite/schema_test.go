package sqlite

import (
	"fmt"
	"testing"

	"github.com/webx-top/com"
)

var schemaData = `CREATE TABLE "forever_process2" (
	"id" integer,
	"name" varchar(60) NOT NULL COLLATE NOCASE,
	"disabled" char(9) NOT NULL DEFAULT 'N' COLLATE NOCASE,
	"updated" integer NOT NULL DEFAULT '0',
	"description" varchar(500) NOT NULL DEFAULT '' COLLATE NOCASE,
	"created" integer NOT NULL,
	"idx" integer,
	"wid" integer,
	PRIMARY KEY ("id"),
	CONSTRAINT "wid" FOREIGN KEY ("wid") REFERENCES "forever_process" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION,
	CONSTRAINT "idx" UNIQUE ("idx") ON CONFLICT FAIL
);
` + "CREATE UNIQUE INDEX `UNQ_forever_process3_name` ON `forever_process3`(`name`)\n" + "CREATE UNIQUE INDEX `UNQ_forever_process_name` ON `forever_process`(`name`)\n"

func TestDBSchema(t *testing.T) {
	schema := NewSchemaData(schemaData, `source`)
	fmt.Printf(`tables: %#v`+"\n", schema.GetTableNames())
	fmt.Printf(`schema: %s`+"\n", schema.GetTableSchema(`forever_process2`))
	sc := ParseSchema(schema.GetTableSchema(`forever_process2`))
	com.Dump(sc)
	//panic(schemaData)
}
