package two

import "embed"

//go:embed hello.txt world.txt
var col embed.FS
