package main

import (
	"os"
	"path/filepath"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var version = "1.0.0"

func main() {
	cfg := struct {
		Listen string
	}{}

	kp := kingpin.New(filepath.Base(os.Args[0]), "Demo of create-react-app integration into golang http server")
	kp.Version(version)
	kp.Flag("listen", "Which address should be listened").Required().StringVar(&cfg.Listen)
	kp.HelpFlag.Short('h')

	if _, err := kp.Parse(os.Args[1:]); err != nil {
		kp.Usage(os.Args[1:])
		os.Exit(1)
	}
}
