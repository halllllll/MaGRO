-- name: AddLog :exec

INSERT INTO logs(
  user_id, action_id
) VALUES($1, $2);