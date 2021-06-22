package main

import (
        "bj38_cli/handler"
        bj38_cli "bj38_cli/proto/bj38"
        context2 "context"
        "fmt"
        "github.com/gin-gonic/gin"
        "github.com/micro/go-micro/v2"
        //"github.com/micro/go-micro/v2/client"
        log "github.com/micro/go-micro/v2/logger"
        "github.com/micro/go-micro/v2/registry"
        "github.com/micro/go-micro/v2/web"
        "github.com/micro/go-plugins/registry/consul/v2"
        "net/http"
)
func CallRemote(context *gin.Context){
        consulReg:=consul.NewRegistry(func(options *registry.Options){
                //options.Addrs=[]string{
                //	"127.0.0.1:8800",
                //}
        })
        // New Service
        service := micro.NewService(
                micro.Registry(consulReg),
        )
        microClient:=bj38_cli.NewBj38Service("go.micro.service.bj38",service.Client())
        resp,err:=microClient.Call(context2.TODO(),&bj38_cli.Request{
                Name: "chenbin",
        })
        if err!=nil{
                fmt.Println("call err:",err)
                return
        }
        context.Writer.WriteString(resp.Msg)
}
func main() {
        router:=gin.Default()
        router.GET("/",CallRemote)
        router.Run(":8080")
        // create new web service
        service := web.NewService(
                web.Name("go.micro.web.bj38_cli"),
                web.Version("latest"),
        )

        // initialise service
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

        // register html handler
        service.Handle("/", http.FileServer(http.Dir("html")))

        // register call handler
        service.HandleFunc("/bj38_cli/call", handler.Bj38_cliCall)

        // run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
func main_backup() {
	// create new web service
        service := web.NewService(
                web.Name("go.micro.web.bj38_cli"),
                web.Version("latest"),
        )

	// initialise service
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/bj38_cli/call", handler.Bj38_cliCall)

	// run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
