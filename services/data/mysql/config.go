package mysql

// Config ...
type Config struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DbName   string `json:"db_name"`
}
