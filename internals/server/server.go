package server

/*
* -> Creates Redis client
* -> Creates rate limiter
* -> Builds router
* -> Attach Middlewares
* -> Handle Graceful shutdown
* -> Expose Start()
 */
import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ndbeloved/rate-limiter-go/internals/cache"
	"github.com/Ndbeloved/rate-limiter-go/internals/config"
	"github.com/Ndbeloved/rate-limiter-go/internals/middleware"
	"github.com/Ndbeloved/rate-limiter-go/internals/ratelimit"
	"github.com/Ndbeloved/rate-limiter-go/internals/router"
)

type Server struct {
	cfg  *config.Config
	http *http.Server
	ctx  context.Context
}

func New(cfg *config.Config, ctx context.Context) *Server {
	return &Server{
		cfg: cfg,
		ctx: ctx,
	}
}

func (s *Server) Start() error {
	//Create Redis Client
	rdb, err := cache.New(s.cfg.RedisAddr, s.ctx)
	if err != nil {
		return err
	}
	//Create rate limiter
	limiter := ratelimit.NewSlidingWindow(rdb, s.cfg.Limit, s.cfg.Window)

	//Create middleware
	rlMiddleware := middleware.RateLimitMiddleware(limiter)

	//Build router
	router := router.NewRouter(rlMiddleware)

	//Configure http.Server
	s.http = &http.Server{
		Addr:    s.cfg.Port,
		Handler: router,
	}
	//Start server in Goroutine
	go func() {
		if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v\n", err)
		}
	}()
	log.Printf("Server running on %s\n", s.cfg.Port)
	//Handle graceful shutdown
	return s.gracefulShutdown()
}

func (s *Server) gracefulShutdown() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.http.Shutdown(ctx)
}
