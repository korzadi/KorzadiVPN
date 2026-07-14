package com.korzadi.vpn.server;


import java.util.ArrayList;


public class ServerManager {


    public ArrayList<Server> getServers(){

        ArrayList<Server> list =
                new ArrayList<>();


        list.add(
                new Server(
                        "Korzadi USA",
                        "United States",
                        "server1"
                )
        );


        list.add(
                new Server(
                        "Korzadi Europe",
                        "Europe",
                        "server2"
                )
        );


        return list;
    }
}
