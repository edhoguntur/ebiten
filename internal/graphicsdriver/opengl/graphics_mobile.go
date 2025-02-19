// Copyright 2018 The Ebiten Authors
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

//go:build android || ios

package opengl

import (
	mgl "golang.org/x/mobile/gl"

	"github.com/hajimehoshi/ebiten/v2/internal/graphicsdriver"
	"github.com/hajimehoshi/ebiten/v2/internal/graphicsdriver/opengl/gl"
)

// NewGraphics creates an implementation of graphicsdriver.Graphics for OpenGL.
// The returned graphics value is nil iff the error is not nil.
func NewGraphics(context mgl.Context) (graphicsdriver.Graphics, error) {
	g := &Graphics{}
	if context != nil {
		g.context.ctx = gl.NewGomobileContext(context.(mgl.Context))
	} else {
		g.context.ctx = gl.DefaultContext{}
	}
	return g, nil
}
