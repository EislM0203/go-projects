# Go Project: Event Management REST API

This is a Go-based REST API for managing events, users, and registrations. It supports full CRUD operations on events, allows user registration and login, and enables users to register and unregister from events. The API uses JWT tokens for authentication and interacts with a database for data storage.

## API Routes

The following routes are available in the API:

- **GET /events**  
  Retrieves a list of all events.  
  Handlers: `traunseenet.com/rest-api/routes.getEvents`

- **GET /events/:id**  
  Retrieves details of a specific event by its ID.  
  Handlers: `traunseenet.com/rest-api/routes.getEvent`

- **POST /events**  
  Creates a new event. Requires authentication.  
  Handlers: `traunseenet.com/rest-api/routes.createEvent`

- **PUT /events/:id**  
  Updates an existing event by its ID. Requires authentication.  
  Handlers: `traunseenet.com/rest-api/routes.updateEvent`

- **DELETE /events/:id**  
  Deletes an event by its ID. Requires authentication.  
  Handlers: `traunseenet.com/rest-api/routes.deleteEvent`

- **POST /events/:id/register**  
  Registers a user for an event by its ID. Requires authentication.  
  Handlers: `traunseenet.com/rest-api/routes.registerForEvent`

- **DELETE /events/:id/unregister**  
  Unregisters a user from an event by its ID. Requires authentication.  
  Handlers: `traunseenet.com/rest-api/routes.unregisterFromEvent`

- **POST /signup**  
  Registers a new user.  
  Handlers: `traunseenet.com/rest-api/routes.signup`

- **POST /login**  
  Logs in a user and returns a JWT token.  
  Handlers: `traunseenet.com/rest-api/routes.login`

### Authentication

Event creation, updating, deletion, and user registration for events can only be performed by authenticated users. The authentication is handled via JWT tokens.

### Testing the API

If you don't want to build the executable, after downloading the dependencies, you can run the Go application directly with: `go run .` from the rest-api folder. Alternatively, run `go mod tidy` & `go build` which will result in a binary you can run. There are `.http` files in the `api-test` folder which you can use to test the API endpoints.
