package main

import (
	"flag"
	"fmt"

	"github.com/mizy/llama-cpp-gpt-api/internal/config"
	"github.com/mizy/llama-cpp-gpt-api/internal/handler"
	"github.com/mizy/llama-cpp-gpt-api/internal/svc"
	"github.com/mizy/llama-cpp-gpt-api/pkg/model"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/gpt-api.yaml", "the config file")

func main() {
	flag.Parse()

	conf.MustLoad(*configFile, &config.C)

	model.LoadModel()
	server := rest.MustNewServer(config.C.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(config.C)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", config.C.Host, config.C.Port)
	server.Start()
}
