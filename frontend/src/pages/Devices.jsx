import { useEffect, useState } from "react";
import api from "../api/api";
import { useNavigate } from "react-router-dom";
import "../cyberpunk.css";

const ICONS = {
  mobile: "📱",
  desktop: "🖥️",
  laptop: "💻",
  router: "📡",
};

export default function Devices() {
  const navigate = useNavigate();
  const [devices, setDevices] = useState([]);
  const [loading, setLoading] = useState(true);
  const [deletingId, setDeletingId] = useState(null);
  const [error, setError] = useState("");

  useEffect(() => {
    loadDevices();
  }, []);

  async function loadDevices() {
    setLoading(true);
    try {
      const res = await api.get("/api/user/devices");
      setDevices(res.data || []);
    } catch (err) {
      setError("No se pudieron cargar tus dispositivos.");
    } finally {
      setLoading(false);
    }
  }

  async function remove(id) {
    if (!confirm("¿Eliminar este dispositivo? Se cerrará su sesión VPN.")) return;
    setDeletingId(id);
    try {
      await api.delete(`/api/user/device/${id}`);
      setDevices((prev) => prev.filter((d) => d.id !== id));
    } catch (err) {
      setError("No se pudo eliminar el dispositivo.");
    } finally {
      setDeletingId(null);
    }
  }

  return (
    <div className="cyber-page">
      <div className="cyber-shell">
        <button className="cyber-back" onClick={() => navigate("/")}>← volver al dashboard</button>

        <h1 className="cyber-h1">Dispositivos</h1>
        <p className="cyber-sub">&gt; dispositivos_vinculados_a_tu_cuenta</p>

        {loading && <p className="cyber-empty">Cargando dispositivos…</p>}
        {!loading && error && <p className="cyber-empty" style={{ color: "var(--hot)" }}>{error}</p>}
        {!loading && !error && devices.length === 0 && (
          <p className="cyber-empty">Aún no has conectado ningún dispositivo.</p>
        )}

        {!loading && devices.length > 0 && (
          <div className="cyber-card">
            {devices.map((d) => (
              <div className="cyber-row" key={d.id}>
                <div style={{ display: "flex", alignItems: "center", gap: 12 }}>
                  <span style={{ fontSize: 22 }}>{ICONS[d.device_type] || "🔌"}</span>
                  <div>
                    <div style={{ fontWeight: 600, fontSize: 14 }}>{d.device_name}</div>
                    <div className="cyber-mono" style={{ fontSize: 11, color: "var(--fog)" }}>
                      {d.last_server ? `${d.last_server} · ` : ""}{d.last_ip || "sin IP registrada"}
                    </div>
                    <div className="cyber-mono" style={{ fontSize: 11, color: "var(--fog)" }}>
                      últ. conexión: {d.last_seen || "—"}
                    </div>
                  </div>
                </div>

                <div style={{ display: "flex", alignItems: "center", gap: 14 }}>
                  <span className={`signal ${d.status === "online" ? "online" : "offline"}`} />
                  <button
                    className="cyber-btn danger"
                    style={{ padding: "8px 14px", fontSize: 11 }}
                    onClick={() => remove(d.id)}
                    disabled={deletingId === d.id}
                  >
                    {deletingId === d.id ? "..." : "Eliminar"}
                  </button>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
