package cmd

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "Manage vault",
	Long:  "This cmd will manage secret information of the vault",
}

func init() {
	vaultCmd.RunE = func(cmd *cobra.Command, args []string) error {
		text := []byte("Hello world!")
		secretKey := []byte("passphrasewhichneedstobe32bytes!")
		encryptedText, err := encrypeText(text, secretKey)
		if err != nil {
			return err
		}
		fmt.Printf("ecrypted text: %s\n", string(encryptedText))
		plainText, err := decrypeText(encryptedText, secretKey)
		if err != nil {
			return err
		}
		fmt.Printf("plain text: %s\n", plainText)
		return nil
	}
	rootCmd.AddCommand(vaultCmd)
}

func encrypeText(text []byte, secretKey []byte) ([]byte, error) {
	gcm, err := createGCM(secretKey)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter additional text: ")
	if scanner.Scan() {
		return gcm.Seal(nonce, nonce, text, []byte(scanner.Text())), nil
	}
	return nil, nil
}

func decrypeText(ciphertext []byte, secretKey []byte) ([]byte, error) {
	gcm, err := createGCM(secretKey)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, err
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter additional text: ")
	if scanner.Scan() {
		return gcm.Open(nil, nonce, ciphertext, []byte(scanner.Text()))
	}
	return nil, nil
}

func createGCM(secretKey []byte) (cipher.AEAD, error) {
	c, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(c)
}
