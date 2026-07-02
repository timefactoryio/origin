package main

import (
	"os"

	"github.com/timefactoryio/pathless"
)

func main() {
	p := pathless.NewPathless(os.Getenv("ORIGIN"), os.Getenv("CIRCUIT"))
	p.Home(os.Getenv("LOGO"), os.Getenv("TITLE"))
	p.Slides(os.Getenv("SLIDES"))
	p.Text(os.Getenv("TEXT"))
	p.Serve()
}
