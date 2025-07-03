# 🖼️ ASCII Art Web Application

`ascii-art-web` is an extended version of the original `ascii-art` project, which generates ASCII art in the terminal. This version is a Go-based web application that allows users to generate ASCII art from text through a user-friendly graphical interface in a web browser. Users can select banner styles (`standard`, `shadow`, `thinkertoy`) and view the results on a responsive, stylized webpage. The implementation follows best practices in Go development and modern web design.

## 🧩 Features

- 💻 Web interface with HTML/CSS
- 🧠 Banner selection (`standard`, `shadow`, `thinkertoy`)
- ✅ Valid HTTP status support (200, 400, 404, 500)
- 🐳 Fully dockerized
- 🎨 Stylized and responsive design
- 🛡 Secure input validation
- 📁 Standard project structure
- 🚫 Only Go standard library is used

---

## 📦 Project Structure

```
ascii-art-web/
│
├── templates/          # HTML templates (index.html, result.html, error.html)
├── static/             # CSS files (style.css, etc.)
├── banner.txt files    # standard.txt, shadow.txt, thinkertoy.txt (in root)
├── cmd/                # main.go (entry point)
├── internal/           # handlers, ASCII generation logic
├── Dockerfile          # Docker image
├── *.sh                # build/run/stop/cleanup scripts
└── README.md
```

---

## 🚀 How to Run

### 🛠️ Prerequisites

- Go `1.16+` (for local run)
- Docker `20.10+` (for containerization)
- Bash (if using shell scripts)

---

### 🔧 Run Locally

```bash
git clone https://github.com/Baktiyar88/ascii-art-web.git
cd ascii-art-web
go run .
```

Open your browser at: `http://localhost:8080`

---

### 🐳 Run with Docker

#### Option 1: With helper scripts

```bash
bash build.sh      # build the image
bash run.sh        # run the container (port 7777)
bash stop.sh       # stop the container
bash cleanup.sh    # remove image
```

#### Option 2: Manually via Docker CLI

```bash
docker build -f Dockerfile -t ascii-art-web .
docker run -d -p 7777:7777 --name ascii-art ascii-art-web
```

---

## 🌐 How to Use

1. Enter text in the input field.
2. Select a banner style: `standard`, `shadow`, `thinkertoy`.
3. Click `Generate`.
4. The ASCII art will appear on the same page or a new result page.

---

## 🧠 Example

**Input:**  
```
Hello
```

**Banner:** `shadow`  
**Result:**  
(ASCII characters displayed in a `<pre>` block)

---

## 🔍 HTTP Endpoints

| Method | URL          | Description                              |
|--------|--------------|------------------------------------------|
| GET    | `/`          | Main page with form                      |
| POST   | `/ascii-art` | Generate ASCII art from submitted input  |

---

## ⚙️ Implementation

### 📡 Server

- Built using Go's `net/http` package.
- Templates rendered with `html/template`.
- ASCII generation reads banner files line by line.
- Static files served with `http.FileServer`.

### 🎨 Styling (Stylized Version)

- CSS stored in `static/style.css`.
- **Responsive design** with media queries.
- Animations, hover effects, and user feedback messages.
- Monospace font for ASCII, sans-serif for UI elements.

---

## 📤 HTTP Status Codes

| Code | Meaning                             |
|------|-------------------------------------|
| 200  | OK (successful request)             |
| 400  | Bad Request (empty input, invalid)  |
| 404  | Not Found (missing file/template)   |
| 500  | Internal Server Error               |

---

## 🧼 Dockerfile Highlights

- **Multi-stage build**: `golang:1.24.2` for building, `alpine` for runtime
- Metadata included:
  ```Dockerfile
  LABEL maintainer="Baktiyar <bzhaksyb>"
  LABEL version="1.0"
  LABEL description="ASCII Art Generator web application in Go"
  ```
- **Security**: Runs as non-root user
- **Small final image**: only binary and essential files
