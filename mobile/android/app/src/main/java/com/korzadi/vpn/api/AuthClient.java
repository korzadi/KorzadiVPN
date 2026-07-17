package com.korzadi.vpn.api;

import java.io.*;
import java.net.*;
import org.json.*;

public class AuthClient {

    private String token;

    public boolean login(String email, String password) {

        try {
            URL url = new URL(API.BASE_URL + API.LOGIN);

            HttpURLConnection conn =
                    (HttpURLConnection) url.openConnection();

            conn.setRequestMethod("POST");
            conn.setRequestProperty(
                    "Content-Type",
                    "application/json"
            );
            conn.setDoOutput(true);

            String body =
                    "{\"email\":\"" + email +
                    "\",\"password\":\"" + password + "\"}";

            OutputStream os = conn.getOutputStream();
            os.write(body.getBytes());
            os.close();

            int code = conn.getResponseCode();

            if(code != 200){
                return false;
            }

            BufferedReader br =
                    new BufferedReader(
                    new InputStreamReader(
                    conn.getInputStream()));

            StringBuilder response = new StringBuilder();

            String line;
            while((line = br.readLine()) != null){
                response.append(line);
            }

            JSONObject json =
                    new JSONObject(response.toString());

            token = json.getString("token");

            return true;

        } catch(Exception e){

            e.printStackTrace();
            return false;
        }
    }


    public String getToken(){
        return token;
    }
}
