package com.gochat.client.service;

import com.google.gson.Gson;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Scope;
import org.springframework.stereotype.Component;

import java.util.concurrent.LinkedBlockingDeque;

@Component
@Scope(value = "prototype")
public class ServerMessageConsumer implements Runnable {

    private LinkedBlockingDeque<String> serverMessageQueue;

    @Autowired
    Gson gson;

    public ServerMessageConsumer(LinkedBlockingDeque<String> serverMessageQueue) {
        this.serverMessageQueue = serverMessageQueue;
    }

    @Override
    public void run() {
        while (true){
            String message = null;
            try {
                message = serverMessageQueue.take();
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            Response response = gson.fromJson(message, Response.class);
            System.out.println(response.getRestag());
        }
    }
}
