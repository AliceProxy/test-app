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
	"github.com/plombardi89/gozeug/randomzeug"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_GetQuote(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	s := Server{
		quotes: []string{"A principal idea is omnipresent, much like candy."},
		random: randomzeug.NewRandom(),
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.GetQuote)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("content-type"))
}

func TestServer_GetOpenAPIDocument(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	s := Server{
		quotes: []string{},
		random: randomzeug.NewRandom(),
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.GetOpenAPIDocument)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("content-type"))
	assert.Equal(t, openapiDocument, rr.Body.String())
}
