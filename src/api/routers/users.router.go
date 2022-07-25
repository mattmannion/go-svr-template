package routers

import (
	"root/src/api/controllers/users"
	"root/src/api/middleware"
)

func (r *router) UserRouter() {

	u := r.eng.Group("/users")
	{
		u.GET("", users.GetUsers)
		u.POST("", users.PostUser)

		id := u.Group("/:id")
		{
			id.GET("", users.GetUser)
			id_auth := id.Group("")

			id_auth.Use(middleware.Auth)
			{
				id_auth.PUT("", users.UpdateUser)
				id_auth.DELETE("", users.DeleteUser)
			}
		}
	}
}
