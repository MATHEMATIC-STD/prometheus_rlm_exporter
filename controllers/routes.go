package controllers

import "github.com/OlivierArgentieri/NukeRlmPrometheusExporter/middlewares"

func (server *Server) initRoutes() {
	// route

	// rlm routes
	server.Router.HandleFunc("/metrics", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareCORS(server.GetLicenses))).Methods("GET")
}
