# Overview
It's a simple lab designed to test your capability for discovering SSRF vulnerability.

# lab Description
- The external service allows you to send a get request to a specific target (URL/IP). [http://your_interface_ip:8080/check-response?url=<your_target>]
- There's another service that runs on localhost and has a secret flag the path of the flag is `/ctf`

# Install & Run
- Clone the repository.
- [Update the external service IP]([https://www.google.com](https://github.com/bassammaged/gorobindns/blob/master/external.go#L17)https://github.com/bassammaged/gorobindns/blob/master/external.go#L17) with your interface IP.
- Run the script: `go run .`

# Hint
- Round Robin DNS
