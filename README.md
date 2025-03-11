# 🚀 Chatbot con DeepSeek Coder, Go y Vite

Este proyecto es un chatbot basado en **DeepSeek Coder** utilizando **Ollama**, con una API en **Go (Gin)** y un frontend en **Vite + React**. Permite interactuar con un modelo de IA para responder preguntas técnicas o generales.

## 📌 Características
- 🔥 **Backend en Go** con **Gin**.
- 🤖 **Modelo DeepSeek Coder** corriendo en **Ollama**.
- 🌐 **Frontend en Vite + React** para una interfaz rápida y moderna.
- 📡 **Comunicación mediante API REST**.
- 🎨 **Markdown en respuestas** para mejor legibilidad.
- 🎭 **Manejo de estados** y animaciones de "escribiendo" en el frontend.

---

## 🛠️ Instalación y Configuración

### 🔹 1. Clonar el repositorio
```sh
git clone https://github.com/lugonesnicolas/chatbot_deepseek
cd chatbot-deepseek
```

### 🔹 2. Instalar y configurar **Ollama**
1. [Descarga e instala Ollama](https://ollama.com/download).
2. Ejecuta el servidor:
   ```sh
   ollama serve
   ```
3. Descarga el modelo **DeepSeek Coder**:
   ```sh
   ollama pull deepseek-r1:7b
   ```

### 🔹 3. Backend en Go
1. Ejecutar la API:
   ```sh
   cd cmd
   go run main.go
   ```
   - La API se ejecutará en `http://localhost:8080`.

### 🔹 4. Frontend en Vite
1. Instalar dependencias:
   ```sh
   cd frontend
   npm install
   ```
2. Ejecutar el frontend:
   ```sh
   npm run dev
   ```
   - La aplicación se ejecutará en `http://localhost:3000`.

---

## 🔗 **Endpoints de la API**
### 📌 POST `/chat`
- **Descripción:** Envía un mensaje al modelo de IA y recibe una respuesta.
- **Ejemplo de request:**
  ```json
  {
    "message": "¿Qué es Go?"
  }
  ```
- **Ejemplo de respuesta:**
  ```json
  {
    "response": "Go es un lenguaje de programación creado por Google."
  }
  ```

### 📌 GET `/history`
- **Descripción:** Obtiene el historial de conversaciones guardado en la base de datos.

---

## 📌 Estructura del Proyecto
```
chatbot-deepseek/
│── api/                 # Backend en Go
│   ├── internal/        # Lógica interna
│   ├── main.go          # Punto de entrada
│   ├── go.mod           # Dependencias
│── frontend/            # Frontend en Vite + React
│   ├── src/
│   │   ├── App.jsx      # Interfaz del chatbot
│   │   ├── services/    # Llamadas a la API
│   ├── package.json     # Dependencias de Node.js
│── README.md            # Documentación
```

---

## 🛠️ **Tecnologías utilizadas**
- **Go (Gin)** 🟡
- **Ollama** 🤖
- **DeepSeek Coder** 🧠
- **React + Vite** ⚛️
- **Markdown Rendering** 📝
- **SQLite** (para historial) 🗄️

---

## 💡 Mejoras futuras
✅ Agregar autenticación de usuario.
✅ Guardar historial de conversaciones en localStorage.
✅ Mejorar diseño del frontend.

Si tienes sugerencias o mejoras, ¡envía un PR! 🚀

---

## 👨‍💻 **Autor**
Desarrollado por **Lugones Nicolás**.

📌 _¡Espero que te sirva! Si tienes dudas, abre un issue en el repo._

