package repository

func (r *repository) CloseConnection() {
	r.db.Close()
}
