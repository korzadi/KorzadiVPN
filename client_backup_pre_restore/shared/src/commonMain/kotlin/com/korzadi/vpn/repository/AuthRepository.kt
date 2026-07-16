package com.korzadi.vpn.repository

import com.korzadi.vpn.api.ApiClient
import com.korzadi.vpn.model.*
import io.ktor.client.call.*
import io.ktor.client.request.*
import io.ktor.http.*

class AuthRepositoryImpl(private val api: ApiClient) : AuthRepository {
    override suspend fun register(email: String, pass: String): AuthResponse {
        return api.client.post("${api.baseUrl}/api/register") {
            contentType(ContentType.Application.Json)
            setBody(AuthRequest(email, pass))
        }.body()
    }

    override suspend fun login(email: String, pass: String): AuthResponse {
        return api.client.post("${api.baseUrl}/api/login") {
            contentType(ContentType.Application.Json)
            setBody(AuthRequest(email, pass))
        }.body()
    }

    override suspend fun logout() {
        // Implementar lógica de logout
    }
}
