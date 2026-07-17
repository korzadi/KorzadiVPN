package com.korzadi.vpn.data

import android.content.Context

class TokenManager(context: Context) {

    private val prefs =
        context.getSharedPreferences(
            "korzadi",
            Context.MODE_PRIVATE
        )


    fun saveToken(token: String) {

        prefs.edit()
            .putString(
                "token",
                token
            )
            .apply()

    }


    fun getToken(): String? {

        return prefs.getString(
            "token",
            null
        )

    }


    fun hasToken(): Boolean {

        return !getToken().isNullOrEmpty()

    }


    fun logout() {

        prefs.edit()
            .remove("token")
            .apply()

    }

}
