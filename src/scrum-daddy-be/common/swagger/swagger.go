package swagger

import (
	"scrum-daddy-be/docs"
)

import (
	"github.com/swaggo/http-swagger"
	"net/http"
)

// SetupSwagger sets up the Swagger UI
func SetupSwagger(mux *http.ServeMux) {
	setSwaggerInfo()
	mux.Handle("/swagger-ui/", httpSwagger.WrapHandler)
}

func setSwaggerInfo() {
	docs.SwaggerInfo.Title = "Scrum Daddy"
	docs.SwaggerInfo.Description = "Scrum Daddy API"
	docs.SwaggerInfo.Version = "1.0"
}
