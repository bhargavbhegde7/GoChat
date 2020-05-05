package com.gochat.client.service;

import java.util.Base64;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.LinkedBlockingDeque;

import com.google.gson.Gson;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

public class ServerMessageConsumer implements Runnable {

    private static final Logger LOGGER = LogManager.getLogger(ServerMessageConsumer.class);

    private LinkedBlockingDeque<String> serverMessageQueue;

	private Client client;

    public ServerMessageConsumer(LinkedBlockingDeque<String> serverMessageQueue, Client client) {
        this.serverMessageQueue = serverMessageQueue;
        this.client = client;
    }

	public void setClient(Client client) {
		this.client = client;
	}

	@Override
    public void run() {
		Gson gson = new Gson();
		Base64.Decoder decoder = Base64.getDecoder();
        while (true){
            String message;
            try {
                message = serverMessageQueue.take();

				Response response = gson.fromJson(message, Response.class);
				switch (response.getRestag()){
					case "connection_successful":
						LOGGER.info("Connected to the server "+response.toString());
						client.setServerPubKey(decoder.decode(response.getMessage()));
						startServerKeyExchange(client);
						break;
					case "NONE":
						LOGGER.error("Request tag did not match any in the server : "+response.toString());
						break;
					default:
						LOGGER.error("Received an unrecognised tag from server : "+response.toString());
				}
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }

    private void startServerKeyExchange(Client client){
    	ServerKeyExchanger serverKeyExchanger = new ServerKeyExchanger(client);
		Executors.newSingleThreadExecutor().submit(serverKeyExchanger);
	}
}
