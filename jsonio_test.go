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

package jsonio_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edwarnicke/jsonio"
)

type TestStruct struct {
	A string
	B string
}

func TestReaderToWriter(t *testing.T) {
	in := &TestStruct{A: "a", B: "b"}
	var out *TestStruct
	r := jsonio.Reader()
	w := jsonio.Writer()
	assert.NoError(t, r.Encode(&in))
	_, err := io.Copy(w, r)
	assert.NoError(t, err)
	assert.NoError(t, w.Decode(&out))
	assert.Equal(t, in, out)
}
