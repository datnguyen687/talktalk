package authentication

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"talktalk/controllers"
	authController "talktalk/controllers/authentication"
	handlers "talktalk/handlers"
	"talktalk/models"
	"time"

	"talktalk/controllers/authentication"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Server ...
type server struct {
	router     chi.Router
	config     ServerConfig
	controller controllers.ControllerInterface
}

// NewAuthorizationServer ...
func NewAuthorizationServer(cfg ServerConfig) (handlers.HandlerInterface, error) {
	controller, err := authController.NewAuthenticationController(authentication.AuthenticationConfig{
		EmailConfig: cfg.Email,
		MySQLConfig: cfg.SQL,
	})
	if err != nil {
		return nil, err
	}
	srv := &server{
		router:     chi.NewRouter(),
		config:     cfg,
		controller: controller,
	}

	// basic middleware
	srv.router.Use(middleware.RequestID)
	srv.router.Use(middleware.RealIP)
	srv.router.Use(middleware.Logger)
	srv.router.Use(middleware.Recoverer)

	// 1 minute timeout for middleware
	srv.router.Use(middleware.Timeout(time.Minute))

	// Set up routes
	srv.router.Get("/", srv.handlePing)
	srv.router.Route("/user", func(r chi.Router) {
		r.Post("/signup", srv.handleSignUp)
		r.Get("/activate", srv.handleActivate)
	})

	return srv, nil
}

func (srv *server) Init(config interface{}) error {
	return nil
}

func (srv *server) Run() error {
	addr := fmt.Sprintf(`:%d`, srv.config.Port)
	log.Println("listening at", addr)
	return http.ListenAndServe(addr, srv.router)
}

func (srv *server) sendResponse(w http.ResponseWriter, err error, httpStatus int) {
	resp := UserSignUpResponse{}
	resp.Error = err

	data, err := jsoniter.Marshal(&resp)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	if _, err = w.Write(data); err != nil {
		log.Println(err)
	}
}

func (srv *server) handlePing(w http.ResponseWriter, r *http.Request) {
	resp := handlers.BasicJSONResponse{}

	data, _ := json.Marshal(&resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (srv *server) handleSignUp(w http.ResponseWriter, r *http.Request) {
	in, err := ioutil.ReadAll(r.Body)
	if err != nil {
		srv.sendResponse(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	dto := models.UserDTO{}
	if err := jsoniter.Unmarshal(in, &dto); err != nil {
		srv.sendResponse(w, err, http.StatusBadRequest)
		return
	}

	err = srv.controller.SignUp(&dto)
	if err != nil {
		srv.sendResponse(w, err, http.StatusInternalServerError)
		return
	}

	srv.sendResponse(w, nil, http.StatusOK)
}

func (srv *server) handleActivate(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	emails, ok := query["email"]
	if !ok || len(emails) <= 0 || emails[0] == "" {
		srv.sendResponse(w, errors.New("missing email"), http.StatusBadRequest)
		return
	}

	codes, ok := query["code"]
	if !ok || len(codes) <= 0 || codes[0] == "" {
		srv.sendResponse(w, errors.New("missing code"), http.StatusBadRequest)
		return
	}

	err := srv.controller.ActivateUser(emails[0], codes[0])
	if err != nil {
		srv.sendResponse(w, err, http.StatusInternalServerError)
		return
	}

	srv.sendResponse(w, nil, http.StatusOK)
}
