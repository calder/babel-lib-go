# Babel

Babel is a Go utility library for the [Vita](https://github.com/Calder/Vita) protocol.

## Coding

* **AddType(typeSig,encodeFunc,DecodeFunc)->bits**
* **Encode(anything)->bits**
* **Decode(bits)->anything**

## Crypto

* **PubKey.Encrypt(bits)->bits**
* **PriKey.Decrypt(bits)->bits**
* **PriKey.Pub()->PubKey**

## Network

* **Send(toId,message)**
* **Receive(toId,callback)**
