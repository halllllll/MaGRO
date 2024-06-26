-- dockerイメージ作成時に初期データを読み込ませている

-- name: UpdateSystem :one
UPDATE system
SET version = $1
WHERE id = 1 -- 常に一行目に期待されるデータが格納されているとする
RETURNING *;

-- name: GetSystemInfo :one
SELECT * FROM system WHERE id = 1;