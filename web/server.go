package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello world!"
  })

  m.Group("/books", func(r martini.Router) {
      r.Get("/:id", func(params martini.Params) string {
        return "get book " + params["id"]
      })
      //r.Post("/new", NewBook)
      r.Get("/delete/:id", func(params martini.Params) string {
        return "delete book " + params["id"]
      })
  })

  m.RunOnAddr(":8080")
}
