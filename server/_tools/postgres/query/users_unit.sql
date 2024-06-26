-- name: MapUserUnit :exec
INSERT INTO users_unit(
  user_id, unit_id
) VALUES($1, $2);
