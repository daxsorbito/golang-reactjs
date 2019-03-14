package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"syscall"
	"time"

	"github.com/daxsorbito/golang-reactjs/api"
	"github.com/daxsorbito/golang-reactjs/server"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var version = "1.0.0"

func main() {
	cfg := struct {
		BuildPath string
		Listen    string
	}{}

	kp := kingpin.New(filepath.Base(os.Args[0]), "Demo of create-react-app integration into golang http server")
	kp.Version(version)
	kp.Flag("listen", "Which address should be listened").Required().StringVar(&cfg.Listen)
	kp.Flag("build", "Path to the build directory of the project created using create-react-app").Required().StringVar(&cfg.BuildPath)
	kp.HelpFlag.Short('h')

	if _, err := kp.Parse(os.Args[1:]); err != nil {
		kp.Usage(os.Args[1:])
		os.Exit(1)
	}

	buildPath := path.Clean(cfg.BuildPath)
	buildURL := fmt.Sprintf("/%s/", buildPath)

	mux := http.NewServeMux()
	mux.Handle(buildURL, http.StripPrefix(buildURL, http.FileServer(http.Dir(buildPath))))
	mux.Handle("/api", api.Handler())
	mux.Handle("/", server.Handler(buildPath))

	srv := &http.Server{
		Addr:    cfg.Listen,
		Handler: mux,
	}

	errs := make(chan error, 1)
	go func() {
		fmt.Println("Starting", cfg.Listen)
		errs <- srv.ListenAndServe()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case <-stop:
		fmt.Println("Shutting down...")
		os.Exit(0)
	case err := <-errs:
		fmt.Println("Failed to start server:", err.Error())
		os.Exit(1)
	}

	shutdown, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdown); err != nil {
		fmt.Println("Failed to shutdown server:", err.Error())
		os.Exit(1)
	}
}
