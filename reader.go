// Copyright (c) 2020 Cisco and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jsonio

import (
	"bytes"
	"encoding/json"
	"io"
)

// JSONReader - json.Encode output is read from JSONReader.Read
type JSONReader interface {
	io.Reader
	// Encode - same as json.Encoder.Encode
	Encode(v interface{}) error
	// SetIndent - same as json.Encoder.SetIndent
	SetIndent(prefix, indent string)
	// SetEscapeHTML - same as json.Encoder.SetEscapeHTML
	SetEscapeHTML(on bool)
}

type reader struct {
	*json.Encoder
	io.Reader
}

// Reader - creates a new JSONReader
func Reader() JSONReader {
	buffer := bytes.NewBuffer([]byte{})
	return &reader{Encoder: json.NewEncoder(buffer), Reader: buffer}
}
