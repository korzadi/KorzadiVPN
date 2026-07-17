package com.korzadi.vpn.wireguard

import android.content.Context
import android.content.Intent
import android.net.VpnService
import com.wireguard.android.backend.GoBackend
import com.wireguard.android.backend.Tunnel
import com.wireguard.android.backend.Tunnel.State
import com.wireguard.config.Config
import java.io.ByteArrayInputStream

class WireGuardManager(
    private val context: Context
) {

    private val backend = GoBackend(context)

    private var currentTunnel: Tunnel? = null

    fun prepare(): Intent? {
        return VpnService.prepare(context)
    }

    fun connect(configText: String): Boolean {

        return try {

            val config = Config.parse(
                ByteArrayInputStream(configText.toByteArray())
            )

            val tunnel = object : Tunnel {

                override fun getName(): String {
                    return "KorzadiVPN"
                }

                override fun onStateChange(newState: State) {
                }
            }

            backend.setState(
                tunnel,
                State.UP,
                config
            )

            currentTunnel = tunnel

            true

        } catch (e: Exception) {
            e.printStackTrace()
            false
        }
    }


    fun disconnect(): Boolean {

        return try {

            currentTunnel?.let {
                backend.setState(
                    it,
                    State.DOWN,
                    null
                )
            }

            currentTunnel = null

            true

        } catch (e: Exception) {
            false
        }
    }


    fun isConnected(): Boolean {

        return try {

            currentTunnel != null &&
            backend.getState(currentTunnel!!) == State.UP

        } catch (e: Exception) {
            false
        }
    }
}
