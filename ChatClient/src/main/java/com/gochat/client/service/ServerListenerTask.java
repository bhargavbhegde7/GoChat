package com.gochat.client.service;

import java.io.BufferedReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.net.Socket;
import java.util.concurrent.LinkedBlockingDeque;

public class ServerListenerTask implements Runnable {

    private LinkedBlockingDeque<String> serverMessageQueue;

    public ServerListenerTask(LinkedBlockingDeque<String> serverMessageQueue) {
        this.serverMessageQueue = serverMessageQueue;
    }

    @Override
    public void run() {

        try{
            Socket socket = new Socket("127.0.0.1", 8080);
            InputStream input = socket.getInputStream();
            BufferedReader reader = new BufferedReader(new InputStreamReader(input));

            while(true){
                String line;
                try {
                    line = reader.readLine();
                    serverMessageQueue.put(line);
                } catch (Exception e) {
                    e.printStackTrace();
                }
            }
        }catch (Exception e){
            e.printStackTrace();
        }
    }
}
