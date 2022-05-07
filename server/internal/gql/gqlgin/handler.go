package gqlgin

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Router struct {
	log *zap.Logger
}

// NewRouter ハンドラ作成
func NewRouter(log *zap.Logger) *Router {
	return &Router{log}
}

// GraghQL ハンドラ
func (r *Router) GraphQL(schema graphql.ExecutableSchema) gin.HandlerFunc {
	h := handler.NewDefaultServer(schema)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Playground Playgroundハンドラ
func (r *Router) Playground() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
