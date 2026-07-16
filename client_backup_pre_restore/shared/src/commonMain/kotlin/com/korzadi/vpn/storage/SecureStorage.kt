package com.korzadi.vpn.storage

interface SecureStorage {
    fun saveToken(token: String)
    fun getToken(): String?
    fun clear()
}
