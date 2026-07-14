package com.korzadi.vpn.api;


public class Request {


    public String endpoint;


    public String method;


    public Request(
            String endpoint,
            String method
    ){

        this.endpoint = endpoint;

        this.method = method;
    }
}
