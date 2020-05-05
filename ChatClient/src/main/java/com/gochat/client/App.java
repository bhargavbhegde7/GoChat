package com.gochat.client;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.LinkedBlockingDeque;

import javax.swing.SwingUtilities;

import com.gochat.client.service.Client;
import com.gochat.client.service.CustomThreadFactory;
import com.gochat.client.service.ServerListenerTask;
import com.gochat.client.service.ServerMessageConsumer;
import com.gochat.client.service.UIBuilder;

import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.context.ConfigurableApplicationContext;

@SpringBootApplication
public class App {
	public static void main(String[] args) {

		ConfigurableApplicationContext ctx = new SpringApplicationBuilder(App.class).headless(false).run(args);

		startListeningToServer();

		SwingUtilities.invokeLater(()->ctx.getBean(UIBuilder.class).build());

	}

	private static void startListeningToServer(){
		try {
			LinkedBlockingDeque<String> serverMessageBuffer = new LinkedBlockingDeque<>(100);

			ServerListenerTask serverListenerTask = new ServerListenerTask(serverMessageBuffer);
			ServerMessageConsumer serverMessageConsumer = new ServerMessageConsumer(serverMessageBuffer, new Client());

			ExecutorService serverListenerExecutor = Executors.newSingleThreadExecutor(new CustomThreadFactory("Server Listener Task"));
			ExecutorService serverMsgConsumerExecutor = Executors.newSingleThreadExecutor(new CustomThreadFactory("Server Message Consumer"));

			serverListenerExecutor.submit(serverListenerTask);
			serverMsgConsumerExecutor.submit(serverMessageConsumer);

			serverListenerExecutor.shutdown();
			serverMsgConsumerExecutor.shutdown();
		}
		catch (Exception e) {
			e.printStackTrace();
		}
	}
}
