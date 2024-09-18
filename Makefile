# makefile

# 项目名称
PROJECT_NAME := cgncode

# 执行文件
EXECUTABLE := $(PROJECT_NAME)

# go 项目目录
SRC_DIR := $(CURRENT_DIR)

# 参数
CONFIG_FILE := ./config.yaml

# Swagger 命令
SWAG := swag

# Go 命令
GO := go

# 帮助命令
.PHONY: help
help:
	$(GO) run main.go help

# 生成代码
.PHONY: code
code:
	$(GO) run main.go code

# 生成handler代码
.PHONY: h
h:
	$(GO) run main.go handler

# 生成service代码
.PHONY: s
s:
	$(GO) run main.go service

# 生成dto代码
.PHONY: d
d:
	$(GO) run main.go dto
