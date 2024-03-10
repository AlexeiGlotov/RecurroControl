package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"math/rand"
)

// Ключ для шифрования и дешифрования
var encryptionKey = []byte("supersecretkey12")

// Шифрование строки
func encryptString(value string) (string, error) {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	// Заполняем буфер до длины блока
	valueBytes := []byte(value)
	blockSize := block.BlockSize()
	padding := blockSize - len(valueBytes)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	valueBytes = append(valueBytes, padText...)

	// Создаем криптографический блочный шифр в режиме шифрования CBC
	ciphertext := make([]byte, blockSize+len(valueBytes))
	iv := ciphertext[:blockSize]
	/*if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}*/

	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[blockSize:], valueBytes)

	return hex.EncodeToString(ciphertext), nil
}

// Дешифрование строки
func decryptString(encryptedValue string) (string, error) {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	// Декодируем строку из шестнадцатеричного формата
	encryptedBytes, err := hex.DecodeString(encryptedValue)
	if err != nil {
		return "", err
	}

	// Создаем криптографический блочный шифр в режиме дешифрования CBC
	iv := encryptedBytes[:block.BlockSize()]
	encryptedBytes = encryptedBytes[block.BlockSize():]
	mode := cipher.NewCBCDecrypter(block, iv)

	// Расшифровываем и удаляем дополнение
	mode.CryptBlocks(encryptedBytes, encryptedBytes)
	length := len(encryptedBytes)
	unpadding := int(encryptedBytes[length-1])
	return string(encryptedBytes[:(length - unpadding)]), nil
}

// Генерация 32-символьного ключа
func generateKey(creator, holder, cheat string, number int) (string, error) {
	// Шифруем значения переменных
	encryptedCreator, err := encryptString(creator)
	if err != nil {
		return "", err
	}

	encryptedHolder, err := encryptString(holder)
	if err != nil {
		return "", err
	}

	encryptedCheat, err := encryptString(cheat)
	if err != nil {
		return "", err
	}

	// Преобразуем число в строку и шифруем
	encryptedNumber, err := encryptString(fmt.Sprintf("%d", number))
	if err != nil {
		return "", err
	}

	// Объединяем переменные и случайную строку для формирования ключа
	fullKey := fmt.Sprintf("%s%s%s%s", encryptedCreator, encryptedHolder, encryptedCheat, encryptedNumber)

	return fullKey, nil
}

// Дешифровка ключа и извлечение оригинальных значений
func decryptKey(key string) (creator, holder, cheat string, number int, err error) {
	if len(key)%32 != 0 {
		return "", "", "", 0, fmt.Errorf("неверная длина ключа")
	}

	// Извлекаем части ключа и дешифруем
	encryptedCreator := key[:32]
	encryptedHolder := key[32:64]
	encryptedCheat := key[64:96]
	encryptedNumber := key[96:]

	// Дешифруем значения переменных
	creator, err = decryptString(encryptedCreator)
	if err != nil {
		return "", "", "", 0, err
	}

	holder, err = decryptString(encryptedHolder)
	if err != nil {
		return "", "", "", 0, err
	}

	cheat, err = decryptString(encryptedCheat)
	if err != nil {
		return "", "", "", 0, err
	}

	// Дешифруем число
	decryptedNumber, err := decryptString(encryptedNumber)
	if err != nil {
		return "", "", "", 0, err
	}

	// Преобразуем строку в число
	fmt.Sscanf(decryptedNumber, "%d", &number)

	return creator, holder, cheat, number, nil
}

func main() {
	// Пример использования функции генерации ключа
	creator := "abcd"
	holder := "efgh"
	cheat := "ijkl"
	number := 42

	key, err := generateKey(creator, holder, cheat, number)
	if err != nil {
		fmt.Println("Ошибка при генерации ключа:", err)
		return
	}

	fmt.Println("Сгенерированный ключ:", key)

	// Пример использования функции дешифровки
	decryptedCreator, decryptedHolder, decryptedCheat, decryptedNumber, err := decryptKey(key)
	if err != nil {
		fmt.Println("Ошибка при дешифровке ключа:", err)
		return
	}

	fmt.Printf("Расшифрованные значения:\nCreator: %s\nHolder: %s\nCheat: %s\nNumber: %d\n",
		decryptedCreator, decryptedHolder, decryptedCheat, decryptedNumber)
}
