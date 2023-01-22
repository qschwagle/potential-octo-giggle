package webserver


import (
	"bytes"
	"testing"
	"net/http/httptest"
	"encoding/json"
	"github.com/qschwagle/potential-octo-giggle/internal/application_mock"
	"github.com/qschwagle/potential-octo-giggle/internal/application"
	"github.com/stretchr/testify/assert"
)

func setupMockWebApp() (*WebServer, application.IApplication) {
	webBuilder := newWebServerBuilder()
	webBuilder.Testing = true
	web := webBuilder.Create()
	app := application_mock.NewApplicationMock(web)
	app.Setup()

	return web, &app
}

func TestUserGet(t* testing.T) {
	web, app := setupMockWebApp()

	app.Run()

	// check and get valid user

	req := httptest.NewRequest("GET", "/api/user/6c41e5a5-842a-4298-b0c3-7f244e4cd9c0", nil)

	resp, _ := web.Server.Test(req, 1)

	respBody  := make([]byte, 4096)

	read, _ := resp.Body.Read(respBody)

	text := string(respBody[:read])

	expectedData := &struct {
		Username string
	}{ Username: "mary" }
	expected, _ := json.Marshal(expectedData)

	assert.Equalf(t, 200, resp.StatusCode, "GetValidUser: status was not 200")
	assert.Equalf(t, string(expected), text, "GetValidUser: JSON incorrect")

	// get invalid id

	req = httptest.NewRequest("GET", "/api/user/HelloWorld1224332", nil)


	resp, _ = web.Server.Test(req, 1)

	assert.Equalf(t, 400, resp.StatusCode, "GetInvalidId: status was not 400")

	// get valid id, but unknown user

	req = httptest.NewRequest("GET", "/api/user/342ef4e4-4668-47f5-b3eb-05505f0fd6d6", nil)

	resp, _ = web.Server.Test(req, 1)

	assert.Equalf(t, 404, resp.StatusCode, "GetUnknwonUser: status was not 404")

}

func TestUserPost(t* testing.T) {
	web, app := setupMockWebApp()

	app.Run()

	// no body

	req := httptest.NewRequest("POST", "/api/user/new", nil)

	resp, _ := web.Server.Test(req, 1)

	assert.Equalf(t, 400, resp.StatusCode, "PostNewUserNoBody: statuscode is not 400")

	// proper body

	payload := &struct {
		Username string
		Email string
		Password string
	}{
		Username: "joan",
		Email: "joan@example.com",
		Password: "joan2020",
	}

	payloadJson, _ := json.Marshal(payload)

	payloadReader := bytes.NewReader(payloadJson)

	req = httptest.NewRequest("POST", "/api/user/new", payloadReader)

	req.Header.Add("content-type", "application/json")

	resp, _ = web.Server.Test(req, 1)

	assert.Equalf(t, 201, resp.StatusCode, "PostNewUserCreated: statuscode is not 201")
}

func TestUserPatch(t* testing.T) {
	web, app := setupMockWebApp()

	app.Run()

	req := httptest.NewRequest("PATCH", "/api/user", nil)

	resp, _ := web.Server.Test(req, 1)

	assert.Equalf(t, 200, resp.StatusCode, "Hello, World")
}

func TestUserDelete(t* testing.T) {
	web, app := setupMockWebApp()

	app.Run()

	req := httptest.NewRequest("DELETE", "/api/user", nil)

	resp, _ := web.Server.Test(req, 1)

	respBody  := make([]byte, 4096)

	read, _ := resp.Body.Read(respBody)

	text := string(respBody[:read])

	assert.Equalf(t, 200, resp.StatusCode, "Expected 200 Status Code")
	assert.Equal(t, "Hello, World", text)
}
