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

```mermaid
sequenceDiagram
    participant U as User
    participant A as App (Frontend)
    participant B as Backend API
    participant Auth as Authorization Service
    participant AI as AI Recommendation Service

    U->>A: Request recommendations
    A->>B: Send "Get Recommendations" request
    B->>Auth: Request access token (Authorization)
    Auth-->>B: Return access token
    B->>AI: Request top 5 recommendations<br/>(with access token)
    AI-->>B: Return 5 recommended options
    B-->>A: Send recommendations to app
    A-->>U: Display top 5 recommended option
 