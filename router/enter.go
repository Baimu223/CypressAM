package router

import (
	"xkginweb/router/bbs"
	"xkginweb/router/code"
	"xkginweb/router/course"
	"xkginweb/router/login"
	"xkginweb/router/state"
	"xkginweb/router/sys"
	"xkginweb/router/video"
)

type WebRouterGroup struct {
	Course course.WebRouterGroup
	Video  video.WebRouterGroup
	Sys    sys.WebRouterGroup
	State  state.WebRouterGroup
	BBs    bbs.WebRouterGroup
	Login  login.WebRouterGroup
	Code   code.WebRouterGroup
}

var RouterWebGroupApp = new(WebRouterGroup)
