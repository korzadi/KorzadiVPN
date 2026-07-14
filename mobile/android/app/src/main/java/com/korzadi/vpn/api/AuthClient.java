package com.korzadi.vpn.api;


public class AuthClient {


    private String token;


    public void saveToken(String value){

        token = value;
    }


    public String getToken(){

        return token;
    }
}
