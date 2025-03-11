import { useState } from "react";
import ReactMarkdown from "react-markdown";
import "./App.css";

function App() {
  const [messages, setMessages] = useState([]);
  const [message, setMessage] = useState("");
  const [isTyping, setIsTyping] = useState(false); // Estado para animación

  const sendMessage = async () => {
    if (!message.trim()) return;

    const userMessage = { text: message, sender: "user" };
    setMessages((prevMessages) => [...prevMessages, userMessage]);
    setMessage("");

    setIsTyping(true); // Mostrar animación de escribiendo...

    try {
      const res = await fetch("http://localhost:8080/chat", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ message }),
      });

      const data = await res.json();
      const botMessage = { text: data.response, sender: "bot" };

      setMessages((prevMessages) => [...prevMessages, botMessage]);
    } catch (error) {
      console.error("Error al conectar con la API:", error);
      setMessages((prevMessages) => [
        ...prevMessages,
        { text: "⚠️ Error: No se pudo conectar con el servidor.", sender: "bot" },
      ]);
    } finally {
      setIsTyping(false);
    }
  };

  return (
    <div className="app-container">
      {/* 📌 Encabezado fijo en la parte superior */}
      <header className="chat-header">DeepSeek R1 - Lugones Nicolas</header>

      {/* 🗨️ Contenedor del chat */}
      <div className="chat-container">
        {messages.map((msg, index) => (
          <div
            key={index}
            className={`message ${msg.sender === "user" ? "user-message" : "bot-message"}`}
          >
            <ReactMarkdown>{msg.text}</ReactMarkdown>
          </div>
        ))}

        {/* Animación cuando la IA está "escribiendo" */}
        {isTyping && (
          <div className="bot-message typing-indicator">
            <span></span>
            <span></span>
            <span></span>
          </div>
        )}
      </div>

      {/* 📩 Input y botón de enviar */}
      <div className="input-container">
        <input
          type="text"
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          onKeyDown={(e) => {
            console.log("Tecla presionada:", e.key); // DEBUG
            if (e.key === "Enter") sendMessage();
          }}
          placeholder="Escribe un mensaje..."
          className="chat-input"
        />
        <button 
          onClick={sendMessage} 
          onTouchStart={sendMessage} 
          className="chat-button"
        >
          Enviar
        </button>
      </div>
    </div>
  );
}

export default App;
