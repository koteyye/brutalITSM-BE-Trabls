package store

import "brutalITSM-BE-Trabls/internal/app/model"

type TrablRepository struct {
	store *Store
}

func (t TrablRepository) SelectService() (*model.Services, error) {
	s := &model.Services{}
	if err := t.store.db.QueryRow(
		"select id, \"name\", description, (select s3path from files f where f.id = s.logo_file_id) from services s").Scan(
		s.ID, s.Name, s.Description, s.LogoFileS3); err != nil {
		return nil, err
	}

	return nil, nil
}
