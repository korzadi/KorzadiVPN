package com.korzadi.vpn.model

import kotlinx.serialization.Serializable

@Serializable
data class Server(val id: Int, val name: String, val server_ip: String, val status: String)

@Serializable
data class Device(val email: String, val device_name: String, val device_type: String, val status: String)

@Serializable
data class VPNClient(
    val email: String,
    val client_ip: String,
    val public_key: String,
    val status: String
)

@Serializable
data class VPNClientResponse(
    val message: String,
    val client: VPNClient,
    val server: Server,
    val device: Device
)

@Serializable
data class VPNConnectionResponse(val message: String, val status: String)

@Serializable
data class VPNStatusResponse(val status: String, val connection_status: String)
