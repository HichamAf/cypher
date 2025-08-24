package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// main() changed: now get all input in one call
func main() {
	fmt.Println("\n=== Welcome to the Cypher Tool ===")
	// Changed: replaced 3 separate calls by one call to getInput()
	toEncrypt, encoding, message := getInput()

	var result string
	switch encoding {
	case "rot13":
		if toEncrypt {
			result = encrypt_rot13(message)
		} else {
			result = decrypt_rot13(message)
		}
	case "reverse":
		if toEncrypt {
			result = encrypt_reverse(message)
		} else {
			result = decrypt_reverse(message)
		}
	case "caesar":
		shift := askShift()
		if toEncrypt {
			result = encrypt_caesar(message, shift)
		} else {
			result = decrypt_caesar(message, shift)
		}
	}

	opStr := "Encrypted"
	if !toEncrypt {
		opStr = "Decrypted"
	}
	fmt.Printf("\n%s message using %s:\n%s\n", opStr, encoding, result)
}

func getInput() (toEncrypt bool, encoding string, message string) {
	r := bufio.NewReader(os.Stdin) // reuse reader for all input
	toEncrypt = askOperation(r)
	encoding = askCypher(r)
	message = askMessage(r)
	return
}

func askOperation(r *bufio.Reader) bool {
	for {
		fmt.Println("\nSelect operation (1/2):")
		fmt.Println("1. Encrypt.")
		fmt.Println("2. Decrypt.")
		op, _ := r.ReadString('\n')
		op = strings.TrimSpace(op)
		if op == "1" {
			return true
		} else if op == "2" {
			return false
		}
		fmt.Println("Invalid input. Please enter 1 or 2.")
	}
}
// Changed: accept *bufio.Reader instead of creating inside
func askCypher(r *bufio.Reader) string {
	for {
		fmt.Println("\nSelect cypher (1/2/3):")
		fmt.Println("1. ROT13.")
		fmt.Println("2. Reverse.")
		fmt.Println("3. Caesar.")
		enc, _ := r.ReadString('\n')
		enc = strings.TrimSpace(enc)
		if enc == "1" {
			return "rot13"
		} else if enc == "2" {
			return "reverse"
		} else if enc == "3" {
			return "caesar"
		}
		fmt.Println("Invalid input. Please enter 1, 2, or 3.")
	}
}
// 🔹 Changed: accept *bufio.Reader instead of creating inside
func askMessage(r *bufio.Reader) string {
	fmt.Println("\nEnter the message:")
	msg, _ := r.ReadString('\n')
	return strings.TrimSpace(msg)
}

func askShift() int {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nEnter shift value (e.g. 3):")
		shiftStr, _ := r.ReadString('\n')
		shiftStr = strings.TrimSpace(shiftStr)
		if shift, err := strconv.Atoi(shiftStr); err == nil {
			return shift % 26 // normalize shift to alphabet size
		}
		fmt.Println("Invalid number, please try again.")
	}
}

func encrypt_rot13(s string) string  { return rot13(s) }
func decrypt_rot13(s string) string  { return rot13(s) }
func encrypt_reverse(s string) string { return atbash(s) }
func decrypt_reverse(s string) string { return atbash(s) }
func encrypt_caesar(s string, shift int) string { return caesar(s, shift) }
func decrypt_caesar(s string, shift int) string { return caesar(s, -shift) }

func rot13(s string) string {
	var b strings.Builder
	for _, c := range s {
		switch {
		case c >= 'A' && c <= 'Z':
			b.WriteRune((c-'A'+13)%26 + 'A')
		case c >= 'a' && c <= 'z':
			b.WriteRune((c-'a'+13)%26 + 'a')
		default:
			b.WriteRune(c)
		}
	}
	return b.String()
}

func atbash(s string) string {
	var b strings.Builder
	for _, c := range s {
		switch {
		case c >= 'A' && c <= 'Z':
			b.WriteRune('Z' - (c - 'A'))
		case c >= 'a' && c <= 'z':
			b.WriteRune('z' - (c - 'a'))
		default:
			b.WriteRune(c)
		}
	}
	return b.String()
}

func caesar(s string, shift int) string {
	var b strings.Builder
	for _, c := range s {
		switch {
		case c >= 'A' && c <= 'Z':
			b.WriteRune(((c-'A')+rune(shift)+26)%26 + 'A')
		case c >= 'a' && c <= 'z':
			b.WriteRune(((c-'a')+rune(shift)+26)%26 + 'a')
		default:
			b.WriteRune(c)
		}
	}
	return b.String()
}