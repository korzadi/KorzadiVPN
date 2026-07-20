package com.korzadi.vpn.viewmodel

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.korzadi.vpn.data.local.TokenManager
import com.korzadi.vpn.data.repository.VPNRepository
import com.korzadi.vpn.vpn.WireGuardManager
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import kotlinx.coroutines.runBlocking
import javax.inject.Inject

@HiltViewModel
class VPNViewModel @Inject constructor(
    private val repository: VPNRepository,
    private val wireGuardManager: WireGuardManager,
    private val tokenManager: TokenManager
) : ViewModel() {

    private val _isAuthenticated = MutableStateFlow(tokenManager.getToken() != null)
    val isAuthenticated = _isAuthenticated.asStateFlow()

    private val _connectionState = MutableStateFlow(WireGuardManager.ConnectionState.DISCONNECTED)
    val connectionState = _connectionState.asStateFlow()

    fun register(username: String, email: String, password: String): Boolean {
        var success = false
        runBlocking {
            success = repository.register(username, email, password)
        }
        return success
    }

    fun connectAuth(username: String, password: String): Boolean {
        var success = false
        runBlocking {
            if (repository.login(username, password)) {
                _isAuthenticated.value = true
                success = true
            }
        }
        return success
    }

    fun startVpn() {
        viewModelScope.launch {
            _connectionState.value = WireGuardManager.ConnectionState.CONNECTING
            val config = repository.getVpnConfig()
            if (config != null) {
                wireGuardManager.startTunnel(config.config, "KorzadiTunnel")
                _connectionState.value = WireGuardManager.ConnectionState.CONNECTED
            } else {
                _connectionState.value = WireGuardManager.ConnectionState.ERROR
            }
        }
    }

    fun disconnect() {
        wireGuardManager.stopTunnel()
        _connectionState.value = WireGuardManager.ConnectionState.DISCONNECTED
    }
}
