package main

import (
	"flag"
	"fmt"

	"go-zero-demo/mall/order/api/internal/config"
	"go-zero-demo/mall/order/api/internal/handler"
	"go-zero-demo/mall/order/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "E:\\project\\go-playground\\go-zero-demo\\mall\\order\\api\\etc\\order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
