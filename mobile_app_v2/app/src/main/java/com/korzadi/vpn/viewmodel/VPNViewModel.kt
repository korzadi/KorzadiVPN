package com.korzadi.vpn.viewmodel

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.korzadi.vpn.data.repository.VPNRepository
import com.korzadi.vpn.vpn.WireGuardManager
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import javax.inject.Inject

@HiltViewModel
class VPNViewModel @Inject constructor(
    private val repository: VPNRepository,
    private val wireGuardManager: WireGuardManager
) : ViewModel() {

    private val _connectionState = MutableStateFlow(WireGuardManager.ConnectionState.DISCONNECTED)
    val connectionState = _connectionState.asStateFlow()

    fun connect(username: String, password: String) {
        viewModelScope.launch {
            _connectionState.value = WireGuardManager.ConnectionState.CONNECTING
            if (repository.login(username, password)) {
                val config = repository.getVpnConfig()
                if (config != null) {
                    wireGuardManager.startTunnel(config.config)
                    _connectionState.value = WireGuardManager.ConnectionState.CONNECTED
                } else {
                    _connectionState.value = WireGuardManager.ConnectionState.ERROR
                }
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
