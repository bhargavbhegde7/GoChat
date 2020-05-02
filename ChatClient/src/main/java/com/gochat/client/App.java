package com.gochat.client;

import com.gochat.client.service.AppBuilder;
import com.gochat.client.service.ServerListenerTask;
import com.gochat.client.service.ServerMessageConsumer;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.context.ConfigurableApplicationContext;

import javax.swing.*;
import java.awt.*;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.LinkedBlockingDeque;

@SpringBootApplication
@EnableAutoConfiguration
public class App {
	public static void main(String[] args) {

		ConfigurableApplicationContext ctx = new SpringApplicationBuilder(App.class).headless(false).run(args);

		try {
			LinkedBlockingDeque<String> serverMessageBuffer = new LinkedBlockingDeque<>(100);
			ExecutorService executorService = Executors.newFixedThreadPool(2);

			ServerListenerTask serverListenerTask = ctx.getBean(ServerListenerTask.class, serverMessageBuffer);
			ServerMessageConsumer serverMessageConsumer = ctx.getBean(ServerMessageConsumer.class, serverMessageBuffer);

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
