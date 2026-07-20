package com.korzadi.vpn.vpn

import kotlinx.coroutines.flow.MutableStateFlow

actual class WireGuardManager actual constructor() {
    private val _state = MutableStateFlow(VPNState.DISCONNECTED)
    
    actual fun connect(config: VPNConfig) {
        _state.value = VPNState.CONNECTING
        // Lógica de wg-quick
        _state.value = VPNState.CONNECTED
    }

    actual fun disconnect() {
        _state.value = VPNState.DISCONNECTING
        // Lógica de desconexión
        _state.value = VPNState.DISCONNECTED
    }

    actual fun status(): VPNState {
        return _state.value
    }
}
