# Admin

The admin control panel can be found at [http://localhost:3006](http://localhost:3006)
The default email is: `admin@example.com` and the default password: `admin`

# API Specification

| Endpoint                                   | Payload                     | Auth  | Return                   | Description                                              |
|--------------------------------------------|-----------------------------|-------|--------------------------|----------------------------------------------------------|
| /api/v1/config                             |                             | true  | current config           | Getting the current config for the frontend              |
| /api/v1/config/update                      | Updated Config              | true  |                          | Updating the current config                              |
| /api/v1/enabled                            |                             | false | if enabled: 200 else 423 | Returns if the app is enabled                            |
| /api/v1/name                               |                             | false | name of event            | Get the name of the current Event                        |
| /api/v1/rescueURL                          |                             | false | rescue url               | To get the url to the rescue standings                   |
| /api/v1/leagues/update                     | leaguesBody                 | yes   |                          | To update all leagues                                    |
| /api/v1/teams                              |                             | false | teamsBody                | Get all teams                                            |
|                                            |                             |       |                          |                                                          |
| /api/v1/login                              | Email & Password & DeviceId | false | BrowserToken             | Auth Handler for the admin login                         |
| /api/v1/logout                             |                             | false | 200                      | deletes the browser Token from the database              |
| /api/v1/checkLogin                         |                             | false | 200                      | Returns 200, if the code is still valid                  |
|                                            |                             |       |                          |                                                          |
| /api/v1/users                              |                             | true  | UsersBody                | Returns all users                                        |
| /api/v1/users/create                       | UsersBody with Password     | true  |                          | Create a new user                                        |
| /api/v1/users/update/:id                   | UsersBody with Password     | true  |                          | Update an existing owner                                 |
| /api/v1/users/delete/:id                   |                             | true  | 200                      | Delete a user (can be restored)                          |
|                                            |                             |       |                          |                                                          |
| /api/v1/teams                              |                             | false | teamsBody                | Get all teams                                            |
| /api/v1/teams/create                       | teamsBody                   | true  |                          | Create a new team                                        |
| /api/v1/teams/update/:id                   | teamsBody                   | true  |                          | Update an existing team                                  |
| /api/v1/teams/delete/:id                   | id                          | true  |                          | Delete a team (can be restored)                          |
|                                            |                             |       |                          |                                                          |
| /api/v1/institutions                       |                             | false | institutionsBody         | Get all institutions                                     |
| /api/v1/institution/create                 | institutionBody             | true  |                          | Create a new institution                                 |
| /api/v1/institution/update/:id             | institutionBody             | true  |                          | Update a existing institution                            |
| /api/v1/institution/delete/:id             | id                          | true  |                          | Delete an institution (with all teams) (can be restored) |
|                                            |                             |       |                          |                                                          |
| /api/v1/fields                             |                             | true  | fieldsBody               | Get all fields                                           |
| /api/v1/fields/create                      | fieldsBody                  | true  |                          | Create a new field                                       |
| /api/v1/fields/update/:id                  | id, fieldsBody              | true  |                          | Update an existing field                                 |
| /api/v1/fields/delete/:id                  | id                          | true  |                          | Delete an existing field                                 |
|                                            |                             |       |                          |                                                          |
| /api/v1/matches                            |                             | false | matchesBody              | Get all matches                                          |
| /api/v1/matches/league/:league             | league                      | false | matchesBody              | Get all matches by league                                |
| /api/v1/matches/team/:teamID               | teamID                      | false | matchesBody              | Get all matches by team                                  |
| /api/v1/matches/institution/:institutionID | institutionID               | false | matchesBody              | Get all matches by institution                           |
| /api/v1/matches/field/:fieldID             | fieldID                     | false | matchesBody              | Get all matches by field                                 |
| /api/v1/matches/upload/:league             | league, file                | true  |                          | Upload a finished spreadsheet                            |
| /api/v1/matches/create/:league             | league                      | true  |                          | Generate a new spreadsheet for the specified league      |
| /api/v1/matches/update/:id                 | id, matchesBody             | true  |                          | Update a match by its id                                 |
| /api/v1/matches/delete/:id                 | id                          | true  |                          | Delete a match by its id                                 |

"Can be restored": [Gorm](https://gorm.io/) (the database interface used by me) sets the `deleted_at` column to the
current time on
a delete action, so you can go into the database and set it to null again to restore the lost data.

We don't plan any major changes in the future, but just to be sure, there will be legacy endpoints.

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
  "other": [ // For rescue and onstage matches
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
  ],
  "soccer": { // For all soccer matches
    "tournament_abbrev": "string",
    "tournament_name": "string",
    "tournament_header_type": "string",
    "last_updated": "time.Time",
    "matches": [
      {
        "number": "number",
        "last_published": "string", // Format: yyyy-mm-ddThh:mm:ss
        "league": "string",
        "league_stage": "string",
        "group_name": "string",
        "start": "string", // Format: yyyy-mm-ddThh-mm-ss
        "duration": "string", // Format: hh:mm:ss
        "pitch": "string", // Field
        "goals1": "number",
        "points1": "number",
        "team1": {
          "id": "number",
          "name": "string",
          "affiliation": "string", // Institution
          "startnumber": "number",
          "external_key": "string"
        },
        "goals2": "number",
        "points2": "number",
        "team2": {
          "id": "number",
          "name": "string",
          "affiliation": "string",
          "startnumber": "number",
          "external_key": "string"
        },
        "referees": [
          {
            "first_name": "string",
            "last_name": "string"
          }
        ]
      }
    ]
  }
}
````