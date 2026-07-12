import api from "../api/api";

export default function Dashboard() {

async function conectar(){

try{

await api.post("/api/connect",{
server_id:1,
device:"Android"
});

alert("VPN Conectada");

}catch(e){

alert("Error al conectar");

}

}

async function desconectar(){

try{

await api.post("/api/disconnect");

alert("VPN Desconectada");

}catch(e){

alert("Error al desconectar");

}

}

async function descargar(){

window.open(
"http://localhost:8080/api/vpn/profile/download?token="+localStorage.getItem("token")
);

}

return(

<div style={{
background:"#0f172a",
minHeight:"100vh",
padding:"30px",
color:"white",
fontFamily:"Arial"
}}>

<h1>KorzadiVPN</h1>

<div style={{
background:"#1e293b",
padding:"20px",
borderRadius:"15px",
marginTop:"20px"
}}>

<h2>Estado VPN</h2>

<button onClick={conectar}>
Conectar
</button>

<button
onClick={desconectar}
style={{marginLeft:"10px"}}
>
Desconectar
</button>

<button
onClick={descargar}
style={{marginLeft:"10px"}}
>
Descargar WireGuard
</button>

</div>

</div>

);

}
