package com.korzadi.vpn

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.viewModels
import androidx.compose.material3.MaterialTheme
import androidx.compose.runtime.*
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.korzadi.vpn.ui.LoginScreen
import com.korzadi.vpn.ui.RegisterScreen
import com.korzadi.vpn.ui.VPNMainScreen
import com.korzadi.vpn.viewmodel.VPNViewModel
import dagger.hilt.android.AndroidEntryPoint

@AndroidEntryPoint
class MainActivity : ComponentActivity() {

    private val viewModel: VPNViewModel by viewModels()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            val navController = rememberNavController()
            val isAuthenticated by viewModel.isAuthenticated.collectAsState()
            
            MaterialTheme {
                NavHost(navController = navController, startDestination = if (isAuthenticated) "vpn" else "login") {
                    composable("login") {
                        LoginScreen(viewModel, 
                            onNavigateToRegister = { navController.navigate("register") }, 
                            onLoginSuccess = { navController.navigate("vpn") { popUpTo("login") { inclusive = true } } }
                        )
                    }
                    composable("register") {
                        RegisterScreen(viewModel, 
                            onRegisterSuccess = { navController.navigate("login") { popUpTo("register") { inclusive = true } } }
                        )
                    }
                    composable("vpn") {
                        VPNMainScreen(viewModel)
                    }
                }
            }
        }
    }
}
