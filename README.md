# Simple Cypher Tool

What does the tool do?
This tool allows you to **encrypt** and **decrypt** messages.  
It can be useful if you want to hide information or decode a message you’ve received.  
Example: if you get a message that is written using reversed characters, you can run it through our tool to reveal the original content.

## How to use the tool

1. **Choose your action:**
   <br>`1` Encrypt
   <br>`2` Decrypt

2. **Choose the encryption/decryption method:**
    <br> `1` **ROT13**  
      Shifts every letter by 13 positions in the alphabet.
    <br> `2` **Reverse (Atbash)**  
      Replaces each letter with its "mirror" in the alphabet (A ↔ Z, B ↔ Y, etc.).
    <br> `3` **Caesar cipher**  
      Asks you for a shift number and shifts every letter by that number.

3. **Enter your message** (non-alphabet characters will remain unchanged).

## Examples

**Example 1: ROT13 encryption**

Operation: Encrypt<br>
Cypher: ROT13<br>
Message: Hello World<br>
Result: Uryyb Jbeyq<br>


**Example 2: Caesar cipher with shift 3**

Operation: Encrypt<br>
Cypher: Caesar<br>
Shift: 3<br>
Message: Hello World<br>
Result: Khoor Zruog