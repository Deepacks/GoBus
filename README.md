# GoBus

A Go library for making Modbus calls, supporting seamless communication via IPC. It provides a simple and efficient interface for integrating Modbus devices in industrial and automation applications, designed for scalability and ease of use.

<!-- GETTING STARTED -->

## Concepts

The **concepts** folder is designed to showcase and explain key components of the project, isolating different functionalities to provide a clear understanding of how each part contributes to the overall system.

### Contents

1. **Modbus Client**: A Modbus client implementation that demonstrates sending read and write requests to a Modbus server, specifically for interacting with holding registers.

2. **Modbus Server**: A Modbus server that handles read and write requests for holding registers.

3. **UDS IPC Client**: A UNIX Domain Socket client that reads requests and sends an ACK response.
