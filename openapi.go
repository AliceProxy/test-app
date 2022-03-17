// Copyright 2019 Philip Lombardi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"log"
	"net/http"
)

var openapiDocument = `
{
	"openapi": "3.0.0",
	"info": {
		"title": "Quote Service API",
		"description": "Quote Service API",
		"version": "0.1.0"
	},
	"servers": [
		{
			"url": "http://api.example.com"
		}
	],
	"paths": {
		"/": {
			"get": {
				"summary": "Return a randomly selected quote.",
				"responses": {
					"200": {
						"description": "A JSON object with a quote and some additional metadata.",
						"content": {
							"application/json": {
								"schema": {
									"type": "object",
									"properties": {
										"server": {"type": "string"},
										"quote": {"type": "string"},
										"time": {"type": "string"}
									}
								}
							}
						}
					}
				}
			}
		},
		"/debug/": {
			"get": {
				"summary": "Return debug information about the request.",
				"responses": {
					"200": {
						"description": "A JSON object with debug information about the request and additional metadata.",
						"content": {
							"application/json" : {
								"schema": {
									"type": "object",
									"properties": {
										"server": {"type": "string"},
										"time": {"type": "string"},
										"host": {"type": "string"},
										"proto": {"type": "string"},
										"url":  {"type": "object"},
										"remoteaddr": {"type": "string"},
										"headers": {"type": "object"}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
`

func (s *Server) GetOpenAPIDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte(openapiDocument)); err != nil {
		log.Panicln(err)
	}
}
