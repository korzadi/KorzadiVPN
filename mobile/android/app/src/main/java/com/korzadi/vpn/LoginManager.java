package com.korzadi.vpn;


public class LoginManager {


    public boolean login(
            String email,
            String password
    ){

        if(email == null || password == null){

            return false;
        }


        return true;
    }
}
