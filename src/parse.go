package elysium

import (
	"log"
	_ "regexp"
	"strings"

	"github.com/VojtechVitek/mention"
	"github.com/frustra/bbcode"
	"github.com/russross/blackfriday"
)

func nl2br(t string) string {
	t = strings.Replace(t, "\n", "<br />", -1)
	return t
}
func parseText(text string) string {
	log.Println("")
	compiler := bbcode.NewCompiler(true, true)
	//text = qr.ReplaceAllString(text, "<div class=\"quote\">$3</div>")
	text = compiler.Compile(text)
	return text
}

func mentions(text string) string {
	tags := mention.GetTags('@', strings.NewReader(text))
	log.Println(tags)

	return ""
	// placeholder
}

func markdown(text string) string {
	output := blackfriday.MarkdownBasic([]byte(text))
	return string(output)
}
