package elysium

import (
	html "html/template"
)

type Page struct {
	User    User
	Payload interface{}
	CSRF    html.HTML
}
