package projects

import (
	"github.com/zviedris/echo-rest-template/modules/core"
	"github.com/zviedris/echo-rest-template/modules/projects/handlers"
	"github.com/zviedris/echo-rest-template/modules/projects/repositories"
	"github.com/zviedris/echo-rest-template/modules/projects/usecases"
	"github.com/zviedris/echo-rest-template/pkg/middlewares"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Module core.ModuleInstance = &projectModule{}

type projectModule struct{}

func (projectModule) RegisterRepositories(container *dig.Container) error {
	container.Provide(repositories.NewPgsqlProjectRepository)
	return nil
}

func (projectModule) RegisterUseCases(container *dig.Container) error {
	container.Provide(usecases.NewProjectUsecase)
	return nil
}

func (projectModule) RegisterHandlers(g *echo.Group, container *dig.Container) error {
	return container.Invoke(func(
		middManager *middlewares.MiddlewareManager,
		projectUsecase usecases.ProjectUsecase,
	) {
		handlers.NewProjectHandler(g, middManager, projectUsecase)
	})
}
