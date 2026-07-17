package com.korzadi.vpn;

import com.korzadi.vpn.api.AuthClient;

public class LoginManager {

    private AuthClient authClient;

    public LoginManager(){
        authClient = new AuthClient();
    }


    public boolean login(
            String email,
            String password
    ){

        return authClient.login(
                email,
                password
        );
    }


    public String getToken(){

        return authClient.getToken();

    }
}
