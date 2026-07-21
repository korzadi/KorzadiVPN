package com.korzadi.vpn.data.api

import com.korzadi.vpn.data.model.LoginRequest
import com.korzadi.vpn.data.model.LoginResponse
import com.korzadi.vpn.data.model.RegisterRequest
import com.korzadi.vpn.data.model.RegisterResponse
import com.korzadi.vpn.data.model.VpnProfile
import com.korzadi.vpn.data.model.WireGuardConfig
import retrofit2.http.Body
import retrofit2.http.GET
import retrofit2.http.POST

interface KorzadiApi {

    @POST("/api/login")
    suspend fun login(
        @Body request: LoginRequest
    ): LoginResponse


    @POST("/api/register")
    suspend fun register(
        @Body request: RegisterRequest
    ): RegisterResponse


    @GET("/api/vpn/profile")
    suspend fun getVpnProfile(): VpnProfile


    @POST("/api/vpn/client/create")
    suspend fun createClientConfig(): WireGuardConfig
}
