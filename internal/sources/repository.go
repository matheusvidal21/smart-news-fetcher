package sources

import (
	"database/sql"
	"github.com/matheusvidal21/smart-news-fetcher/pkg/utils"
)

type SourceRepositoryInterface interface {
	FindAll(page, limit int, sort string) ([]Source, error)
	FindOne(id int) (Source, error)
	Create(source Source) (Source, error)
	Update(id int, source Source) (Source, error)
	Delete(id int) error
}

type SourceRepository struct {
	db *sql.DB
}

func NewSourceRepository(db *sql.DB) *SourceRepository {
	return &SourceRepository{db: db}
}

func (sr *SourceRepository) FindAll(page, limit int, sort string) ([]Source, error) {
	sql := "SELECT id, name, url, saved_at FROM sources"
	offset := (page - 1) * limit

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		sql = sql + " ORDER BY created_at " + sort + " LIMIT ? OFFSET ? "
	} else {
		sql = sql + " ORDER BY created_at " + sort
	}

	rows, err := sr.db.Query(sql, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sources []Source
	for rows.Next() {
		var source Source
		var savedAt []byte
		err = rows.Scan(&source.ID, &source.Name, &source.Url, &savedAt)
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

func (sr *SourceRepository) FindOne(id int) (Source, error) {
	stmt, err := sr.db.Prepare("SELECT id, name, url, saved_at FROM sources WHERE id = ?")
	if err != nil {
		return Source{}, err
	}
	defer stmt.Close()

	var source Source
	var savedAt []byte
	err = stmt.QueryRow(id).Scan(&source.ID, &source.Name, &source.Url, &savedAt)
	if err != nil {
		return Source{}, err
	}

	source.SavedAt, err = utils.ParseTime(savedAt)
	if err != nil {
		return Source{}, err
	}
	return source, nil
}

func (sr *SourceRepository) Create(source Source) (Source, error) {
	stmt, err := sr.db.Prepare("INSERT INTO sources (name, url, saved_at) VALUES (?, ?, ?)")
	if err != nil {
		return Source{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(source.Name, source.Url, source.SavedAt)
	if err != nil {
		return Source{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return Source{}, err
	}

	return Source{
		ID:      int(id),
		Name:    source.Name,
		Url:     source.Url,
		SavedAt: source.SavedAt,
	}, nil
}

func (sr *SourceRepository) Update(id int, source Source) (Source, error) {
	stmt, err := sr.db.Prepare("UPDATE sources SET name = ?, url = ? WHERE id = ?")
	if err != nil {
		return Source{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(source.Name, source.Url, id)
	if err != nil {
		return Source{}, err
	}

	updatedSource, err := sr.FindOne(id)
	if err != nil {
		return Source{}, err
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
