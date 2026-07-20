package com.korzadi.vpn.vpn

import android.net.VpnService
import android.content.Intent
import android.os.ParcelFileDescriptor

class KorzadiVpnService : VpnService() {
    private var vpnInterface: ParcelFileDescriptor? = null

    override fun onStartCommand(intent: Intent?, flags: Int, startId: Int): Int {
        // Implement VPN setup using WireGuard config
        return START_STICKY
    }

    override fun onDestroy() {
        vpnInterface?.close()
        super.onDestroy()
    }
}
