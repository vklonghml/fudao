package options

import (
	"github.com/kubeedge/beehive/pkg/common/config"
)

type Options struct {
	// server config
	Address string
	Port    int

	// db config
	DBSource  string
	DBName    string
	DBTimeout string

	// work config
	Interval int
}

func GetOptions() *Options {
	opt := Options{}

	config.StringVar(&opt.Address, "server.address", "127.0.0.1", "server ip")
	config.IntVar(&opt.Port, "server.port", 12345, "server port")
	config.StringVar(&opt.DBSource, "database.dbsource", "", "database string")
	config.StringVar(&opt.DBName, "database.dbname", "", "database name")
	config.StringVar(&opt.DBTimeout, "database.dbtimeout", "timeout=15s", "database timout")
	config.IntVar(&opt.Interval, "worker.interval", 1, "worker get courses interval")

	return &opt
}
