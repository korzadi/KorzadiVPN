package com.korzadi.vpn.api

import okhttp3.Interceptor
import okhttp3.Response
import com.korzadi.vpn.data.TokenManager

class AuthInterceptor(
    private val tokenManager: TokenManager
) : Interceptor {

    override fun intercept(chain: Interceptor.Chain): Response {

        val token = tokenManager.getToken()

        val request =
            chain.request()
                .newBuilder()
                .apply {
                    if (!token.isNullOrEmpty()) {
                        addHeader(
                            "Authorization",
                            "Bearer $token"
                        )
                    }
                }
                .build()

        return chain.proceed(request)
    }
}
