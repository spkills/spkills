// This file is automatically generated by qtc from "base.qtpl".
// See https://github.com/valyala/quicktemplate for details.

// This is a base page template. All the other template pages implement this interface.
//

//line base.qtpl:4
package templates

//line base.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line base.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line base.qtpl:5
type Page interface {
	//line base.qtpl:5
	Title() string
	//line base.qtpl:5
	StreamTitle(qw422016 *qt422016.Writer)
	//line base.qtpl:5
	WriteTitle(qq422016 qtio422016.Writer)
	//line base.qtpl:5
	Body() string
	//line base.qtpl:5
	StreamBody(qw422016 *qt422016.Writer)
	//line base.qtpl:5
	WriteBody(qq422016 qtio422016.Writer)
//line base.qtpl:5
}

// Page prints a page implementing Page interface.

//line base.qtpl:13
func StreamPageTemplate(qw422016 *qt422016.Writer, p Page) {
	//line base.qtpl:13
	qw422016.N().S(`
<html>
	<head>
		<title>`)
	//line base.qtpl:16
	p.StreamTitle(qw422016)
	//line base.qtpl:16
	qw422016.N().S(`</title>
	</head>
	<body>
		<div>
			<a href="/">return to main page</a>
		</div>
		`)
	//line base.qtpl:22
	p.StreamBody(qw422016)
	//line base.qtpl:22
	qw422016.N().S(`
	</body>
</html>
`)
//line base.qtpl:25
}

//line base.qtpl:25
func WritePageTemplate(qq422016 qtio422016.Writer, p Page) {
	//line base.qtpl:25
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line base.qtpl:25
	StreamPageTemplate(qw422016, p)
	//line base.qtpl:25
	qt422016.ReleaseWriter(qw422016)
//line base.qtpl:25
}

//line base.qtpl:25
func PageTemplate(p Page) string {
	//line base.qtpl:25
	qb422016 := qt422016.AcquireByteBuffer()
	//line base.qtpl:25
	WritePageTemplate(qb422016, p)
	//line base.qtpl:25
	qs422016 := string(qb422016.B)
	//line base.qtpl:25
	qt422016.ReleaseByteBuffer(qb422016)
	//line base.qtpl:25
	return qs422016
//line base.qtpl:25
}

// Base page implementation. Other pages may inherit from it if they need
// overriding only certain Page methods

//line base.qtpl:30
type BasePage struct{}

//line base.qtpl:31
func (p *BasePage) StreamTitle(qw422016 *qt422016.Writer) {
//line base.qtpl:31
qw422016.N().S(`This is a base title`) }

//line base.qtpl:31
//line base.qtpl:31
func (p *BasePage) WriteTitle(qq422016 qtio422016.Writer) {
	//line base.qtpl:31
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line base.qtpl:31
	p.StreamTitle(qw422016)
	//line base.qtpl:31
	qt422016.ReleaseWriter(qw422016)
//line base.qtpl:31
}

//line base.qtpl:31
func (p *BasePage) Title() string {
	//line base.qtpl:31
	qb422016 := qt422016.AcquireByteBuffer()
	//line base.qtpl:31
	p.WriteTitle(qb422016)
	//line base.qtpl:31
	qs422016 := string(qb422016.B)
	//line base.qtpl:31
	qt422016.ReleaseByteBuffer(qb422016)
	//line base.qtpl:31
	return qs422016
//line base.qtpl:31
}

//line base.qtpl:32
func (p *BasePage) StreamBody(qw422016 *qt422016.Writer) {
//line base.qtpl:32
qw422016.N().S(`This is a base body`) }

//line base.qtpl:32
//line base.qtpl:32
func (p *BasePage) WriteBody(qq422016 qtio422016.Writer) {
	//line base.qtpl:32
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line base.qtpl:32
	p.StreamBody(qw422016)
	//line base.qtpl:32
	qt422016.ReleaseWriter(qw422016)
//line base.qtpl:32
}

//line base.qtpl:32
func (p *BasePage) Body() string {
	//line base.qtpl:32
	qb422016 := qt422016.AcquireByteBuffer()
	//line base.qtpl:32
	p.WriteBody(qb422016)
	//line base.qtpl:32
	qs422016 := string(qb422016.B)
	//line base.qtpl:32
	qt422016.ReleaseByteBuffer(qb422016)
	//line base.qtpl:32
	return qs422016
//line base.qtpl:32
}