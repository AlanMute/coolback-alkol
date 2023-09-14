package repository

func (r *repository) AddCourse(name string, description string, folderName string) error {
	_, err := r.db.Exec("call add_course($1, $2, $3)", name, description, folderName)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) AddModule(name string, description string, id int, folderName string) error {
	_, err := r.db.Exec("call add_course($1, $2, $3, $4)", name, description, id, folderName)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) AddLesson(name string, description string, id int, fileName string) error {
	_, err := r.db.Exec("call add_course($1, $2, $3, $4)", name, description, id, fileName)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) CloseConnection() {
	r.db.Close()
}
