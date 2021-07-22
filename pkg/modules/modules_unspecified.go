// Copyright 2021 Authors of Cilium
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

// +build !linux

package modules

import (
	"errors"
	"io"
)

var ErrNotImplemented = errors.New("not implemented")

func listModules() ([]string, error) {
	return nil, ErrNotImplemented
}

func moduleLoader() string {
	return "unknown-module-loader"
}

func parseModulesFile(r io.Reader) ([]string, error) {
	return nil, ErrNotImplemented
}
