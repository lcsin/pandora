package web

import "embed"

// HTMLFS 嵌入html文件系统
// 使用embed的好处是，打包成二进制文件后，这些静态资源文件会一并打包进去
// 因为embed不支持直接嵌入父目录文件，所以直接在当前目录嵌入模板文件，在其他包下间接引入
//
//go:embed template/*
var HTMLFS embed.FS

// StaticFS 嵌入静态资源文件
//go:embed static/*
var StaticFS embed.FS