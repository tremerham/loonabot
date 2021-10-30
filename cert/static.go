package cert

import (
	"embed"
)

//go:embed *.pem
var Certs embed.FS
