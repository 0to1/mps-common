package cfghelper

import (
	"os"
	"strings"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/env"
	"github.com/micro/go-micro/v2/config/source/etcd"
	"github.com/micro/go-micro/v2/config/source/file"
)

// var builtinConfig = []byte(`{
//     "hosts": {
//         "database": {
//             "address": "10.0.0.1",
//             "port": 3306
//         },
//         "cache": {
//             "address": "10.0.0.2",
//             "port": 6379
//         }
//     }
// }`)

// CfgSource 配置来源
var CfgSource source.Source

func LoadConfig(opts ...Option) error {
	opt := newOptions(opts...)

	//配置3个来源，按照优先级从高到低，从配置中心、配置文件、环境变量、内置变参数中读取配置

	// 检测环境变量，是否设置了CONFIG_SOURCE,
	// 如果etcd:10.0.0.10:8500的形式，则采用配置中心的配置，
	// 如果设置为file，则加载程序工作目录下config/config.json的配置，
	// 如果设置为env，则从环境变量中读取配置，
	// 如果未设置，则尝试使用内置配置

	// 未显式指定配置来源，则读取环境变量中配置的配置来源
	if opt.source == "none" {
		opt.source = os.Getenv("CONFIG_SOURCE")
	}

	if strings.HasPrefix(opt.source, "etcd") {
		// create new source from etcd

		var addr string

		index := strings.Index(opt.source, ":")
		if index == len("etcd") {
			addr = opt.source[index+1:]
		}

		// index = strings.Index(addr, ":")
		// if index <= 0 {
		// 	return errors.New("load config failed, details: etcd address is empty")
		// }

		CfgSource = etcd.NewSource(
			etcd.WithAddress(addr),
			// optionally specify prefix; defaults to /micro/config
			etcd.WithPrefix("/mps"))
	} else if strings.HasPrefix(opt.source, "file") {
		// create new source from json file
		var path string

		index := strings.Index(opt.source, ":")
		if index == len("file") {
			path = opt.source[index+1:]
		} else {
			path = "./config/config.json"
		}

		CfgSource = file.NewSource(file.WithPath(path))
	} else if opt.source == "env" || opt.source == "" {
		// create new source from env
		CfgSource = env.NewSource()
	}
	// else if opt.source == "" {
	// 	CfgSource = memory.NewSource(memory.WithJSON(builtinConfig))
	// }

	if err := config.Load(CfgSource); err != nil {
		return err
	}

	return nil
}
