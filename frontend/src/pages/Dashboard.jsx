import { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../api/api";
import "../cyberpunk.css";

export default function Dashboard() {
  const navigate = useNavigate();
  const [connected, setConnected] = useState(false);
  const [busy, setBusy] = useState(false);
  const [error, setError] = useState("");

  async function conectar() {
    setBusy(true);
    setError("");
    try {
      await api.post("/api/connect", { server_id: 1, device: "Android" });
      setConnected(true);
    } catch (e) {
      setError("No se pudo conectar. Intenta de nuevo.");
    } finally {
      setBusy(false);
    }
  }

  async function desconectar() {
    setBusy(true);
    setError("");
    try {
      await api.post("/api/disconnect");
      setConnected(false);
    } catch (e) {
      setError("No se pudo desconectar.");
    } finally {
      setBusy(false);
    }
  }

  function descargar() {
    window.open(
      "http://localhost:8080/api/vpn/profile/download?token=" + localStorage.getItem("token")
    );
  }

  function logout() {
    localStorage.removeItem("token");
    window.location.reload();
  }

  return (
    <div className="cyber-page">
      <div className="cyber-shell">
        <div style={{ display: "flex", justifyContent: "space-between", alignItems: "flex-start" }}>
          <div>
            <h1 className="cyber-h1">KorzadiVPN</h1>
            <p className="cyber-sub">&gt; panel_de_control</p>
          </div>
          <button className="cyber-btn ghost" style={{ padding: "8px 14px", fontSize: 11 }} onClick={logout}>
            Cerrar sesión
          </button>
        </div>

        {/* Estado de conexión */}
        <div className="cyber-card" style={{ textAlign: "center", padding: 40, marginBottom: 24 }}>
          <div
            style={{
              width: 90,
              height: 90,
              margin: "0 auto 20px",
              borderRadius: "50%",
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              border: `2px solid ${connected ? "var(--acid)" : "var(--fog)"}`,
              boxShadow: connected ? "0 0 30px rgba(57,255,20,0.35)" : "none",
              position: "relative",
            }}
          >
            <span className={`signal ${connected ? "online" : "offline"}`} style={{ width: 16, height: 16 }} />
          </div>

          <p className="cyber-mono" style={{ fontSize: 13, letterSpacing: 1, color: connected ? "var(--acid)" : "var(--fog)", marginBottom: 24 }}>
            {connected ? "TÚNEL ACTIVO" : "DESCONECTADO"}
          </p>

          <div style={{ display: "flex", gap: 12, justifyContent: "center", flexWrap: "wrap" }}>
            {!connected ? (
              <button className="cyber-btn" onClick={conectar} disabled={busy}>
                {busy ? "Conectando..." : "Conectar"}
              </button>
            ) : (
              <button className="cyber-btn danger" onClick={desconectar} disabled={busy}>
                {busy ? "Desconectando..." : "Desconectar"}
              </button>
            )}
            <button className="cyber-btn ghost" onClick={descargar}>
              Descargar perfil WireGuard
            </button>
          </div>

          {error && (
            <p className="cyber-mono" style={{ color: "var(--hot)", fontSize: 12, marginTop: 16 }}>
              ⚠ {error}
            </p>
          )}
        </div>

        {/* Accesos rápidos */}
        <div className="cyber-grid">
          <div className="cyber-card" style={{ cursor: "pointer" }} onClick={() => navigate("/servers")}>
            <p className="cyber-mono" style={{ color: "var(--violet)", fontSize: 12, marginBottom: 6 }}>01</p>
            <h2 style={{ fontFamily: "'Orbitron',sans-serif", fontSize: 15, margin: "0 0 4px" }}>Servidores</h2>
            <p style={{ fontSize: 12, color: "var(--fog)", margin: 0 }}>Elige el nodo al que te conectas</p>
          </div>

          <div className="cyber-card" style={{ cursor: "pointer" }} onClick={() => navigate("/devices")}>
            <p className="cyber-mono" style={{ color: "var(--cyan)", fontSize: 12, marginBottom: 6 }}>02</p>
            <h2 style={{ fontFamily: "'Orbitron',sans-serif", fontSize: 15, margin: "0 0 4px" }}>Dispositivos</h2>
            <p style={{ fontSize: 12, color: "var(--fog)", margin: 0 }}>Gestiona tus sesiones vinculadas</p>
          </div>

          <div className="cyber-card" style={{ cursor: "pointer" }} onClick={() => navigate("/plans")}>
            <p className="cyber-mono" style={{ color: "var(--acid)", fontSize: 12, marginBottom: 6 }}>03</p>
            <h2 style={{ fontFamily: "'Orbitron',sans-serif", fontSize: 15, margin: "0 0 4px" }}>Plan</h2>
            <p style={{ fontSize: 12, color: "var(--fog)", margin: 0 }}>Revisa o cambia tu suscripción</p>
          </div>
        </div>
      </div>
    </div>
  );
}
