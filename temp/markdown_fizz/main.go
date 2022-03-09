// https://github.com/wI2L/fizz/tree/master/markdown

package main

import (
	"fmt"
	"github.com/wI2L/fizz/markdown"
)

func main() {

	builder := markdown.Builder{}

	builder.
		H1("Markdown Builder").
		P("A simple builder to help your write Markdown in Go").
		H2("Installation").
		Code("go get -u github.com/wI2L/fizz", "bash").
		H2("Todos").
		BulletedList(
			"write tests",
			builder.Block().NumberedList("A", "B", "C"),
			"add more markdown features",
		)

	builder.Table(
		[][]string{
			[]string{"Letter", "Title", "ID"},
			[]string{"A", "The Good", "500"},
			[]string{"B", "The Very very Bad Man", "2885645"},
			[]string{"C", "The Ugly"},
			[]string{"D", "The\nGopher", "800"},
		}, []markdown.TableAlignment{
			markdown.AlignCenter,
			markdown.AlignCenter,
			markdown.AlignRight,
		},
	)

	md := builder.String()

	fmt.Println(md)

}
