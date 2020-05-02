package com.gochat.client;

import java.awt.Button;
import java.awt.Dialog;
import java.awt.FlowLayout;
import java.awt.Frame;
import java.awt.Label;

/**
 * This class gives a dialog with two configurable buttons, meant for positive and negative response
 */
public class BooleanPrompt extends Dialog {
	private Button yesButton;
	private Button noButton;
	private String labelMessage;
	private String positiveBtnLabel;
	private String negativeBtnLabel;

	public BooleanPrompt(Frame frame, String title, String labelMessage, String positiveBtnLabel, String negativeBtnLabel) {
		super(frame, title);
		this.labelMessage = labelMessage;
		this.positiveBtnLabel = positiveBtnLabel;
		this.negativeBtnLabel = negativeBtnLabel;
		init();
	}

	public Button getPositiveButton() {
		return yesButton;
	}

	public Button getNegativeButton() {
		return noButton;
	}

	private void init(){
		yesButton = new Button(positiveBtnLabel);
		noButton = new Button(negativeBtnLabel);

		setLayout( new FlowLayout() );
		add( new Label(labelMessage));
		setSize(400,100);
		add(yesButton);
		add(noButton);
	}
}
