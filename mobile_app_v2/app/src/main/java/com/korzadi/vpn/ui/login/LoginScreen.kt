package com.korzadi.vpn.ui.login

import androidx.compose.foundation.layout.Column
import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.material3.TextField
import androidx.compose.runtime.*
import com.korzadi.vpn.api.RetrofitClient
import com.korzadi.vpn.data.TokenManager
import com.korzadi.vpn.model.LoginRequest
import kotlinx.coroutines.launch

@Composable
fun LoginScreen(
    onLoginSuccess: () -> Unit
) {

    val context = androidx.compose.ui.platform.LocalContext.current
    val scope = rememberCoroutineScope()

    var email by remember {
        mutableStateOf("")
    }

    var password by remember {
        mutableStateOf("")
    }

    var result by remember {
        mutableStateOf("Esperando login")
    }


    Column {

        TextField(
            value = email,
            onValueChange = {
                email = it
            },
            label = {
                Text("Email")
            }
        )


        TextField(
            value = password,
            onValueChange = {
                password = it
            },
            label = {
                Text("Password")
            }
        )


        Button(
            onClick = {

                scope.launch {

                    try {

                        val api = RetrofitClient.create(context)

                        val response = api.login(
                            LoginRequest(
                                email,
                                password
                            )
                        )


                        response.token?.let {

                            TokenManager(context)
                                .saveToken(it)

                        }


                        result = "Login correcto"


                        onLoginSuccess()


                    } catch (e: Exception) {

                        result =
                            "Error: ${e.message}"

                    }

                }

            }
        ) {

            Text("LOGIN")

        }


        Text(result)

    }

}
