# crypto
加解密
==

<pre>
b := crypto.KeyBuilder{}
key, _ := b.SetEncrypt("123456").SetSignature("abc").Build()
c, _ := crypto.NewCipher(key)
cipherText := c.Encrypt("books", 0)
plainText, err, isExpired, expiredAt :=c.Decrypt(cipherText)
fmt.Println(plainText, err, isExpired, expiredAt)
</pre>