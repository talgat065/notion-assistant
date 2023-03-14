package main

import (
	"github.com/talgat065/notion-assistant/internal/webserver"
)

func main() {
	webserver.NewServer().Run()
}
