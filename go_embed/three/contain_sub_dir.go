package three

import "embed"

//go:embed stubs world.txt
var Sub embed.FS

//go:embed [ab]?.txt
var T embed.FS

//go:embed [ab].txt
var T1 embed.FS

//go:embed [ab]?.txt
var T2 embed.FS

//go:embed [ab]*.txt
var T3 embed.FS

//go:embed a.txt b.txt
var T4 embed.FS

//go:embed stub/*
var T5 embed.FS

//go:embed stubs/*
var T6 embed.FS
