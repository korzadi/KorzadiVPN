import { useState, useEffect } from "react";
import api from "../api/api";

export default function ServerList() {
  const [servers, setServers] = useState([]);
  const [selectedServer, setSelectedServer] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchServers();
  }, []);

  async function fetchServers() {
    try {
      // Simulamos servidores locales - reemplazar con API real
      setServers([
        { id: 1, name: "🇺🇸 Estados Unidos - NYC", country: "US", ping: "12ms", load: "45%" },
        { id: 2, name: "🇬🇧 Reino Unido - Londres", country: "UK", ping: "25ms", load: "60%" },
        { id: 3, name: "🇯🇵 Japón - Tokio", country: "JP", ping: "98ms", load: "30%" },
        { id: 4, name: "🇩🇪 Alemania - Berlín", country: "DE", ping: "18ms", load: "55%" },
        { id: 5, name: "🇦🇺 Australia - Sídney", country: "AU", ping: "145ms", load: "40%" },
      ]);
      setLoading(false);
    } catch (err) {
      console.error("Error al cargar servidores:", err);
      setLoading(false);
    }
  }

  if (loading) {
    return <p style={{ color: "#cbd5e1" }}>Cargando servidores...</p>;
  }

  return (
    <div style={{
      background: "#1e293b",
      padding: "20px",
      borderRadius: "10px",
      border: "1px solid #334155"
    }}>
      <h3 style={{ color: "#38bdf8", margin: "0 0 15px 0" }}>
        Servidores disponibles
      </h3>

      <div style={{
        display: "grid",
        gridTemplateColumns: "repeat(auto-fill, minmax(250px, 1fr))",
        gap: "10px"
      }}>
        {servers.map((server) => (
          <div
            key={server.id}
            onClick={() => setSelectedServer(server.id)}
            style={{
              padding: "15px",
              background: selectedServer === server.id ? "#0ea5e9" : "#0f172a",
              border: "1px solid " + (selectedServer === server.id ? "#38bdf8" : "#334155"),
              borderRadius: "8px",
              cursor: "pointer",
              transition: "all 0.3s"
            }}
          >
            <p style={{ color: "#e2e8f0", margin: "0 0 8px 0", fontWeight: "bold" }}>
              {server.name}
            </p>
            <p style={{ color: "#cbd5e1", margin: "5px 0", fontSize: "12px" }}>
              ⏱️ Ping: {server.ping}
            </p>
            <p style={{ color: "#cbd5e1", margin: "5px 0", fontSize: "12px" }}>
              📊 Carga: {server.load}
            </p>
          </div>
        ))}
      </div>

      {selectedServer && (
        <button style={{
          width: "100%",
          marginTop: "15px",
          padding: "12px",
          background: "#38bdf8",
          color: "white",
          border: "none",
          borderRadius: "8px",
          fontWeight: "bold",
          cursor: "pointer"
        }}>
          Conectar a {servers.find(s => s.id === selectedServer)?.name}
        </button>
      )}
    </div>
  );
}
