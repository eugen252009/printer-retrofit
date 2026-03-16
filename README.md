# Anycubic Kobra Plus Retrofit for Cura

This project implements a bridge between an Anycubic Kobra Plus 3D printer and Cura, enabling communication over TCP and Serial. It serves as a retrofit solution to seamlessly integrate the printer with Cura.

## Features

- TCP <-> Serial bridge for real-time communication
- Support for GCode commands via Protobuf (binarygcode)
- Easy initialization of the printer via standard USB serial port
- Allows Cura to control the printer without native driver modifications

## Requirements

- Go 1.25.5 or higher
- Anycubic Kobra Plus printer
- Cura (or compatible software that can send GCode over TCP)
- USB connection or TCP/IP network

## Installation

1. Clone the repository:
   `bash
   git clone https://github.com/eugen252009/printer-retrofit.git
   `
2. Change into the project directory:
   `bash
   cd printer-retrofit
   `
3. Install modules:
   `bash
   go mod tidy
   `
4. Build the program:
   `bash
   go build -o printer-retrofit main.go
   `
5. Start the bridge:
   `bash
   ./printer-retrofit
   `

## Usage

1. Run the program on the computer connected to the printer via USB.
2. Connect Cura or other GCode-capable software via TCP port 5000.
3. The bridge automatically forwards GCode commands to the printer and relays responses back.

## License

MIT
