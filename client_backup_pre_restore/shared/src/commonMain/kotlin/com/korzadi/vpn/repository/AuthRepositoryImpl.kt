package com.korzadi.vpn.repository

import com.korzadi.vpn.api.ApiClient
import com.korzadi.vpn.model.*
import com.korzadi.vpn.storage.SecureStorage
import io.ktor.client.call.*
import io.ktor.client.request.*
import io.ktor.http.*

class AuthRepositoryImpl(
    private val api: ApiClient,
    private val storage: SecureStorage
) : AuthRepository {
    override suspend fun register(email: String, pass: String): AuthResponse {
        val response: AuthResponse = api.client.post("${api.baseUrl}/api/register") {
            contentType(ContentType.Application.Json)
            setBody(AuthRequest(email, pass))
        }.body()
        return response
    }

    override suspend fun login(email: String, pass: String): AuthResponse {
        val response: AuthResponse = api.client.post("${api.baseUrl}/api/login") {
            contentType(ContentType.Application.Json)
            setBody(AuthRequest(email, pass))
        }.body()
        
        storage.saveToken(response.token)
        return response
    }

    override suspend fun logout() {
        storage.clear()
    }
}
