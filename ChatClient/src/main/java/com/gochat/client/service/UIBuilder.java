package com.gochat.client.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class UIBuilder {

    @Autowired
    HomePanel homePanel;

    @Autowired
    MainAppFrame mainAppFrame;

    public void build(){
        mainAppFrame.add(homePanel);
        mainAppFrame.setVisible(true);
    }
}
