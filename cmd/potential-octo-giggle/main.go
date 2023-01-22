package main

import (
	"github.com/qschwagle/potential-octo-giggle/internal/webserver"
	"github.com/qschwagle/potential-octo-giggle/internal/application"
)

func main() {
	webserverBuilder := webserver.WebServerBuilder {}
	webserver := webserverBuilder.Create()

	applicationBuilder := application.ApplicationBuilder {}
	applicationBuilder.Frontend = webserver
	application := applicationBuilder.Create()

	application.Setup()
	application.Run()
}
