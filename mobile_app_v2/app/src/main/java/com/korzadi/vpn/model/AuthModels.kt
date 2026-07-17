package com.korzadi.vpn.model

data class LoginRequest(
    val email: String,
    val password: String
)

data class LoginResponse(
    val message: String,
    val token: String?,
    val plan: String?,
    val status: String?
)

data class RegisterRequest(
    val email: String,
    val password: String
)

data class RegisterResponse(
    val message: String,
    val email: String?,
    val plan: String?,
    val status: String?,
    val vpn: String?
)
