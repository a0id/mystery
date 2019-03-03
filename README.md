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

## The attempt struct
The attempt struct is a data structure accepted by the server (when properly encrypted).

The attempt struct has the following fields:
```
Username* Username // The username and 4-digit pin
byte[] Payload     // The payload
string Origin      // The ip address of the host on which the attempt was created
string Timestamp   // The time when the attempt was created
byte[] Hash        // The SHA3 hash of the attempt
```

The username type contains a four-digit integer pin and a string.

### Generation and further explanation
When generating an attempt struct using the `--generate` flag, you will be prompted to enter a username, pin, and payload filename. You will also be prompted to enter a passphrase and an export filename. This passphrase is used to encrypt the struct, which is then exported to the filename supplied. Then, when running the client, the user will be prompted to enter a payload filename. When prompted to do so, enter the name of the exported attempt. If the server is able to decrypt the attempt using the passphrase on the server, then the client's attempt struct will be accepted.