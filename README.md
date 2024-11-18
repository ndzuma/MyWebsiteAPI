# Portfolio Projects API 🚀

A fast API built with Go and Echo to manage my portfolio projects. Previously I was using Supabase, I switched to a custom API for better performance, cost-effectiveness, and consistent uptime. Plus, building things is fun!

## Why Go? 🤔

I chose Go for this API because:
- It's really fast, but simple at the same time
- Has great concurrency support
- Simple to deploy and maintain
- Amazing standard library
- Strong typing helps catch errors early
- Fast builds

## API Endpoints 🛣️

All API endpoints (except static files) are under `/api` and require Basic Authentication.

### Public Routes
- `GET /` - Nothing yet, but will display basic api docs (soon)
- `GET /admin` - Admin panel for managing projects
- `GET /api/ping` - Health check endpoint (returns a friendly message)

### Protected Routes (Basic Auth Required)
- `GET /api/projects` - Get all projects with full details
- `GET /api/projects/list` - Get a simplified list of projects
- `GET /api/projects/:name` - Get a specific project by name
- `POST /api/projects` - Create a new project
- `PUT /api/projects` - Update an existing project
- `DELETE /api/projects/:name` - Delete a project by name

## Deployment 🚂

This API is deployed on Railway using Nixpacks. Configuration files:
- `railway.toml` - Railway-specific configuration
- `.nixpacks.toml` - Build instructions for Nixpacks

## Environment Variables 🌳

Required environment variables:
```env
DATABASE_PUBLIC_URL=your_postgresql_url
DATABASE_URL=your_postgresql_url (for prod)
API_USERNAME=your_admin_username
API_PASSWORD=your_admin_password
PORT=8080 (default)
MODE=pro/dev
```

## Future Improvements 🎯

Some areas I'm looking to improve:

### Security 🔒
- [ ] Switch from Basic Auth to JWT
- [ ] Add rate limiting
- [ ] Implement request validation
- [ ] Add CORS configuration
- [ ] Implement proper password hashing

### Testing 🧪
- [ ] Add unit tests for handlers
- [ ] Add integration tests for API endpoints (for fun)
- [ ] Set up CI/CD pipeline with GitHub Actions (for fun)
- [ ] Add API documentation using Swagger/OpenAPI (for fun)
- [ ] Implement end-to-end testing
- [ ] Add performance benchmarks (for fun)

### Features 🌟
- [x] Implement caching layer for faster responses
- [ ] Add comprehensive logging system (extremely needed)
- [ ] Add image upload support (Cloudinary integration, because they have generous free tier)
- [ ] Create automated backup strategy (for fun)

### Code Quality 📝
- [ ] Add more comprehensive error handling
- [ ] Implement request validation middleware
- [ ] Add code documentation
- [ ] Implement proper logging levels

## Local Development 💻

1. Clone the repository
2. Create a `.env` file with required environment variables:
    ```env
    DATABASE_PUBLIC_URL=postgresql://username:password@localhost:5432/dbname
    DATABASE_URL=postgresql://username:password@localhost:5432/dbname (if using in prod)
    API_USERNAME=your_username
    API_PASSWORD=your_password
    PORT=8080
    MODE=dev
    ```
3. Ensure PostgreSQL is running locally, or on the cloud.
4. Run the API:
    ```bash
    go run api/main.go
    ```

*Note: This API is currently in active development. Features and endpoints may change.*\
*Note: Everything labelled "for fun" is primarily for learning purposes, the API works well enough to fit my needs.*

