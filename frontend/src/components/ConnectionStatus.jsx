import { useState, useEffect } from "react";

export default function ConnectionStatus() {
  const [connected, setConnected] = useState(false);
  const [ip, setIp] = useState("--");
  const [speed, setSpeed] = useState("0 Mbps");

  useEffect(() => {
    // Simular obtener IP pública
    fetch("https://api.ipify.org?format=json")
      .then(res => res.json())
      .then(data => setIp(data.ip))
      .catch(() => setIp("No disponible"));
  }, []);

  function toggleConnection() {
    setConnected(!connected);
    setSpeed(connected ? "0 Mbps" : Math.floor(Math.random() * 100) + " Mbps");
  }

  return (
    <div style={{
      background: "#1e293b",
      padding: "20px",
      borderRadius: "10px",
      border: "1px solid #334155",
      marginBottom: "20px"
    }}>
      <div style={{
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center"
      }}>
        <div>
          <h3 style={{ color: "#38bdf8", margin: "0 0 10px 0" }}>
            Estado de conexión
          </h3>
          <p style={{ color: "#cbd5e1", margin: "5px 0" }}>
            IP: <strong>{ip}</strong>
          </p>
          <p style={{ color: "#cbd5e1", margin: "5px 0" }}>
            Velocidad: <strong>{speed}</strong>
          </p>
        </div>

        <button
          onClick={toggleConnection}
          style={{
            width: "120px",
            height: "120px",
            borderRadius: "50%",
            border: "3px solid " + (connected ? "#22c55e" : "#ef4444"),
            background: connected ? "rgba(34, 197, 94, 0.1)" : "rgba(239, 68, 68, 0.1)",
            color: connected ? "#22c55e" : "#ef4444",
            fontSize: "18px",
            fontWeight: "bold",
            cursor: "pointer",
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            transition: "all 0.3s"
          }}
        >
          {connected ? "🔓\nCONECTADO" : "🔒\nDESCONECTADO"}
        </button>
      </div>
    </div>
  );
}
