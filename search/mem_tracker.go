//  Copyright (c) 2018 Couchbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package search

import ()

type MemTracker struct {
	bytes uint64
}

func NewMemTracker() *MemTracker {
	return &MemTracker{}
}

func (m *MemTracker) Add(add uint64) {
	m.bytes += add
}

func (m *MemTracker) Usage() uint64 {
	return m.bytes
}
