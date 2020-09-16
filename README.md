# docx
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/gingfrederik/docx?sort=semver)
[![Go Report Card](https://goreportcard.com/badge/github.com/gingfrederik/docx)](https://goreportcard.com/report/github.com/gingfrederik/docx)
[![GoDoc](https://pkg.go.dev/badge/github.com/gingfrederik/docx?status.svg)](https://pkg.go.dev/github.com/gingfrederik/docx?tab=doc)
## Introduction
docx is a simple library to creating DOCX file in Go.

## Getting Started
### Install
Go modules supported
```sh
go get github.com/gingfrederik/docx
```
Import:
```sh
import "github.com/gingfrederik/docx"
```

### Usage
**Example:**
```go
package main

import (
	"github.com/gingfrederik/docx"
)

func main() {
	f := docx.NewFile()
	// add new paragraph
	para := f.AddParagraph()
	// add text
	para.AddText("test")

	para.AddText("test font size").Size(22)
	para.AddText("test color").Color("808080")
	para.AddText("test font size and color").Size(22).Color("121212")

	nextPara := f.AddParagraph()
	nextPara.AddLink("google", `http://google.com`)

	f.Save("./test.docx")
}

```
