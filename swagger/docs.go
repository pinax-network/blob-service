// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

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
        "/blobs/by_slot/{slot}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blobs"
                ],
                "summary": "Get Blobs by slot number",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Slot Number",
                        "name": "slot",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.ApiDataResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/controllers.BlobsResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "No blobs in this slot",
                        "schema": {
                            "$ref": "#/definitions/response.ApiErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ApiErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.BlobsResponse": {
            "type": "object",
            "properties": {
                "blobs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/pbbmsrv.Blob"
                    }
                }
            }
        },
        "pbbmsrv.Blob": {
            "type": "object",
            "properties": {
                "blob": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "extra": {
                    "$ref": "#/definitions/pbbmsrv.Extra"
                },
                "index": {
                    "type": "integer"
                },
                "kzg_commitment": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "kzg_commitment_inclusion_proof": {
                    "type": "array",
                    "items": {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        }
                    }
                },
                "kzg_proof": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "signed_block_header": {
                    "$ref": "#/definitions/pbbmsrv.SignedBlockHeader"
                }
            }
        },
        "pbbmsrv.Extra": {
            "type": "object",
            "properties": {
                "block_number": {
                    "type": "integer"
                },
                "timestamp": {
                    "$ref": "#/definitions/timestamppb.Timestamp"
                }
            }
        },
        "pbbmsrv.Message": {
            "type": "object",
            "properties": {
                "body_root": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "parent_root": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "proposer_index": {
                    "type": "integer"
                },
                "slot": {
                    "type": "integer"
                },
                "state_root": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "pbbmsrv.SignedBlockHeader": {
            "type": "object",
            "properties": {
                "message": {
                    "$ref": "#/definitions/pbbmsrv.Message"
                },
                "signature": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "response.ApiDataResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "meta": {}
            }
        },
        "response.ApiError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "detail": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "response.ApiErrorResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.ApiError"
                    }
                }
            }
        },
        "timestamppb.Timestamp": {
            "type": "object",
            "properties": {
                "nanos": {
                    "description": "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.",
                    "type": "integer"
                },
                "seconds": {
                    "description": "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.",
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{"http", "https"},
	Title:            "Ethereum Consensus Layer Blobs REST API",
	Description:      "This is the API for Ethereum Blobs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
