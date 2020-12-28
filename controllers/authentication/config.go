package authentication

import (
	mysqlDS "talktalk/services/data/mysql"
	sendGrid "talktalk/services/email/sendgrid"
)

type AuthenticationConfig struct {
	MySQLConfig mysqlDS.Config
	EmailConfig sendGrid.Config
}
