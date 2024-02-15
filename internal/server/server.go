package server

import (
	"time"

	"github.com/TimothyYe/godns/internal/server/controllers"
	"github.com/TimothyYe/godns/internal/settings"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	addr       string
	username   string
	password   string
	app        *fiber.App
	controller *controllers.Controller
	config     *settings.Settings
	configPath string
}

func (s *Server) SetAddress(addr string) *Server {
	s.addr = addr
	return s
}

func (s *Server) SetAuthInfo(username, password string) *Server {
	s.username = username
	s.password = password
	return s
}

func (s *Server) SetConfig(config *settings.Settings) *Server {
	s.config = config
	return s
}

func (s *Server) SetConfigPath(configPath string) *Server {
	s.configPath = configPath
	return s
}

func (s *Server) Build() {
	config := fiber.Config{}
	s.app = fiber.New(config)
	s.controller = controllers.NewController(s.config, s.configPath)
}

func (s *Server) Start() error {
	s.initRoutes()

	log.Infof("Server is listening on port: %s", s.addr)
	return s.app.Listen(s.addr)
}

func (s *Server) Stop() error {
	return s.app.ShutdownWithTimeout(200 * time.Millisecond)
}

func (s *Server) initRoutes() {
	// set cross domain access rules
	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST ,PUT ,DELETE, OPTIONS",
		AllowHeaders: "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
	}))

	// Create routes group.
	route := s.app.Group("/api/v1")
	route.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			s.username: s.password,
		},
	}))

	// Register routes
	route.Get("/auth", s.controller.Auth)

	// Get basic info
	route.Get("/info", s.controller.GetBasicInfo)

	// Domain related routes
	route.Get("/domains", s.controller.GetDomains)
	route.Post("/domains", s.controller.AddDomain)
	route.Delete("/domains/:name", s.controller.DeleteDomain)
}
