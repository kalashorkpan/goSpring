package main

import (
	"fmt"
	"github.com/kalashorkpan/msgo"
	"net/http"
)

func main() {

	engine := msgo.New()

	engine.Add("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world 你好世界")
	})

	engine.Run()
}
