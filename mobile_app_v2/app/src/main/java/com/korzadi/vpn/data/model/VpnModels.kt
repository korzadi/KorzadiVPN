package com.korzadi.vpn.data.model

data class VpnProfile(
    val id: String,
    val name: String,
    val serverAddress: String
)

data class WireGuardConfig(
    val config: String
)
