package com.gochat.client.service;

import java.net.Socket;

public class Client {
	private Socket socker;
	byte[] Targetpubkey;
	byte[] Username;
	byte[] ServerPubKey;
	byte[] ServerKey;
	byte[] PubKey;
	byte[] PrivKey;

	public Client() {
	}

	public Socket getSocker() {
		return socker;
	}

	public void setSocker(Socket socker) {
		this.socker = socker;
	}

	public byte[] getTargetpubkey() {
		return Targetpubkey;
	}

	public void setTargetpubkey(byte[] targetpubkey) {
		Targetpubkey = targetpubkey;
	}

	public byte[] getUsername() {
		return Username;
	}

	public void setUsername(byte[] username) {
		Username = username;
	}

	public byte[] getServerPubKey() {
		return ServerPubKey;
	}

	public void setServerPubKey(byte[] serverPubKey) {
		ServerPubKey = serverPubKey;
	}

	public byte[] getServerKey() {
		return ServerKey;
	}

	public void setServerKey(byte[] serverKey) {
		ServerKey = serverKey;
	}

	public byte[] getPubKey() {
		return PubKey;
	}

	public void setPubKey(byte[] pubKey) {
		PubKey = pubKey;
	}

	public byte[] getPrivKey() {
		return PrivKey;
	}

	public void setPrivKey(byte[] privKey) {
		PrivKey = privKey;
	}
}
