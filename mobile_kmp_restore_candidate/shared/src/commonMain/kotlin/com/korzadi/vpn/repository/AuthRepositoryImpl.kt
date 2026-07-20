package com.korzadi.vpn.repository

import com.korzadi.vpn.api.ApiClient
import com.korzadi.vpn.model.AuthRequest
import com.korzadi.vpn.model.AuthResponse
import com.korzadi.vpn.storage.SecureStorage
import io.ktor.client.call.body
import io.ktor.client.request.post
import io.ktor.client.request.setBody
import io.ktor.http.ContentType
import io.ktor.http.contentType

class AuthRepositoryImpl(
    private val api: ApiClient,
    private val storage: SecureStorage
) : AuthRepository {

    override suspend fun register(email: String, pass: String): AuthResponse {
        return api.client.post("${api.baseUrl}/api/register") {
            contentType(ContentType.Application.Json)
            setBody(AuthRequest(email, pass))
        }.body()
    }

    override suspend fun login(email: String, pass: String): AuthResponse {
        val response: AuthResponse =
            api.client.post("${api.baseUrl}/api/login") {
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
