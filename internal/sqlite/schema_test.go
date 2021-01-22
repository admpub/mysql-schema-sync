package sqlite

import (
	"fmt"
	"testing"

	"github.com/webx-top/com"
)

var schemaData = `CREATE TABLE "forever_process2" (
	"id" integer NOT NULL,
	"name" varchar(60) NOT NULL COLLATE NOCASE,
	"disabled" char(9) NOT NULL DEFAULT 'N' COLLATE NOCASE,
	"updated" integer NOT NULL DEFAULT '0',
	"description" varchar(500) NOT NULL DEFAULT '' COLLATE NOCASE,
	"created" integer NOT NULL,
	"idx" integer,
	"wid" integer,
	PRIMARY KEY ("id", "name"),
	CONSTRAINT "wid" FOREIGN KEY ("wid") REFERENCES "forever_process" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION,
	CONSTRAINT "idx" UNIQUE ("idx" ASC) ON CONFLICT FAIL
  );
` + "CREATE UNIQUE INDEX `UNQ_forever_process3_name` ON `forever_process2`(`name`);\n" + "CREATE UNIQUE INDEX `UNQ_forever_process_name` ON `forever_process2`(`name2`);\n"

var schemaData2 = "CREATE TABLE `nging_db_sync_log` (`id` integer PRIMARY KEY ,`sync_id` int NOT NULL,`created` int NOT NULL,`result` text COLLATE NOCASE NOT NULL,`change_tables` text COLLATE NOCASE NOT NULL,`change_table_num` int NOT NULL DEFAULT '0',`elapsed` bigint NOT NULL DEFAULT '0',`failed` int NOT NULL DEFAULT '0')"

func TestDBSchema(t *testing.T) {
	schema := NewSchemaData(schemaData, `source`)
	fmt.Printf(`tables: %#v`+"\n", schema.GetTableNames())
	fmt.Printf(`schema: %s`+"\n", schema.GetTableSchema(`forever_process2`))
	sc := ParseSchema(schema.GetTableSchema(`forever_process2`))
	com.Dump(sc)
	//panic(``)
}

func TestDBSchema2(t *testing.T) {
	schema := NewSchemaData(schemaData2, `source`)
	fmt.Printf(`tables: %#v`+"\n", schema.GetTableNames())
	fmt.Printf(`schema: %s`+"\n", schema.GetTableSchema(`nging_db_sync_log`))
	sc := ParseSchema(schema.GetTableSchema(`nging_db_sync_log`))
	com.Dump(sc)
	//panic(``)
}
