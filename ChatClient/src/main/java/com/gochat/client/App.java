package com.gochat.client;

import javax.swing.SwingUtilities;

public class App {
	public static void main(String[] args) {
		SwingUtilities.invokeLater(() -> new App().buildMainApp());
	}

	private void buildMainApp(){

		MainAppFrame mainAppFrame = new MainAppFrame("Secure Messenger");

		mainAppFrame.add(new HomePanel());

		mainAppFrame.setVisible(true);

	}
}
