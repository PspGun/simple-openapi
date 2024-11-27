package di

import (
	"backend/internal/routes"
	"github.com/gin-gonic/gin"
)

type IPillar interface {
	RunApp() error
}

type Pillar struct {
	app   *gin.Engine
	route routes.IRouter
}

func (p *Pillar) RunApp() error {
	p.app.Use(gin.Recovery())
	p.app.Use(gin.Logger())
	p.route.InitializedRouter(p.app)

	return p.app.Run(":8000")
}

func ProvideGinEngine(route routes.IRouter) IPillar {
	app := gin.New()

	return &Pillar{
		app:   app,
		route: route,
	}
}
