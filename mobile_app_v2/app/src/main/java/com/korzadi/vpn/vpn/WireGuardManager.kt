package com.korzadi.vpn.vpn

import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class WireGuardManager @Inject constructor() {
    enum class ConnectionState {
        DISCONNECTED, CONNECTING, CONNECTED, ERROR
    }

    // In a real app, integrate with WireGuard's JNI/Go library
    fun startTunnel(config: String) {
        // Implementation
    }

    fun stopTunnel() {
        // Implementation
    }
}
