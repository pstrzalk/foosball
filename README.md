# Foosball Lambda API

### Lambdas

##### CreateMatch
Creates a match, with corresponding scores

```sh
# Example
./lambda_invoke.sh CreateMatch "{\"scores\":[{\"player_ids\":[1,2],\"score\":1},{\"player_ids\":[4,6],\"score\":3}]}"

# Response
"{\"id\":22,\"scores\":[{\"id\":42,\"players\":[{\"id\":1,\"name\":\"\"},{\"id\":2,\"name\":\"\"}],\"score\":1},{\"id\":43,\"players\":[{\"id\":4,\"name\":\"\"},{\"id\":6,\"name\":\"\"}],\"score\":3}]}"
```

##### CreatePlayer
Creates a player

Parameters:
- name

```sh
# Example (do not use lambda_invoke.sh when passing parameters with spaces)
./lambda_invoke.sh CreatePlayer "{\"name\":\"Samuel\"}"

# Response
"{\"id\":17,\"name\":\"Samuel\"}"
```

##### GetPlayers
Lists all players

```sh
# Example
./lambda_invoke.sh GetPlayers

# Response
"[{\"id\":3,\"name\":\"A\"},{\"id\":4,\"name\":\"B\"},{\"id\":5,\"name\":\"C\"},{\"id\":6,\"name\":\"D\"},{\"id\":7,\"name\":\"E\"},{\"id\":8,\"name\":\"F\"},{\"id\":9,\"name\":\"G\"},{\"id\":10,\"name\":\"H\"},{\"id\":1,\"name\":\"Z\"},{\"id\":2,\"name\":\"X\"},{\"id\":11,\"name\":\"Y\"},{\"id\":12,\"name\":\"Sam Wise\"}]"
```


##### GetMatches
Lists matches

Parameters:
- per_page [10|25|50], default = 10
- page, default = 1

```sh
# Example
./lambda_invoke.sh GetMatches "{\"per_page\":10,\"page\":1}"

# Response
"[{\"id\":13,\"scores\":[{\"id\":24,\"players\":[{\"id\":1,\"name\":\"Z\"},{\"id\":2,\"name\":\"X\"}],\"score\":1},{\"id\":25,\"players\":[{\"id\":4,\"name\":\"B\"},{\"id\":6,\"name\":\"D\"}],\"score\":3}]},{\"id\":14,\"scores\":[{\"id\":26,\"players\":[{\"id\":1,\"name\":\"Z\"},{\"id\":2,\"name\":\"X\"}],\"score\":1},{\"id\":27,\"players\":[{\"id\":4,\"name\":\"B\"},{\"id\":6,\"name\":\"D\"}],\"score\":3}]}]"
```

### Database

You may populate database with dump file in [sql|sql] folder.
