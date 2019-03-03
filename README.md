# Mystery Project
This is a pointless project that taught me how to use encryption, concurrency, and networking in Golang. Here's how it works:

The client can send files, "payloads", to a server which processes these payloads to see if they can be decrypted with the server's passphrase. The client can also create "attempts", encrypted payloads with a username, pin, and timestamp.

## Running
To start a server, first create a file in the same directory as the binary called `passphrase.sec`. This will be the passphrase that the client's payload must be encrypted with in order for the server to accept the payload. Then run
<br>
`./main --server <port>`

To start the client, run
<br>
`./main --client <ip>:<port>`

To generate an attempt, run
<br>
`./main --generate`