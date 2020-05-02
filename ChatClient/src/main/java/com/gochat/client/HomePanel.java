package com.gochat.client;

import java.awt.GridLayout;

import javax.swing.JPanel;

public class HomePanel extends JPanel {
	public HomePanel() {
		init();
	}

	private void init(){
		int rows = 1, columns = 2;
		this.setLayout(new GridLayout(rows, columns));

		UserListPanel usersPanel = new UserListPanel();
		ChatPanel chatPanel = new ChatPanel();

		this.add(usersPanel);
		this.add(chatPanel);
	}
}
