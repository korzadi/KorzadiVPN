import { useEffect, useState } from "react";
import api from "../api/api";
import { useNavigate } from "react-router-dom";
import "../cyberpunk.css";

const ACCENTS = {
  free: "var(--fog)",
  premium: "var(--violet)",
  enterprise: "var(--cyan)",
};

export default function Plans() {
  const navigate = useNavigate();
  const [plans, setPlans] = useState([]);
  const [currentPlan, setCurrentPlan] = useState(null);
  const [loading, setLoading] = useState(true);
  const [changingTo, setChangingTo] = useState(null);
  const [error, setError] = useState("");

  useEffect(() => {
    load();
  }, []);

  async function load() {
    setLoading(true);
    try {
      const [plansRes, profileRes] = await Promise.all([
        api.get("/api/plans"),
        api.get("/api/profile"),
      ]);
      setPlans(plansRes.data || []);
      setCurrentPlan(profileRes.data?.plan || null);
    } catch (err) {
      setError("No se pudo cargar la información de planes.");
    } finally {
      setLoading(false);
    }
  }

  async function changePlan(name) {
    setChangingTo(name);
    setError("");
    try {
      // Verifica en change_plan.go el nombre exacto del campo esperado
      // (aquí se asume "plan"); ajusta si tu handler usa otro.
      await api.post("/api/change-plan", { plan: name });
      setCurrentPlan(name);
    } catch (err) {
      setError(err?.response?.data?.error || "No se pudo cambiar de plan.");
    } finally {
      setChangingTo(null);
    }
  }

  return (
    <div className="cyber-page">
      <div className="cyber-shell">
        <button className="cyber-back" onClick={() => navigate("/")}>← volver al dashboard</button>

        <h1 className="cyber-h1">Planes</h1>
        <p className="cyber-sub">&gt; selecciona_tu_nivel_de_acceso</p>

        {loading && <p className="cyber-empty">Cargando planes…</p>}
        {!loading && error && <p className="cyber-empty" style={{ color: "var(--hot)" }}>{error}</p>}

        {!loading && (
          <div className="cyber-grid">
            {plans.map((p) => {
              const accent = ACCENTS[p.name] || "var(--violet)";
              const active = currentPlan === p.name;
              return (
                <div
                  className="cyber-card"
                  key={p.name}
                  style={active ? { borderColor: accent, boxShadow: `0 0 24px ${accent}33` } : {}}
                >
                  <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center" }}>
                    <span
                      className="cyber-mono"
                      style={{ textTransform: "uppercase", letterSpacing: 1, fontSize: 13, color: accent }}
                    >
                      {p.name}
                    </span>
                    {active && <span className="cyber-badge acid">actual</span>}
                  </div>

                  <p style={{ fontSize: 13, color: "#c8cad6", margin: "14px 0" }}>{p.description}</p>

                  <p className="cyber-mono" style={{ fontSize: 12, color: "var(--fog)", marginBottom: 20 }}>
                    hasta {p.max_devices} dispositivo{p.max_devices > 1 ? "s" : ""}
                  </p>

                  <button
                    className={`cyber-btn full ${active ? "ghost" : ""}`}
                    onClick={() => changePlan(p.name)}
                    disabled={active || changingTo === p.name}
                  >
                    {active ? "Plan activo" : changingTo === p.name ? "Cambiando..." : "Elegir plan"}
                  </button>
                </div>
              );
            })}
          </div>
        )}
      </div>
    </div>
  );
}
