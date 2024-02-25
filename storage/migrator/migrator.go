package migrator

import (
	"database/sql"
	"fmt"
	"gameapp/storage/mysql"
	migrate "github.com/rubenv/sql-migrate"
)

type Migrator struct {
	dialect    string
	config     mysql.Config
	migrations *migrate.FileMigrationSource
}

func New(cfg mysql.Config) Migrator {
	migrations := &migrate.FileMigrationSource{
		Dir: "./storage/mysql/migrations",
	}
	return Migrator{dialect: "mysql", config: cfg, migrations: migrations}
}

func (m Migrator) Up() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s(%s:%d)/%s?parseTime=true", m.config.User, m.config.Password, m.config.Address, m.config.Port, m.config.DataBaseName))
	if err != nil {
		panic(fmt.Sprintf("error happend when open mysql in migrator :%v", err))
	}
	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Up)
	if err != nil {
		panic(fmt.Sprintf("error happend when excute up migarte:%v", err))
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func (m Migrator) Down() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s(%s:%d)/%s?parseTime=true", m.config.User, m.config.Password, m.config.Address, m.config.Port, m.config.DataBaseName))
	if err != nil {
		panic(fmt.Sprintf("error happend when open mysql in migrator :%v", err))
	}
	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Down)
	if err != nil {
		panic(fmt.Sprintf("error happend when excute up migarte:%v", err))
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func (m Migrator) Status() {
	panic("not implemented")
}
