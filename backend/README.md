# Admin

The admin control panel can be found at [http://localhost:3006](http://localhost:3006)
The default email is: `admin@example.com` and the default password: `admin`

# API Specification

| Endpoint                                | Payload                     | Auth  | Return                   | Description                                              |
|-----------------------------------------|-----------------------------|-------|--------------------------|----------------------------------------------------------|
| /api/config                             |                             | true  | current config           | Getting the current config for the frontend              |
| /api/config/update                      | Updated Config              | true  |                          | Updating the current config                              |
| /api/enabled                            |                             | false | if enabled: 200 else 423 | Returns if the app is enabled                            |
| /api/name                               |                             | false | name of event            | Get the name of the current Event                        |
| /api/leagues/update                     | leaguesBody                 | yes   |                          | To update all leagues                                    |
| /api/teams                              |                             | false | teamsBody                | Get all teams                                            |
|                                         |                             |       |                          |                                                          |
| /api/login                              | Email & Password & DeviceId | false | BrowserToken             | Auth Handler for the admin login                         |
| /api/logout                             |                             | false | 200                      | deletes the browser Token from the database              |
| /api/checkLogin                         |                             | false | 200                      | Returns 200, if the code is still valid                  |
|                                         |                             |       |                          |                                                          |
| /api/users                              |                             | true  | UsersBody                | Returns all users                                        |
| /api/users/create                       | UsersBody with Password     | true  |                          | Create a new user                                        |
| /api/users/update/:id                   | UsersBody with Password     | true  |                          | Update an existing owner                                 |
| /api/users/delete/:id                   |                             | true  | 200                      | Delete a user (can be restored)                          |
|                                         |                             |       |                          |                                                          |
| /api/teams                              |                             | false | teamsBody                | Get all teams                                            |
| /api/teams/create                       | teamsBody                   | true  |                          | Create a new team                                        |
| /api/teams/update/:id                   | teamsBody                   | true  |                          | Update an existing team                                  |
| /api/teams/delete/:id                   | id                          | true  |                          | Delete a team (can be restored)                          |
|                                         |                             |       |                          |                                                          |
| /api/institutions                       |                             | false | institutionsBody         | Get all institutions                                     |
| /api/institution/create                 | institutionBody             | true  |                          | Create a new institution                                 |
| /api/institution/update/:id             | institutionBody             | true  |                          | Update a existing institution                            |
| /api/institution/delete/:id             | id                          | true  |                          | Delete an institution (with all teams) (can be restored) |
|                                         |                             |       |                          |                                                          |
| /api/fields                             |                             | true  | fieldsBody               | Get all fields                                           |
| /api/fields/create                      | fieldsBody                  | true  |                          | Create a new field                                       |
| /api/fields/update/:id                  | id, fieldsBody              | true  |                          | Update an existing field                                 |
| /api/fields/delete/:id                  | id                          | true  |                          | Delete an existing field                                 |
|                                         |                             |       |                          |                                                          |
| /api/matches                            |                             | false | matchesBody              | Get all matches                                          |
| /api/matches/league/:league             | league                      | false | matchesBody              | Get all matches by league                                |
| /api/matches/team/:teamID               | teamID                      | false | matchesBody              | Get all matches by team                                  |
| /api/matches/institution/:institutionID | institutionID               | false | matchesBody              | Get all matches by institution                           |
| /api/matches/field/:fieldID             | fieldID                     | false | matchesBody              | Get all matches by field                                 |
| /api/matches/upload/:league             | league, file                | true  |                          | Upload a finished spreadsheet                            |
| /api/matches/create/:league             | league                      | true  |                          | Generate a new spreadsheet for the specified league      |
| /api/matches/update/:id                 | id, matchesBody             | true  |                          | Update a match by its id                                 |
| /api/matches/delete/:id                 | id                          | true  |                          | Delete a match by its id                                 |

"Can be restored": [Gorm](https://gorm.io/) (the database interface used by me) sets the `deleted_at` column to the
current time on
a delete action, so you can go into the database and set it to null again to restore the lost data.

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
[
  {
    "id": "number",
    "username": "string",
    "email": "string",
    "password": "string | only on create & edit"
  }
]
````

### Teams Body

```json
{
  "lastUpdate": "time.Time",
  "data": [
    {
      "id": "number",
      "name": "string",
      "league": "league | corresponds to the naming convention from the leagues body",
      "institution": "string"
    }
  ]
}
```

### Institutions Body

```json
{
  "lastUpdate": "time.Time",
  "data": [
    {
      "id": "number",
      "name": "string",
      "numberTeams": "number"
    }
  ]
}
```

### Matches Body

````json
{
  "lastRequested": "time.Time",
  "matches": [
    {
      "id": "number",
      "updatedAt": "time.Time",
      "league": "string | corresponds to the naming convention from the leagues body",
      "name": "string",
      "startTime": "time.Time",
      "duration": "time.Duration",
      "field": "string",
      
      "institutionID": "number",
      "institutionName": "string",
      
      "teamID": "number",
      "teamName": "string"
    }
  ]
}
````