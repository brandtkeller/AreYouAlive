# AreYouAlive
Application/Endpoint status dashboard

## Purpose
Hosted location for presenting current state of target applications or services to present to end-users.
End-users can check this dashboard during percieved outages.

(Future - Owners can configure webhooks for sending Alerts to Chat applications)

## Current Work-In-Progress
- Conversion to a proper project structure
- consume a json file and iterate over it
- Re-Introduce the REST server 

### Education
This project also serves as an introductory project for learning some Golang concepts.

### Structure
Structure of the project will be within compliance with [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

## Current Thoughts
Space to Brain Storm thoughts and ideas for execution

### Application Structure
- Main Thread
    - REST API
        - '/target'
    - Web Server
        - '/'
    - Web Socket Connector
        - '/ws'
- Execution Thread
    - Target Update

### Initial Functionality (Iteration #1)
The focus of the first iteration will be a service that consumes a configuration file of target information and delivers a user interface for the current state of the targets. State of the targets will be queried via a goroutine and will have a channel back to send updates / update notices.
- Create the `target` object struct
    - Create a test data-set
- Create `/target` CRUD API's
    - GET `/target`
    - GET `/target/{id}`
    - Confirm functionality via `CURL`
    - Write Initial tests
- Create execution goroutine
    - Consume targets
    - goroutine per target?
        - Loop
            - Execute target query (How-many of these can we execute sequentially?)
            - Sleep for interval time `time.Sleep( interval * time.Second)`
- React frontend web application served via `/` endpoint (GET)
- Websocket established between backend/frontend for updates from execution thread
### Iteration #2
This is till to-be-determined. I would like to implement a full CRUD API for targets and allow the service to have a database for which to store state. An administrative dashboard would then allow admins to add targets dynamically.
- Configurable target health query parameters would be a nice to have. 
- Webhooks for MM integrations or other integrations would be a nice to have.

### Other Considerations
- Test Driven Development
    - Do I know enough from the start to drive this?
- Standard Go Project structure
    - Should I just create a single-file webapp and refactor?