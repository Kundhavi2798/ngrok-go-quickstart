ngrok-go-quickstart

This project demonstrates how to expose a local Go web server to the internet using [ngrok](https://ngrok.com/) and the official [ngrok-go SDK](https://github.com/ngrok/ngrok-go).  
It’s ideal for testing webhooks, APIs, or any local application without deploying it to a public server.

## Features

- Embed ngrok directly into your Go application
- Automatically create secure tunnels to your local server
- Simplify testing of webhooks and APIs during development

## Prerequisites

- Go 1.18 or higher
- An [ngrok account](https://dashboard.ngrok.com/signup) to obtain an authtoken

## Getting Started

### 1. Clone the Repository


git clone https://github.com/Kundhavi2798/ngrok-go-quickstart.git
cd ngrok-go-quickstart

2. Set Up ngrok Authtoken
   
   Get your ngrok authtoken from here and set it in your environment:
            export NGROK_AUTHTOKEN=your-ngrok-authtoken

3. Run the Application:
   
   go run main.go
   
5. You will see output like:
   
   2025/06/05 19:09:34 Endpoint online: https://<random-subdomain>.ngrok.io
   
Visit and click see the result.



