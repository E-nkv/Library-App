package models

import "fmt"

func createQueryGetUsersWithTags(limit int, lastID int64) string {
	var query string
	if lastID > 0 {
		query = fmt.Sprintf(`
		WITH users AS (
			SELECT * FROM USERS u WHERE u.id > %d
		)`, lastID)
	}
	query += fmt.Sprintf(`
		SELECT 
			u.id, 
			u.full_name, 
			email,
			COALESCE(
				JSON_AGG(
					JSON_BUILD_OBJECT('tag_id', ut.tag_id, 'tag_name', t.name)
				) FILTER (WHERE ut.tag_id IS NOT NULL),'[]'::JSON
			) AS tags
		FROM USERS u 
		LEFT JOIN users_tags ut ON(u.id = ut.user_id)
		LEFT JOIN tags t ON (t.id = ut.tag_id)
		GROUP BY (u.id, u.full_name, email)
		ORDER BY u.id
		LIMIT %d
		`, limit)
	return query
}

func createQueryGetUserWithTags(id int64) string {
	var query = fmt.Sprintf(`
		WITH selected_user AS (
			SELECT * FROM users u WHERE u.id = %d
		)
		SELECT 
			u.id, 
			full_name, 
			email, 
			hash_pass, 
			COALESCE(
				JSON_AGG(
					JSON_BUILD_OBJECT('id', ut.tag_id, 'name', t.name)
				) FILTER (WHERE ut.tag_id IS NOT NULL), 
				'[]'::JSON
			) AS tags
		FROM selected_user u 
		LEFT JOIN users_tags  ut ON (u.id = ut.user_id)
		LEFT JOIN tags t ON (t.id = ut.tag_id)
		GROUP BY (u.id, u.full_name, u.email, u.hash_pass)
	`, id)
	return query
}

func createQueryGetUser(id int64) string {
	return fmt.Sprintf("SELECT id, full_name, email, hash_pass, is_active, is_verified, role FROM USERS u WHERE u.id = %d", id)
}
func createQueryGetUserByEmail(email string) string {
	return fmt.Sprintf("SELECT id, full_name, email, hash_pass, is_active, is_verified, role FROM USERS u WHERE u.email = '%s'", email)
}

func createQueryGetUsers(limit int, lastID int64) string {
	var q string
	if lastID > 0 {
		q = fmt.Sprintf("WITH users AS (SELECT id, full_name, email, is_active, is_verified, role FROM users u WHERE u.id > %d)", lastID)
	}
	q += fmt.Sprintf("SELECT id, full_name, email, is_active, is_verified, role FROM users LIMIT %d", limit)
	return q
}
