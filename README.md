The Task

Given API endpoint:

https://api-origin.onefootball.com/score-one-proxy/api/teams/en/{team_id}.json
with the following characteristics:

The team ID (team_id) is an unsigned integer.
The team IDs are sequential.
Some team IDs respond with errors.
Using the above API endpoint, find the following teams by name:

Germany
England
France
Spain
Manchester United
Arsenal
Chelsea
Barcelona
Real Madrid
Bayern Munich
Extract all the players from the given teams and render to stdout the information about players ordered by their playerID.

Each player entry should contain the following information:

playerID; full name; age; list of teams
The goal of the task is to print the players of the given teams. If there is a team overlap (e.g. a player plays in Germany and Chelsea), then print both teams.

Output Example:

1. 6232; Bernd Leno; 29; Arsenal, Germany
2. 6683; Alexander Mustermann; 25; France, Manchester Utd
3. 8010; Marcelo; 33; Real Madrid
4. ...
Delivery of the task:

Push your code to the main branch of this repository. Feel free to makes any changes you want in this repo (including the README.md).

FAQs

Submission

The solution should use the latest version of Go.
Feel free to use third-party libraries for the solution.
Code should be tested and well documentated.
Using concurrency in the solution is optional.
The API Endpoint URL and the list of teams should be easily configurable in the solution.
API

Discovering the valid team IDs is a part of the task.
All teams on the requested list can be found through the API endpoint.
The solution should not scan all the IDs on the API; only enough to get information about all the teams listed.
Some API calls might return errors, your code should handle that gracefully.
It's safe to assume that a player can be a part of maximum of 2 teams.
