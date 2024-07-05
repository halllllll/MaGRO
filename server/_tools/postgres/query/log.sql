-- name: AddLog :exec

INSERT INTO log(
  user_id, action_id
) VALUES($1, $2);