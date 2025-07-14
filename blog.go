// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embed"
	"image/color"

	"cogentcore.org/core/content"
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/htmlcore"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/tree"
	_ "cogentcore.org/lab/yaegilab"
)

//go:embed content
var econtent embed.FS

func main() {
	b := core.NewBody("Cogent Blog")
	ct := content.NewContent(b).SetContent(econtent)
	ctx := ct.Context
	ctx.AddWikilinkHandler(htmlcore.GoDocWikilink("doc", "cogentcore.org"))
	b.AddTopBar(func(bar *core.Frame) {
		tb := core.NewToolbar(bar)
		tb.Maker(ct.MakeToolbar)
		tb.Maker(func(p *tree.Plan) {
			tree.Add(p, func(w *core.Button) {
				ctx.LinkButton(w, "https://youtube.com/@CogentCore")
				w.SetText("Videos").SetIcon(icons.VideoLibrary)
			})
			tree.Add(p, func(w *core.Button) {
				ctx.LinkButton(w, "https://github.com/cogentcore")
				w.SetText("GitHub").SetIcon(icons.GitHub)
			})
			tree.Add(p, func(w *core.Button) {
				ctx.LinkButton(w, "https://cogentcore.org/community")
				w.SetText("Community").SetIcon(icons.Forum)
			})
			tree.Add(p, func(w *core.Button) {
				ctx.LinkButton(w, "https://github.com/sponsors/cogentcore")
				w.SetText("Sponsor").SetIcon(icons.Favorite)
			})
		})
	})

	ctx.ElementHandlers["color-scheme-control"] = func(ctx *htmlcore.Context) bool {
		type theme struct {
			Theme core.Themes `default:"Auto"`
			Color color.RGBA  `default:"#4285f4"`
		}
		th := &theme{core.AppearanceSettings.Theme, core.AppearanceSettings.Color}
		fm := core.NewForm(ctx.BlockParent).SetStruct(th)
		fm.OnChange(func(e events.Event) {
			core.AppearanceSettings.Theme = th.Theme
			core.AppearanceSettings.Color = th.Color
			core.UpdateSettings(ctx.BlockParent, core.AppearanceSettings)
		})
		return true
	}

	b.RunMainWindow()
}
