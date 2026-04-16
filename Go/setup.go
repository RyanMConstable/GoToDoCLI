package main

import (
	"github.com/BurntSushi/toml"
)

func Setup(d *Data) {
	d.setup.dataFile = "/etc/todo/todo.json"
	d.setup.confFile = "/etc/todo/todo.conf"
}

func LoadConfig(d *Data) error {
	if d.flags.config != "" {
		d.setup.confFile = d.flags.config
	}

	_, err := toml.DecodeFile(d.setup.confFile, &d.config)
	if err != nil {
		return err
	}

	return nil

}

func IsDatabaseConfigured(d *Data) error {
	c := d.config

	if (len(c.DB_HOST) + len(c.DB_USER) + len(c.DB_PASSWORD) + len(c.DB_PORT) + len(c.DB_NAME)) != 0 {
		d.flags.db = true
	}

	return nil
}
