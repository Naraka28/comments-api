package comments

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type sqlCommentRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &sqlCommentRepository{db: db}
}

func (r *sqlCommentRepository) FindAll() ([]Comment, error) {
	rows, err := r.db.Query("SELECT id, username, message, date FROM comments ORDER BY date DESC")
	if err != nil {
		return []Comment{}, fmt.Errorf("db: error al obtener comentarios: %w", err)
	}
	defer rows.Close()

	comments := []Comment{}

	for rows.Next() {
		var c Comment
		if err := rows.Scan(&c.Id, &c.Username, &c.Message, &c.Date); err != nil {
			return comments, err
		}
		comments = append(comments, c)
	}

	return comments, nil
}

func (r *sqlCommentRepository) Save(data NewComment) (Comment, error) {
    now := time.Now().UTC()
    result, err := r.db.Exec(
        "INSERT INTO comments (username, message, date) VALUES (?, ?, ?)",
        data.Username, data.Message, now,
    )
    if err != nil {
        return Comment{}, fmt.Errorf("db: error al guardar: %w", err)
    }
    id, _ := result.LastInsertId()
    return r.FindById(int(id))
}

func (r *sqlCommentRepository) FindById(id int) (Comment, error) {
	var c Comment
	err := r.db.QueryRow("SELECT id, username, message, date FROM comments WHERE id = ?", id).
		Scan(&c.Id, &c.Username, &c.Message, &c.Date)

	if err != nil {
		if err == sql.ErrNoRows {
			return c, errors.New("comentario no encontrado")
		}
		return c, err
	}
	return c, nil
}

func (r *sqlCommentRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM comments WHERE id = ?", id)
	return err
}