package repositories

import (
	"api/src/models"
	"database/sql"
)

// Publications - Publications Repository
type Publications struct {
	db *sql.DB
}

// NewPublicationRepository - Create new publication repository
func NewPublicationRepository(db *sql.DB) *Publications {
	return &Publications{db}
}

// CreatePublication - Create one publication in database
func (repository Publications) CreatePublication(publication models.Publication) (uint64, error) {
	statement, err := repository.db.Prepare("insert into publications (title, content, author_id) values (?, ?, ?)")
	if err != nil {
		return 0, nil
	}
	defer statement.Close()

	result, err := statement.Exec(publication.Title, publication.Content, publication.AuthorID)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(lastID), nil
}

// GetPublicationByID - Get one publication based on id
func (repository Publications) GetPublicationByID(publicationID uint64) (models.Publication, error) {
	lines, err := repository.db.Query(`
		select p.*, u.nick from 
		publications p inner join users u
		on u.id = p.author_id where p.id = ?
	`, publicationID)
	if err != nil {
		return models.Publication{}, err
	}
	defer lines.Close()

	var publication models.Publication

	if lines.Next() {
		if err = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return models.Publication{}, err
		}
	}

	return publication, nil
}

// GetPublications - Return all publication with user and your followers
func (repository Publications) GetPublications(userID uint64) ([]models.Publication, error) {
	lines, err := repository.db.Query(`
		select distinct p.*, u.nick from publications p
		inner join users u on u.id = p.author_id
		inner join followers f on p.author_id = f.user_id
		where u.id = ? or f.follower_id = ?
	`, userID, userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var publications []models.Publication

	for lines.Next() {
		var publication models.Publication

		if err := lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// UpdatePublication - Update the publication
func (repository Publications) UpdatePublication(publicationID uint64, publication models.Publication) error {
	statement, err := repository.db.Prepare("update publications set title = ?, content = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publication.Title, publication.Content, publicationID); err != nil {
		return err
	}

	return nil
}

// DeletePublication - Delete publication in database
func (repository Publications) DeletePublication(publicationID uint64) error {
	statement, err := repository.db.Prepare("delete from publications where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicationID); err != nil {
		return err
	}
	return nil
}

// GetPublicationsForUser - Get all publications based on user ID
func (repository Publications) GetPublicationsForUser(userID uint64) ([]models.Publication, error) {
	lines, err := repository.db.Query(`
		select p.*, u.nick from publications p
		join users u on u.id = p.author_id
		where p.author_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var publications []models.Publication

	for lines.Next() {
		var publication models.Publication
		if err = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// LikePublication - Method to like publication
func (repository Publications) LikePublication(publicationID uint64) error {
	statement, err := repository.db.Prepare("update publications set likes = likes + 1 where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}

// DeslikePublication - Remove like from publication
func (repository Publications) DeslikePublication(publicationID uint64) error {
	statement, err := repository.db.Prepare(`
		update publications set likes = 
		CASE WHEN likes > 0 THEN likes - 1
		ELSE likes END
		where id = ?
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}
