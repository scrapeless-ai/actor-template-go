# Actor Template (Go)

## 项目描述
这是一个基于 Go 语言的 Actor 模式模板项目，用于快速启动支持 Actor 模型的服务开发。项目包含基础配置、Docker 支持及输入输出规范。

## 前置要求
- Go 1.20+（[安装指南](https://go.dev/doc/install)）

## 安装与运行
### 本地运行
1. 克隆仓库：
```bash
git clone https://github.com/scrapeless-ai/actor-template-go.git
cd actor-template-go
```
2. 安装依赖：
```bash
go mod tidy
```
3. 启动服务（使用示例输入）：
```bash
go run main.go
```

## 配置说明
- `.env.example`：环境变量模板，重命名为 `.env` 并根据需要修改
- `.actor/actor.json`：Actor 元数据配置（名称、版本等）
- `.actor/input_schema.json`：输入数据校验模式

## 贡献与反馈
欢迎提交 Issues 或 Pull Requests 参与改进！

## 许可协议
本项目采用 [MIT 许可协议](LICENSE)。
        