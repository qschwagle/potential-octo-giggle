package application_mock

import (
	"github.com/qschwagle/potential-octo-giggle/internal/application"
	"github.com/google/uuid"
)

type ApplicationMock struct { 
	Frontend application.IFrontend
}

func NewApplicationMock(frontend application.IFrontend) ApplicationMock {
	return ApplicationMock {
		Frontend: frontend,
	}
}

func (app* ApplicationMock) Setup() {
	app.Frontend.Setup(app)
}

func (app* ApplicationMock) Run() {

}

/// Gets the user by id. if User == nil, then the user was not found
func (app* ApplicationMock) GetUser(id uuid.UUID) (*application.User, error) {

	if id == uuid.MustParse("6c41e5a5-842a-4298-b0c3-7f244e4cd9c0") {
		return &application.User { Username: "mary" }, nil
	}

	return nil, nil
}

func (app* ApplicationMock) CreateUser(username string, password string, email string) (*uuid.UUID, error) {
	id := uuid.New()
	if username == "mary" && email == "may@example.com" {
		return nil, nil
	}
	return &id, nil
}
