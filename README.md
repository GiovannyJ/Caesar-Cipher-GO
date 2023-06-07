# CAESAR CIPHER GO
Implementation of the Caesar Cipher made in golang

# What is Caesar Cipher
The Caesar Cipher is a simple form of encryption that shifts letters by a key value to become a new letter. For example starting at 'A' if the key for the shift is 4 then 'A' -> 'E'. The calculation is done in mod 26 for the 26 characters of the alphabet. If a shift is done to 'Z' then it wraps around to the beginning of the alphabet. This form of encryption is obviously not secure as it is easy to brute force the key. There can be added complexity if This shift is done with all ASCII characters, but even then it is still not secure

# How To Run
1. git clone github.com/GiovannyJ/Caesar-Cipher-GO
2. cd caesar-cipher-go
3. go run main.go enc|dec -t TEXT -s NUM
