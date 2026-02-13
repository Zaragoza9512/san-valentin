# 💝 Propuesta San Valentín - Servidor Web en Go

> *"No necesito una fecha en el calendario para recordarte cuánto te amo, pero sí quiero aprovecharla para consentirte como te mereces"*

Aplicación web completa con **backend en Go** y frontend interactivo, creada para pedirle a alguien muy especial ser mi San Valentín 💕

## 🌐 Demo en Vivo

**[→ Ver sitio web](http://acceso-privado-yus-95.onrender.com/)**

## 📸 Preview

(screenshot.png.jpeg)

---

## 🛠️ Stack Técnico

### Backend
- **Go 1.23** - Servidor HTTP nativo
- **net/http** - Servidor web sin frameworks
- **html/template** - Template engine server-side

### Frontend
- **HTML5** - Estructura semántica
- **CSS3** - Animaciones y diseño responsive
- **JavaScript** - Interactividad y efectos

### DevOps
- **Render** - Deployment automático
- **Git/GitHub** - Control de versiones

---

## ✨ Características Principales

### Backend (Go)
- ✅ Servidor HTTP desde cero (sin frameworks)
- ✅ Template engine con inyección de variables
- ✅ Configuración con variables de entorno
- ✅ Single binary deployment

### Frontend
- 💝 Diseño romántico responsive
- 🎨 Animaciones CSS suaves
- 💕 Botón "No" que se reduce progresivamente
- 💖 Botón "Sí" que crece con cada rechazo
- 🎉 Sistema de confetti (50 corazones animados)

---

## 🚀 Instalación y Uso

### Prerrequisitos
```bash
go version  # Necesitas Go 1.23+
```

### Clonar y ejecutar
```bash
git clone https://github.com/Zaragoza9512/[TU-REPO].git
cd [TU-REPO]
go run main.go
```

Abre en tu navegador: **http://localhost:8080**

### Personalizar
Edita estas líneas en `main.go`:
```go
const (
    nombreDeTuNovia = "Tu Nombre Aquí 💝"
    tuNombre        = "Tu Firma"
)
```

---

## 📁 Estructura del Proyecto
```
propuesta-san-valentin/
│
├── main.go              # Servidor completo (Go + HTML + CSS + JS)
├── README.md            # Este archivo
└── screenshot.png       # Preview del sitio
```

---

## 🎯 Motivación

Como desarrollador backend especializado en Go, quería crear algo que combinara:
- ✅ Habilidades técnicas (servidor HTTP en Go)
- ✅ Creatividad (propósito personal)
- ✅ Deployment real (producción en Render)

**Resultado:** Una aplicación web completa que demuestra capacidad full-stack.

---

## 📚 Aprendizajes Técnicos

### Go Backend
- Servidor HTTP nativo sin frameworks
- Template engine server-side (`html/template`)
- Manejo de variables de entorno
- Single binary deployment

### Frontend
- Animaciones CSS con keyframes
- Manipulación del DOM con JavaScript
- Sistema de confetti dinámico
- Diseño responsive

### DevOps
- Deployment de aplicaciones Go en Render
- Build automático desde GitHub
- Configuración de PORT dinámico

---

## 💡 Por Qué Go

Ventajas sobre HTML estático:
- ✅ **Template dinámico:** Variables inyectadas desde backend
- ✅ **Single binary:** Un archivo ejecutable contiene todo
- ✅ **Performance:** Servidor ultra-rápido (~20MB RAM)
- ✅ **Zero dependencies:** Solo stdlib de Go

---

## 🔧 Deployment en Render

### Configuración
```
Build Command: go build -o server
Start Command: ./server
```

Cada `git push` despliega automáticamente.

---

## 📊 Métricas

- **Líneas de código:** ~500 (Go + HTML + CSS + JS)
- **Binary size:** ~10MB
- **Memory usage:** ~20MB en producción
- **Cold start:** <100ms
- **Costo:** $0 (Render free tier)

---

## 👨‍💻 Autor

**Luis Eduardo Zaragoza Hernández**

Desarrollador Backend especializado en Go, Docker y Kubernetes.

- 💻 **GitHub:** [@Zaragoza9512](https://github.com/Zaragoza9512)
- 💼 **LinkedIn:** [luis-zaragoza95](https://linkedin.com/in/luis-zaragoza95)
- 📧 **Email:** zaragoza95.luis@gmail.com

### Otros proyectos:
- 🚀 [Go E-commerce API](https://github.com/Zaragoza9512/go-api-chi) - REST API con Kubernetes y CI/CD

---

## 💕 Dedicatoria

*Para Yus, quien hace que valga la pena escribir código.*  
*Este servidor web es una pequeña expresión digital de lo mucho que significas.* 💖

---

## ⭐ ¿Te gustó?

Si este proyecto te pareció creativo:
- Dale una ⭐ al repositorio
- Compártelo con otros developers
- Crea tu propia versión

---

<div align="center">

### 💝 Hecho con Go, amor y código 💻

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white)
![CSS3](https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white)
![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black)

**[Ver proyecto en vivo →](http://acceso-privado-yus-95.onrender.com/)**

</div>