# File-System Symmetric Encryptor (TEA)

A complete implementation of TEA (Tiny Encryption Algorithm) encryption and decryption for files.

## Workflow
<img width="1485" alt="Screenshot 2025-06-25 at 18 51 43" src="https://github.com/user-attachments/assets/78cc9342-4263-4073-8b4c-e36f33f36a7e" />

## Features

- **TEA Encryption/Decryption**: Full implementation of the Tiny Encryption Algorithm
- **File-based Operations**: Encrypt and decrypt entire files
- **Custom Key Support**: Use your own key files (converted to 128-bit hash)
- **Block Processing**: Handles files of any size by processing in 8-byte blocks
- **Automatic Padding**: Handles files that aren't multiples of 8 bytes
- **SHA-256 Key Hashing**: Converts any key file to a 128-bit hash for consistent key size

## How it Works

The implementation follows the standard TEA algorithm:

1. **Key Processing**: Converts input key to SHA-256 hash and uses first 16 bytes
2. **Block Processing**: Processes data in 8-byte blocks
3. **Encryption**: 32 rounds of TEA encryption with delta constant
4. **Padding**: Automatically adds null bytes to make data length multiple of 8
5. **Decryption**: 32 rounds of TEA decryption with automatic padding removal

### Core Components

- **TEA Algorithm**: Complete implementation of Tiny Encryption Algorithm
- **Key Hashing**: SHA-256 hashing for consistent 128-bit keys
- **Block Processing**: 8-byte block encryption/decryption
- **Padding Management**: Automatic padding during encryption and removal during decryption

## Usage

### Basic Commands

```bash
# Encrypt a file
go run main.go -e file=<file to encrypt> key=<keyfile>

# Decrypt a file
go run main.go -d file=<file to decrypt> key=<keyfile>
```

### Command Line Arguments

- `-e`: Encrypt the specified file
- `-d`: Decrypt the specified file
- `file=<path>`: Path to the input file
- `key=<path>`: Path to the key file

### Key File Format

The key file can contain any content. The system will:

1. Read the entire key file
2. Generate a SHA-256 hash of the content
3. Use the first 16 bytes of the hash as the TEA key

Example of creating a key file:

```bash
# Create a key file with any content
echo "MySecretKey12345" > mykey.txt

# Or use a random key
openssl rand -base64 32 > mykey.txt
```

### Output Files

- **Encryption**: Overwrites the input file with encrypted content
- **Decryption**: Overwrites the input file with decrypted content

## Implementation Details

### TEA Algorithm

The implementation includes all standard TEA operations:

1. **Key Setup**: 4 uint32 values derived from 16-byte key
2. **Block Processing**: 8-byte blocks with 32 rounds
3. **Delta Constant**: Uses 0x9E3779B9 as the delta value
4. **Padding**: Null byte padding for non-8-byte aligned data

### Key Features

- **Complete TEA Implementation**: Full encryption and decryption
- **SHA-256 Key Hashing**: Consistent 128-bit key generation
- **Automatic Padding**: Handles files of any size
- **Error Handling**: Input validation and error reporting
- **Overflow Protection**: Proper handling of large calculations

### Security Notes

- Uses TEA algorithm with 32 rounds
- SHA-256 key hashing for consistent key size
- Processes data in 8-byte blocks
- Automatic padding management

## Building and Running

```bash
# Build the application
go build -o file-encryptor

# Run with arguments
./file-encryptor -e file=document.txt key=secret.key
```

## Example Workflow

```bash
# 1. Create a test file
echo "Hello, this is a secret message!" > secret.txt

# 2. Create a key file
echo "MySecretKey12345" > mykey.txt

# 3. Encrypt the file
go run main.go -e file=secret.txt key=mykey.txt

# 4. Verify encryption (file should be unreadable)
cat secret.txt

# 5. Decrypt the file
go run main.go -d file=secret.txt key=mykey.txt

# 6. Verify decryption
cat secret.txt
```

## Technical Implementation

The codebase is organized into several modules:

- `tea/encrypt.go`: TEA encryption implementation
- `tea/decrypt.go`: TEA decryption implementation
- `cmd/args.go`: Command-line argument parsing
- `cmd/io.go`: File I/O operations
- `main.go`: Application entry point and workflow

All cryptographic operations are implemented from scratch, ensuring no external cryptographic dependencies beyond Go's standard library.

## File Structure

```
file-encryptor/
├── main.go          # Main application entry point
├── tea/
│   ├── encrypt.go   # TEA encryption implementation
│   └── decrypt.go   # TEA decryption implementation
├── cmd/
│   ├── args.go      # Command-line argument parsing
│   └── reader.go    # File I/O operations
├── go.mod           # Go module definition
└── readme.md        # This file
```
