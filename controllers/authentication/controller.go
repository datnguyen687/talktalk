package authentication

import (
	"errors"
	"math/rand"
	"talktalk/controllers"
	"talktalk/models"
	dataservice "talktalk/services/data"
	mysqlDS "talktalk/services/data/mysql"
	emailservice "talktalk/services/email"
	sendGrid "talktalk/services/email/sendgrid"
	"time"

	"gorm.io/gorm"
)

type controller struct {
	ds dataservice.ServiceInterface
	es emailservice.ServiceInterface
}

// NewAuthenticationController ...
func NewAuthenticationController(cfg AuthenticationConfig) (controllers.ControllerInterface, error) {
	c := &controller{}

	var err error
	if c.ds, err = mysqlDS.NewMySQLDataService(cfg.MySQLConfig); err != nil {
		return nil, err
	}

	c.es = sendGrid.NewEmailService(cfg.EmailConfig)

	return c, nil
}

func (c *controller) generateActivationCode() string {
	tokens := []rune("123456789")
	length := len(tokens)
	code := make([]rune, models.ActivationCodeLength)
	for i := 0; i < models.ActivationCodeLength; i++ {
		index := rand.Int() % length
		code[i] = tokens[index]
	}

	return string(code)
}

func (c *controller) SignUp(dto *models.UserDTO) error {
	found, err := c.ds.GetUser(dto.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if found != nil {
		return errors.New("user already signed up")
	}

	model := models.User{
		CreatedAt:      time.Now().UTC(),
		Email:          dto.Email,
		Password:       dto.Password,
		UserStatusesID: models.NotActivated,
	}
	if err = c.ds.InsertUser(&model); err != nil {
		return err
	}

	// generate activation code and store
	// expire after 15 mins
	code := c.generateActivationCode()
	now := time.Now()
	activation := models.ActivationCode{
		Code:      code,
		CreatedAt: now,
		ExpiredAt: now.Add(time.Minute * time.Duration(MaxActivationCodeLifeSpanInMinutes)),
		UserEmail: dto.Email,
	}

	go c.es.SendActivationCode(dto.Email, code)

	return c.ds.InsertActivationCode(&activation)
}

func (c *controller) ActivateUser(email, code string) error {
	_, err := c.ds.GetUser(email)
	if err != nil {
		return errors.New("unable to find user")
	}

	ac, err := c.ds.GetActivationCode(email)
	if err != nil {
		return err
	}
	if ac == nil {
		return errors.New("code not found")
	}

	if code != ac.Code {
		return errors.New("code mismatched")
	}

	now := time.Now().UTC()
	deadLine := ac.ExpiredAt.UTC()
	if deadLine.Before(now) {
		go c.ds.DeleteActivationCode(email, code)
		return errors.New("code has been expired")
	}

	ac.User.UserStatusesID = models.Activated
	err = c.ds.UpdateUser(&ac.User)
	if err != nil {
		return err
	}

	return c.ds.DeleteActivationCode(email, code)
}

func (c *controller) ResendCode(email string) (string, error) {
	_, err := c.ds.GetUser(email)
	if err != nil {
		return "", errors.New("unable to find user")
	}

	ac, err := c.ds.GetActivationCode(email)
	now := time.Now()
	if err == nil && ac != nil {
		deadLine := ac.ExpiredAt.UTC()
		if deadLine.After(now) {
			return ac.Code, nil
		}

		err = c.ds.DeleteActivationCode(email, ac.Code)
		if err != nil {
			return "", err
		}
	}

	code := c.generateActivationCode()
	activation := models.ActivationCode{
		Code:      code,
		CreatedAt: now,
		ExpiredAt: now.Add(time.Minute * time.Duration(MaxActivationCodeLifeSpanInMinutes)),
		UserEmail: email,
	}

	if err = c.ds.InsertActivationCode(&activation); err != nil {
		return "", err
	}

	return code, nil
}
