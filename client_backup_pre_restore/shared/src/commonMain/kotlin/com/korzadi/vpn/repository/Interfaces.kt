package com.korzadi.vpn.repository

import com.korzadi.vpn.model.*

interface AuthRepository {
    suspend fun register(email: String, pass: String): AuthResponse
    suspend fun login(email: String, pass: String): AuthResponse
    suspend fun logout()
}

interface VPNRepository {
    suspend fun createClient(): VPNClientResponse
    suspend fun connect(clientId: Int): VPNConnectionResponse
    suspend fun disconnect(clientId: Int): VPNConnectionResponse
    suspend fun getStatus(clientId: Int): VPNStatusResponse
}

interface WireGuardManager {
    fun connect(config: String)
    fun disconnect()
    fun status(): String
}
