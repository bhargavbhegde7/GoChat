package com.gochat.client;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.net.Socket;

import javax.swing.SwingUtilities;

public class App {
	public static void main(String[] args) {

		try {
			Socket socket = new Socket("127.0.0.1", 8080);
			InputStream input = socket.getInputStream();
			BufferedReader reader = new BufferedReader(new InputStreamReader(input));

			while(true){
				String line = reader.readLine();
				System.out.println(line);
			}

		}
		catch (Exception e) {
			e.printStackTrace();
		}

		SwingUtilities.invokeLater(() -> new App().buildMainApp());
	}

	private void buildMainApp(){

		MainAppFrame mainAppFrame = new MainAppFrame("Secure Messenger");

		mainAppFrame.add(new HomePanel());

		mainAppFrame.setVisible(true);

	}
}
