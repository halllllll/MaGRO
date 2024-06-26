-- name: MapUserSubunit :exec
INSERT INTO users_subunit(
  user_id, subunit_id
) VALUES($1, $2);
