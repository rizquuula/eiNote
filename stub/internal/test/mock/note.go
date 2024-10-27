package mock

import (
	"stub/internal/domain/model/note"
	"time"

	"github.com/google/uuid"
)

var titleOne = "A demo of `react-markdown`"
var contentOne = "# A demo of `react-markdown`\n\n" +
	"`react-markdown` is a markdown component for React.\n\n" +
	"👉 Changes are re-rendered as you type.\n" +
	"[![Preview](https://user-images.githubusercontent.com/3472373/51319377-26fe6e00-1a5d-11e9-8cc6-3137a566796d.png)](https://stackblitz.com/edit/easymde/)\n" +
	"👈 Try writing some markdown on the left.\n\n" +
	"> hmm?"

var titleTwo = "Second Demo"
var contentTwo = "# Second Demo\n\n" +
	"## Overview\n\n" +
	"* Follows [CommonMark](https://commonmark.org)\n" +
	"* Optionally follows [GitHub Flavored Markdown](https://github.github.com/gfm/)\n" +
	"* Renders actual React elements instead of using `dangerouslySetInnerHTML`\n" +
	"* Lets you define your own components (to render `MyHeading` instead of `'h1'`)\n" +
	"* Has a lot of plugins\n"

var titleThree = "Third Demo"
var contentThree = "# Third Demo\n\n" +
	"## Syntax highlighting\n\n" +

	"Here is an example of a plugin to highlight code:\n" +
	"[`rehype-highlight`](https://github.com/rehypejs/rehype-highlight).\n\n" +

	"```js\n" +
	"import React from 'react'\n" +
	"import ReactDOM from 'react-dom'\n" +
	"import Markdown from 'react-markdown'\n" +
	"import rehypeHighlight from 'rehype-highlight'\n\n" +

	"const markdown = `\n" +
	"# Your markdown here\n" +
	"`\n\n" +

	"ReactDOM.render(\n" +
	"  <Markdown rehypePlugins={[rehypeHighlight]}>{markdown}</Markdown>,\n" +
	"  document.querySelector('#content')\n" +
	")\n" +
	"```\n"

var NoteMock = note.Note{
	ID:        uuid.New(),
	Title:     titleOne,
	Content:   contentOne,
	UpdatedAt: time.Now(),
}

var NotesMock = note.Notes{
	Notes: []note.Note{
		{
			ID:        uuid.New(),
			Title:     titleOne,
			Content:   contentOne,
			UpdatedAt: time.Now().Add(time.Hour * -3),
		},
		{
			ID:        uuid.New(),
			Title:     titleTwo,
			Content:   contentTwo,
			UpdatedAt: time.Now().Add(time.Minute * -31),
		},
		{
			ID:        uuid.New(),
			Title:     titleThree,
			Content:   contentThree,
			UpdatedAt: time.Now(),
		},
	},
}
