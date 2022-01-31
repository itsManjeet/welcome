package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/itsmanjeet/welcome/app"
)

func main() {
	HOME := os.Getenv("HOME")
	doneFile := path.Join(HOME, ".config", "welcome")

	if _, err := os.Stat(doneFile); err == nil && len(os.Getenv("FORCE_START")) == 0 {
		return
	}

	application, err := gtk.ApplicationNew("dev.rlxos.Welcome", glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatal(err)
	}

	application.Connect("activate", func() {
		builder, err := gtk.BuilderNewFromString(app.UI)
		if err != nil {
			log.Fatal(err)
		}

		window, err := app.New(builder)
		if err != nil {
			log.Fatal(err)
		}

		window.SetApplication(application)
		window.Present()
	})

	application.Connect("shutdown", func() {

		ioutil.WriteFile(doneFile, []byte("done"), 0644)
	})

	os.Exit(application.Run(os.Args))
}
