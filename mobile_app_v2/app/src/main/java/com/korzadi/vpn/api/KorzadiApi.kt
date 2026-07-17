package com.korzadi.vpn.api

import com.korzadi.vpn.model.*
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


    @GET("/api/user/dashboard")
    suspend fun dashboard(): DashboardResponse


    @POST("/api/vpn/client/create")
    suspend fun createVPN(
    ): VPNResponse

}
