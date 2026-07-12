import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import api from "../api/api";
import Header from "../components/Header";
import ConnectionStatus from "../components/ConnectionStatus";
import ServerList from "../components/ServerList";
import "../cyberpunk.css";

export default function Dashboard() {
  const navigate = useNavigate();
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);
  const [connected, setConnected] = useState(false);
  const [busy, setBusy] = useState(false);
  const [error, setError] = useState("");

  useEffect(() => {
    fetchUser();
  }, []);

  async function fetchUser() {
    try {
      const res = await api.get("/api/user");
      setUser(res.data);
    } catch (err) {
      console.error("Error al obtener usuario:", err);
      localStorage.removeItem("token");
      window.location.reload();
    } finally {
      setLoading(false);
    }
  }

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

  if (loading) {
    return (
      <div style={{
        minHeight: "100vh",
        background: "#0f172a",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        color: "#38bdf8"
      }}>
        Cargando...
      </div>
    );
  }

  return (
    <div style={{
      minHeight: "100vh",
      background: "#0f172a"
    }}>
      <Header user={user} />

      <div style={{
        maxWidth: "1200px",
        margin: "0 auto",
        padding: "30px 20px"
      }}>
        <h1 style={{ color: "#e2e8f0", marginBottom: "30px" }}>
          Bienvenido, {user?.name || user?.email}! 👋
        </h1>

        {/* Estado de conexión mejorado */}
        <div style={{
          background: "#1e293b",
          padding: "30px",
          borderRadius: "10px",
          border: "1px solid #334155",
          marginBottom: "20px",
          textAlign: "center"
        }}>
          <h3 style={{ color: "#38bdf8", margin: "0 0 20px 0" }}>
            Estado de conexión
          </h3>

          <div style={{
            width: "120px",
            height: "120px",
            margin: "0 auto 20px",
            borderRadius: "50%",
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            border: `3px solid ${connected ? "#22c55e" : "#ef4444"}`,
            background: connected ? "rgba(34, 197, 94, 0.1)" : "rgba(239, 68, 68, 0.1)",
          }}>
            <span style={{
              fontSize: "48px"
            }}>
              {connected ? "🔓" : "🔒"}
            </span>
          </div>

          <p style={{
            color: connected ? "#22c55e" : "#ef4444",
            fontSize: "18px",
            fontWeight: "bold",
            margin: "0 0 20px 0"
          }}>
            {connected ? "CONECTADO" : "DESCONECTADO"}
          </p>

          <div style={{ display: "flex", gap: "10px", justifyContent: "center", flexWrap: "wrap" }}>
            {!connected ? (
              <button
                onClick={conectar}
                disabled={busy}
                style={{
                  padding: "12px 30px",
                  background: "#38bdf8",
                  color: "white",
                  border: "none",
                  borderRadius: "8px",
                  fontWeight: "bold",
                  cursor: "pointer"
                }}
              >
                {busy ? "Conectando..." : "Conectar"}
              </button>
            ) : (
              <button
                onClick={desconectar}
                disabled={busy}
                style={{
                  padding: "12px 30px",
                  background: "#ef4444",
                  color: "white",
                  border: "none",
                  borderRadius: "8px",
                  fontWeight: "bold",
                  cursor: "pointer"
                }}
              >
                {busy ? "Desconectando..." : "Desconectar"}
              </button>
            )}

            <button
              onClick={descargar}
              style={{
                padding: "12px 30px",
                background: "#10b981",
                color: "white",
                border: "none",
                borderRadius: "8px",
                fontWeight: "bold",
                cursor: "pointer"
              }}
            >
              Descargar WireGuard
            </button>
          </div>

          {error && (
            <p style={{
              color: "#ef4444",
              fontSize: "14px",
              marginTop: "15px"
            }}>
              ⚠️ {error}
            </p>
          )}
        </div>

        <ConnectionStatus />
        <ServerList />

        {/* Stats */}
        <div style={{
          display: "grid",
          gridTemplateColumns: "repeat(auto-fit, minmax(200px, 1fr))",
          gap: "20px",
          marginTop: "30px"
        }}>
          <div style={{
            background: "#1e293b",
            padding: "20px",
            borderRadius: "10px",
            border: "1px solid #334155"
          }}>
            <p style={{ color: "#cbd5e1", margin: "0 0 10px 0" }}>Datos enviados</p>
            <h2 style={{ color: "#38bdf8", margin: 0 }}>0 GB</h2>
          </div>

          <div style={{
            background: "#1e293b",
            padding: "20px",
            borderRadius: "10px",
            border: "1px solid #334155"
          }}>
            <p style={{ color: "#cbd5e1", margin: "0 0 10px 0" }}>Datos recibidos</p>
            <h2 style={{ color: "#38bdf8", margin: 0 }}>0 GB</h2>
          </div>

          <div style={{
            background: "#1e293b",
            padding: "20px",
            borderRadius: "10px",
            border: "1px solid #334155"
          }}>
            <p style={{ color: "#cbd5e1", margin: "0 0 10px 0" }}>Tiempo conectado</p>
            <h2 style={{ color: "#38bdf8", margin: 0 }}>0h 00m</h2>
          </div>

          <div style={{
            background: "#1e293b",
            padding: "20px",
            borderRadius: "10px",
            border: "1px solid #334155"
          }}>
            <p style={{ color: "#cbd5e1", margin: "0 0 10px 0" }}>Plan</p>
            <h2 style={{ color: "#38bdf8", margin: 0 }}>{user?.plan || "Gratis"}</h2>
          </div>
        </div>

        {/* Accesos rápidos */}
        <div style={{
          display: "grid",
          gridTemplateColumns: "repeat(auto-fit, minmax(200px, 1fr))",
          gap: "20px",
          marginTop: "30px"
        }}>
          <div
            onClick={() => navigate("/servers")}
            style={{
              background: "#1e293b",
              padding: "20px",
              borderRadius: "10px",
              border: "1px solid #334155",
              cursor: "pointer",
              transition: "all 0.3s"
            }}
          >
            <p style={{ color: "#38bdf8", fontSize: "12px", margin: "0 0 10px 0" }}>01</p>
            <h3 style={{ color: "#e2e8f0", margin: "0 0 8px 0" }}>Servidores</h3>
            <p style={{ fontSize: "12px", color: "#cbd5e1", margin: 0 }}>Elige el nodo al que te conectas</p>
          </div>

          <div
            onClick={() => navigate("/devices")}
            style={{
              background: "#1e293b",
              padding: "20px",
              borderRadius: "10px",
              border: "1px solid #334155",
              cursor: "pointer",
              transition: "all 0.3s"
            }}
          >
            <p style={{ color: "#38bdf8", fontSize: "12px", margin: "0 0 10px 0" }}>02</p>
            <h3 style={{ color: "#e2e8f0", margin: "0 0 8px 0" }}>Dispositivos</h3>
            <p style={{ fontSize: "12px", color: "#cbd5e1", margin: 0 }}>Gestiona tus sesiones vinculadas</p>
          </div>

          <div
            onClick={() => navigate("/plans")}
            style={{
              background: "#1e293b",
              padding: "20px",
              borderRadius: "10px",
              border: "1px solid #334155",
              cursor: "pointer",
              transition: "all 0.3s"
            }}
          >
            <p style={{ color: "#38bdf8", fontSize: "12px", margin: "0 0 10px 0" }}>03</p>
            <h3 style={{ color: "#e2e8f0", margin: "0 0 8px 0" }}>Plan</h3>
            <p style={{ fontSize: "12px", color: "#cbd5e1", margin: 0 }}>Revisa o cambia tu suscripción</p>
          </div>
        </div>
      </div>
    </div>
  );
}
