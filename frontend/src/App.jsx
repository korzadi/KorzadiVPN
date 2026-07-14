import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Dashboard from "./pages/Dashboard";
import Servers from "./pages/Servers";
import Devices from "./pages/Devices";
import Plans from "./pages/Plans";

function Private({ children }) {
  const token = localStorage.getItem("token");
  return token ? children : <Navigate to="/login" replace />;
}

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />

        <Route path="/" element={<Private><Dashboard /></Private>} />
        <Route path="/servers" element={<Private><Servers /></Private>} />
        <Route path="/devices" element={<Private><Devices /></Private>} />
        <Route path="/plans" element={<Private><Plans /></Private>} />

        <Route path="*" element={<Navigate to="/" replace />} />
      </Routes>
    </BrowserRouter>
  );
}
