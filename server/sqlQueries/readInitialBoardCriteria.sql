SELECT c.* FROM criteria c
INNER JOIN gametocriteria g
ON c.id = g.criteria_id
WHERE g.game_id = $1
ORDER BY random() LIMIT 25;