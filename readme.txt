
Mini service for registration and authorization of users.
This is an educational project. Created for training and self-education.


Мини сервис регистрации и авторизации пользователей.
Это учебный проект. Создан для обучения и самообразования.


🧭 Карта функций — Регистрация (шаг за шагом)
ShowRegisterPage(w, r)
🔹 Показывает начальную страницу регистрации (первое поле: никнейм)

CheckUsername(w, r)
🔹 Проверяет никнейм (валидный? занят?)
🔹 Возвращает поле никнейма + статус (ошибка/успех)
🔸 При успехе → HTMX подгружает ShowPasswordStep(w, r)

ShowPasswordStep(w, r)
🔹 Показывает второе поле — пароль
🔸 При вводе → CheckPassword(w, r) (опционально, если хочешь валидировать)

ShowEmailStep(w, r)
🔹 Показывает поле для email
🔸 При вводе → CheckEmail(w, r) (если валидируешь)

RegisterUser(w, r)
🔹 После ввода всех данных — создаёт пользователя в БД
🔸 Успех → редирект на личный кабинет: ShowDashboard(w, r)

🔐 Карта функций — Авторизация (вход)
ShowLoginPage(w, r)
🔹 Показывает форму входа (только поле никнейм)

CheckLoginUsername(w, r)
🔹 Проверяет, существует ли пользователь с таким ником
🔸 При успехе → ShowPasswordLoginStep(w, r)

CheckLoginPassword(w, r)
🔹 Проверяет пароль
🔸 Успех → редирект в ShowDashboard(w, r)

🧩 Дополнительные функции:
IsUsernameTaken(username string)
🔹 Проверка — занят ли ник

ValidatePassword(password string)
🔹 Проверка надёжности пароля (опционально)

CreateUser(username, password, email)
🔹 Добавляет пользователя в БД

Authenticate(username, password)
🔹 Проверяет пару логин/пароль