package com.gochat.client.service;

import java.util.concurrent.ThreadFactory;

public class CustomThreadFactory implements ThreadFactory
{
    private String name;

    public CustomThreadFactory(String nameOfThreadToCreate)
    {
        this.name = nameOfThreadToCreate;
    }

    @Override
    public Thread newThread(Runnable runnable)
    {
        Thread t = new Thread(runnable, name);
        return t;
    }
}
