package main

import (
	"net/http"
)

func main() {
	http.Get("https://pingpongtester.onrender.com/ping")
}
