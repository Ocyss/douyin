package conf

var (
	Conf      Config
	Address   = Conf.Address
	Port      = Conf.Port
	JwtSecret = Conf.JwtSecret
	Scheme    = Conf.Scheme
	Database  = Conf.Database
	Log       = Conf.Log
)

func DefaultConfig() Config {
	return Config{
		Address: "0.0.0.0",
		Port:    23724, // 2023-07-24
		Database: confDatabase{
			Type:   "sqlite3",
			DbFile: "data/data.db",
		},
		Log: confLog{
			Enable:     true,
			Level:      "info",
			Name:       "data/log/log.log",
			MaxSize:    10,
			MaxBackups: 5,
			MaxAge:     28,
			Compress:   false,
		},
	}
}
