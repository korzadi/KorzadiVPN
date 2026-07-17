package com.korzadi.vpn.wireguard

import android.net.VpnService
import android.content.Intent


class KorzadiVpnService : VpnService() {


    override fun onStartCommand(
        intent: Intent?,
        flags: Int,
        startId: Int
    ): Int {


        return START_STICKY

    }


    override fun onDestroy() {

        super.onDestroy()

    }

}
