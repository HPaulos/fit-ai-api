# Fit AI API

A Go-based fitness API service built with Gin framework, PostgreSQL, and Firebase Firestore.

## Tech Stack

- **Framework**: [Gin](https://github.com/gin-gonic/gin) - Fast HTTP web framework
- **Database**: PostgreSQL with [GORM](https://gorm.io/) ORM
- **Firebase**: Firestore for document storage
- **Environment**: [godotenv](https://github.com/joho/godotenv) for configuration
- **Containerization**: Docker & Docker Compose for easy development

## Quick Start

### Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- Git
- Firebase project with Firestore enabled

### Installation

1. **Clone the repository**
   ```bash
   git clone <your-repo-url>
   cd fit-ai-api
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables**
   ```bash
   cp env.example .env
   # Edit .env with your database credentials and Firebase config
   ```

4. **Set up Firebase**
   - Download your Firebase service account key from Firebase Console
   - Save it as `serviceAccountKey.json` in the project root
   - Or set `GOOGLE_APPLICATION_CREDENTIALS` environment variable

5. **Start the database with Docker**
   ```bash
   make db-up
   # or manually:
   # docker-compose up -d postgres
   ```

6. **Run the application**
   ```bash
   go run main.go
   ```

The API will be available at `http://localhost:8080`

## Firebase Setup

### 1. Get Firebase Service Account Key

1. Go to [Firebase Console](https://console.firebase.google.com/)
2. Select your project
3. Go to Project Settings > Service Accounts
4. Click "Generate new private key"
5. Save the JSON file as `serviceAccountKey.json` in your project root

### 2. Environment Configuration

Add to your `.env` file:
```
GOOGLE_APPLICATION_CREDENTIALS=serviceAccountKey.json
```

## API Endpoints

### Health Check
- `GET /health` - Check if the API is running

### User Management (PostgreSQL)
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get user by ID
- `POST /api/v1/users` - Create new user
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### Firestore Document Retrieval
- `GET /api/v1/firestore/:id` - Get document by ID (defaults to "users" collection)
- `GET /api/v1/firestore/:collection/:id` - Get document by ID from specific collection

### AI Workout Plan Generation
- `POST /api/v1/ai/workout-plan/:user_id` - Generate personalized workout plan for user
- `GET /api/v1/ai/workout-plan/:plan_id` - Get specific workout plan by ID
- `PUT /api/v1/ai/workout-plan/:plan_id` - Update workout plan
- `DELETE /api/v1/ai/workout-plan/:plan_id` - Delete workout plan
- `GET /api/v1/ai/workout-plans/:user_id` - Get all workout plans for a user

### Example Firestore Requests

```bash
# Get document with ID "i05zVUkMmkabNryrIdD4vwnBPkO2" from "users" collection
curl http://localhost:8080/api/v1/firestore/i05zVUkMmkabNryrIdD4vwnBPkO2

# Get document with ID "i05zVUkMmkabNryrIdD4vwnBPkO2" from "profiles" collection
curl http://localhost:8080/api/v1/firestore/profiles/i05zVUkMmkabNryrIdD4vwnBPkO2

# Get document with custom collection via query parameter
curl "http://localhost:8080/api/v1/firestore/i05zVUkMmkabNryrIdD4vwnBPkO2?collection=profiles"
```

### Example AI Workout Plan Requests

```bash
# Generate personalized workout plan for user
curl -X POST http://localhost:8080/api/v1/ai/workout-plan/i05zVUkMmkabNryrIdD4vwnBPkO2

# Get specific workout plan by ID
curl http://localhost:8080/api/v1/ai/workout-plan/1

# Get all workout plans for a user
curl http://localhost:8080/api/v1/ai/workout-plans/i05zVUkMmkabNryrIdD4vwnBPkO2

# Update workout plan
curl -X PUT http://localhost:8080/api/v1/ai/workout-plan/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Plan","description":"Updated description"}'

# Delete workout plan
curl -X DELETE http://localhost:8080/api/v1/ai/workout-plan/1
```

## Docker Setup

### Using Makefile (Recommended)

```bash
# Complete setup for new developers
make setup

# Start development environment (database + API)
make dev

# Database management
make db-up      # Start database
make db-down    # Stop database
make db-reset   # Reset database (delete all data)
make db-logs    # View database logs

# Application commands
make run        # Run the API
make build      # Build the API
make test       # Run tests
make deps       # Install dependencies

# View all available commands
make help
```

### Manual Docker Commands

```bash
# Start PostgreSQL database
docker-compose up -d postgres

# Start all services (including pgAdmin)
docker-compose up -d

# Stop all services
docker-compose down

# View logs
docker-compose logs -f postgres
```

## Database Access

- **PostgreSQL**: `localhost:5432`
  - Database: `fit_ai_db`
  - Username: `postgres`
  - Password: `postgres`

- **pgAdmin** (Web UI): `http://localhost:8081`
  - Email: `admin@fitai.com`
  - Password: `admin`

## Project Structure

```
fit-ai-api/
├── main.go              # Application entry point
├── go.mod               # Go module dependencies
├── env.example          # Environment variables template
├── docker-compose.yml   # Docker services configuration
├── init.sql             # Database initialization script
├── Makefile             # Development commands
├── .gitignore           # Git ignore rules
├── serviceAccountKey.json # Firebase service account key
├── models/              # Database models
├── handlers/            # API handlers
├── services/            # Business logic services
└── README.md            # This file
```

## Database Schema

The database includes tables for:
- **Users** - User profiles and authentication (PostgreSQL)
- **Firestore Collections** - Document storage (Firebase)

## Development

### Adding New Endpoints

1. Create models in `models/` directory
2. Create handlers in `handlers/` directory  
3. Add routes in `main.go` or create separate route files

### Database Migrations

GORM will automatically handle migrations when you define your models.

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DATABASE_URL` | PostgreSQL connection string | `postgres://postgres:postgres@localhost:5432/fit_ai_db?sslmode=disable` |
| `PORT` | Server port | `8080` |
| `GIN_MODE` | Gin mode (debug/release) | `debug` |
| `JWT_SECRET` | JWT signing secret | - |
| `GOOGLE_APPLICATION_CREDENTIALS` | Firebase service account key path | `serviceAccountKey.json` |
| `GOOGLE_CLOUD_PROJECT` | Firebase project ID | - |
| `AI_API_KEY` | AI service API key for workout plan generation | - |

## Troubleshooting

### Database Connection Issues
```bash
# Check if database is running
docker-compose ps

# View database logs
make db-logs

# Reset database if needed
make db-reset
```

### Firebase Connection Issues
```bash
# Check if service account key exists
ls -la serviceAccountKey.json

# Verify Firebase credentials
gcloud auth application-default print-access-token
```

### Port Conflicts
If port 5432 or 8081 is already in use:
```bash
# Stop existing PostgreSQL service
sudo service postgresql stop

# Or change ports in docker-compose.yml
```

## AI Workout Plan Features

The API now includes AI-powered workout plan generation with the following features:

### Personalized Workout Plans
- Analyzes user profile data from Firestore
- Considers fitness level, goals, available equipment
- Generates appropriate exercise recommendations
- Adjusts difficulty based on user experience

### Workout Plan Structure
- Multiple workout sessions per plan
- Detailed exercise specifications (sets, reps, weight)
- Support for both weighted and bodyweight exercises
- Progressive overload principles

### OpenAI Integration
- **Real AI-powered workout plan generation using GPT-4**
- Personalized based on user profile from Firestore
- Considers fitness level, goals, and available equipment
- Generates appropriate exercises, sets, reps, and weights
- Uses structured prompts for consistent workout plan generation
- High-quality fitness recommendations from OpenAI's advanced language model

## Next Steps

- [x] Add AI-powered workout plan generation
- [x] Integrate with OpenAI GPT-4 for real AI responses
- [ ] Add user authentication
- [ ] Create workout tracking and progress analytics
- [ ] Add exercise library and variations
- [ ] Implement workout plan scheduling
- [ ] Add nutrition recommendations
- [ ] Add workout plan persistence to database
- [ ] Implement AI feedback and plan optimization 