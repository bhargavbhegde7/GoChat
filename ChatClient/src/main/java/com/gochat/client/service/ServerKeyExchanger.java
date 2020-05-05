package com.gochat.client.service;

import java.security.InvalidKeyException;
import java.security.KeyFactory;
import java.security.NoSuchAlgorithmException;
import java.security.PublicKey;
import java.security.spec.X509EncodedKeySpec;
import java.util.Base64;
import java.util.UUID;

import javax.crypto.BadPaddingException;
import javax.crypto.Cipher;
import javax.crypto.IllegalBlockSizeException;
import javax.crypto.NoSuchPaddingException;

public class ServerKeyExchanger implements Runnable {

	private Client client;

	Base64.Decoder decoder = Base64.getDecoder();

	public ServerKeyExchanger(Client client) {
		this.client = client;
	}

	@Override
	public void run() {

		try {

			client.setServerKey(UUID.randomUUID().toString().getBytes());

			byte[] publicBytes = decoder.decode(client.getServerPubKey());
			X509EncodedKeySpec keySpec = new X509EncodedKeySpec(publicBytes);
			KeyFactory keyFactory = KeyFactory.getInstance("RSA");
			PublicKey pubKey = keyFactory.generatePublic(keySpec);

			byte[] encryptedKey = encrypt(pubKey, client.getServerKey());


		}catch (Exception e){
			e.printStackTrace();
		}
	}

	public byte[] encrypt(PublicKey key, byte[] plaintext) throws NoSuchAlgorithmException, NoSuchPaddingException, InvalidKeyException, IllegalBlockSizeException, BadPaddingException
	{
		Cipher cipher = Cipher.getInstance("RSA/ECB/OAEPWithSHA1AndMGF1Padding");
		cipher.init(Cipher.ENCRYPT_MODE, key);
		return cipher.doFinal(plaintext);
	}
}
