// Copyright 2020 The Ebiten Authors
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

// +build ignore

package main

var Time float
var Cursor vec2
var ScreenSize vec2

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	center := ScreenSize / 2
	// As texel coodinates should be image0's texture texels, image0TextureSize should be used.
	// TODO: This seems too tricky. Improve the API.
	amount := (center - Cursor) / image0TextureSize() / 10
	var clr vec3
	clr.r = image2TextureBoundsAt(texCoord + amount).r
	clr.g = image2TextureAt(texCoord).g
	clr.b = image2TextureBoundsAt(texCoord - amount).b
	return vec4(clr, 1.0)
}
