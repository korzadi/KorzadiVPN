package com.korzadi.vpn.navigation

import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import androidx.compose.ui.platform.LocalContext
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.korzadi.vpn.data.TokenManager
import com.korzadi.vpn.ui.dashboard.DashboardScreen
import com.korzadi.vpn.ui.login.LoginScreen
import com.korzadi.vpn.ui.vpn.VpnScreen


sealed class Screen(val route: String) {

    object Login : Screen("login")

    object Dashboard : Screen("dashboard")

    object VPN : Screen("vpn")

}


@Composable
fun AppNavigation() {

    val context = LocalContext.current

    val tokenManager = remember {
        TokenManager(context)
    }

    val navController = rememberNavController()


    val startDestination =
        if (tokenManager.hasToken()) {
            Screen.Dashboard.route
        } else {
            Screen.Login.route
        }


    NavHost(
        navController = navController,
        startDestination = startDestination
    ) {


        composable(Screen.Login.route) {

            LoginScreen(
                onLoginSuccess = {

                    navController.navigate(
                        Screen.Dashboard.route
                    )

                }
            )

        }


        composable(Screen.Dashboard.route) {

            DashboardScreen(

                onOpenVPN = {

                    navController.navigate(
                        Screen.VPN.route
                    )

                },


                onLogout = {

                    navController.navigate(
                        Screen.Login.route
                    ) {

                        popUpTo(
                            Screen.Dashboard.route
                        ) {

                            inclusive = true

                        }

                    }

                }

            )

        }


        composable(Screen.VPN.route) {

            VpnScreen()

        }


    }

}
