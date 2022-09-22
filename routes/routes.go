package routes

import (
	// "fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

    "crud-echo/models"
    // "github.com/golang/protobuf/jsonpb"
)
type User struct {
  Name  string `json:"name" xml:"name"`
  Email string `json:"email" xml:"email"`
}

var customer = &models.Customer{
    Id   :  "1",
    Name : "Thomas",
    Email: "thomas@gmail.com",
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
    var customerList = &models.ListCustomer{
        List :  []*models.Customer{
            customer,
        },
    }
    return c.JSON(http.StatusOK, customerList)
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