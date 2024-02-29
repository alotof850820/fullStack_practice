package router

import (
	"ginchat/docs"
	"ginchat/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default() // 默认gin Engine
	r.Use(cors.Default())
	//設定api接口
	//swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 靜態資源
	//將 ./views/assets 目錄下的靜態檔案映射到 /assets 路徑。
	r.Static("/assets", "./views/assets")
	r.LoadHTMLGlob("views/**/*")

	// 渲染前端首頁
	r.GET("/", service.GetVueHtml)
	// 用戶
	r.GET("/user/getUserList", service.GetUserList)
	r.GET("/user/createUser", service.CreateUser)
	r.DELETE("/user/deleteUser", service.DeleteUser)
	r.POST("/user/updateUser", service.UpdateUser)
	r.POST("/user/findUserByNameAndPassword", service.FindUserByNameAndPassword)
	r.POST("/searchFriends", service.SearchFriends)

	// 發消息
	r.GET("/user/sendMsg", service.SendMsg)
	r.GET("/user/sendUserMsg", service.SendUserMsg)
	r.POST("/attach/upload", service.Upload)
	r.POST("/contact/addFriend", service.AddFriend)

	return r
}
