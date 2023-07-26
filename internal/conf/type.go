package conf

type Config struct {
	Address   string       `json:"address"`
	Port      int          `json:"port"`
	JwtSecret string       `json:"jwt_secret"`
	Scheme    confScheme   `json:"scheme"`
	Database  confDatabase `json:"database"`
	Log       confLog      `json:"log"`
}
type confScheme struct {
	Https    bool   `json:"https"`
	CertFile string `json:"cert_file"`
	KeyFile  string `json:"key_file"`
}
type confDatabase struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	DbFile   string `json:"db_file"`
}
type confLog struct {
	Enable     bool   `json:"enable"`
	Level      string `json:"level"`
	Name       string `json:"name"`
	MaxSize    int    `json:"max_size"`
	MaxBackups int    `json:"max_backups"`
	MaxAge     int    `json:"max_age"`
	Compress   bool   `json:"compress"`
}
