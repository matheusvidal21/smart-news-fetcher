package database

import (
	"database/sql"
	"github.com/matheusvidal21/smart-news-fetcher/internal/models"
	"github.com/matheusvidal21/smart-news-fetcher/pkg/utils"
)

type SourceRepository struct {
	db *sql.DB
}

func NewSourceRepository(db *sql.DB) *SourceRepository {
	return &SourceRepository{db: db}
}

func (sr *SourceRepository) FindAll(page, limit int, sort string) ([]models.Source, error) {
	sql := "SELECT id, name, url, saved_at, user_id, update_interval FROM sources"
	offset := (page - 1) * limit

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		sql = sql + " ORDER BY saved_at " + sort + " LIMIT ? OFFSET ? "
	} else {
		sql = sql + " ORDER BY saved_at " + sort
	}

	rows, err := sr.db.Query(sql, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sources []models.Source
	for rows.Next() {
		var source models.Source
		var savedAt []byte
		err = rows.Scan(&source.ID, &source.Name, &source.Url, &savedAt, &source.UserID, &source.UpdateInterval)
		if err != nil {
			return nil, err
		}

		source.SavedAt, err = utils.ParseTime(savedAt)
		if err != nil {
			return nil, err
		}
		sources = append(sources, source)
	}
	return sources, nil
}

func (sr *SourceRepository) FindOne(id int) (models.Source, error) {
	stmt, err := sr.db.Prepare("SELECT id, name, url, saved_at, user_id, update_interval FROM sources WHERE id = ?")
	if err != nil {
		return models.Source{}, err
	}
	defer stmt.Close()

	var source models.Source
	var savedAt []byte
	err = stmt.QueryRow(id).Scan(&source.ID, &source.Name, &source.Url, &savedAt, &source.UserID, &source.UpdateInterval)
	if err != nil {
		return models.Source{}, err
	}

	source.SavedAt, err = utils.ParseTime(savedAt)
	if err != nil {
		return models.Source{}, err
	}
	return source, nil
}

func (sr *SourceRepository) Create(source models.Source) (models.Source, error) {
	stmt, err := sr.db.Prepare("INSERT INTO sources (name, url, saved_at, user_id, update_interval) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return models.Source{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(source.Name, source.Url, source.SavedAt, source.UserID, source.UpdateInterval)
	if err != nil {
		return models.Source{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return models.Source{}, err
	}

	return models.Source{
		ID:             int(id),
		Name:           source.Name,
		Url:            source.Url,
		UpdateInterval: source.UpdateInterval,
		SavedAt:        source.SavedAt,
		UserID:         source.UserID,
	}, nil
}

func (sr *SourceRepository) Update(id int, source models.Source) (models.Source, error) {
	stmt, err := sr.db.Prepare("UPDATE sources SET name = ?, url = ?, update_interval = ? WHERE id = ?")
	if err != nil {
		return models.Source{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(source.Name, source.Url, source.UpdateInterval, id)
	if err != nil {
		return models.Source{}, err
	}

	updatedSource, err := sr.FindOne(id)
	if err != nil {
		return models.Source{}, err
	}

	return updatedSource, nil
}

func (sr *SourceRepository) Delete(id int) error {
	stmt, err := sr.db.Prepare("DELETE FROM sources WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (sr *SourceRepository) FindByUrl(url string) (models.Source, error) {
	stmt, err := sr.db.Prepare("SELECT id, name, url, saved_at, user_id, update_interval FROM sources WHERE url = ?")
	if err != nil {
		return models.Source{}, err
	}
	defer stmt.Close()

	var source models.Source
	var savedAt []byte
	err = stmt.QueryRow(url).Scan(&source.ID, &source.Name, &source.Url, &savedAt, &source.UserID, &source.UpdateInterval)
	if err != nil {
		return models.Source{}, err
	}
	source.SavedAt, err = utils.ParseTime(savedAt)
	if err != nil {
		return models.Source{}, err
	}
	return source, nil
}

func (sr *SourceRepository) FindByUserId(userId int) ([]models.Source, error) {
	rows, err := sr.db.Query("SELECT id, name, url, saved_at, update_interval FROM sources WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sources []models.Source
	for rows.Next() {
		var source models.Source
		var savedAt []byte
		err = rows.Scan(&source.ID, &source.Name, &source.Url, &savedAt, &source.UpdateInterval)
		if err != nil {
			return nil, err
		}
		source.SavedAt, err = utils.ParseTime(savedAt)
		if err != nil {
			return nil, err
		}
		sources = append(sources, source)
	}
	return sources, nil
}
