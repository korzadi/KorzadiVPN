export default function Header({ user }) {
  function logout() {
    localStorage.removeItem("token");
    window.location.reload();
  }

  return (
    <header style={{
      background: "#1e293b",
      borderBottom: "1px solid #334155",
      padding: "15px 30px",
      display: "flex",
      justifyContent: "space-between",
      alignItems: "center"
    }}>
      <div>
        <h2 style={{ color: "#38bdf8", margin: 0 }}>
          🔒 KorzadiVPN
        </h2>
      </div>

      <div style={{
        display: "flex",
        alignItems: "center",
        gap: "20px"
      }}>
        <span style={{ color: "#cbd5e1" }}>
          {user?.email}
        </span>
        
        <button
          onClick={logout}
          style={{
            padding: "8px 15px",
            background: "#ef4444",
            color: "white",
            border: "none",
            borderRadius: "5px",
            cursor: "pointer"
          }}
        >
          Cerrar sesión
        </button>
      </div>
    </header>
  );
}
