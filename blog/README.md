# 博客系统

## prompt

```
开发一个完整的个人博客系统，该系统需包含后台登录和前台显示两大核心功能模块。技术栈要求如下：前端采用Vue框架结合Vite构建工具，并使用tailwindcss实现UI界面；后端采用Go 语言，数据库使用 sqlite存储，支持 markdown 编写博客文章，支持复制粘贴上传图片
```

## 测试模型

> 一开始打算是测多个模型的效果，后来觉得主要是为了测试 kontext 环境下的效果，所以就只测了 minimax-2.7 和 claude-opus-4-6 ，一个国内一个国外的模型效果

## 测试环境介绍

Claude Code，安装了 superpowers ，但由于 kontext 环境下提供了多轮对话澄清，所以在非 kontext 环境下也使用该 skills，本次测试主要用于测试 kontext 环境下的效果，kontext 对于开发过程的约束效果

## 产物目录

### kontext 环境下的产物

所在目录： [kontext](./kontext)，相关截图在 [screenshots](./kontext/screenshots)

生成报告：

耗时：13m左右

生成产物遇到了sqlite3库报错，整体界面UI还凑合，无法新建文章

```
[error] failed to initialize database, got error Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub
2026/04/01 00:46:07 failed to connect database: Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub
exit status 1
```

### 非 kontext 环境下的产物

所在目录： [non-kontext](./non-kontext)，相关截图在 [screenshots](./non-kontext/screenshots)
