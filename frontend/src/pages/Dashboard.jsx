import { useState, useEffect } from "react";
import api from "../api/api";
import Header from "../components/Header";
import ConnectionStatus from "../components/ConnectionStatus";
import ServerList from "../components/ServerList";

export default function Dashboard() {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

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

  async function descargar() {
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

        {/* Descargar */}
        <div style={{
          marginTop: "30px",
          background: "#1e293b",
          padding: "20px",
          borderRadius: "10px",
          border: "1px solid #334155",
          textAlign: "center"
        }}>
          <h3 style={{ color: "#38bdf8", margin: "0 0 15px 0" }}>
            Configuración de clientes
          </h3>
          <button
            onClick={descargar}
            style={{
              padding: "12px 30px",
              background: "#10b981",
              color: "white",
              border: "none",
              borderRadius: "8px",
              fontWeight: "bold",
              cursor: "pointer",
              fontSize: "16px"
            }}
          >
            📥 Descargar WireGuard
          </button>
        </div>
      </div>
    </div>
  );
}
