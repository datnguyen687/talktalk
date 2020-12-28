package authentication

import (
	mysqlDS "talktalk/services/data/mysql"
	sendGridES "talktalk/services/email/sendgrid"
)

// ServerConfig ...
type ServerConfig struct {
	Port  int               `json:"port"`
	SQL   mysqlDS.Config    `json:"mysql"`
	Email sendGridES.Config `json:"email"`
}
