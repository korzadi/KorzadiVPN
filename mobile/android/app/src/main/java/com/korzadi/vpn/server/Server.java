package com.korzadi.vpn.server;


public class Server {


    public String name;

    public String country;

    public String ip;

    public int ping;


    public Server(
            String name,
            String country,
            String ip
    ){

        this.name = name;

        this.country = country;

        this.ip = ip;

        this.ping = 0;
    }
}
