package com.korzadi.vpn.api

import android.content.Context
import com.korzadi.vpn.data.TokenManager
import okhttp3.OkHttpClient
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory

object RetrofitClient {

    private const val BASE_URL =
        "https://vpn.korzadi.com/"

    fun create(context: Context): KorzadiApi {

        val client =
            OkHttpClient.Builder()
                .addInterceptor(
                    AuthInterceptor(
                        TokenManager(context)
                    )
                )
                .build()

        return Retrofit.Builder()
            .baseUrl(BASE_URL)
            .client(client)
            .addConverterFactory(
                GsonConverterFactory.create()
            )
            .build()
            .create(KorzadiApi::class.java)
    }
}
