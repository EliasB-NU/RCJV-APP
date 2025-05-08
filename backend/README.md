# API Specification

| Endpoint              | Payload                     | Auth  | Return                   | Description                                 |
|-----------------------|-----------------------------|-------|--------------------------|---------------------------------------------|
| /api/config           |                             | true  | current config           | Getting the current config for the frontend |
| /api/config/update    | Updated Config              | true  |                          | Updating the current config                 |
| /api/enabled          |                             | false | if enabled: 200 else 423 | Returns if the app is enabled               |
| /api/name             |                             | false | name of event            | Get the name of the current Event           |
| /api/login            | Email & Password & DeviceId | false | BrowserToken             | Auth Handler for the admin login            |
| /api/logout           |                             | false | 200                      | deletes the browser Token from the database |
| /api/checkLogin       |                             | false | 200                      | Returns 200, if the code is still valid     |
| /api/users            |                             | true  | UsersBody                | Returns all users                           |
| /api/users/create     | UsersBody with Password     | true  |                          | Create a new user                           |
| /api/users/update/:id | UsersBody with Password     | true  |                          | Update an existing owner                    |
| /api/users/delete/:id |                             | true  | 200                      | Delete a user (can be restored)             |
| /api/leagues          |                             | false | leaguesBody              | Returns a list of active leagues            |
| /api/leagues/update   | leaguesBody                 | yes   |                          | To update all leagues                       |
|                       |                             |       |                          |                                             |

# JSON Bodies
These are the bodies returned by the backend

### Leagues Body
```json
{
  "rescueLineEntry": "bool",
  "rescueLine": "bool",
  "rescueMazeEntry": "bool",
  "rescueMaze": "bool",
  
  "soccerEntry": "bool",
  "soccerLightWeightEntry": "bool",
  "soccerLightWeight": "bool",
  "soccerOpen": "bool",
  
  "onStageEntry": "bool",
  "onStage": "bool"
}
```

### Users Body

````json
{
  "id": "number",
  "username": "string",
  "email": "string",
  "password": "string | only on create & edit"
}
````