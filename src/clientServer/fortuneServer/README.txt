This project implements a simple fortune server that uses nonce authentication.

Step 1: Obtaining a nonce (https://en.wikipedia.org/wiki/Cryptographic_nonce)
Clients first connect to an authentication server to obtain a unique nonce.

Step 2: Authenticating using a shared secret
The client then computes an MD5 hash of the (nonce + the secret) and sends this to the authentication server as the authentication step.

The authentication server verifies the hash, and if correct, it returns information for contacting the fortune server.

Step 3: Retrieving the fortune
The client can now send a fortune request to the fortune server, and receives its fortune as a string.

Usage:
#PROJECT_DIRECTORY: the fortuneServer project directory, containing src
set GOPATH=<PROJECT_DIRECTORY>
cd <PROJECT_DIRECTORY>
go run src/aserver/auth-server.go [aserver UDP ip:port] [fserver RPC ip:port] [secret]
go run src/fserver/fortune-server.go [fserver RPC ip:port] [fserver UDP ip:port] [fortune-string]
go run src/testClients/client.go [local UDP ip:port] [aserver UDP ip:port] [secret]

Example:
go run src/fserver/fortune-server.go localhost:16806 localhost:15826 "MyFortune"
go run src/aserver/auth-server.go localhost:16210 localhost:16806 2016
go run src/testClients/client.go localhost:16066 localhost:16210 2016
