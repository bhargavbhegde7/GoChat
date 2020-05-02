package com.gochat.client.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.awt.GridLayout;

import javax.annotation.PostConstruct;
import javax.swing.JPanel;

@Component
public class HomePanel extends JPanel {

	@Autowired
	UserListPanel userListPanel;

	@Autowired
	ChatPanel chatPanel;

	@PostConstruct
	public void init(){
		int rows = 1, columns = 2;
		this.setLayout(new GridLayout(rows, columns));

		this.add(userListPanel);
		this.add(chatPanel);
	}
}
