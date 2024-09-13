package routes

import (
	"s3service/pkg/boot"
	"s3service/pkg/user/handler"
)

type UserRoutesstruct struct {
	Server *boot.Server
	User   handler.Handler
}

func (ur *UserRoutesstruct) Routes() {

	ur.Server.Engine.POST("/upload", ur.User.Upload)
	ur.Server.Engine.GET("/Image/:name", ur.User.Retrive)

}
