import { useEffect, useState } from "react";
import api from "../api/api";
import { useNavigate } from "react-router-dom";
import "../cyberpunk.css";

const FLAGS = {
  "Estados Unidos": "🇺🇸",
  "España": "🇪🇸",
  "Brasil": "🇧🇷",
};

export default function Servers() {
  const navigate = useNavigate();
  const [servers, setServers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [connectingId, setConnectingId] = useState(null);
  const [error, setError] = useState("");

  useEffect(() => {
    loadServers();
  }, []);

  async function loadServers() {
    setLoading(true);
    try {
      const res = await api.get("/api/servers");
      setServers(res.data || []);
    } catch (err) {
      setError("No se pudieron cargar los servidores.");
    } finally {
      setLoading(false);
    }
  }

  async function connect(server) {
    setConnectingId(server.id);
    setError("");
    try {
      await api.post("/api/connect", { server_id: server.id, device: "Android" });
      alert(`Conectado a ${server.name}`);
    } catch (err) {
      setError(err?.response?.data?.error || "No se pudo conectar al servidor.");
    } finally {
      setConnectingId(null);
    }
  }

  function latencyColor(ms) {
    if (ms <= 25) return "var(--acid)";
    if (ms <= 60) return "var(--cyan)";
    return "var(--hot)";
  }

  return (
    <div className="cyber-page">
      <div className="cyber-shell">
        <button className="cyber-back" onClick={() => navigate("/")}>← volver al dashboard</button>

        <h1 className="cyber-h1">Servidores</h1>
        <p className="cyber-sub">&gt; escaneando_nodos_disponibles...</p>

        {loading && <p className="cyber-empty">Cargando servidores…</p>}
        {!loading && error && <p className="cyber-empty" style={{ color: "var(--hot)" }}>{error}</p>}
        {!loading && !error && servers.length === 0 && (
          <p className="cyber-empty">No hay servidores disponibles ahora mismo.</p>
        )}

        <div className="cyber-grid">
          {servers.map((s) => (
            <div className="cyber-card" key={s.id}>
              <div style={{ display: "flex", justifyContent: "space-between", alignItems: "flex-start" }}>
                <div>
                  <div style={{ fontSize: 22, marginBottom: 4 }}>
                    {FLAGS[s.country] || "🌐"} <strong style={{ fontFamily: "'Orbitron',sans-serif", fontSize: 15 }}>{s.name}</strong>
                  </div>
                  <p className="cyber-mono" style={{ color: "var(--fog)", fontSize: 12, margin: 0 }}>
                    {s.city}, {s.country}
                  </p>
                </div>
                <span className={`signal ${s.status === "online" ? "online" : "offline"}`} title={s.status} />
              </div>

              <div style={{ display: "flex", gap: 8, margin: "14px 0" }}>
                <span className="cyber-badge violet">{s.protocol}</span>
                <span className="cyber-badge cyan cyber-mono" style={{ color: latencyColor(s.latency), borderColor: latencyColor(s.latency) }}>
                  {s.latency} ms
                </span>
              </div>

              <div className="cyber-mono" style={{ fontSize: 11, color: "var(--fog)", marginBottom: 14 }}>
                {s.current_users}/{s.max_users} usuarios conectados
              </div>

              <button
                className="cyber-btn full"
                onClick={() => connect(s)}
                disabled={connectingId === s.id || s.status !== "online"}
              >
                {connectingId === s.id ? "Conectando..." : "Conectar"}
              </button>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
