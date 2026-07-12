import { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../api/api";
import "../cyberpunk.css";

export default function Register() {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirm, setConfirm] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  async function register() {
    setError("");

    if (!email || !password) {
      setError("Completa correo y contraseña.");
      return;
    }
    if (password !== confirm) {
      setError("Las contraseñas no coinciden.");
      return;
    }

    setLoading(true);
    try {
      await api.post("/api/register", { email, password });
      const res = await api.post("/api/login", { email, password });
      localStorage.setItem("token", res.data.token);
      navigate("/");
    } catch (err) {
      setError(
        err?.response?.data?.error ||
        "No se pudo crear la cuenta. Intenta con otro correo."
      );
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="cyber-page" style={{ display: "flex", justifyContent: "center", alignItems: "center" }}>
      <div className="cyber-card" style={{ width: 360 }}>
        <h1 className="cyber-h1" style={{ fontSize: 22, textAlign: "center" }}>KorzadiVPN</h1>
        <p className="cyber-sub" style={{ textAlign: "center", marginBottom: 24 }}>
          &gt; crear_cuenta.exe
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
        />

        <div style={{ height: 16 }} />

        <label className="cyber-label">Confirmar contraseña</label>
        <input
          className="cyber-input"
          type="password"
          placeholder="••••••••"
          value={confirm}
          onChange={(e) => setConfirm(e.target.value)}
          onKeyDown={(e) => e.key === "Enter" && register()}
        />

        {error && (
          <p className="cyber-mono" style={{ color: "var(--hot)", fontSize: 12, marginTop: 14 }}>
            ⚠ {error}
          </p>
        )}

        <div style={{ height: 22 }} />

        <button className="cyber-btn full" onClick={register} disabled={loading}>
          {loading ? "Creando cuenta..." : "Crear cuenta"}
        </button>

        <p
          className="cyber-mono"
          style={{ textAlign: "center", fontSize: 12, color: "var(--fog)", marginTop: 18, cursor: "pointer" }}
          onClick={() => navigate("/login")}
        >
          ¿Ya tienes cuenta? <span style={{ color: "var(--cyan)" }}>Inicia sesión</span>
        </p>
      </div>
    </div>
  );
}
