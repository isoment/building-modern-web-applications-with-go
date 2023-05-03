package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/isoment/booking-app/internal/config"
	"github.com/isoment/booking-app/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	// Taken from main.go and modified for testing
	gob.Register(models.Reservation{})
	testApp.InProduction = false

	// Set up logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testApp.InfoLog = infoLog
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false
	testApp.Session = session
	app = &testApp

	os.Exit(m.Run())
}

// A stub writer we can use to substitute for the http.ResponseWriter
type myWriter struct{}

// Now we just need to add these three methods to our writer above that satisfy
// the http.ResponseWriter interface
func (m *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (m *myWriter) WriteHeader(statusCode int) {}

func (m *myWriter) Write(a []byte) (int, error) {
	length := len(a)
	return length, nil
}
