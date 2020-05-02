package com.gochat.client.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.stereotype.Component;

import javax.swing.*;
import java.awt.event.WindowAdapter;
import java.awt.event.WindowEvent;

import static com.gochat.client.service.AwtUtils.enableComponents;

@Component
public class AppCloseListener extends WindowAdapter {

	@Autowired
	ConfigurableApplicationContext configurableApplicationContext;

	private JFrame parentFrame;

	public void setParentFrame(JFrame parentFrame) {
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
			booleanPrompt.getPositiveButton().addActionListener(actionEvent -> {
				parentFrame.dispose();

				// do something
				int exitCode = SpringApplication.exit(configurableApplicationContext, () -> 0);

				System.exit(exitCode);

			});
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
