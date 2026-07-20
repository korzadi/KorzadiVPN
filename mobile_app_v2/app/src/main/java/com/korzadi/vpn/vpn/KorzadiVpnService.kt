package com.korzadi.vpn.vpn

import android.content.Intent
import android.net.VpnService
import dagger.hilt.android.AndroidEntryPoint
import javax.inject.Inject

@AndroidEntryPoint
class KorzadiVpnService : VpnService() {

    @Inject
    lateinit var wireGuardManager: WireGuardManager

    companion object {
        const val ACTION_CONNECT = "com.korzadi.vpn.ACTION_CONNECT"
        const val ACTION_DISCONNECT = "com.korzadi.vpn.ACTION_DISCONNECT"
        const val EXTRA_CONFIG = "extra_config"
        const val EXTRA_NAME = "extra_name"
    }

    override fun onStartCommand(intent: Intent?, flags: Int, startId: Int): Int {
        when (intent?.action) {
            ACTION_CONNECT -> {
                val config = intent.getStringExtra(EXTRA_CONFIG)
                val name = intent.getStringExtra(EXTRA_NAME)
                if (config != null && name != null) {
                    wireGuardManager.startTunnel(config, name)
                }
            }
            ACTION_DISCONNECT -> {
                wireGuardManager.stopTunnel()
                stopSelf()
            }
        }
        return START_STICKY
    }

    override fun onDestroy() {
        wireGuardManager.stopTunnel()
        super.onDestroy()
    }
}
