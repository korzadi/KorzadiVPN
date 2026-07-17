package com.korzadi.vpn.ui.vpn

import android.app.Activity
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.unit.dp
import com.korzadi.vpn.api.RetrofitClient
import com.korzadi.vpn.data.TokenManager
import com.korzadi.vpn.wireguard.WireGuardManager
import kotlinx.coroutines.launch


private const val VPN_REQUEST_CODE = 100


@Composable
fun VpnScreen() {

    val context = LocalContext.current
    val activity = context as Activity
    val scope = rememberCoroutineScope()

    val manager = remember {
        WireGuardManager(context)
    }

    var config by remember {
        mutableStateOf("")
    }

    var result by remember {
        mutableStateOf("VPN desconectada")
    }


    Column(
        modifier = Modifier.padding(24.dp)
    ) {

        Text("KorzadiVPN")

        Text(result)


        Button(
            onClick = {

                scope.launch {

                    try {

                        val token =
                            TokenManager(context).getToken()


                        if(token.isNullOrEmpty()) {

                            result =
                                "Error: no hay sesión iniciada"

                            return@launch
                        }


                        val response =
                            RetrofitClient
                                .create(context)
                                .createVPN()


                        config =
                            response.client?.config ?: ""


                        result =
                            if(config.isNotEmpty())
                                "Perfil VPN listo"
                            else
                                "Servidor creado, falta configuración"


                    } catch(e: Exception) {

                        result =
                            "Error: ${e.message}"

                    }

                }

            }
        ) {

            Text("CREAR PERFIL")

        }



        Button(
            onClick = {

                if(config.isEmpty()) {

                    result =
                        "Primero crea el perfil VPN"

                    return@Button
                }


                val permission =
                    manager.prepare()


                if(permission != null) {

                    activity.startActivityForResult(
                        permission,
                        VPN_REQUEST_CODE
                    )

                } else {

                    val ok =
                        manager.connect(config)


                    result =
                        if(ok)
                            "VPN conectada"
                        else
                            "Error VPN"
                }

            }
        ) {

            Text("CONECTAR VPN")

        }



        Button(
            onClick = {

                manager.disconnect()

                result =
                    "VPN desconectada"

            }
        ) {

            Text("DESCONECTAR VPN")

        }

    }

}
