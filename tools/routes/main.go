package main

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func main() {
	generateRoutes()
}

func generateRoutes() {
	e := echo.New()

	mapRoutes := map[string]map[string]string{}
	count := 0
	for _, r := range e.Routes() {
		if strings.HasPrefix(r.Name, "github.com") {
			continue
		}
		count++
		acl := mapRoutes[r.Path]
		if len(acl) == 0 {
			acl = map[string]string{}
		}
		acl[r.Method] = r.Name
		mapRoutes[r.Path] = acl
	}

	log.Log().Msgf("Generated routes: %d", count)
	data, err := json.MarshalIndent(mapRoutes, "", "  ")
	if err != nil {
		log.Fatal().Msgf("error json marshal: %v", err)
	}
	os.WriteFile("./pkg/authz/routes.json", data, 0644)
}	
