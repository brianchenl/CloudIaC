CloudIaC
================
> Cloud Infrastructure as Code

CloudIaC 是基于基础设施即代码构建的云环境自动化管理平台。
CloudIaC 将易于使用的界面与强大的治理工具相结合，让您和您团队的成员可以快速轻松的在云中部署和管理环境。

通过将 CloudIaC 集成到您的流程中，您可以获得对组织的云使用情况的可见性、可预测性和更好的治理。

- [快速入门](docs/mkdocs/user-guide/quick-start.md)
- [产品文档](docs/mkdocs/intra/)
- [容器化部署](docs/mkdocs/deploy/container.md)
- [文档索引](docs/mkdocs/)

## 编译
依赖 go 1.16 及以上版本

```
go get -x github.com/swaggo/swag/cmd/swag
go mod download -x
make build
```

*如果下载较慢可以设置 GOPROXY:*
```bash
go env -w GOPROXY="https://goproxy.io,direct"
```

