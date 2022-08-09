package rekordbox

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// DjmdSongPlaylist represents a row from 'djmdSongPlaylist'.
type DjmdSongPlaylist struct {
	ID                sql.NullString `json:"ID"`                   // ID
	PlaylistID        sql.NullString `json:"PlaylistID"`           // PlaylistID
	ContentID         sql.NullString `json:"ContentID"`            // ContentID
	TrackNo           sql.NullInt64  `json:"TrackNo"`              // TrackNo
	UUID              sql.NullString `json:"UUID"`                 // UUID
	RbDataStatus      sql.NullInt64  `json:"rb_data_status"`       // rb_data_status
	RbLocalDataStatus sql.NullInt64  `json:"rb_local_data_status"` // rb_local_data_status
	RbLocalDeleted    sql.NullInt64  `json:"rb_local_deleted"`     // rb_local_deleted
	RbLocalSynced     sql.NullInt64  `json:"rb_local_synced"`      // rb_local_synced
	Usn               sql.NullInt64  `json:"usn"`                  // usn
	RbLocalUsn        sql.NullInt64  `json:"rb_local_usn"`         // rb_local_usn
	CreatedAt         Time           `json:"created_at"`           // created_at
	UpdatedAt         Time           `json:"updated_at"`           // updated_at
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the DjmdSongPlaylist exists in the database.
func (dsp *DjmdSongPlaylist) Exists() bool {
	return dsp._exists
}

// Deleted returns true when the DjmdSongPlaylist has been marked for deletion from
// the database.
func (dsp *DjmdSongPlaylist) Deleted() bool {
	return dsp._deleted
}

// Insert inserts the DjmdSongPlaylist to the database.
func (c *Client) InsertDjmdSongPlaylist(ctx context.Context, dsp *DjmdSongPlaylist) error {
	db := c.db

	switch {
	case dsp._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case dsp._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO djmdSongPlaylist (` +
		`ID, PlaylistID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`)`
	// run
	logf(sqlstr, dsp.ID, dsp.PlaylistID, dsp.ContentID, dsp.TrackNo, dsp.UUID, dsp.RbDataStatus, dsp.RbLocalDataStatus, dsp.RbLocalDeleted, dsp.RbLocalSynced, dsp.Usn, dsp.RbLocalUsn, dsp.CreatedAt, dsp.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dsp.ID, dsp.PlaylistID, dsp.ContentID, dsp.TrackNo, dsp.UUID, dsp.RbDataStatus, dsp.RbLocalDataStatus, dsp.RbLocalDeleted, dsp.RbLocalSynced, dsp.Usn, dsp.RbLocalUsn, dsp.CreatedAt, dsp.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dsp._exists = true
	return nil
}

// Update updates a DjmdSongPlaylist in the database.
func (c *Client) UpdateDjmdSongPlaylist(ctx context.Context, dsp *DjmdSongPlaylist) error {
	db := c.db

	switch {
	case !dsp._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case dsp._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE djmdSongPlaylist SET ` +
		`PlaylistID = $1, ContentID = $2, TrackNo = $3, UUID = $4, rb_data_status = $5, rb_local_data_status = $6, rb_local_deleted = $7, rb_local_synced = $8, usn = $9, rb_local_usn = $10, created_at = $11, updated_at = $12 ` +
		`WHERE ID = $13`
	// run
	logf(sqlstr, dsp.PlaylistID, dsp.ContentID, dsp.TrackNo, dsp.UUID, dsp.RbDataStatus, dsp.RbLocalDataStatus, dsp.RbLocalDeleted, dsp.RbLocalSynced, dsp.Usn, dsp.RbLocalUsn, dsp.CreatedAt, dsp.UpdatedAt, dsp.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dsp.PlaylistID, dsp.ContentID, dsp.TrackNo, dsp.UUID, dsp.RbDataStatus, dsp.RbLocalDataStatus, dsp.RbLocalDeleted, dsp.RbLocalSynced, dsp.Usn, dsp.RbLocalUsn, dsp.CreatedAt, dsp.UpdatedAt, dsp.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the DjmdSongPlaylist to the database.
func (c *Client) SaveDjmdSongPlaylist(ctx context.Context, dsp *DjmdSongPlaylist) error {
	if dsp.Exists() {
		return c.UpdateDjmdSongPlaylist(ctx, dsp)
	}
	return c.InsertDjmdSongPlaylist(ctx, dsp)
}

// Upsert performs an upsert for DjmdSongPlaylist.
func (c *Client) UpsertDjmdSongPlaylist(ctx context.Context, dsp *DjmdSongPlaylist) error {
	db := c.db

	switch {
	case dsp._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO djmdSongPlaylist (` +
		`ID, PlaylistID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`)` +
		` ON CONFLICT (ID) DO ` +
		`UPDATE SET ` +
		`PlaylistID = EXCLUDED.PlaylistID, ContentID = EXCLUDED.ContentID, TrackNo = EXCLUDED.TrackNo, UUID = EXCLUDED.UUID, rb_data_status = EXCLUDED.rb_data_status, rb_local_data_status = EXCLUDED.rb_local_data_status, rb_local_deleted = EXCLUDED.rb_local_deleted, rb_local_synced = EXCLUDED.rb_local_synced, usn = EXCLUDED.usn, rb_local_usn = EXCLUDED.rb_local_usn, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, dsp.ID, dsp.PlaylistID, dsp.ContentID, dsp.TrackNo, dsp.UUID, dsp.RbDataStatus, dsp.RbLocalDataStatus, dsp.RbLocalDeleted, dsp.RbLocalSynced, dsp.Usn, dsp.RbLocalUsn, dsp.CreatedAt, dsp.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dsp.ID, dsp.PlaylistID, dsp.ContentID, dsp.TrackNo, dsp.UUID, dsp.RbDataStatus, dsp.RbLocalDataStatus, dsp.RbLocalDeleted, dsp.RbLocalSynced, dsp.Usn, dsp.RbLocalUsn, dsp.CreatedAt, dsp.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dsp._exists = true
	return nil
}

// Delete deletes the DjmdSongPlaylist from the database.
func (c *Client) DeleteDjmdSongPlaylist(ctx context.Context, dsp *DjmdSongPlaylist) error {
	db := c.db

	switch {
	case !dsp._exists: // doesn't exist
		return nil
	case dsp._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM djmdSongPlaylist ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, dsp.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dsp.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	dsp._deleted = true
	return nil
}

func scanDjmdSongPlaylistRows(rows *sql.Rows) ([]*DjmdSongPlaylist, error) {
	var res []*DjmdSongPlaylist
	for rows.Next() {
		dsp := DjmdSongPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dsp.ID, &dsp.PlaylistID, &dsp.ContentID, &dsp.TrackNo, &dsp.UUID, &dsp.RbDataStatus, &dsp.RbLocalDataStatus, &dsp.RbLocalDeleted, &dsp.RbLocalSynced, &dsp.Usn, &dsp.RbLocalUsn, &dsp.CreatedAt, &dsp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dsp)
	}

	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

func (c *Client) AllDjmdSongPlaylist(ctx context.Context) ([]*DjmdSongPlaylist, error) {
	db := c.db

	const sqlstr = `SELECT * FROM DjmdSongPlaylist`
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()
	res, err := scanDjmdSongPlaylistRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongPlaylistByContentID retrieves a row from 'djmdSongPlaylist' as a DjmdSongPlaylist.
//
// Generated from index 'djmd_song_playlist__content_i_d'.
func (c *Client) DjmdSongPlaylistByContentID(ctx context.Context, contentID sql.NullString) ([]*DjmdSongPlaylist, error) {
	// func DjmdSongPlaylistByContentID(ctx context.Context, db DB, contentID sql.NullString) ([]*DjmdSongPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, PlaylistID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongPlaylist ` +
		`WHERE ContentID = $1`
	// run
	logf(sqlstr, contentID)
	rows, err := db.QueryContext(ctx, sqlstr, contentID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongPlaylist
	for rows.Next() {
		dsp := DjmdSongPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dsp.ID, &dsp.PlaylistID, &dsp.ContentID, &dsp.TrackNo, &dsp.UUID, &dsp.RbDataStatus, &dsp.RbLocalDataStatus, &dsp.RbLocalDeleted, &dsp.RbLocalSynced, &dsp.Usn, &dsp.RbLocalUsn, &dsp.CreatedAt, &dsp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dsp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongPlaylistByContentIDRbLocalDeleted retrieves a row from 'djmdSongPlaylist' as a DjmdSongPlaylist.
//
// Generated from index 'djmd_song_playlist__content_i_d_rb_local_deleted'.
func (c *Client) DjmdSongPlaylistByContentIDRbLocalDeleted(ctx context.Context, contentID sql.NullString, rbLocalDeleted sql.NullInt64) ([]*DjmdSongPlaylist, error) {
	// func DjmdSongPlaylistByContentIDRbLocalDeleted(ctx context.Context, db DB, contentID sql.NullString, rbLocalDeleted sql.NullInt64) ([]*DjmdSongPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, PlaylistID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongPlaylist ` +
		`WHERE ContentID = $1 AND rb_local_deleted = $2`
	// run
	logf(sqlstr, contentID, rbLocalDeleted)
	rows, err := db.QueryContext(ctx, sqlstr, contentID, rbLocalDeleted)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongPlaylist
	for rows.Next() {
		dsp := DjmdSongPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dsp.ID, &dsp.PlaylistID, &dsp.ContentID, &dsp.TrackNo, &dsp.UUID, &dsp.RbDataStatus, &dsp.RbLocalDataStatus, &dsp.RbLocalDeleted, &dsp.RbLocalSynced, &dsp.Usn, &dsp.RbLocalUsn, &dsp.CreatedAt, &dsp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dsp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongPlaylistByPlaylistID retrieves a row from 'djmdSongPlaylist' as a DjmdSongPlaylist.
//
// Generated from index 'djmd_song_playlist__playlist_i_d'.
func (c *Client) DjmdSongPlaylistByPlaylistID(ctx context.Context, playlistID sql.NullString) ([]*DjmdSongPlaylist, error) {
	// func DjmdSongPlaylistByPlaylistID(ctx context.Context, db DB, playlistID sql.NullString) ([]*DjmdSongPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, PlaylistID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongPlaylist ` +
		`WHERE PlaylistID = $1`
	// run
	logf(sqlstr, playlistID)
	rows, err := db.QueryContext(ctx, sqlstr, playlistID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongPlaylist
	for rows.Next() {
		dsp := DjmdSongPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dsp.ID, &dsp.PlaylistID, &dsp.ContentID, &dsp.TrackNo, &dsp.UUID, &dsp.RbDataStatus, &dsp.RbLocalDataStatus, &dsp.RbLocalDeleted, &dsp.RbLocalSynced, &dsp.Usn, &dsp.RbLocalUsn, &dsp.CreatedAt, &dsp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dsp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongPlaylistByUUID retrieves a row from 'djmdSongPlaylist' as a DjmdSongPlaylist.
//
// Generated from index 'djmd_song_playlist__u_u_i_d'.
func (c *Client) DjmdSongPlaylistByUUID(ctx context.Context, uuid sql.NullString) ([]*DjmdSongPlaylist, error) {
	// func DjmdSongPlaylistByUUID(ctx context.Context, db DB, uuid sql.NullString) ([]*DjmdSongPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, PlaylistID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongPlaylist ` +
		`WHERE UUID = $1`
	// run
	logf(sqlstr, uuid)
	rows, err := db.QueryContext(ctx, sqlstr, uuid)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongPlaylist
	for rows.Next() {
		dsp := DjmdSongPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dsp.ID, &dsp.PlaylistID, &dsp.ContentID, &dsp.TrackNo, &dsp.UUID, &dsp.RbDataStatus, &dsp.RbLocalDataStatus, &dsp.RbLocalDeleted, &dsp.RbLocalSynced, &dsp.Usn, &dsp.RbLocalUsn, &dsp.CreatedAt, &dsp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dsp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongPlaylistByRbDataStatus retrieves a row from 'djmdSongPlaylist' as a DjmdSongPlaylist.
//
// Generated from index 'djmd_song_playlist_rb_data_status'.
func (c *Client) DjmdSongPlaylistByRbDataStatus(ctx context.Context, rbDataStatus sql.NullInt64) ([]*DjmdSongPlaylist, error) {
	// func DjmdSongPlaylistByRbDataStatus(ctx context.Context, db DB, rbDataStatus sql.NullInt64) ([]*DjmdSongPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, PlaylistID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongPlaylist ` +
		`WHERE rb_data_status = $1`
	// run
	logf(sqlstr, rbDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongPlaylist
	for rows.Next() {
		dsp := DjmdSongPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dsp.ID, &dsp.PlaylistID, &dsp.ContentID, &dsp.TrackNo, &dsp.UUID, &dsp.RbDataStatus, &dsp.RbLocalDataStatus, &dsp.RbLocalDeleted, &dsp.RbLocalSynced, &dsp.Usn, &dsp.RbLocalUsn, &dsp.CreatedAt, &dsp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dsp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongPlaylistByRbLocalDataStatus retrieves a row from 'djmdSongPlaylist' as a DjmdSongPlaylist.
//
// Generated from index 'djmd_song_playlist_rb_local_data_status'.
func (c *Client) DjmdSongPlaylistByRbLocalDataStatus(ctx context.Context, rbLocalDataStatus sql.NullInt64) ([]*DjmdSongPlaylist, error) {
	// func DjmdSongPlaylistByRbLocalDataStatus(ctx context.Context, db DB, rbLocalDataStatus sql.NullInt64) ([]*DjmdSongPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, PlaylistID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongPlaylist ` +
		`WHERE rb_local_data_status = $1`
	// run
	logf(sqlstr, rbLocalDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongPlaylist
	for rows.Next() {
		dsp := DjmdSongPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dsp.ID, &dsp.PlaylistID, &dsp.ContentID, &dsp.TrackNo, &dsp.UUID, &dsp.RbDataStatus, &dsp.RbLocalDataStatus, &dsp.RbLocalDeleted, &dsp.RbLocalSynced, &dsp.Usn, &dsp.RbLocalUsn, &dsp.CreatedAt, &dsp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dsp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongPlaylistByRbLocalDeleted retrieves a row from 'djmdSongPlaylist' as a DjmdSongPlaylist.
//
// Generated from index 'djmd_song_playlist_rb_local_deleted'.
func (c *Client) DjmdSongPlaylistByRbLocalDeleted(ctx context.Context, rbLocalDeleted sql.NullInt64) ([]*DjmdSongPlaylist, error) {
	// func DjmdSongPlaylistByRbLocalDeleted(ctx context.Context, db DB, rbLocalDeleted sql.NullInt64) ([]*DjmdSongPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, PlaylistID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongPlaylist ` +
		`WHERE rb_local_deleted = $1`
	// run
	logf(sqlstr, rbLocalDeleted)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDeleted)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongPlaylist
	for rows.Next() {
		dsp := DjmdSongPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dsp.ID, &dsp.PlaylistID, &dsp.ContentID, &dsp.TrackNo, &dsp.UUID, &dsp.RbDataStatus, &dsp.RbLocalDataStatus, &dsp.RbLocalDeleted, &dsp.RbLocalSynced, &dsp.Usn, &dsp.RbLocalUsn, &dsp.CreatedAt, &dsp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dsp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongPlaylistByRbLocalUsnID retrieves a row from 'djmdSongPlaylist' as a DjmdSongPlaylist.
//
// Generated from index 'djmd_song_playlist_rb_local_usn__i_d'.
func (c *Client) DjmdSongPlaylistByRbLocalUsnID(ctx context.Context, rbLocalUsn sql.NullInt64, id sql.NullString) ([]*DjmdSongPlaylist, error) {
	// func DjmdSongPlaylistByRbLocalUsnID(ctx context.Context, db DB, rbLocalUsn sql.NullInt64, id sql.NullString) ([]*DjmdSongPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, PlaylistID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongPlaylist ` +
		`WHERE rb_local_usn = $1 AND ID = $2`
	// run
	logf(sqlstr, rbLocalUsn, id)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalUsn, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongPlaylist
	for rows.Next() {
		dsp := DjmdSongPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dsp.ID, &dsp.PlaylistID, &dsp.ContentID, &dsp.TrackNo, &dsp.UUID, &dsp.RbDataStatus, &dsp.RbLocalDataStatus, &dsp.RbLocalDeleted, &dsp.RbLocalSynced, &dsp.Usn, &dsp.RbLocalUsn, &dsp.CreatedAt, &dsp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dsp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongPlaylistByID retrieves a row from 'djmdSongPlaylist' as a DjmdSongPlaylist.
//
// Generated from index 'sqlite_autoindex_djmdSongPlaylist_1'.
func (c *Client) DjmdSongPlaylistByID(ctx context.Context, id sql.NullString) (*DjmdSongPlaylist, error) {
	// func DjmdSongPlaylistByID(ctx context.Context, db DB, id sql.NullString) (*DjmdSongPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, PlaylistID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongPlaylist ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, id)
	dsp := DjmdSongPlaylist{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&dsp.ID, &dsp.PlaylistID, &dsp.ContentID, &dsp.TrackNo, &dsp.UUID, &dsp.RbDataStatus, &dsp.RbLocalDataStatus, &dsp.RbLocalDeleted, &dsp.RbLocalSynced, &dsp.Usn, &dsp.RbLocalUsn, &dsp.CreatedAt, &dsp.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &dsp, nil
}
