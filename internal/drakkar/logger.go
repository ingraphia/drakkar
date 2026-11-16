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

package drakkar

import (
	"io"
	"log/slog"
	"os"

	"github.com/ingraphia/drakkar/internal/config"
)

// newLogger
func newLogger(w io.Writer, level string) (*slog.Logger, error) {
	lvl, err := parseLevel(level)
	if err != nil {
		// Fallback
		// Level: INFO
		lvl = slog.Level(0)
	}

	h := slog.NewJSONHandler(w, &slog.HandlerOptions{Level: lvl})
	logger := slog.New(h)

	return logger, err
}

// parseLevel
func parseLevel(level string) (slog.Level, error) {
	var lvl slog.Level
	err := lvl.UnmarshalText([]byte(level))

	return lvl, err
}

// SetupLoggers
func SetupLoggers(c *config.Config) {
	system, err := newLogger(os.Stdout, c.LogLevel())
	if err != nil {
		system.Warn("Provided invalid log level, fallback to INFO", "provided", c.LogLevel(), "error", err)
	}

	slog.SetDefault(system)
}
