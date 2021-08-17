# CommonConfig
go语言一些常用配置封装

## 项目结构

```shell
├── api
│   └── v1
├── config
├── core
├── docs
├── global
├── initialize
│   └── internal
├── log
├── middleware
├── model
│   ├── request
│   └── response
├── router
├── service
└── utils
    └── upload
```

| 文件夹       | 说明                    | 描述                        |
| ------------ | ----------------------- | --------------------------- |
| `api`        | api层                   | api层 |
| `--v1`       | v1版本接口              | v1版本接口                  |
| `config`     | 配置包                  | config.toml或其他类型配置文件对应的配置结构体 |
| `core`       | 核心文件                | 核心组件(zap, viper, server)的初始化 |
| `docs`       | swagger文档目录         | swagger文档目录 |
| `global`     | 全局对象                | 全局对象 |
| `initialize` | 初始化 | router,redis,gorm的初始化 |
| `--internal` | 初始化内部函数 | gorm 的 longger 自定义,在此文件夹的函数只能由 `initialize` 层进行调用 |
| `log`        | 日志    | 用于存放 日志文件 |
| `middleware` | 中间件层 | 用于存放 `gin` 中间件代码 |
| `model`      | 模型层                  | 模型对应数据表              |
| `--request`  | 入参结构体              | 接收前端发送到后端的数据。  |
| `--response` | 出参结构体              | 返回给前端的数据结构体      |
| `router`     | 路由层                  | 路由层 |
| `service`    | service层               | 存放业务逻辑问题 |
| `utils`      | 工具包                  | 工具函数封装            |
| `--upload`      | oss                  | oss接口封装        |

