package com.korzadi.vpn.vpn

import android.app.Notification
import android.app.NotificationChannel
import android.app.NotificationManager
import android.content.Intent
import android.net.VpnService
import android.os.Build
import androidx.core.app.NotificationCompat
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
        private const val NOTIFICATION_CHANNEL_ID = "KorzadiVpnChannel"
        private const val NOTIFICATION_ID = 1
    }

    override fun onStartCommand(intent: Intent?, flags: Int, startId: Int): Int {
        when (intent?.action) {
            ACTION_CONNECT -> {
                val config = intent.getStringExtra(EXTRA_CONFIG)
                val name = intent.getStringExtra(EXTRA_NAME)
                if (config != null && name != null) {
                    startForeground(NOTIFICATION_ID, createNotification())
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

    private fun createNotification(): Notification {
        val manager = getSystemService(NotificationManager::class.java)
            ?: return NotificationCompat.Builder(this, NOTIFICATION_CHANNEL_ID)
                .setContentTitle("KorzadiVPN")
                .setContentText("Túnel activo")
                .setSmallIcon(android.R.drawable.ic_lock_lock)
                .build()

        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            val channel = NotificationChannel(
                NOTIFICATION_CHANNEL_ID,
                "VPN Service",
                NotificationManager.IMPORTANCE_LOW
            )
            manager.createNotificationChannel(channel)
        }

        return NotificationCompat.Builder(this, NOTIFICATION_CHANNEL_ID)
            .setContentTitle("KorzadiVPN")
            .setContentText("Túnel activo")
            .setSmallIcon(android.R.drawable.ic_lock_lock)
            .setPriority(NotificationCompat.PRIORITY_LOW)
            .build()
    }

    override fun onDestroy() {
        wireGuardManager.stopTunnel()
        super.onDestroy()
    }
}
