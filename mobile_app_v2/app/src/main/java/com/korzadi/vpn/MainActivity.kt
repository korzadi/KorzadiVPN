package com.korzadi.vpn

import android.app.Activity
import android.content.Intent
import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import com.korzadi.vpn.data.TokenManager
import com.korzadi.vpn.ui.login.LoginScreen
import com.korzadi.vpn.ui.vpn.VpnScreen


class MainActivity : ComponentActivity() {


    override fun onCreate(savedInstanceState: Bundle?) {

        super.onCreate(savedInstanceState)


        setContent {

            val hasToken =
                TokenManager(this).hasToken()


            if(hasToken) {

                VpnScreen()

            } else {

                LoginScreen(
                    onLoginSuccess = {

                        recreate()

                    }
                )

            }

        }

    }



    override fun onActivityResult(
        requestCode: Int,
        resultCode: Int,
        data: Intent?
    ) {

        super.onActivityResult(
            requestCode,
            resultCode,
            data
        )


        if(
            requestCode == 100 &&
            resultCode == Activity.RESULT_OK
        ) {

        }

    }

}
