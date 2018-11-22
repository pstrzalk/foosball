Player
- id
- name

Team
- id

PlayerTeam
- player_id
- team_id

Match
- played_at

MatchTeamScore
- score
- team_id
- match_id


SEND
team[0]players[0] = PLAYER_ID_1
team[0]players[1] = PLAYER_ID_2
team[0]score = SCORE_1

team[1]players[0] = PLAYER_ID_3
team[1]players[1] = PLAYER_ID_4
team[1]score = SCORE_2

SAVE

1a/ Get Team 1 Id
SELECT pt_1.team_id
  FROM player_teams pt_1,
       player_teams pt_2,
  WHERE pt_1.player_id = PLAYER_ID_1 AND
        pt_2.player_id = PLAYER_ID_2

2a/ If Team 1 id blank, create team and get ID. Create Team Players
INSERT INTO team (created_at) VALUES(NOW()) RETURNING id
INSERT INTO player_teams (player_id, team_id) VALUES($player_id_1, $team_id)
INSERT INTO player_teams (player_id, team_id) VALUES($player_id_2, $team_id)



1b/ Get Team 2 Id
SELECT pt_1.team_id
  FROM player_teams pt_1,
       player_teams pt_2,
  WHERE pt_1.player_id = PLAYER_ID_1 AND
        pt_2.player_id = PLAYER_ID_2

2b/ If Team 2 id blank, create team and get ID. Create Team Players
INSERT INTO team (created_at) VALUES(NOW()) RETURNING id
INSERT INTO player_teams (player_id, team_id) VALUES($player_id_1, $team_id)
INSERT INTO player_teams (player_id, team_id) VALUES($player_id_2, $team_id)



3/ Create match
INSERT INTO MATCH (played_at) VALUES(NOW()) RETURNING id
INSERT INTO match_team_score (match_id, team_id, score) VALUES($match_id, $team_id_1, score_2)
INSERT INTO match_team_score (match_id, team_id, score) VALUES($match_id, $team_id_2, score_2)
