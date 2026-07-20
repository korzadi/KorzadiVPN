package com.korzadi.vpn.vpn

import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow

actual class WireGuardManager actual constructor() {
    private val _state = MutableStateFlow(VPNState.DISCONNECTED)
    
    actual fun connect(config: VPNConfig) {
        // 1. Usar SecureStorage para obtener claves (mocked/plataforma)
        // 2. Invocar Android VPNService vía JNI o plataforma
        _state.value = VPNState.CONNECTING
        // Lógica nativa de túnel aquí...
        _state.value = VPNState.CONNECTED
    }

    actual fun disconnect() {
        _state.value = VPNState.DISCONNECTING
        // Lógica nativa de desconexión aquí...
        _state.value = VPNState.DISCONNECTED
    }

    actual fun status(): VPNState {
        return _state.value
    }
}
