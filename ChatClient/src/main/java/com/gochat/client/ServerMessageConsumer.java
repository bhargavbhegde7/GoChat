package com.gochat.client;

import java.util.concurrent.LinkedBlockingDeque;

public class ServerMessageConsumer implements Runnable {

    private LinkedBlockingDeque<String> serverMessageQueue;

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
            System.out.println(message);
        }
    }
}
