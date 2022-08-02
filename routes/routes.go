package routes

import(
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)
type User struct {
  Name  string `json:"name" xml:"name"`
  Email string `json:"email" xml:"email"`
}

func Routes() *echo.Echo {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    	AllowOrigins: []string{"*"},
    	AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
    }))

    v1 := e.Group("/v1")
    groupACustomer(v1)
    return e
}

func handler(c echo.Context) error {
// 	return c.String(http.StatusOK, c.Request().RequestURI)
    u := &User{
        Name:  "Jon",
        Email: "jon@labstack.com",
      }
return c.JSON(http.StatusOK, u)
}

func handler2(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("id"))
}



//CUSTOMERS
func groupACustomer(e *echo.Group) {
	grA := e.Group("/customer")

	grA.POST("/submit", handler)
	grA.GET("/all", handler)
	grA.PATCH("/by:id", handler2)
    grA.GET("/by:id", handler2)
}