package com.gochat.client;

import javax.swing.*;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.LinkedBlockingDeque;

public class App {
	public static void main(String[] args) {

		try {
			LinkedBlockingDeque<String> serverMessageBuffer = new LinkedBlockingDeque<>(100);
			ExecutorService executorService = Executors.newFixedThreadPool(2);

			ServerListenerTask serverListenerTask = new ServerListenerTask(serverMessageBuffer);
			ServerMessageConsumer serverMessageConsumer = new ServerMessageConsumer(serverMessageBuffer);

			executorService.submit(serverListenerTask);
			executorService.submit(serverMessageConsumer);

			executorService.shutdown();
		}
		catch (Exception e) {
			e.printStackTrace();
		}

		SwingUtilities.invokeLater(AppBuilder::build);
	}
}
