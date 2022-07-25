package routers

import "core/src/controllers/users"

func (r *router) UserRouter() {

	u := r.eng.Group("/users")
	{
		u.GET("", users.GetUsers)
		u.POST("", users.PostUser)

		id := u.Group("/:id")
		{
			id.GET("", users.GetUser)
			id_auth := id.Group("")
			// auth middleware here
			// id_auth.Use()
			{
				id_auth.PUT("", users.UpdateUser)
				id_auth.DELETE("", users.DeleteUser)
			}
		}
	}
}

// session middlware example
// func AuthRequired(c *gin.Context) {
// 	session := sessions.Default(c)
// 	user := session.Get(userkey)
// 	if user == nil {
// 		// Abort the request with the appropriate error code
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
// 		return
// 	}
// 	// Continue down the chain to handler etc
// 	c.Next()
// }
