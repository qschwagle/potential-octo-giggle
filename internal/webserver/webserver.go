package webserver

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/qschwagle/potential-octo-giggle/internal/application"
	"github.com/google/uuid"

)


type WebServer struct {
	Server *fiber.App
	testing bool
}

func (web *WebServer) Setup(app application.IApplication) {

	web.Server.Use(logger.New())

	web.Server.Post("/api/login", func (c *fiber.Ctx) error {

		payload := struct {
			Email string `json:"email"`
			Password string `json:"password"`
		}{ }

		if err := c.BodyParser(&payload); err != nil {
			return c.SendStatus(400)
		}

		id, err := app.Login(payload.Email, payload.Password)

		if err != nil {
			return c.SendStatus(500)
		}

		if id == nil {
			return c.SendStatus(401)
		}

		cookie := new(fiber.Cookie)

		cookie.Name = "auth_key"
		cookie.Value = id.String()

		c.Cookie(cookie)

		return c.SendStatus(201)
	})

	web.Server.Delete("/api/logout", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})


	// USER ENDPOINT

	web.Server.Get("/api/user/:id", func (c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		
		if err != nil {
			return c.SendStatus(400)
		}

		user, app_err := app.GetUser(id)

		if app_err != nil {
			return c.SendStatus(500)
		}


		if user != nil {
			ret := struct {
				Username string
			}{
				Username: user.Username,
			}

			return c.JSON(ret)
		} 

		return c.SendStatus(404)
	})

	web.Server.Patch("/api/user", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	web.Server.Delete("/api/user", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	web.Server.Post("/api/user/new", func (c *fiber.Ctx) error {
		payload := struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Email string `json:"email"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return c.SendStatus(400)
		}

		id, err := app.CreateUser(payload.Username, payload.Password, payload.Email)

		if err != nil {
			// server error
			return c.SendStatus(500)
		}

		if id ==  nil {
			// account exists already
			return c.SendStatus(400)
		}

		return c.SendStatus(201)
	})

	// POST ENDPOINT

	web.Server.Get("/api/post", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	web.Server.Post("/api/post", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	web.Server.Patch("/api/post", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	web.Server.Delete("/api/post", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	// POSTS aggregate ENDPOINT

	web.Server.Get("/api/posts/latest", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	web.Server.Get("/api/user/posts", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

}

func (web *WebServer) Run(app application.IApplication) {
	if !web.testing {
		log.Fatal(web.Server.Listen("localhost:3000"))
	}
}


type WebServerBuilder struct {
	Testing bool
}

func newWebServerBuilder() WebServerBuilder {
	return WebServerBuilder { 
		Testing: false,
	}
}

func (builder *WebServerBuilder) Create() *WebServer {
	return &WebServer {
		Server: fiber.New(),
		testing: builder.Testing,
	}
}

