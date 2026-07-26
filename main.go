package main

import (
	"github.com/timefactoryio/pathless"
)

// func main() {
// 	p := pathless.NewPathless(os.Getenv("ORIGIN"), os.Getenv("CIRCUIT"))
// 	p.Home(os.Getenv("LOGO"), os.Getenv("TITLE"))
// 	p.Text(os.Getenv("TEXT"))
// 	p.Slides(os.Getenv("SLIDES"))
// 	p.Serve()
// }

func main() {
	p := pathless.NewPathless()
	p.Home("https://zero.s3.timefactory.io/timefactory.svg", "the point of origin")
	p.Text("../pathless/README.md")
	p.Text("../pathless/mechanics.md")
	p.Text("../pathless/zero/README.md")
	p.Text("../pathless/fx/README.md")
	p.Text("../pathless/one/README.md")
	p.Keyboard()
	p.Serve()
}
