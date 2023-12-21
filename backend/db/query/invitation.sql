-- name: CreateGameInvitationWithUsername :exec
INSERT INTO game_invitations (inviter_player_id, invitee_player_id, game_id)
VALUES (
    $1, 
    (SELECT id FROM users WHERE username = $2),
    $3
);

-- name: ListGameInvitationsForUser :many
SELECT * FROM game_invitations
WHERE invitee_player_id = $1;

-- name: AcceptGameInvitation :exec
INSERT INTO player_game (player_id, game_id)
SELECT users.id, game_invitations.game_id
FROM users
INNER JOIN game_invitations ON users.id = game_invitations.invitee_player_id
WHERE users.id = $1 AND game_invitations.game_id = $2;

-- name: DeleteGameInvitation :exec
DELETE FROM game_invitations
WHERE invitee_player_id = $1 AND game_id = $2;

-- name: IsUserGameCreator :one
SELECT EXISTS (
  SELECT 1 FROM games
  WHERE created_by = $1 AND game_id = $2
) AS is_creator;

-- name: DoesInvitationExist :one
SELECT EXISTS (
    SELECT 1 FROM game_invitations
    WHERE invitee_player_id = $1 AND game_id = $2
) AS exists;
