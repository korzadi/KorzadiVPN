import { useState } from "react";
import api from "../api/api";

export default function Login() {

  const [email,setEmail]=useState("");
  const [password,setPassword]=useState("");

  async function login(){

    try{

      const res=await api.post("/api/login",{
        email,
        password
      });

      localStorage.setItem("token",res.data.token);

      window.location.reload();

    }catch(err){

      alert("Correo o contraseña incorrectos");

    }

  }

  return(

    <div style={{
      minHeight:"100vh",
      background:"#0f172a",
      display:"flex",
      justifyContent:"center",
      alignItems:"center"
    }}>

      <div style={{
        width:"340px",
        background:"#1e293b",
        padding:"30px",
        borderRadius:"20px"
      }}>

        <h1 style={{
          color:"#38bdf8",
          textAlign:"center"
        }}>
          KorzadiVPN
        </h1>

        <input
          placeholder="Correo electrónico"
          value={email}
          onChange={(e)=>setEmail(e.target.value)}
          style={{
            width:"100%",
            padding:"15px",
            marginTop:"20px"
          }}
        />

        <input
          type="password"
          placeholder="Contraseña"
          value={password}
          onChange={(e)=>setPassword(e.target.value)}
          style={{
            width:"100%",
            padding:"15px",
            marginTop:"15px"
          }}
        />

        <button
          onClick={login}
          style={{
            width:"100%",
            marginTop:"20px",
            padding:"15px",
            background:"#38bdf8",
            color:"white",
            border:"none",
            borderRadius:"10px"
          }}
        >
          Iniciar sesión
        </button>

      </div>

    </div>

  );

}
