package com.korzadi.vpn.auth

import com.korzadi.vpn.model.*
import com.korzadi.vpn.repository.AuthRepository
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.launch
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers

class AuthViewModel(private val repository: AuthRepository) {
    private val _uiState = MutableStateFlow<AuthResult>(AuthResult.Idle)
    val uiState: StateFlow<AuthResult> = _uiState
    private val scope = CoroutineScope(Dispatchers.Main)

    fun login(email: String, pass: String) {
        scope.launch {
            _uiState.value = AuthResult.Loading
            try {
                val response = repository.login(email, pass)
                _uiState.value = AuthResult.Success(response.token)
            } catch (e: Exception) {
                _uiState.value = AuthResult.Error(e.message ?: "Error desconocido")
            }
        }
    }
}
