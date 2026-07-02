package main

import (
	"os"

	"github.com/timefactoryio/pathless"
)

// func main() {
// 	p := pathless.NewPathless()
// 	p.Home("https://zero.s3.timefactory.io/timefactory.svg", "the point of origin")
// 	p.Slides("https://zero.s3.timefactory.io/slides")
// 	p.Text("../pathless/mechanics.md")
// 	p.CustomHTML("./custom.html")
// 	p.Serve()
// }

func main() {
	p := pathless.NewPathless(os.Getenv("ORIGIN"), os.Getenv("CIRCUIT"))
	p.Home(os.Getenv("LOGO"), os.Getenv("TITLE"))
	p.Slides(os.Getenv("SLIDES"))
	p.Text(os.Getenv("TEXT"))
	p.Serve()
}
