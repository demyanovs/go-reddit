package postgres

import (
	"fmt"

	goreddit "github.com/demyanovs/go-reddit"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewThreadStore(db *sqlx.DB) *ThreadStore {
	return &ThreadStore{
		DB: db,
	}
}

type ThreadStore struct {
	*sqlx.DB
}

func (s *PostStore) Post(id uuid.UUID) (goreddit.Post, error) {
	var p goreddit
	if err := s.Get(&p, `SELECT * FROM posts WHERE id = $1`, id); err != nil {
		return goreddit.Post{}, fmt.Errorf("error getting post: %w", err)
	}
	return p, nil
}

func (s *PostStore) PostsByThread(threadID uuid.UUID) ([]goreddit.Post, error) {
	var pp []goreddit.Post
	if err := s.Select(&pp, `SELECT * FROM posts WHERE thread_id = $1`, threadID); err != nil {
		return []goreddit.Post{}, fmt.Errorf("error getting posts: %w", err)
	}
	return pp, nil
}

func (s *PostStore) CreatePost(p *goreddit.Post) error {
	if err := s.Get(t, `INSERT INTO posts VALUES ($1, $2, $3, $4, $5) RETURNING *`,
		p.ID,
		p.threadID,
		p.Title,
		p.Content,
		p.Votes); err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}
	return nil
}

func (s *PostStore) UpdatePost(p *goreddit.Post) error {
	if err := s.Get(t, `UPDATE posts SET thread_id = $1, title = $2, content = $3, votes =$4 WHERE id = $5 RETURNING *`,
		p.threadID
		p.Title,
		p.Content,
		p.Votes,
		p.ID); err != nil {
		return fmt.Errorf("error updating post: %w", err)
	}
	return nil
}

func (s *PostStore) DeletePost(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM posts WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting post: %w", err)
	}
	return nil
}
