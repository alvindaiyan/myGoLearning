package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
)

func main() {
	key := []byte("XY0nG86PSRJqMGz957Yza1D34393MPII") // 32 bytes
	plaintext := "Password1"
	fmt.Printf("%s\n", plaintext)
	ciphertext, err := encrypt(plaintext, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("string: %s\n", ciphertext)
	result, err := decrypt(ciphertext, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", result)
}

func encrypt(str string, key []byte) (string, error) {
	text := []byte(str)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return base64.StdEncoding.EncodeToString(ciphertext[:]), nil
}

func decrypt(str string, key []byte) (string, error) {
	text, err := base64.StdEncoding.DecodeString(str)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	fmt.Printf("aes.BolockSize = %d\n", aes.BlockSize)
	if len(text) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return "", err
	}
	return string(data[:]), nil
}

func CToGoString(c []byte) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}

// import (
// 	"fmt"
// 	"log"
// 	"time"
// )

// func main() {
// 	t := time.Now().Local()
// 	const layout = "2006-01-02"
// 	fmt.Println(t.Format(layout))
// 	log.Println("time to log")
// }
