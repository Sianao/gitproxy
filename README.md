# :tada: git proxy  :tada:
> :bookmark:
>
> 一款高颜值 高性能的github代理软件   支持二进制文件下载 支持git clone代理  ui界面友好  交互简单

## 支持的文件形式

>https://github.com/laurent22/joplin/releases/download/v3.1.24/Joplin-3.1.24-arm64.DMG
>https://github.com/laurent22/joplin/archive/refs/tags/v3.1.24.zip
>https://github.com/laurent22/joplin/blob/dev/.gitignore
>https://raw.githubusercontent.com/laurent22/joplin/dev/.gitignore
>https://raw.githubusercontent.com/laurent22/joplin/e652db05e1ba47725249a6ff543628aeeb32fad7/.gitignore
>https://raw.githubusercontent.com/laurent22/joplin/android-v3.2.2/.gitignore

## 部署

Dockerfile:

直接使用Dockerfile进行打包

```bash
docker build . -t gitproxy:latest
```

裸机部署：

按照dockerfile编译规则 首先编译静态文件 再编译`main.go`入口 运行 `go run main.go`
