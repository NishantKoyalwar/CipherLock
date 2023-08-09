# CipherLock
CipherLock is CLI based password manager which uses Pebble DB for data storage, Cobra for UI/CLI, and Advanced Encryption Standard (AES) via Crypto to ensure robust security


--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------


## Install
1. Install go
2. Clone this repo
```
go get https://github.com/NishantKoyalwar/CipherLock
```

## Usage
```
Usage:
vaultX [command]

Available Commands:
add         Adds website credentials to the password manager, providing the website name and password.
delete      Deletes stored credentials from the password manager for a specified website or account.
generate    Generates a strong and secure password with specified length and complexity
get         Retrieves stored credentials for a specific website or account from the password manager.
help        Help about any command
showall     shows all the keys stored in database

Flags:
-h, --help   help for vaultX

Use "vaultX [command] --help" for more information about a command.
```

