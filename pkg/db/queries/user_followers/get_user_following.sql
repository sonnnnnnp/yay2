-- name: GetUserFollowing :many
SELECT
    users.*,
    user_followers.created_at AS followed_at
FROM users
INNER JOIN user_followers ON users.id = user_followers.following_id
WHERE user_followers.follower_id = @user_id::uuid
ORDER BY user_followers.created_at DESC;