package models

import (
	"html/template"
	"server/config"
	"time"
)

type Post struct {
	Pid        int       `json:"pid"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Content    string    `json:"content"`
	MarkDown   string    `json:"markdown"`
	CategoryId int       `json:"categoryId"`
	UserId     int       `json:"userId"`
	ViewCount  int       `json:"viewCount"`
	Type       int       `json:"type"` // 文章类型 0普通 1自定义
	CreateAt   time.Time `json:"createAt"`
	UpdateAt   time.Time `json:"updateAt"`
}

type PostMore struct {
	Pid          int           `json:"pid"`
	Title        string        `json:"title"`
	Slug         string        `json:"slug"`
	Content      template.HTML `json:"content"`
	CategoryId   int           `json:"categoryId"`
	CategoryName string        `json:"categoryName"`
	UserId       int           `json:"userId"`
	UserName     string        `json:"userName"`
	ViewCount    int           `json:"viewCount"`
	Type         int           `json:"type"` // 文章类型 0普通 1自定义
	CreateAt     time.Time     `json:"createAt"`
	UpdateAt     time.Time     `json:"updateAt"`
}

type PostReq struct {
	Pid        int    `json:"pid"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Content    string `json:"content"`
	MarkDown   string `json:"markdown"`
	CategoryId int    `json:"categoryId"`
	UserId     int    `json:"userId"`
	Type       int    `json:"type"` // 文章类型 0普通 1自定义
}

type SearchResp struct {
	Pid   int    `orm:"pid" json:"pid"`
	Title string `orm:"title" json:"title"`
}

type PostResp struct {
	config.Viewer
	config.SystemConfig
	Article PostMore
}
