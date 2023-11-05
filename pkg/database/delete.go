package database

import "log/slog"

func (d Database) Delete(what string) {
	slog.Debug("ID passed to delete", what)
	data, err := d.db.Delete(what)
	if err != nil {
		slog.Error("Unable to delete data in table: %v", err)
	}
	slog.Info("Successfully deleted data in table", data)
	return
}
