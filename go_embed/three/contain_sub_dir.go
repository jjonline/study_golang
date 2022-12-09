package three

import "embed"

//go:embed stubs world.txt
var Sub embed.FS

//go:embed [ab]?.txt
var T embed.FS
