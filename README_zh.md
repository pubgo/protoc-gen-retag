# protoc-gen-retag

[![Go Reference](https://pkg.go.dev/badge/github.com/pubgo/protoc-gen-retag.svg)](https://pkg.go.dev/github.com/pubgo/protoc-gen-retag)
[![CI](https://github.com/pubgo/protoc-gen-retag/actions/workflows/ci.yml/badge.svg)](https://github.com/pubgo/protoc-gen-retag/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/pubgo/protoc-gen-retag)](https://goreportcard.com/report/github.com/pubgo/protoc-gen-retag)

[English](README.md)

一个 protoc 插件，用于在生成的 protobuf Go 文件中添加自定义 struct tags。

## 功能特性

- 为 protobuf 生成的 Go 结构体添加自定义 struct tags（如 `graphql`、`xml`、`validate` 等）
- 修改已有的 tags（如覆盖默认的 `json` tag）
- 支持 `oneof` 字段
- 保留原有的 protobuf/json tags，同时添加新的 tags

## 安装

```bash
go install github.com/pubgo/protoc-gen-retag@latest
```

## 使用方法

### 1. 导入 retag proto 文件

```proto
import "proto/retag/retag.proto";
```

### 2. 为字段添加自定义 tags

```proto
syntax = "proto3";

package example;

import "proto/retag/retag.proto";

message User {
    // 添加新的 graphql tag
    string name = 1 [
        (retag.tags) = {name: "graphql", value: "userName,optional"}
    ];

    // 添加多个自定义 tags
    string email = 2 [
        (retag.tags) = {name: "graphql", value: "userEmail"},
        (retag.tags) = {name: "validate", value: "required,email"}
    ];

    // 覆盖默认的 json tag
    string user_id = 3 [
        (retag.tags) = {name: "json", value: "id,omitempty"}
    ];

    // 支持 oneof 字段
    oneof contact {
        option (retag.oneof_tags) = {name: "graphql", value: "contact,optional"};
        string phone = 4;
        string address = 5;
    }
}
```

### 3. 生成 Go 代码

```bash
protoc --go_out=. --retag_out=. your.proto
```

或者使用 buf：

```yaml
# buf.gen.yaml
version: v1
plugins:
  - plugin: go
    out: gen
    opt: paths=source_relative
  - plugin: retag
    out: gen
    opt: paths=source_relative
```

## 生成结果

生成的 Go 结构体将包含你的自定义 tags：

```go
type User struct {
    Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" graphql:"userName,optional"`
    Email  string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty" graphql:"userEmail" validate:"required,email"`
    UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"id,omitempty"`
    // ...
}
```

## Tag 选项

| 扩展 | 作用域 | 描述 |
|-----|-------|------|
| `retag.tags` | 字段 | 在 message 字段上添加/修改 tags |
| `retag.oneof_tags` | Oneof | 在 oneof 字段上添加 tags |

## 致谢

本项目的灵感来源于 [protoc-gen-go-tag](https://github.com/searKing/golang/tree/master/tools/protoc-gen-go-tag)。

## 许可证

MIT License
