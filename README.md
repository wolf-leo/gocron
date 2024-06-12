# gocron - 总控任务调度 + 定时任务管理系统

> 此仓库为 fork 原始仓库后升级改良版
>
> 原始地址 [https://github.com/ouqiang/gocron](https://github.com/ouqiang/gocron)

# 项目简介

使用Go语言开发的轻量级定时任务集中调度和管理系统, 用于替代Linux-crontab [查看文档](https://github.com/ouqiang/gocron/wiki)

原有的延时任务拆分为独立项目[延迟队列](https://github.com/ouqiang/delay-queue)

## 功能特性

* Web界面管理定时任务
* crontab时间表达式, 精确到秒
* 任务执行失败可重试
* 任务执行超时, 强制结束
* 任务依赖配置, A任务完成后再执行B任务
* 账户权限控制
* 任务类型
    * shell任务
  > 在任务节点上执行shell命令, 支持任务同时在多个节点上运行
    * HTTP任务
  > 访问指定的URL地址, 由调度器直接执行, 不依赖任务节点
* 查看任务执行结果日志
* 任务执行结果通知, 支持邮件、Slack、Webhook

### 截图

![流程图](/assets/screenshot/scheduler.png)
![任务](/assets/screenshot/task.png)
![Slack](/assets/screenshot/notification.png)

### 支持平台

> Windows、Linux、Mac OS

### 环境要求

> MySQL

## 下载

请下载源码后自行go build

## 安装

### 二进制安装

1. 目录 cmd 下
2. 启动

* 调度器启动:
    * cd cmd/gocron
    * Windows、Linux、Mac OS:  `go run ./gocron.go web`
    * 浏览器访问 http://localhost:5920

* 任务节点启动, 默认监听0.0.0.0:5921
    * cd cmd/node
    * Windows、Linux、Mac OS:  `go run ./node.go`

### 开发

1. 安装Go1.11+, Node.js, Yarn
2. 安装前端依赖 `yarn insall`
3. 启动gocron, gocron-node `yarn run dev`
4. 访问地址 http://localhost:8080

访问http://localhost:8080, API请求会转发给gocron

### 命令

* gocron
    * -v 查看版本

* gocron web
    * --host 默认0.0.0.0
    * -p 端口, 指定端口, 默认5920
    * -e 指定运行环境, dev|test|prod, dev模式下可查看更多日志信息, 默认prod
    * -h 查看帮助
* gocron-node
    * -allow-root *nix平台允许以root用户运行
    * -s ip:port 监听地址
    * -enable-tls 开启TLS
    * -ca-file CA证书文件* -cert-file 证书文件
    * -key-file 私钥文件
    * -h 查看帮助
    * -v 查看版本

## To Do List

- [x] 版本升级
- [x] 批量开启、关闭、删除任务
- [x] 调度器与任务节点通信支持https
- [x] 任务分组
- [x] 多用户
- [x] 权限控制

## 程序使用的组件

* Web框架 [Macaron](http://go-macaron.com/)
* 定时任务调度 [Cron](https://github.com/robfig/cron)
* ORM [Xorm](https://github.com/go-xorm/xorm)
* UI框架 [Element UI](https://github.com/ElemeFE/element)
* 依赖管理 [Govendor](https://github.com/kardianos/govendor)
* RPC框架 [gRPC](https://github.com/grpc/grpc)

## ChangeLog

v2.0
原作者多年未更新，拿来稍微改了下

--------

* 任务可同时在多个节点上运行
* *nix平台默认禁止以root用户运行任务节点
* 子任务命令中增加预定义占位符, 子任务可根据主任务运行结果执行相应操作
* 删除守护进程模块
* Web访问日志输出到终端
