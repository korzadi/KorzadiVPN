package com.korzadi.vpn.data.model

data class LoginRequest(
    val username: String,
    val password: String
)

data class LoginResponse(
    val token: String,
    val userId: String
)
