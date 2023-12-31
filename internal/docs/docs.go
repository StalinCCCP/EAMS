// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "tags": [
                    "公共方法"
                ],
                "summary": "根据用户名和密码登录，并返回token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "Username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "Pwd",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"} data中包含了用户token",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "tags": [
                    "公共方法"
                ],
                "summary": "输入用户名、密码、电子邮箱、手机号来创建用户（默认创建普通用户，提权在其他地方实现），并返回用户token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "Username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "Pwd",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "全名",
                        "name": "Full_name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "电子邮箱",
                        "name": "Email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "手机号",
                        "name": "Phone_number",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"} data中包含了用户token",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user-detail": {
            "get": {
                "tags": [
                    "公共方法"
                ],
                "summary": "根据用户ID找到用户，显示用户详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "User_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"} data中包含了用户ID、用户名、用户密码、用户角色、全名、电子邮箱、电话号码和注册日期",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
