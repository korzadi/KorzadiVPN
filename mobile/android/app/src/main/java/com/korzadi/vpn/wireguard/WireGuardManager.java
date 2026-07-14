package com.korzadi.vpn.wireguard;


public class WireGuardManager {


    private boolean connected = false;


    public boolean connect(
            WireGuardProfile profile
    ){

        connected = true;

        return connected;
    }


    public boolean disconnect(){

        connected = false;

        return true;
    }


    public boolean isConnected(){

        return connected;
    }
}
