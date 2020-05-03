package com.gochat.client.service;

import com.google.gson.Gson;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Scope;
import org.springframework.stereotype.Component;

import java.util.concurrent.LinkedBlockingDeque;

@Component
@Scope(value = "prototype")
public class ServerMessageConsumer implements Runnable {

    private static final Logger LOGGER = LogManager.getLogger(ServerMessageConsumer.class);

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
            switch (response.getRestag()){
                case "connection_successful":
                    LOGGER.info("Connection successful "+response.toString());
                    break;
                case "NONE":
                    LOGGER.error("Request tag did not match any in the server : "+response.toString());
                    break;
                default:
                    LOGGER.error("Received an unrecognised tag from server : "+response.toString());
            }
        }
    }
}
