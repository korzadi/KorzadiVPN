import { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../api/api";
import "../cyberpunk.css";

export default function Login() {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  async function login() {
    setError("");
    setLoading(true);
    try {
      const res = await api.post("/api/login", { email, password });
      localStorage.setItem("token", res.data.token);
      navigate("/");
    } catch (err) {
      setError("Correo o contraseña incorrectos.");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="cyber-page" style={{ display: "flex", justifyContent: "center", alignItems: "center" }}>
      <div className="cyber-card" style={{ width: 360 }}>
        <h1 className="cyber-h1" style={{ fontSize: 22, textAlign: "center" }}>KorzadiVPN</h1>
        <p className="cyber-sub" style={{ textAlign: "center", marginBottom: 24 }}>
          &gt; iniciar_sesion.exe
        </p>

        <label className="cyber-label">Correo electrónico</label>
        <input
          className="cyber-input"
          placeholder="tu@correo.com"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />

        <div style={{ height: 16 }} />

        <label className="cyber-label">Contraseña</label>
        <input
          className="cyber-input"
          type="password"
          placeholder="••••••••"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          onKeyDown={(e) => e.key === "Enter" && login()}
        />

        {error && (
          <p className="cyber-mono" style={{ color: "var(--hot)", fontSize: 12, marginTop: 14 }}>
            ⚠ {error}
          </p>
        )}

        <div style={{ height: 22 }} />

        <button className="cyber-btn full" onClick={login} disabled={loading}>
          {loading ? "Verificando..." : "Iniciar sesión"}
        </button>

        <p
          className="cyber-mono"
          style={{ textAlign: "center", fontSize: 12, color: "var(--fog)", marginTop: 18, cursor: "pointer" }}
          onClick={() => navigate("/register")}
        >
          ¿No tienes cuenta? <span style={{ color: "var(--cyan)" }}>Regístrate</span>
        </p>
      </div>
    </div>
  );
}
