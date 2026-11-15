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

package app

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/ingraphia/drakkar/internal/config"
	"github.com/ingraphia/drakkar/internal/server"
)

func Start() {
	c := config.Setup()
	api := *server.New(c)

	go func() {
		fmt.Println("Drakkar server starting...")
		api.Run()
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()
    defer fmt.Println("\nDrakkar server shutdown...")

	<-ctx.Done()
}
