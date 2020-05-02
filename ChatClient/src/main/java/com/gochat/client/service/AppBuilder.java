package com.gochat.client.service;

public class AppBuilder {
    public static void build(){
        MainAppFrame mainAppFrame = new MainAppFrame("Secure Messenger");

        mainAppFrame.add(new HomePanel());

        mainAppFrame.setVisible(true);
    }
}
