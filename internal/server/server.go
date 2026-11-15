// Copyright 2026 InGraphia authors
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

package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ingraphia/drakkar/internal/config"
)

type Server struct {
	Conf *config.Config
}

func New(c *config.Config) *Server {
	return &Server{
		Conf: c,
	}
}

func (s *Server) Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Drakkar!")
	})

	addr := ":" + strconv.FormatUint(uint64(s.Conf.Port()), 10)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println("Error: ", err)
	}
}
