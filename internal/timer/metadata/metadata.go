// Copyright 2022 Linkall Inc.
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

package metadata

type PointerMeta struct {
	Layer   int64 `json:"layer"`
	Pointer int64 `json:"pointer"`
}

type OffsetMeta struct {
	Layer    int64  `json:"layer"`
	Slot     int64  `json:"slot"`
	Offset   int64  `json:"offset"`
	Eventbus string `json:"eventbus"`
}
