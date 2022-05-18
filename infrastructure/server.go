package infrastructure

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"gitlab.com/yars-ai/yars/services/engine"
	"golang.org/x/sync/errgroup"
)

type Server interface {
	Run(addr string)
}

func NewServer(engines []engine.Engine) Server {
	log.Print("start router initialization")

	projects := make(map[string]Project)
	router := mux.NewRouter()

	for i := range engines {
		projectName := engines[i].GetName()
		_, ok := projects[projectName]
		if ok {
			panic("duplicate project names")
		} else {
			projects[projectName] = newProject(
				router.PathPrefix(fmt.Sprintf("/%s", projectName)).Subrouter(),
				engines[i],
			)
		}
	}

	log.Print("router initialization done")
	return &server{
		router:   router,
		projects: projects,
	}
}

type server struct {
	router   *mux.Router
	projects map[string]Project
}

func (s *server) Run(addr string) {
	log.Printf("start server on addr %s", addr)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	httpServer := &http.Server{
		Addr:    addr,
		Handler: s.router,
	}

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return httpServer.ListenAndServe()
	})

	g.Go(func() error {
		<-gCtx.Done()
		return httpServer.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		log.Printf("server stopped")
		log.Printf("stop reason: %s", err)
	}
}
