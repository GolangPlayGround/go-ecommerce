package router

import "github.com/gin-gonic/gin"

func (r *Router) GET(path string, handler gin.HandlerFunc) gin.IRoutes {
	return r.engine.GET(path, handler)
}

func (r *Router) POST(path string, handler gin.HandlerFunc) gin.IRoutes {
	return r.engine.POST(path, handler)
}

func (r *Router) PUT(path string, handler gin.HandlerFunc) gin.IRoutes {
	return r.engine.PUT(path, handler)
}

func (r *Router) DELETE(path string, handler gin.HandlerFunc) gin.IRoutes {
	return r.engine.DELETE(path, handler)
}
