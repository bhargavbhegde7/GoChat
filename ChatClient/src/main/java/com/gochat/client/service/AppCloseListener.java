package com.gochat.client.service;

import java.awt.event.WindowAdapter;
import java.awt.event.WindowEvent;

import javax.swing.JFrame;

import static com.gochat.client.service.AwtUtils.enableComponents;

public class AppCloseListener extends WindowAdapter {

	private JFrame parentFrame;

	public AppCloseListener(JFrame parentFrame) {
		this.parentFrame = parentFrame;
	}

	@Override
	public void windowClosing(WindowEvent windowEvent) {
		super.windowClosing(windowEvent);

		//we do not want to launch multiple prompt dialogs even when one is already launched and this frame is in disabled state and still someone clicks on the cross button.
		if(parentFrame.isEnabled()){
			//disable frame
			parentFrame.setEnabled(false);

			//disable inner components recursively
			enableComponents(parentFrame.getContentPane(), false);

			BooleanPrompt booleanPrompt = new BooleanPrompt(parentFrame, "Confirm Exit", "Are you sure you want to exit?", "Yes", "No");
			booleanPrompt.getPositiveButton().addActionListener(actionEvent -> parentFrame.dispose());
			booleanPrompt.getNegativeButton().addActionListener(actionEvent -> {
				booleanPrompt.dispose();

				//disable frame
				parentFrame.setEnabled(true);

				//disable inner components recursively
				enableComponents(parentFrame.getContentPane(), true);
			});
			booleanPrompt.setVisible(true);
		}
	}
}
