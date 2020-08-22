package main

import (
	"fmt"
	"os"

	"github.com/micro-in-cn/starter-kit/pkg/plugin/wrapper/client/router_filter"
	_ "github.com/micro-in-cn/starter-kit/profile"
	_ "github.com/micro/go-micro/v3"
	"github.com/micro/go-micro/v3/client"
	"github.com/micro/go-micro/v3/util/log"
	"github.com/micro/micro/v3/client/cli/util"
	"github.com/micro/micro/v3/client/gateway"
	"github.com/micro/micro/v3/cmd"
	microClient "github.com/micro/micro/v3/service/client"
)

const (
	// EnvLocal is a builtin environment, it represents your local `micro server`
	EnvDev = "dev"
)

var envs = map[string]util.Env{
	EnvDev: {
		Name: EnvDev,
	},
}

func init() {
	// 默认有个 EnvLocal，并且 ProxyAddress = 127.0.0.1:8081
	for _, env := range envs {
		util.AddEnv(env)
	}
}

func main() {
	// 流量染色
	// TODO micro默认的api和web网关均不支持服务筛选，需要自己改造
	// https://micro.mu/blog/cn/2019/12/09/go-micro-service-chain.html
	// 这个方案不可取，查考 asim 在 pull#1388 给的反馈意见，
	// https://github.com/micro/go-micro/pull/1388
	// 自定义 Router 参考 fork 的分支版本
	// https://github.com/hb-chen/micro/tree/gateway/gateway
	// Router services filter
	cmd.Register(
		gateway.Commands(
		//router.WithFilter(chain.NewRouterFilter()),
		),
	)
	microClient.DefaultClient.Init(
		client.WrapCall(router_filter.NewCallWrapper()),
	)

	log.Infof("client options.proxy: %s", microClient.DefaultClient.Options().Proxy)

	if err := cmd.DefaultCmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
