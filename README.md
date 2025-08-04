# productivity-pal

Productivity-pal is an app that is designed to be your personal assistant.
Upon signing up, you will be prompted to enter your daily schedule, options you wish to utilize an assistant eg. Health and Fitness, Increase productivity.


CORE Tech Stack

Mobile: React Native with TypeScript
Backend: Go with Gin/Echo framework
Database: PostgreSQL (primary) + Redis (cache)
AI Integration: OpenAI API + vector database (Pinecone/Weaviate)
Authentication: JWT + OAuth2
Push Notifications: Firebase Cloud Messaging
File Storage: AWS S3 


Folder structure

productivity-pal/
├── mobile/                 # React Native app
│   ├── src/
│   ├── android/
│   ├── ios/
│   └── package.json
├── backend/               # Go API server
│   ├── cmd/
│   ├── handlers/
│   ├── models/
│   ├── server/
│   └── go.mod
├── ai-service/           # AI processing service
├── infrastructure/       # Docker, K8s configs
├── docs/                # API docs, architecture
└── scripts/             # Deployment scripts


Fisrt Steps:

Install React Native CLI, Go, Docker
Set up development databases (PostgreSQL, Redis)
Configure environment variables
Set up version control with proper .gitignore


TLS client certification for MTLS 
openssl genrsa -out client.key 2048                                              
romittajale@Mac certs % openssl req -new -key client.key -out client.csr -subj "/CN=client"              
romittajale@Mac certs % openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365


# Ollama is not able to pull the AI image from docker command: 