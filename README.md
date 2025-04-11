# PhotonAPI

Server side API for the PhotonBoard system.

## Install and run

 - Clone repository and cd into it
 - go mod tidy
 - go run cmd/photonapi.go

You can pass the -testdata argument if you want the testdata to be created on the database


## Endpoints

All data sent and received is using JSON format. Binary data line wall images are base64 encoded. All uids are UUIDs V4.

- /api (GET) : Get api info data, version, etc.
- /api/walls (GET) : Get all walls. You can pass the parameter "fields" to get just the fields you need (api/walls?fields=uid,name for example).
- /api/wall/{uid} (GET) : Get wall by uid.
- /api/newwall (POST) : Upload a new wall. Uses HTTP Basic Auth.
- /api/problems (GET) : Get all problems.
- /api/problem/{uid} (GET) : Get problem by uid.
- /api/newproblem (POST) : Upload a new problem. Uses HTTP Basic Auth.