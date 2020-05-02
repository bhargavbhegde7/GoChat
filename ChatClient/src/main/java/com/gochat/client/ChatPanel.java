package com.gochat.client;

import java.awt.GridBagLayout;
import java.awt.List;

import javax.swing.BorderFactory;
import javax.swing.JPanel;

public class ChatPanel extends JPanel {
	public ChatPanel() {
		init();
	}

	private void init(){
		this.setLayout(new GridBagLayout());
		this.setBorder(BorderFactory.createTitledBorder("Chat"));
		List l1=new List(5);
		//l1.setBounds(100,100, 75,75);
		l1.add("Item 1");
		l1.add("Item 2");
		l1.add("Item 3");
		l1.add("Item 4");
		l1.add("Item 5");
		this.add(l1);
	}
}
