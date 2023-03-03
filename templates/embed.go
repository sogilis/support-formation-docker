package templates

import _ "embed"

//go:embed index.template.html
var IndexPage string

//go:embed uploaded.template.html
var UploadedPage string
