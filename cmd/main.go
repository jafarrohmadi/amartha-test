package main

import (
	"os"

	"github.com/amartha-test/generated"
	"github.com/amartha-test/handler"
	"github.com/amartha-test/repository"
	"github.com/amartha-test/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})

	var useCase usecase.UseCaseInterface = usecase.NewUseCase(usecase.NewUseCaseOptions{
		Repository: repo,
	})

	opts := handler.NewServerOptions{
		UseCase: useCase,
	}

	return handler.NewServer(opts)
}
