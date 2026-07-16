package com.korzadi.vpn.vpn

expect class WireGuardManager() {
    fun connect(config: VPNConfig)
    fun disconnect()
    fun status(): VPNState
}
