package com.korzadi.vpn.model


data class VPNResponse(

    val client: VPNClient?,

    val device: VPNDevice?,

    val server: VPNServer?,

    val message: String?

)


data class VPNClient(

    val id: Int?,

    val email: String?,

    val client_ip: String?,

    val public_key: String?,

    val private_key: String?,

    val config: String?,

    val endpoint: String?,

    val dns: String?,

    val allowed_ips: String?

)


data class VPNDevice(

    val device_name: String?,

    val device_type: String?,

    val status: String?

)


data class VPNServer(

    val name: String?,

    val city: String?,

    val protocol: String?,

    val server_ip: String?,

    val wireguard_port: Int?

)
