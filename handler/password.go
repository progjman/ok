package handler

import (
	"fmt"
	"html"
	"net/http"
	"regexp"
	"strings"
)

func RegisterPassword(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	// Рендерим HTML для второго шага — ввод пароля (не скрыт)
	fmt.Fprintf(w, `
	<form id="register-form">
		<!-- Скрытое поле с никнеймом -->
		<input type="hidden" name="username" value="%s" />

		<div id="password-wrapper">
			<label for="password">Your password</label>
			<input type="text" name="password" id="password" class="input"
				hx-get="/check-password"
				hx-trigger="keyup changed delay:300ms"
				hx-target="#password-wrapper"
				hx-swap="outerHTML"
			/>
			<div class="status-msg">Please enter a strong password</div>
		</div>

		<br>
		<button type="submit">Register</button>
	</form>
	`, html.EscapeString(username))
}

func CheckPassword(w http.ResponseWriter, r *http.Request) {
	password := r.URL.Query().Get("password")
	password = strings.TrimSpace(password)

	// Проверка сложности пароля
	if len(password) < 8 {
		writeInvalidPassword(w, password, " allowed A-Z  a-z  0-9  ! @  #  %  &  *  (  )  _   +  =  -")
		return
	}

	if !isStrongPassword(password) {
		writeInvalidPassword(w, password, " allowed A-Z  a-z  0-9  ! @  #  %  &  *  (  )  _   +  =  -")
		return
	}

	writeValidPassword(w, password, "Password is strong")
}

func isStrongPassword(password string) bool {
	// Проверки на заглавные буквы, цифры, спецсимволы
	uppercase := regexp.MustCompile(`[A-Z]`)
	lowercase := regexp.MustCompile(`[a-z]`)
	number := regexp.MustCompile(`[0-9]`)
	special := regexp.MustCompile(`[!@#%&*()_+=-]`)

	return len(password) >= 8 &&
		uppercase.MatchString(password) &&
		lowercase.MatchString(password) &&
		number.MatchString(password) &&
		special.MatchString(password)
}

func writeInvalidPassword(w http.ResponseWriter, value, msg string) {
	fmt.Fprintf(w, `
	<div id="password-wrapper">
		<label for="password">Password</label>
		<input type="text" name="password" id="password" value="%s" class="input invalid"
			hx-get="/check-password"
			hx-trigger="keyup changed delay:300ms"
			hx-target="#password-wrapper"
			hx-swap="outerHTML"
		/>
		<div class="status-msg error">%s</div>
	</div>
	`, html.EscapeString(value), html.EscapeString(msg))
}

func writeValidPassword(w http.ResponseWriter, value, msg string) {
	fmt.Fprintf(w, `
	<div id="password-wrapper">
		<label for="password">Password</label>
		<input type="text" name="password" id="password" value="%s" class="input valid"
			hx-get="/check-password"
			hx-trigger="keyup changed delay:300ms"
			hx-target="#password-wrapper"
			hx-swap="outerHTML"
		/>
		<div class="status-msg success">%s</div>
	</div>
	`, html.EscapeString(value), html.EscapeString(msg))
}
