package com.korzadi.vpn.model

import com.google.gson.JsonElement

data class DashboardResponse(
    val user: UserInfo?,
    val plan: JsonElement?,
    val statistics: Statistics?,
    val sessions: JsonElement?,
    val activity: JsonElement?,
    val vpn: JsonElement?
)

data class UserInfo(
    val email: String?,
    val plan: String?,
    val status: String?
)

data class Statistics(
    val total_devices: Int?,
    val active_connections: Int?
)
