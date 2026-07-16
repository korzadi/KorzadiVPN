# KorzadiVPN - Producción 100%

## Estructura
- cmd/korzadivpn: Punto de entrada.
- internal/: Lógica de negocio, base de datos, handlers.
- pkg/: Utilidades reutilizables, WireGuard real.

## Administración
- Iniciar API: sudo systemctl start korzadivpn
- Logs: sudo journalctl -u korzadivpn -f

## WireGuard Real
- El backend gestiona peers reales en wg0.
- Requisito: El usuario que ejecuta el binario debe tener privilegios para ejecutar `wg`.

## APK Android
- API lista para integración.
- Usar endpoints definidos en internal/routes/routes.go.
