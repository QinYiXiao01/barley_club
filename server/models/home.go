package models

import "server/config"

type HomeResp struct {
	config.Viewer
	Categories []Category
	Posts      []PostMore
	Total      int
	Page       int
	Pages      []int
	PageEnd    bool
}
