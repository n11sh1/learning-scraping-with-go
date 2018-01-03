package main

import (
	"github.com/microcosm-cc/bluemonday"
	"fmt"
)

func main() {
	p := bluemonday.UGCPolicy()
	html := p.Sanitize(
		`<a onblur="alert(secret)" href="http://www.google.com">Google</a>`,
	)

	// Output:
	// <a href="http://www.google.com" rel="nofollow">Google</a>
	fmt.Println(html)


	p.AllowElements("li").AllowElements("ul")

	html2 := p.Sanitize(
		`<ul>
<li class="toclevel-2 tocsection-2"><a href="#.E5.AD.97.E6.BA.90"><span class="tocnumber">1.1</span> <span class="toctext">字源</span></a></li>
<li class="toclevel-2 tocsection-3"><a href="#.E9.9F.B3"><span class="tocnumber">1.2</span> <span class="toctext">音</span></a></li>
</ul>`,
	)

	fmt.Println(html2)


	p.AllowElements("li").AllowElements("ul")
	p.AllowStandardURLs()
	p.AllowAttrs("href").OnElements("a")
	html3 := p.Sanitize(
		`<ul>
<li class="toclevel-2 tocsection-2"><a href="scp://wiki.jp/hoge"><span class="tocnumber">1.1</span> <span class="toctext">字源</span></a></li>
<li class="toclevel-2 tocsection-3"><a href="scp://wiki.jphoge"><span class="tocnumber">1.2</span> <span class="toctext">音</span></a></li>
<li class="toclevel-2 tocsection-4"><a href="https://google.com"><span class="tocnumber">1.3</span> <span class="toctext">意義</span></a></li>
</ul>
`,
	)

	fmt.Println(html3)
}