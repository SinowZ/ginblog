package routers

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)

	r := gin.Default()

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//User模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		//Category模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)

		//Artical模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)
	}

	public := r.Group("api/v1")
	{
		//User模块的路由接口
		public.GET("users", v1.GetUsers)
		public.POST("login", v1.Login)
		public.POST("user/add", v1.AddUser)

		//Category模块的路由接口
		public.GET("category", v1.GetCate)

		//Artical模块的路由接口
		public.GET("article/list/:id", v1.GetCateArt)
		public.GET("article/info/:id", v1.GetArtInfo)
		public.GET("article", v1.GetArt)
	}

	r.Run(utils.HttpPort)
}
