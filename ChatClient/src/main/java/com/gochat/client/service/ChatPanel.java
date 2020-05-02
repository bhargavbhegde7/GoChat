package com.gochat.client.service;

import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import javax.swing.*;
import java.awt.*;

@Component
public class ChatPanel extends JPanel {

	@PostConstruct
	public void init(){
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
