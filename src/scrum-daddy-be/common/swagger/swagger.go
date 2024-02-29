package swagger

import (
	_ "scrum-daddy-be/docs"
)

import (
	"github.com/swaggo/http-swagger"
	"net/http"
)

// SetupSwagger sets up the Swagger UI
func SetupSwagger(mux *http.ServeMux) {
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
}
