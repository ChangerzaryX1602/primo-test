# Project: Go Logic with TypeScript gRPC Client

> **Note:** ผมไม่ค่อยถนัดภาษา TypeScript และเวลาที่จำกัดเพียง 120 นาทีไม่พอที่จะเรียนรู้ TypeScript ดังนั้นผมจึงใช้ภาษา **Go** ในการเขียน logic และ unit test ก่อน จากนั้นจึง expose ผ่าน **gRPC** ให้ TypeScript client สามารถเรียกใช้งานได้

## Overview

This project demonstrates a hybrid architecture where:
- **Go** is used to implement core business logic and thorough unit testing.
- **gRPC** is employed to expose the Go logic as a service.
- A **TypeScript** client connects to the gRPC service to run the logic.

This approach leverages the strengths of Go for fast and reliable backend logic while still providing a lightweight TypeScript client interface.

## Architecture

- **Go Backend**
  - Implements the main logic (e.g., merging sorted collections).
  - Contains comprehensive unit tests.
  - Exposes functionality via a gRPC server.
  
- **gRPC Service**
  - Uses Protocol Buffers (`.proto` files) to define the service interface.
  - Provides a robust and language-agnostic communication channel.
  
- **TypeScript Client**
  - Connects to the gRPC server.
  - Consumes the service methods defined in the proto file.
  - Handles communication and displays the results.

## Technologies Used

- **Go**: For business logic and unit tests.
- **gRPC & Protocol Buffers**: For service definition and inter-process communication.
- **TypeScript**: For the client-side implementation.
- **Node.js**: To run the TypeScript client.

## Setup Instructions

### 1. Go Server Setup

1. **Install Go:**  
   Ensure you have Go installed (version 1.23.1 or later recommended).

2. **Run server:**
   ```bash
   make server-run

### 1. TypeScript Server Setup

1. **Run client:**
   ```bash
   make client-run