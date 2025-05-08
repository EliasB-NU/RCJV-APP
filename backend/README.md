# API Specification

| Endpoint                        | Payload                     | Auth  | Return       | Description                                 |
|---------------------------------|-----------------------------|-------|--------------|---------------------------------------------|
| /api/login                      | Email & Password & DeviceId | false | BrowserToken | Auth Handler for the admin login            |
| /api/logout                     | Token                       | false | 200          | deletes the browser Token from the database |
| /api/checkLogin                 | Token                       | false | 200          | Returns 200, if the code is still valid     |
| /api/leagues                    |                             | false | leaguesBody  | Returns a list of active leagues            |
| /api/leagues/activate/:league   | league to activate          | yes   |              | To activate a league                        |
| /api/leagues/deactivate/:league | league to deactivate        | yes   |              | To activate a league                        |
|                                 |                             |       |              |                                             |

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