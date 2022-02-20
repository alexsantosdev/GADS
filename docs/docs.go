// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/configuration/build-image": {
            "post": {
                "description": "Starts building the 'ios-appium' image in a goroutine and just returns Accepted",
                "tags": [
                    "configuration"
                ],
                "summary": "Build 'ios-appium' image",
                "responses": {
                    "202": {
                        "description": ""
                    }
                }
            }
        },
        "/configuration/remove-image": {
            "post": {
                "description": "Removes the 'ios-appium' Docker image",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "configuration"
                ],
                "summary": "Remove 'ios-appium' image",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.SimpleResponseJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorJSON"
                        }
                    }
                }
            }
        },
        "/configuration/remove-ios-listener": {
            "post": {
                "description": "Deletes udev rules from /etc/udev/rules.d and reloads udev",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "configuration"
                ],
                "summary": "Removes iOS device listener",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.SimpleResponseJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorJSON"
                        }
                    }
                }
            }
        },
        "/configuration/set-sudo-password": {
            "put": {
                "description": "Sets your sudo password in ./env.json. The password is needed for operations requiring elevated permissions like setting up udev.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "configuration"
                ],
                "summary": "Set sudo password",
                "parameters": [
                    {
                        "description": "Sudo password value",
                        "name": "config",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.SudoPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.SimpleResponseJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorJSON"
                        }
                    }
                }
            }
        },
        "/configuration/setup-ios-listener": {
            "post": {
                "description": "Creates udev rules, moves them to /etc/udev/rules.d and reloads udev. Copies usbmuxd.service to /lib/systemd/system and enables it",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "configuration"
                ],
                "summary": "Sets up iOS device listener",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.SimpleResponseJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorJSON"
                        }
                    }
                }
            }
        },
        "/configuration/update-config": {
            "put": {
                "description": "Updates one  or multiple configuration values",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "configuration"
                ],
                "summary": "Update project configuration",
                "parameters": [
                    {
                        "description": "Update config",
                        "name": "config",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.ProjectConfig"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.SimpleResponseJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorJSON"
                        }
                    }
                }
            }
        },
        "/configuration/upload-wda": {
            "post": {
                "description": "Uploads the provided *.ipa into the ./apps folder with the expected \"WebDriverAgent.ipa\" name",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "configuration"
                ],
                "summary": "Upload WDA",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.SimpleResponseJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorJSON"
                        }
                    }
                }
            }
        },
        "/containers/{container_id}/logs": {
            "get": {
                "description": "Get logs of container by providing container ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "containers"
                ],
                "summary": "Get container logs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Container ID",
                        "name": "container_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.SimpleResponseJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorJSON"
                        }
                    }
                }
            }
        },
        "/containers/{container_id}/remove": {
            "post": {
                "description": "Removes container by provided container ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "containers"
                ],
                "summary": "Remove container",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Container ID",
                        "name": "container_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.SimpleResponseJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorJSON"
                        }
                    }
                }
            }
        },
        "/containers/{container_id}/restart": {
            "post": {
                "description": "Restarts container by provided container ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "containers"
                ],
                "summary": "Restart container",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Container ID",
                        "name": "container_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.SimpleResponseJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorJSON"
                        }
                    }
                }
            }
        },
        "/device-containers/{device_udid}/create": {
            "post": {
                "description": "Creates a container for a connected registered device",
                "tags": [
                    "device-containers"
                ],
                "summary": "Create container for device",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device UDID",
                        "name": "device_udid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": ""
                    }
                }
            }
        },
        "/device-containers/{device_udid}/remove": {
            "post": {
                "description": "Removes a running container for a disconnected registered device by device UDID",
                "tags": [
                    "device-containers"
                ],
                "summary": "Remove container for device",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device UDID",
                        "name": "device_udid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": ""
                    }
                }
            }
        },
        "/device-logs/{log_type}/{device_udid}": {
            "get": {
                "description": "Get logs by type",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "device-logs"
                ],
                "summary": "Get logs for iOS device container",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Log Type",
                        "name": "log_type",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Device UDID",
                        "name": "device_udid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.SimpleResponseJSON"
                        }
                    }
                }
            }
        },
        "/devices/device-control": {
            "post": {
                "description": "Provides the running containers, IOS devices info and apps available for installing",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "devices"
                ],
                "summary": "Get device control info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.DeviceControlInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorJSON"
                        }
                    }
                }
            }
        },
        "/ios-devices": {
            "get": {
                "description": "Returns the connected iOS devices with go-ios",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ios-devices"
                ],
                "summary": "Get connected iOS devices",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.detailsList"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorJSON"
                        }
                    }
                }
            }
        },
        "/ios-devices/register": {
            "post": {
                "description": "Registers a new iOS device in config.json",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ios-devices"
                ],
                "summary": "Register a new iOS device",
                "parameters": [
                    {
                        "description": "Register iOS device",
                        "name": "config",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.registerIOSDevice"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.SimpleResponseJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorJSON"
                        }
                    }
                }
            }
        },
        "/ios-devices/{device_udid}/install-app": {
            "post": {
                "description": "Installs *.ipa or *.app from the './apps' folder with go-ios",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ios-devices"
                ],
                "summary": "Install app on iOS device",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device UDID",
                        "name": "device_udid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Install iOS app",
                        "name": "config",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.iOSAppInstall"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.SimpleResponseJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorJSON"
                        }
                    }
                }
            }
        },
        "/ios-devices/{device_udid}/uninstall-app": {
            "post": {
                "description": "Uninstalls app from iOS device by provided bundleID with go-ios",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ios-devices"
                ],
                "summary": "Uninstall app from iOS device",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device UDID",
                        "name": "device_udid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Uninstall iOS app",
                        "name": "config",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.iOSAppUninstall"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.SimpleResponseJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorJSON"
                        }
                    }
                }
            }
        },
        "/project-logs": {
            "get": {
                "description": "Provides project logs as plain text response",
                "tags": [
                    "project-logs"
                ],
                "summary": "Get project logs",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "main.AndroidDeviceInfo": {
            "type": "object",
            "properties": {
                "deviceConfig": {
                    "$ref": "#/definitions/main.IOSDevice"
                },
                "installedAppsBundleIDs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "main.DeviceControlInfo": {
            "type": "object",
            "properties": {
                "android-devices-info": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.AndroidDeviceInfo"
                    }
                },
                "installable-apps": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "ios-devices-info": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.IOSDeviceInfo"
                    }
                },
                "running-containers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "main.ErrorJSON": {
            "type": "object",
            "properties": {
                "error_message": {
                    "type": "string"
                },
                "event": {
                    "type": "string"
                }
            }
        },
        "main.IOSDevice": {
            "type": "object",
            "properties": {
                "appium_port": {
                    "type": "integer"
                },
                "device_model": {
                    "type": "string"
                },
                "device_name": {
                    "type": "string"
                },
                "device_os_version": {
                    "type": "string"
                },
                "device_udid": {
                    "type": "string"
                },
                "viewport_size": {
                    "type": "string"
                },
                "wda_mjpeg_port": {
                    "type": "integer"
                },
                "wda_port": {
                    "type": "integer"
                },
                "wda_stream_url": {
                    "type": "string"
                },
                "wda_url": {
                    "type": "string"
                }
            }
        },
        "main.IOSDeviceInfo": {
            "type": "object",
            "properties": {
                "deviceConfig": {
                    "$ref": "#/definitions/main.IOSDevice"
                },
                "installedAppsBundleIDs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "main.ProjectConfig": {
            "type": "object",
            "properties": {
                "devices_host": {
                    "type": "string"
                },
                "selenium_hub_host": {
                    "type": "string"
                },
                "selenium_hub_port": {
                    "type": "string"
                },
                "selenium_hub_protocol_type": {
                    "type": "string"
                },
                "wda_bundle_id": {
                    "type": "string"
                }
            }
        },
        "main.SimpleResponseJSON": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "main.SudoPassword": {
            "type": "object",
            "properties": {
                "sudo_password": {
                    "type": "string"
                }
            }
        },
        "main.detailsEntry": {
            "type": "object",
            "properties": {
                "productName": {
                    "type": "string"
                },
                "productType": {
                    "type": "string"
                },
                "productVersion": {
                    "type": "string"
                },
                "udid": {
                    "type": "string"
                }
            }
        },
        "main.detailsList": {
            "type": "object",
            "properties": {
                "deviceList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.detailsEntry"
                    }
                }
            }
        },
        "main.iOSAppInstall": {
            "type": "object",
            "properties": {
                "ipa_name": {
                    "type": "string"
                }
            }
        },
        "main.iOSAppUninstall": {
            "type": "object",
            "properties": {
                "bundle_id": {
                    "type": "string"
                }
            }
        },
        "main.registerIOSDevice": {
            "type": "object",
            "properties": {
                "device_name": {
                    "type": "string"
                },
                "device_os_version": {
                    "type": "string"
                },
                "device_udid": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
