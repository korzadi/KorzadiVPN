# KorzadiVPN 🔐

Una solución VPN completa y segura con backend en Go, frontend en React y aplicación móvil Android.

## 🌟 Características

- ✅ Backend REST API en Go
- ✅ Frontend web responsive en React + Vite
- ✅ Aplicación móvil Android nativa
- ✅ Autenticación JWT
- ✅ Base de datos SQLite
- ✅ Panel de administración
- ✅ Gestión de dispositivos
- ✅ Perfiles VPN personalizados
- ✅ Sistema de planes y suscripciones
- ✅ Estadísticas de conexión en tiempo real

## 📁 Estructura del Proyecto

KorzadiVPN/
├── backend/          # API REST en Go
│   ├── handlers/     # Controladores
│   ├── models/       # Modelos de datos
│   ├── database/     # Capa de datos
│   ├── middleware/   # Middleware (JWT, Auth)
│   ├── routes/       # Rutas
│   ├── utils/        # Utilidades
│   └── main.go
├── frontend/         # Web en React + Vite
├── android/          # App móvil Android
├── api/              # Documentación API
├── dashboard/        # Panel admin
├── database/         # Scripts de DB
├── docs/             # Documentación
├── vpn-core/         # Core VPN
└── website/          # Sitio web

## 🚀 Instalación

### Backend (Go)
cd backend
go mod download
go run main.go

### Frontend (React)
cd frontend
npm install
npm run dev

### Android
Abrir en Android Studio

## 🔧 Requisitos

- Go 1.20+
- Node.js 16+
- Android SDK
- SQLite3

## 📚 API Endpoints

- POST /api/auth/register - Registro de usuario
- POST /api/auth/login - Login
- GET /api/vpn/servers - Listar servidores
- POST /api/vpn/connect - Conectar a VPN
- GET /api/user/dashboard - Dashboard del usuario
- GET /api/admin/stats - Estadísticas (Admin)

## 🔐 Seguridad

- Autenticación JWT
- Middleware de autorización
- Encriptación de contraseñas
- Validación de entrada

## 📄 Licencia

MIT License - Ver archivo LICENSE para más detalles

## 👤 Autor

Korzadi - https://github.com/korzadi

¿Necesitas ayuda? Abre un Issue
