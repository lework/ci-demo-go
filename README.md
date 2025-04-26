# CI Demo Go

一个基于 Gin 框架的轻量级 Go Web 服务，支持多环境配置和结构化日志。

如果你不知道怎么使用，请查看[Dockerfile 中小型企业实战指南](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzU5MzgyMDAyNQ==&action=getalbum&album_id=3950605469210558464&from_itemidx=1&from_msgid=2247484059#wechat_redirect)

## 功能特点

- 使用 Gin 框架构建 RESTful API
- 支持通过 YAML 文件进行多环境配置
- 集成 logrus 日志系统，支持多种日志级别和格式
- 提供健康检查接口，便于监控和部署
- 支持命令行参数配置

## 依赖项

- [Gin](https://github.com/gin-gonic/gin) - Web 框架
- [Logrus](https://github.com/sirupsen/logrus) - 结构化日志
- [yaml.v2](https://gopkg.in/yaml.v2) - YAML 解析

## 安装

```bash
# 克隆代码库
git clone https://github.com/lework/ci-demo-go.git
cd ci-demo-go

# 安装依赖
go mod tidy
```

## 配置

配置文件示例 (`etc/app_dev.yaml`):

```yaml
server:
  port: 8080
  mode: debug
  env: development

app:
  name: ci-demo-go

log:
  level: debug
  format: text
```

可用的日志级别: `debug`, `info`, `warn`, `error`  
日志格式选项: `text`, `json`

## 运行

```bash
# 使用默认配置文件
go run main.go

# 指定配置文件
go run main.go -f etc/app_prod.yaml
```

## API 接口

| 端点        | 方法 | 描述                      |
| ----------- | ---- | ------------------------- |
| `/version`  | GET  | 返回应用版本信息          |
| `/health`   | GET  | 健康检查接口，返回"ok"    |
| `/:message` | GET  | 根据 URL 参数返回欢迎消息 |

## 构建

```bash
# 运行测试
go test

# 构建测试版本
make build-test

# 构建生产版本
make build-prod
```

## 部署

```bash
# 使用Kubernetes部署
kubectl apply -f deployment.yml
```

## CI 集成

本项目支持多种 CI 系统：

- Drone
- Gitlab CI
- Jenkins

## 许可证

MIT
