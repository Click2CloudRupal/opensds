package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import "encoding/json"

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "OpenSDS",
    "version": "v1"
  },
  "paths": {
    "/api/": {
      "get": {
        "description": "get available API versions",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "schemes": [
          "http"
        ],
        "operationId": "ListVersions",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/AvailableVersions"
            }
          }
        }
      }
    },
    "/api/v1": {
      "get": {
        "description": "show API version",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "schemes": [
          "http"
        ],
        "operationId": "GetVersionv1",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/VersionInfo"
            }
          }
        }
      }
    },
    "/api/v1/shares": {
      "get": {
        "description": "list share backend resource types",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "schemes": [
          "http"
        ],
        "operationId": "ListShareResources",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {}
          }
        }
      }
    },
    "/api/v1/shares/{resourceType}": {
      "get": {
        "description": "list shares in specified backend resource",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "schemes": [
          "http"
        ],
        "operationId": "ListShares",
        "parameters": [
          {
            "type": "string",
            "description": "Type of specified share backend resource",
            "name": "resourceType",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ShareResponse"
              }
            }
          }
        }
      },
      "post": {
        "description": "create a share in specified backend resource",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "schemes": [
          "http"
        ],
        "operationId": "CreateShare",
        "parameters": [
          {
            "type": "string",
            "description": "Type of specified share backend resource",
            "name": "resourceType",
            "in": "path",
            "required": true
          },
          {
            "description": "Share request object",
            "name": "shareRequest",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ShareRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/ShareResponse"
            }
          }
        }
      }
    },
    "/api/v1/shares/{resourceType}/{id}": {
      "get": {
        "description": "get specified share in specified backend resource",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "schemes": [
          "http"
        ],
        "operationId": "GetShare",
        "parameters": [
          {
            "type": "string",
            "description": "Type of specified share backend resource",
            "name": "resourceType",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "ID of specified share",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/ShareDetailResponse"
            }
          }
        }
      },
      "put": {
        "description": "update specified share in specified backend resource",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "schemes": [
          "http"
        ],
        "operationId": "UpdateShare",
        "parameters": [
          {
            "type": "string",
            "description": "Type of specified share backend resource",
            "name": "resourceType",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "ID of specified share",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Share request object",
            "name": "shareRequest",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ShareRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/ShareResponse"
            }
          }
        }
      },
      "delete": {
        "description": "delete specified share in specified backend resource",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "schemes": [
          "http"
        ],
        "operationId": "DeleteShare",
        "parameters": [
          {
            "type": "string",
            "description": "Type of specified share backend resource",
            "name": "resourceType",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "ID of specified share",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/DefaultResponse"
            }
          }
        }
      }
    },
    "/api/v1/volumes": {
      "get": {
        "description": "list volume backend resource types",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "schemes": [
          "http"
        ],
        "operationId": "ListVolumeResources",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {}
          }
        }
      }
    },
    "/api/v1/volumes/action/{resourceType}/{id}": {
      "post": {
        "description": "opreate specified volume in specified backend resource",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "schemes": [
          "http"
        ],
        "operationId": "OperateVolume",
        "parameters": [
          {
            "type": "string",
            "description": "Type of specified volume backend resource",
            "name": "resourceType",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "ID of specified volume",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Volume request object",
            "name": "volumeRequest",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/VolumeRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/DefaultResponse"
            }
          }
        }
      }
    },
    "/api/v1/volumes/{resourceType}": {
      "get": {
        "description": "list volumes in specified backend resource",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "schemes": [
          "http"
        ],
        "operationId": "ListVolumes",
        "parameters": [
          {
            "type": "string",
            "description": "Type of specified volume backend resource",
            "name": "resourceType",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/VolumeResponse"
              }
            }
          }
        }
      },
      "post": {
        "description": "create a volume in specified backend resource",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "schemes": [
          "http"
        ],
        "operationId": "CreateVolume",
        "parameters": [
          {
            "type": "string",
            "description": "Type of specified volume backend resource",
            "name": "resourceType",
            "in": "path",
            "required": true
          },
          {
            "description": "Volume request object",
            "name": "volumeRequest",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/VolumeRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/VolumeResponse"
            }
          }
        }
      }
    },
    "/api/v1/volumes/{resourceType}/{id}": {
      "get": {
        "description": "get specified volume in specified backend resource",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "schemes": [
          "http"
        ],
        "operationId": "GetVolume",
        "parameters": [
          {
            "type": "string",
            "description": "Type of specified volume backend resource",
            "name": "resourceType",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "ID of specified volume",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/VolumeDetailResponse"
            }
          }
        }
      },
      "put": {
        "description": "update specified volume in specified backend resource",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "schemes": [
          "http"
        ],
        "operationId": "UpdateVolume",
        "parameters": [
          {
            "type": "string",
            "description": "Type of specified volume backend resource",
            "name": "resourceType",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "ID of specified volume",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Volume request object",
            "name": "volumeRequest",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/VolumeRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/VolumeResponse"
            }
          }
        }
      },
      "delete": {
        "description": "delete specified volume in specified backend resource",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "schemes": [
          "http"
        ],
        "operationId": "DeleteVolume",
        "parameters": [
          {
            "type": "string",
            "description": "Type of specified volume backend resource",
            "name": "resourceType",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "ID of specified volume",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/DefaultResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "AvailableVersions": {
      "type": "object",
      "properties": {
        "versions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/VersionInfo"
          }
        }
      }
    },
    "DefaultResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "message": {
          "type": "string"
        },
        "status": {
          "type": "string"
        }
      }
    },
    "ShareDetailResponse": {
      "type": "object",
      "properties": {
        "accessRulesStatus": {
          "type": "string"
        },
        "availabilityZone": {
          "type": "string"
        },
        "consistencyGroupId": {
          "type": "string"
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "description": {
          "type": "string"
        },
        "exportLocation": {
          "type": "string"
        },
        "exportLocations": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "hasReplicas": {
          "type": "boolean",
          "default": false
        },
        "host": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "isPublic": {
          "type": "boolean",
          "default": false
        },
        "name": {
          "type": "string"
        },
        "projectId": {
          "type": "string"
        },
        "replicationType": {
          "type": "string"
        },
        "resourceType": {
          "type": "string"
        },
        "shareNetworkId": {
          "type": "string"
        },
        "shareProto": {
          "type": "string"
        },
        "shareServerId": {
          "type": "string"
        },
        "shareType": {
          "type": "string"
        },
        "shareTypeName": {
          "type": "string"
        },
        "size": {
          "type": "integer",
          "format": "int32"
        },
        "snapshotSupport": {
          "type": "boolean",
          "default": false
        },
        "sourceCgsnapshotMemberId": {
          "type": "string"
        },
        "status": {
          "type": "string"
        },
        "taskState": {
          "type": "string"
        },
        "volumeType": {
          "type": "string"
        }
      }
    },
    "ShareRequest": {
      "type": "object",
      "properties": {
        "allowDetails": {
          "type": "boolean",
          "default": false
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "resourceType": {
          "type": "string"
        },
        "shareProto": {
          "type": "string"
        },
        "shareType": {
          "type": "string"
        },
        "size": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "ShareResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "resourceType": {
          "type": "string"
        }
      }
    },
    "VersionInfo": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "status": {
          "type": "string"
        },
        "updateAt": {
          "type": "string"
        }
      }
    },
    "VolumeDetailResponse": {
      "type": "object",
      "properties": {
        "availabilityZone": {
          "type": "string"
        },
        "consistencygroupId": {
          "type": "string"
        },
        "createAt": {
          "type": "string",
          "format": "date-time"
        },
        "description": {
          "type": "string"
        },
        "encrypted": {
          "type": "boolean",
          "default": false
        },
        "id": {
          "type": "string"
        },
        "migrationStatus": {
          "type": "string"
        },
        "multiattach": {
          "type": "boolean",
          "default": false
        },
        "name": {
          "type": "string"
        },
        "protected": {
          "type": "boolean",
          "default": false
        },
        "replicationStatus": {
          "type": "string"
        },
        "resourceType": {
          "type": "string"
        },
        "size": {
          "type": "integer",
          "format": "int32"
        },
        "snapshotId": {
          "type": "string"
        },
        "sourceVolid": {
          "type": "string"
        },
        "status": {
          "type": "string"
        },
        "userId": {
          "type": "string"
        },
        "volumeType": {
          "type": "string"
        }
      }
    },
    "VolumeRequest": {
      "type": "object",
      "properties": {
        "actionType": {
          "type": "string"
        },
        "allowDetails": {
          "type": "boolean",
          "default": false
        },
        "attachment": {
          "type": "string"
        },
        "device": {
          "type": "string"
        },
        "fsType": {
          "type": "string"
        },
        "host": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "mountDir": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "resourceType": {
          "type": "string"
        },
        "size": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "VolumeResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "resourceType": {
          "type": "string"
        },
        "size": {
          "type": "integer",
          "format": "int32"
        },
        "status": {
          "type": "string"
        },
        "volumeType": {
          "type": "string"
        }
      }
    }
  }
}`))
}
