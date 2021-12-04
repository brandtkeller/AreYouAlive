# AreYouAlive
Application/Endpoint status dashboard

## Purpose
Hosted location for presenting current state of target applications or services to present to end-users.
End-users can check this dashboard during percieved outages.

(Future - Owners can configure webhooks for sending Alerts to Chat applications)

### Education
This project also serves as an introductory project for learning some Golang concepts. 

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

### Initial Functionality
- Create the `target` object struct
    - Create a test data-set
- Create `/target` CRUD API's
    - POST, GET, PUT, DELETE
    - Confirm functionality via `CURL`
    - Write Initial tests
- Create execution goroutine
    - Consume targets
    - goroutine per target?
        - Loop
            - Execute target query (How-many of these can we execute sequentially?)
            - Sleep for interval time `time.Sleep( interval * time.Second)`

### Other Considerations
- Test Driven Development
    - Do I know enough from the start to drive this?
- Standard Go Project structure
    - Should I just create a single-file webapp and refactor?