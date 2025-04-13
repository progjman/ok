package handler

import (
	"fmt"
	"html"
	"net/http"
	"regexp"
	"strings"

	"github.com/progjman/ok/db"
)

func CheckUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	username = strings.TrimSpace(username)

	if len(username) < 3 || len(username) > 20 {
		writeInvalid(w, username, "From 3 to 20 characters")
		return
	}

	if !isValidNickname(username) {
		writeInvalid(w, username, "Only Latin letters, numbers and _")
		return
	}

	if db.IsUsernameTaken(username) {
		writeInvalid(w, username, "Nick is already taken")
		return
	}

	writeValid(w, username, "Nick is free")
}

func isValidNickname(nick string) bool {
	ok, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, nick)
	return ok
}

func writeInvalid(w http.ResponseWriter, value, msg string) {
	fmt.Fprintf(w, `
	<div id="step-wrapper">
		<label for="username">nickname</label>
		<input type="text" name="username" id="username" value="%s" class="input invalid"
			hx-get="/check-username"
			hx-trigger="keyup changed delay:300ms"
			hx-target="#step-wrapper"
			hx-swap="outerHTML"
		/>
		<div class="status-msg error">%s</div>
	</div>
	`, html.EscapeString(value), html.EscapeString(msg))
}

func writeValid(w http.ResponseWriter, value, msg string) {
	fmt.Fprintf(w, `
	<div id="step-wrapper">
		<label for="username">nickname</label>
		<input type="text" name="username" id="username" value="%s" class="input valid"
			hx-get="/check-username"
			hx-trigger="keyup changed delay:300ms"
			hx-target="#step-wrapper"
			hx-swap="outerHTML"
		/>
		<div class="status-msg success">%s</div>
	</div>
	`, html.EscapeString(value), html.EscapeString(msg))
}
