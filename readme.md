[![GoDoc](https://godoc.org/github.com/cheikhshift/go-securechain?status.svg)](https://godoc.org/github.com/cheikhshift/go-securechain) 

# go-securechain
Set of [SecureChain](https://sc.gophersauce.com) command line tools, this will facilitate working with the API, written in go.

## Install

	go get github.com/cheikhshift/go-securechain

## List of commands
Get documentation on each command by running it with flag `h`.

### sc-login
Generate a new access token for user. 

#### Install

	go get github.com/cheikhshift/go-securechain/cmd/sc-login

### sc-confirm-phone
Confirm a verification code sent to a user's phone.

#### Install

	go get github.com/cheikhshift/go-securechain/cmd/sc-confirm-phone

### sc-create-account
Create a new SecureChain user account. 

#### Install

	go get github.com/cheikhshift/go-securechain/cmd/sc-create-account

### sc-reset-password
Send new password to user's email.

#### Install

	go get github.com/cheikhshift/go-securechain/cmd/sc-reset-password

### sc-logout
Delete specified access token.

#### Install

	go get github.com/cheikhshift/go-securechain/cmd/sc-logout

### sc-encrypt
Encrypt a string or file.

#### Install

	go get github.com/cheikhshift/go-securechain/cmd/sc-encrypt

### sc-decrypt
Decrypt a string or file.

#### Install

	go get github.com/cheikhshift/go-securechain/cmd/sc-decrypt


