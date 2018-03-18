package sync

import (
	"fmt"
	"log"
	"strings"

	"github.com/admpub/mysql-schema-sync/internal"
)

func Sync(c *Config, mc *EmailConfig) (sta *internal.Statics, err error) {
	cfg := new(internal.Config)
	cfg.SourceDSN = c.SourceDSN
	cfg.DestDSN = c.DestDSN
	cfg.Sync = c.Sync
	cfg.Drop = c.Drop
	c.AlterIgnore = strings.TrimSpace(c.AlterIgnore)
	if len(c.AlterIgnore) > 0 {
		to := &map[string]*internal.AlterIgnoreTable{}
		err = internal.ParseJSON(`{`+strings.TrimRight(c.AlterIgnore, `,`)+`}`, &to)
		if err != nil {
			return
		}
		cfg.AlterIgnore = *to
	}

	if mc != nil {
		cfg.Email = &internal.EmailStruct{
			SendMailAble: mc.On,
			SMTPHost:     mc.SMTPHost,
			From:         mc.From,
			Password:     mc.Password,
			To:           mc.To,
		}
	}

	if len(c.MailTo) > 0 && cfg.Email != nil {
		cfg.Email.To = c.MailTo
	}

	if cfg.Tables == nil {
		cfg.Tables = []string{}
	}
	if len(c.Tables) > 0 {
		_ts := strings.Split(c.Tables, ",")
		for _, _name := range _ts {
			_name = strings.TrimSpace(_name)
			if len(_name) > 0 {
				cfg.Tables = append(cfg.Tables, _name)
			}
		}
	}
	if cfg.SkipTables == nil {
		cfg.SkipTables = []string{}
	}
	if len(c.SkipTables) > 0 {
		_ts := strings.Split(c.SkipTables, ",")
		for _, _name := range _ts {
			_name = strings.TrimSpace(_name)
			if len(_name) > 0 {
				cfg.SkipTables = append(cfg.SkipTables, _name)
			}
		}
	}
	defer (func() {
		if e := recover(); e != nil {
			log.Println(e)
			err = fmt.Errorf("%s", e)
			cfg.SendMailFail(err.Error())
		}
	})()

	cfg.Check()
	return internal.CheckSchemaDiff(cfg), nil
}
