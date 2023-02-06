#!/bin/bash

cd `dirname $0`
#sql 辅助生成工具
#1 生成的数据库表实体struct 和基本操作，支持json和sqlx操作，文件默认后缀是 .gen.go
#2 当不满足需要时，要么升级模板，要么新加一个伴生 .man.go 文件 代表开发manual自助开发修改，里面函数通常以 XxxManual 命名
#3 绝对不要修改 .gen.go
mkdir -p ./models

xo --package "models" --suffix ".gen.go" "mysql://[username]:[password]@127.0.0.1:3306/[demo]?charset=utf8&parseTime=true" -o models --template-path templates
