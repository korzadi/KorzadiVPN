package com.korzadi.vpn.vpn

import com.korzadi.vpn.model.*
import com.korzadi.vpn.repository.VPNRepository
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.launch
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers

sealed class DashboardResult {
    object Idle : DashboardResult()
    object Loading : DashboardResult()
    data class Success(val client: VPNClient, val server: Server) : DashboardResult()
    data class Error(val message: String) : DashboardResult()
}

class DashboardViewModel(private val repository: VPNRepository) {
    private val _uiState = MutableStateFlow<DashboardResult>(DashboardResult.Idle)
    val uiState: StateFlow<DashboardResult> = _uiState
    private val scope = CoroutineScope(Dispatchers.Main)

    fun createVPNClient() {
        scope.launch {
            _uiState.value = DashboardResult.Loading
            try {
                val response = repository.createClient()
                _uiState.value = DashboardResult.Success(response.client, response.server)
            } catch (e: Exception) {
                _uiState.value = DashboardResult.Error(e.message ?: "Error al crear cliente VPN")
            }
        }
    }
}
