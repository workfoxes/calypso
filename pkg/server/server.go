package server

import (
	"encoding/json"
	middleware2 "github.com/airstrik/gobase/pkg/server/middleware"
	"github.com/airstrik/gobase/pkg/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	Env           string
	IsMultiTenant bool
)

func init() {
	Env = os.Getenv("ENV")
}

type ApplicationServer interface {
	Start()
	RegisterRoute(pattern string, fn func(r chi.Router))
}

type Server struct {
	ApplicationServer
	Name          string
	Port          int
	Route         *chi.Mux
	IsMultiTenant bool
}

func GetServer(name string, port int) *Server {
	IsMultiTenant = false
	return &Server{
		Name:          name,
		Port:          port,
		Route:         createRoute(),
		IsMultiTenant: false,
	}
}

func GetMultiTenantServer(name string, port int) *Server {
	IsMultiTenant = true
	return &Server{
		Name:          name,
		Port:          port,
		Route:         createRoute(),
		IsMultiTenant: true,
	}
}

// CreateServer : Create Simple Application server for any boot strap Application
func createRoute() *chi.Mux {
	log.Print("Creating Application service for the any service")

	r := chi.NewRouter()
	_cors := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowOriginFunc:    func(r *http.Request, origin string) bool { return true },
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:     []string{"Link"},
		AllowCredentials:   true,
		OptionsPassthrough: true,
		MaxAge:             3599, // Maximum value not ignored by any of major browsers
	})
	// Register All the MiddleWare to used in the Application Server for Security
	r.Use(_cors.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.NoCache)
	r.Use(middleware.StripSlashes)
	r.Use(middleware2.ErrorHandling)
	r.Use(middleware2.PreRequest)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(utils.GenerateSuccessResponse())
	})
	return r
}

func (srv *Server) RegisterRoute(pattern string, fn func(r chi.Router)) chi.Router {
	return srv.Route.Route(pattern, fn)
}

func (srv *Server) Start() {
	err := http.ListenAndServe(":"+strconv.Itoa(srv.Port), srv.Route)
	if err != nil {
		log.Panic(err)
	}
	log.Print("Terminating the Application : " + srv.Name)
}
