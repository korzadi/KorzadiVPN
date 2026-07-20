package com.korzadi.vpn.repository

import com.korzadi.vpn.api.ApiClient
import com.korzadi.vpn.model.*
import com.korzadi.vpn.storage.SecureStorage
import io.ktor.client.call.*
import io.ktor.client.request.*
import io.ktor.http.*

class VPNRepositoryImpl(
    private val api: ApiClient,
    private val storage: SecureStorage
) : VPNRepository {

    override suspend fun createClient(): VPNClientResponse {
        return api.client.post("${api.baseUrl}/api/vpn/client/create") {
            contentType(ContentType.Application.Json)
            header("Authorization", "Bearer ${storage.getToken()}")
        }.body()
    }

    override suspend fun connect(clientId: Int): VPNConnectionResponse {
        return api.client.post("${api.baseUrl}/api/connect") {
            contentType(ContentType.Application.Json)
            header("Authorization", "Bearer ${storage.getToken()}")
        }.body()
    }

    override suspend fun disconnect(clientId: Int): VPNConnectionResponse {
        return api.client.post("${api.baseUrl}/api/disconnect") {
            contentType(ContentType.Application.Json)
            header("Authorization", "Bearer ${storage.getToken()}")
        }.body()
    }

    override suspend fun getStatus(clientId: Int): VPNStatusResponse {
        return api.client.get("${api.baseUrl}/api/vpn/status") {
            header("Authorization", "Bearer ${storage.getToken()}")
        }.body()
    }
}
