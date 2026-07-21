package com.korzadi.vpn.vpn

import android.content.Context
import com.wireguard.android.backend.Backend
import com.wireguard.android.backend.GoBackend
import com.wireguard.android.backend.Tunnel
import com.wireguard.config.Config
import dagger.hilt.android.qualifiers.ApplicationContext
import java.io.ByteArrayInputStream
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class WireGuardManager @Inject constructor(
    @ApplicationContext private val context: Context
) {

    enum class ConnectionState {
        DISCONNECTED,
        CONNECTING,
        CONNECTED,
        ERROR
    }

    private val backend: Backend = GoBackend(context)

    private var activeTunnel: Tunnel? = null

    fun startTunnel(config: String, name: String) {
        try {
            val tunnel = object : Tunnel {
                override fun getName(): String = name

                override fun onStateChange(newState: Tunnel.State) {}
            }

            val parsedConfig = Config.parse(
                ByteArrayInputStream(config.toByteArray())
            )

            backend.setState(
                tunnel,
                Tunnel.State.UP,
                parsedConfig
            )

            activeTunnel = tunnel

        } catch (e: Exception) {
            e.printStackTrace()
        }
    }

    fun stopTunnel() {
        try {
            activeTunnel?.let {
                backend.setState(
                    it,
                    Tunnel.State.DOWN,
                    null
                )
            }

            activeTunnel = null

        } catch (e: Exception) {
            e.printStackTrace()
        }
    }
}
