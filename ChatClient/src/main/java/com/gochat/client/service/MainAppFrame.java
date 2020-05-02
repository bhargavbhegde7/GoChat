package com.gochat.client.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import javax.swing.*;

@Component
public class MainAppFrame extends JFrame {

	@Autowired
	AppCloseListener appCloseListener;

	@PostConstruct
	public void init(){
		this.setSize(1000,600);

		//blocking default behavior since we want to inject our own behavior when close button is clicked.
		this.setDefaultCloseOperation(JFrame.DO_NOTHING_ON_CLOSE);

		appCloseListener.setParentFrame(this);
		this.addWindowListener(appCloseListener);
	}
}