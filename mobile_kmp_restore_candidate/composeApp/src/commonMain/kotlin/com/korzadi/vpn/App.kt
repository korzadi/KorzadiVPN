package com.korzadi.vpn

import androidx.compose.runtime.*
import androidx.compose.material.*
import androidx.compose.foundation.layout.*
import com.korzadi.vpn.api.ApiClient
import com.korzadi.vpn.repository.AuthRepositoryImpl
import com.korzadi.vpn.auth.AuthViewModel
import com.korzadi.vpn.storage.SecureStorage

@Composable
fun App() {
    val apiClient = ApiClient("http://localhost:8080")
    val authRepo = AuthRepositoryImpl(apiClient, object : SecureStorage { 
        override fun saveToken(token: String) {}
        override fun getToken(): String? = null
        override fun clear() {}
    })
    val authViewModel = AuthViewModel(authRepo)

    var currentScreen by remember { mutableStateOf("Splash") }
    
    MaterialTheme {
        when (currentScreen) {
            "Splash" -> SplashScreen(onNavigate = { currentScreen = it })
            "Login" -> LoginScreen(authViewModel, onNavigate = { currentScreen = it })
            "Register" -> RegisterScreen(authViewModel, onNavigate = { currentScreen = it })
            "Dashboard" -> DashboardScreen(onNavigate = { currentScreen = it })
        }
    }
}

@Composable
fun SplashScreen(onNavigate: (String) -> Unit) {
    LaunchedEffect(Unit) { onNavigate("Login") }
    Text("KorzadiVPN - Cargando...")
}

@Composable
fun LoginScreen(viewModel: AuthViewModel, onNavigate: (String) -> Unit) {
    Column {
        Text("Login")
        Button(onClick = { onNavigate("Dashboard") }) { Text("Entrar") }
        Button(onClick = { onNavigate("Register") }) { Text("Registrarse") }
    }
}

@Composable
fun RegisterScreen(viewModel: AuthViewModel, onNavigate: (String) -> Unit) {
    Column {
        Text("Registro")
        Button(onClick = { onNavigate("Login") }) { Text("Ir a Login") }
    }
}

@Composable
fun DashboardScreen(onNavigate: (String) -> Unit) {
    Text("Dashboard - VPN Conectada")
}
