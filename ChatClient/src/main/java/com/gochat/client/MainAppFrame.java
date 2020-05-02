package com.gochat.client;

import java.awt.HeadlessException;
import java.awt.event.WindowListener;

import javax.swing.JFrame;

public class MainAppFrame extends JFrame {

	public MainAppFrame(String title) throws HeadlessException {
		super(title);
		init();
	}

	public void init(){
		this.setSize(1000,600);

		//blocking default behavior since we want to inject our own behavior when close button is clicked.
		this.setDefaultCloseOperation(JFrame.DO_NOTHING_ON_CLOSE);

		WindowListener closeListener = new AppCloseListener(this);
		this.addWindowListener(closeListener);
	}
}