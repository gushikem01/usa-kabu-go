package appgin

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gushikem01/usa-kabu-go/server/internal/gql/gqlgin"
)

// API API
type API struct {
	Router *gin.Engine
}

// NewAPI API作成
func NewAPI(schema graphql.ExecutableSchema, gql *gqlgin.Router) *API {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://127.0.0.1:8080",
			"http://localhost:3000",
			"http://localhost:8080",
			"http://localhost:8888"},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
	r.POST("/graphql", gql.GraphQL(schema))
	r.GET("/", gql.Playground())
	return &API{Router: r}
}
