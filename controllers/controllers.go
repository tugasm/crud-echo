package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/antonlindstrom/pgstore"
	"github.com/labstack/echo"

	"crud-echo/config"
	"crud-echo/models"
)

const SESSION_ID = "id"


func HandlerGetCustomer(c echo.Context) error {

    store := newPostgresStore()

    session, _ := store.Get(c.Request(), SESSION_ID)

    if len(session.Values) == 0 {
        errCust := &models.ErrorCust{
            Message : "Session not found",
            Kode: 400,
        }

        var errCustJson = &models.ListCustomer{
            ErrorCust :  []*models.ErrorCust{
                errCust,
            },
        }

        return c.JSON(http.StatusBadRequest, errCustJson)
    }
    
    sess := &models.Session{
        Session : session.Values["session"].(string),
    }

    db := config.ConnectDB()

    customer := []*models.Customer{}
    err := db.Find(&customer).Error
    if err != nil {
		c.JSON(http.StatusOK, "Error Get Customer")
	}

    var customerList = &models.ListCustomer{
        List :  customer,
        Session: []*models.Session{
            sess,
        },
    } 

    return c.JSON(http.StatusOK, customerList)
}

func HandlerSetSession(c echo.Context) error {

    store := newPostgresStore()
    
    session, _ := store.Get(c.Request(), SESSION_ID)
    // session.Values["session"] = "asd session"
    session.Values["session"] = 20220911 
    sess := &models.Session{
        Session : session.Values["session"].(string),
    }

    sesValues := sess
    
    session.Save(c.Request(), c.Response())
    
    var result = &models.ListCustomer{
        Session :  []*models.Session{
            sesValues,
        },
    }

    return c.JSON(http.StatusOK, result)
}

func newPostgresStore() *pgstore.PGStore {
	url := "postgres://postgres:123456@127.0.0.1:5433/laundryproject?sslmode=disable"
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store, err := pgstore.NewPGStore(url, authKey, encryptionKey)
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(0)
	}

	return store
}

func HandlerDeleteSession(c echo.Context) error {
    store := newPostgresStore()

	session, _ := store.Get(c.Request(), SESSION_ID)
	session.Options.MaxAge = -1
	session.Save(c.Request(), c.Response())

    return c.JSON(http.StatusOK, "Deleted Session")
}