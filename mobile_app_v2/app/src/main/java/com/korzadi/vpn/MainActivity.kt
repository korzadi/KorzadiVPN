package com.korzadi.vpn

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.viewModels
import androidx.compose.foundation.layout.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.korzadi.vpn.viewmodel.VPNViewModel
import com.korzadi.vpn.vpn.WireGuardManager
import dagger.hilt.android.AndroidEntryPoint

@AndroidEntryPoint
class MainActivity : ComponentActivity() {

    private val viewModel: VPNViewModel by viewModels()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            val state by viewModel.connectionState.collectAsState()
            
            MaterialTheme {
                Surface(modifier = Modifier.fillMaxSize()) {
                    Column(
                        modifier = Modifier.fillMaxSize().padding(16.dp),
                        horizontalAlignment = Alignment.CenterHorizontally,
                        verticalArrangement = Arrangement.Center
                    ) {
                        Text(text = "Estado: ${state.name}", style = MaterialTheme.typography.headlineMedium)
                        Spacer(modifier = Modifier.height(16.dp))
                        
                        Button(onClick = { 
                            if (state == WireGuardManager.ConnectionState.CONNECTED) {
                                viewModel.disconnect()
                            } else {
                                viewModel.connect("user", "pass") // Simplified for demo logic
                            }
                        }) {
                            Text(text = if (state == WireGuardManager.ConnectionState.CONNECTED) "Desconectar" else "Conectar")
                        }
                    }
                }
            }
        }
    }
}
