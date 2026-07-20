package com.korzadi.vpn.ui

import androidx.compose.foundation.layout.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.korzadi.vpn.viewmodel.VPNViewModel
import com.korzadi.vpn.vpn.WireGuardManager

@Composable
fun VPNMainScreen(viewModel: VPNViewModel) {
    val state by viewModel.connectionState.collectAsState()

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
                viewModel.startVpn()
            }
        }) {
            Text(text = if (state == WireGuardManager.ConnectionState.CONNECTED) "Desconectar" else "Conectar")
        }
    }
}
