# 🖼️ ASCII Art Web Application

`ascii-art-web` — это веб-приложение на Go, которое позволяет пользователям генерировать ASCII-арт из текста через удобный графический интерфейс в браузере. Приложение поддерживает выбор баннера (`standard`, `shadow`, `thinkertoy`) и предоставляет стилизованную, отзывчивую веб-страницу с подробной обратной связью. Реализация следует лучшим практикам Go и веб-дизайна.

---

## 👤 Авторы

- **Baktiyar Zhaksybay** [`#bzhaksyb`](https://github.com/Baktiyar88)
- **Dana Isabaeva** [`#disabaev`](https://github.com/)
- **Bek Nurekeyev** [`#bnurekey`](https://github.com/)

---

## 🧩 Особенности

- 💻 Веб-интерфейс с HTML/CSS
- 🧠 Выбор баннера (`standard`, `shadow`, `thinkertoy`)
- ✅ Поддержка валидных HTTP-статусов (200, 400, 404, 500)
- 🐳 Полностью dockerized
- 🎨 Стилизация и респонсив-дизайн
- 🛡 Безопасная валидация ввода
- 📁 Стандартная структура проекта
- 🚫 Только стандартная библиотека Go

---

## 📦 Структура проекта

```
ascii-art-web/
│
├── templates/          # HTML-шаблоны (index.html, result.html, error.html)
├── static/             # CSS-файлы (style.css и пр.)
├── banner.txt files    # standard.txt, shadow.txt, thinkertoy.txt (в корне)
├── cmd/                # main.go (точка входа)
├── internal/           # хендлеры, логика генерации ASCII
├── Dockerfile          # Docker-образ
├── *.sh                # build/run/stop/cleanup скрипты
└── README.md
```

---

## 🚀 Как запустить

### 🛠️ Предварительные требования

- Go `1.16+` (для локального запуска)
- Docker `20.10+` (для контейнеризации)
- Bash (если используете shell-скрипты)

---

### 🔧 Запуск локально

```bash
git clone https://github.com/Baktiyar88/ascii-art-web.git
cd ascii-art-web
go run .
```

Откройте браузер: `http://localhost:8080`

---

### 🐳 Запуск с помощью Docker

#### Способ 1: через скрипты

```bash
bash build.sh      # сборка образа
bash run.sh        # запуск контейнера (порт 7777)
bash stop.sh       # остановка контейнера
bash cleanup.sh    # очистка образа
```

#### Способ 2: вручную через Docker CLI

```bash
docker build -f Dockerfile -t ascii-art-web .
docker run -d -p 7777:7777 --name ascii-art ascii-art-web
```

---

## 🌐 Как пользоваться

1. Введите текст в поле ввода.
2. Выберите стиль баннера: `standard`, `shadow`, `thinkertoy`.
3. Нажмите `Generate`.
4. ASCII-арт появится на той же странице или на новой.

---

## 🧠 Пример

**Ввод:**  
```
Hello
```

**Баннер:** `shadow`  
**Результат:**  
(ASCII-арт из символов, отображённый в `<pre>` блоке)

---

## 🔍 HTTP-эндпоинты

| Метод | URL          | Описание                            |
|-------|--------------|-------------------------------------|
| GET   | `/`          | Главная страница                    |
| POST  | `/ascii-art` | Генерация ASCII по введённому тексту |

---

## ⚙️ Реализация

### 📡 Сервер

- Используется стандартный `net/http`.
- Обработка шаблонов — через `html/template`.
- Генерация ASCII — построчная, по баннер-файлам.
- Статические файлы обслуживаются через `http.FileServer`.

### 🎨 Стилизация (Stylized Version)

- `static/style.css` подключён к шаблонам.
- **Responsive design** через media queries.
- Анимации, hover-эффекты, понятные сообщения об ошибках.
- Использование шрифта `monospace` для ASCII, `sans-serif` — для интерфейса.

---

## 📤 HTTP-статусы

| Код | Значение                    |
|-----|-----------------------------|
| 200 | OK (успешная генерация)     |
| 400 | Bad Request (пустой ввод и т.д.) |
| 404 | Not Found (нет шаблона/файла) |
| 500 | Internal Server Error       |

---

## 🧼 Dockerfile Highlights

- **Multi-stage build**: `golang:1.24.2` для сборки, `alpine` для запуска
- Добавлены метаданные:
  ```Dockerfile
  LABEL maintainer="Baktiyar <bzhaksyb>, Bek <bnurekey>, Dana <disabaev>"
  LABEL version="1.0"
  LABEL description="ASCII Art Generator web application in Go"
  ```
- **Безопасность**: Запуск от non-root пользователя
- **Минимальный размер образа**: только бинарник и необходимые файлы

---

## 💡 Принципы

- ✅ Чистый, модульный код
- ✅ Go-конвенции и читаемость
- ✅ Поддержка мобильных устройств
- ✅ Только стандартные библиотеки
- ✅ Обратная связь пользователю

---

## 📚 Чему вы научитесь

- Go + HTML шаблоны
- Работа с HTTP сервером и формами
- CSS стилизация и адаптивная верстка
- Структура веб-приложений на Go
- Docker: multi-stage builds, контейнеризация
- Безопасная обработка пользовательского ввода

---

## 📄 Лицензия

MIT License