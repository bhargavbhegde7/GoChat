package com.gochat.client;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.LinkedBlockingDeque;

import javax.swing.SwingUtilities;

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

		try {
			LinkedBlockingDeque<String> serverMessageBuffer = new LinkedBlockingDeque<>(100);

			ServerListenerTask serverListenerTask = ctx.getBean(ServerListenerTask.class, serverMessageBuffer);
			ServerMessageConsumer serverMessageConsumer = ctx.getBean(ServerMessageConsumer.class, serverMessageBuffer);

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

		SwingUtilities.invokeLater(()->ctx.getBean(UIBuilder.class).build());

	}
}
