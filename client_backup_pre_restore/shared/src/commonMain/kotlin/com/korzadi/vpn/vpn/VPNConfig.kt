package com.korzadi.vpn.vpn

data class VPNConfig(
    val privateKey: String,
    val publicKey: String,
    val serverPublicKey: String,
    val clientIP: String,
    val endpoint: String,
    val allowedIPs: String
)
