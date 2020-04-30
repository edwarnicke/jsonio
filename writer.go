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

// Package jsonio providers io.Reader2 and io.Writer2 wrappers for json
package jsonio

import (
	"bytes"
	"encoding/json"
	"io"
)

// JSONWriter - JSONWriter.Write is decoded by JSONWriter.Decode
type JSONWriter interface {
	io.Writer
	// Buffered - same as json.Decoder.Buffered
	Buffered() io.Reader
	// Decode - same as json.Decoder.Decode
	Decode(v interface{}) error
	// DisallowUnknownFields - same as json.Decoder.DisallowUnknownFields
	DisallowUnknownFields()
	// More - same as json.Decoder.More
	More() bool
	// Token - same as json.Decoder.Token
	Token() (json.Token, error)
	// UseNumber - same as json.Decoder.UseNumber
	UseNumber()
}

type writer struct {
	*json.Decoder
	io.Writer
}

// Writer - creates a new JSONWriter
func Writer() JSONWriter {
	buffer := bytes.NewBuffer([]byte{})
	return &writer{Decoder: json.NewDecoder(buffer), Writer: buffer}
}
