package com.korzadi.vpn.wireguard;


public class WireGuardProfile {


    public String server;

    public String address;

    public String dns;

    public String config;


    public WireGuardProfile(
            String server,
            String address,
            String dns,
            String config
    ){

        this.server = server;

        this.address = address;

        this.dns = dns;

        this.config = config;
    }
}
