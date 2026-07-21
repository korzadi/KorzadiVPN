package com.korzadi.vpn.data.repository

import com.korzadi.vpn.data.api.KorzadiApi
import com.korzadi.vpn.data.local.TokenManager
import com.korzadi.vpn.data.model.LoginRequest
import com.korzadi.vpn.data.model.RegisterRequest
import com.korzadi.vpn.data.model.VpnProfile
import com.korzadi.vpn.data.model.WireGuardConfig
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class VPNRepository @Inject constructor(
    private val api: KorzadiApi,
    private val tokenManager: TokenManager
) {

    suspend fun login(
        username: String,
        password: String
    ): Boolean {
        return try {
            val response = api.login(
                LoginRequest(username, password)
            )

            tokenManager.saveToken(response.token)

            true

        } catch (e: Exception) {
            false
        }
    }


    suspend fun register(
        username: String,
        email: String,
        password: String
    ): Boolean {
        return try {

            val response = api.register(
                RegisterRequest(
                    username,
                    email,
                    password
                )
            )

            response.success

        } catch (e: Exception) {
            false
        }
    }


    suspend fun getProfile(): VpnProfile? {
        return try {
            api.getVpnProfile()
        } catch (e: Exception) {
            null
        }
    }


    suspend fun getVpnConfig(): WireGuardConfig? {
        return try {
            api.createClientConfig()
        } catch (e: Exception) {
            null
        }
    }
}
