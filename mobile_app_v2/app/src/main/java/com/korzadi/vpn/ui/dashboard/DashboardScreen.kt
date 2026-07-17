package com.korzadi.vpn.ui.dashboard

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import androidx.compose.ui.platform.LocalContext
import com.korzadi.vpn.api.RetrofitClient
import com.korzadi.vpn.data.TokenManager
import kotlinx.coroutines.launch


@Composable
fun DashboardScreen(
    onLogout: () -> Unit = {},
    onOpenVPN: () -> Unit = {}
) {

    val context = LocalContext.current

    val scope = rememberCoroutineScope()

    var result by remember {
        mutableStateOf("Cargando dashboard...")
    }


    LaunchedEffect(Unit) {

        scope.launch {

            try {

                val dashboard =
                    RetrofitClient.create(context)
                        .dashboard()


                result =
                    """
Usuario: ${dashboard.user?.email}

Plan: ${dashboard.user?.plan}

Estado: ${dashboard.user?.status}

Dispositivos:
${dashboard.statistics?.total_devices}

Conexiones activas:
${dashboard.statistics?.active_connections}

VPN:
${dashboard.vpn ?: "Sin conexión"}
                    """.trimIndent()


            } catch (e: Exception) {

                result =
                    "Error: ${e.message}"

            }

        }

    }


    Column(
        modifier = Modifier.padding(24.dp)
    ) {


        Text(
            "KorzadiVPN Dashboard"
        )


        Text(
            result
        )


        Button(
            onClick = {

                onOpenVPN()

            }
        ) {

            Text(
                "ABRIR VPN"
            )

        }


        Button(
            onClick = {

                TokenManager(context)
                    .logout()

                onLogout()

            }
        ) {

            Text(
                "CERRAR SESIÓN"
            )

        }


    }

}
