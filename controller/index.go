package controller

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"strconv"
	"github.com/kataras/iris"
)

type IndexController struct {
	mvc.C
	Manager *sessions.Sessions
	Session *sessions.Session
}

func (c *IndexController) BeginRequest(ctx iris.Context) {
	c.C.BeginRequest(ctx)

	if c.Manager == nil {
		ctx.Application().Logger().Errorf(`VisitController: sessions manager is nil, you should bind it`)
		// dont run the main method handler and any "done" handlers.
		ctx.StopExecution()
		return
	}

	c.Session = c.Manager.Start(ctx)
}

func (c *IndexController) Get() string {
	count, _ := c.Session.GetIntDefault("count", 0)
	count++
	c.Session.Set("count", count)
	return "This is my default action..." + strconv.Itoa(count)
}

func (c *IndexController) GetMarkdown() {
	c.Ctx.Markdown(markdown)
}

var markdown = []byte(`## Hello Markdown

This is a sample of Markdown contents



Features
--------

All features of Sundown are supported, including:

*   **Compatibility**. The Markdown v1.0.3 test suite passes with
    the --tidy option.  Without --tidy, the differences are
    mostly in whitespace and entity escaping, where blackfriday is
    more consistent and cleaner.

*   **Common extensions**, including table support, fenced code
    blocks, autolinks, strikethroughs, non-strict emphasis, etc.

*   **Safety**. Blackfriday is paranoid when parsing, making it safe
    to feed untrusted user input without fear of bad things
    happening. The test suite stress tests this and there are no
    known inputs that make it crash.  If you find one, please let me
    know and send me the input that does it.

    NOTE: "safety" in this context means *runtime safety only*. In order to
    protect yourself against JavaScript injection in untrusted content, see
    [this example](https://github.com/russross/blackfriday#sanitize-untrusted-content).

*   **Fast processing**. It is fast enough to render on-demand in
    most web applications without having to cache the output.

*   **Routine safety**. You can run multiple parsers in different
    goroutines without ill effect. There is no dependence on global
    shared state.

*   **Minimal dependencies**. Blackfriday only depends on standard
    library packages in Go. The source code is pretty
    self-contained, so it is easy to add to any project, including
    Google App Engine projects.

*   **Standards compliant**. Output successfully validates using the
    W3C validation tool for HTML 4.01 and XHTML 1.0 Transitional.

	[this is a link](https://github.com/kataras/iris) `)
