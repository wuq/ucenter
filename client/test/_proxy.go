package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/micro/go-micro/errors"
	"github.com/micro/go-web"
)

// exampleCall will handle /example/call
func ac(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// get name
	name := r.Form.Get("name")

	if len(name) == 0 {
		http.Error(
			w,
			errors.BadRequest("go.micro.api.example", "no content").Error(),
			400,
		)
		return
	}

	// marshal response
	b, _ := json.Marshal(map[string]interface{}{
		"message": "got your message " + name,
	})

	// write response
	w.Write(b)
}


func main() {
	// we're using go-web for convenience since it registers with discovery
	service := web.NewService(
		web.Name("go.micro.api.ab"),
	)

	service.HandleFunc("/ab/call", ac)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
