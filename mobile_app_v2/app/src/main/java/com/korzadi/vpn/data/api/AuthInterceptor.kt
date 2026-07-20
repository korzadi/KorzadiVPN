package com.korzadi.vpn.data.api

import okhttp3.Interceptor
import okhttp3.Response

class AuthInterceptor : Interceptor {
    var token: String? = null

    override fun intercept(chain: Interceptor.Chain): Response {
        val originalRequest = chain.request()
        val requestBuilder = originalRequest.newBuilder()

        token?.let {
            requestBuilder.addHeader("Authorization", "Bearer $it")
        }

        return chain.proceed(requestBuilder.build())
    }
}
