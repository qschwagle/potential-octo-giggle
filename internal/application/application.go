package application

import (
	"github.com/google/uuid"
)

type AppErrorType int64

const (
	UnknownError AppErrorType = 0
	AccountExists
	NotImplemented
	UserNotFound
)

type AppError struct {
	Msg string
	Type AppErrorType
}


func (err *AppError) Error() string {
	return err.Msg
}

type IFrontend interface {
	Run(app IApplication)
	Setup(app IApplication)
}

type IStorage interface {

}

type User struct {
	Username string
}


type IApplication interface {
	Run()
	Setup()

	GetUser(id uuid.UUID) (*User, error)
	CreateUser(username string, password string, email string) (*uuid.UUID , error)
}

type Application struct {
	frontend IFrontend
}

type ApplicationBuilder struct {
	Frontend IFrontend
}

func (builder *ApplicationBuilder) Create() *Application {
	return &Application { 
		frontend: builder.Frontend,
	}
}


func (app *Application) Run() {
	app.frontend.Run(app)
}

func (app *Application) Setup() {
	app.frontend.Setup(app)
}

func (app *Application) GetUser(id uuid.UUID) (*User, error) {
	return nil, &AppError { Msg: "NOT IMPLEMENTED", Type: NotImplemented }
}

func (app *Application) CreateUser(username string, password string, email string) (*uuid.UUID, error) {
	return nil, &AppError { Msg: "NOT IMPLEMENTED", Type: NotImplemented }
}

