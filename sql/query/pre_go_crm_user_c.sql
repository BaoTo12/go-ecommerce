-- name: GetUserByEmail :one
SELECT usr_id ,usr_email FROM `pre_go_crm_user_c` WHERE usr_email = ?;


-- name: UpdateUSerStatusByUserId :exec
UPDATE `pre_go_crm_user_c`
SET usr_status = $2,
    usr_updated_at = $3
WHERE usr_id = $1;