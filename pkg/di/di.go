package di

import (
	"s3service/pkg/boot"
	"s3service/pkg/client"
	"s3service/pkg/config"
	"s3service/pkg/routes"
	"s3service/pkg/user/handler"
	"s3service/pkg/user/service"
)

func Init(cnf config.Config) *boot.Server {

	s3initializer := client.S3pathway{Conf: cnf}
	s3initializer.Init()

	server := boot.NewHTTPServer()
	service := service.Service{
		S3: s3initializer,
	}
	handler := handler.Handler{
		Service: service,
	}
	usRoute := routes.UserRoutesstruct{
		Server: server,
		User:   handler,
	}
	usRoute.Routes()

	return server

}
