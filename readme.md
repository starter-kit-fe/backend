# admin 

web api

## 目录结构

以下是项目的目录结构及其说明：

```
project-root/
├── cmd/                    # 应用程序入口
│   └── main.go             # 主启动文件
├── config/                 # 配置文件
│   └── config.go           # 配置加载和管理
├── internal/               # 私有应用程序代码
│   ├── domain/             # 领域模型/实体
│   │   └── user.go         # 用户实体定义
│   ├── repository/         # 数据访问层
│   │   └── user_repository.go # 用户数据操作接口及实现
│   ├── service/            # 业务逻辑层
│   │   └── user_service.go # 用户服务层，处理业务逻辑
│   ├── handler/            # HTTP处理层(控制器)
│   │   └── user_handler.go # 用户相关HTTP请求处理
│   ├── middleware/         # 中间件
│   │   └── auth.go         # 认证授权中间件
│   └── dto/                # 数据传输对象
│       └── user_dto.go     # 用户DTO，用于API输入输出的数据转换
├── pkg/                    # 可以被外部应用程序使用的库代码
│   ├── utils/              # 工具类库
│   └── errors/             # 自定义错误处理
├── api/                    # API文档
│   └── swagger/            # Swagger/OpenAPI规范文档
├── docs/                   # 项目文档
├── test/                   # 测试文件
├── scripts/                # 脚本文件，如部署脚本等
├── go.mod                  # Go模块依赖文件
└── go.sum                  # Go模块校验文件
```

## 快速开始

### 安装依赖

确保你已经安装了 Go 环境。然后运行以下命令来安装项目依赖：

```bash
go mod download
```

### 运行项目

运行项目前，请确保所有配置文件已正确设置。然后执行以下命令启动项目：

```bash
make dev
```

### 构建项目

要构建项目，可以使用以下命令：

```bash
make build
```

这将生成一个名为 `admin` 的可执行文件。

### 运行测试

项目包含了一些单元测试和集成测试。要运行这些测试，可以使用：

```bash
go test ./...
```

## 文档
https://apifox.com/apidoc/shared-66d8bf61-f515-4fa0-8ea0-2c81068aed01

## 贡献

欢迎贡献！请先阅读 [贡献指南](CONTRIBUTING.md)，了解如何参与项目开发。

## 许可证

本项目采用 MIT 许可证，详情见 [LICENSE](LICENSE) 文件。

---

bash <(wget --no-check-certificate -qO- 'https://raw.githubusercontent.com/MoeClub/Note/master/InstallNET.sh') --ip-addr 10.140.0.5 --ip-gate 10.140.0.1 --ip-mask 255.255.255.0 -d 12 -v 64 -p a6WyfFw0v3ShfK6PL3XPJQ== -port 22 