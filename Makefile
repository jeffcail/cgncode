# makefile

# 项目名称
PROJECT_NAME := energy_pledge

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

# 生成代码
.PHONY: code
code:
	$(GO) run main.go code