package rekordbox

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// DjmdSongRelatedTrack represents a row from 'djmdSongRelatedTracks'.
type DjmdSongRelatedTrack struct {
	ID                sql.NullString `json:"ID"`                   // ID
	RelatedTracksID   sql.NullString `json:"RelatedTracksID"`      // RelatedTracksID
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

// Exists returns true when the DjmdSongRelatedTrack exists in the database.
func (dsrt *DjmdSongRelatedTrack) Exists() bool {
	return dsrt._exists
}

// Deleted returns true when the DjmdSongRelatedTrack has been marked for deletion from
// the database.
func (dsrt *DjmdSongRelatedTrack) Deleted() bool {
	return dsrt._deleted
}

// Insert inserts the DjmdSongRelatedTrack to the database.
func (c *Client) InsertDjmdSongRelatedTrack(ctx context.Context, dsrt *DjmdSongRelatedTrack) error {
	db := c.db

	switch {
	case dsrt._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case dsrt._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO djmdSongRelatedTracks (` +
		`ID, RelatedTracksID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`)`
	// run
	logf(sqlstr, dsrt.ID, dsrt.RelatedTracksID, dsrt.ContentID, dsrt.TrackNo, dsrt.UUID, dsrt.RbDataStatus, dsrt.RbLocalDataStatus, dsrt.RbLocalDeleted, dsrt.RbLocalSynced, dsrt.Usn, dsrt.RbLocalUsn, dsrt.CreatedAt, dsrt.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dsrt.ID, dsrt.RelatedTracksID, dsrt.ContentID, dsrt.TrackNo, dsrt.UUID, dsrt.RbDataStatus, dsrt.RbLocalDataStatus, dsrt.RbLocalDeleted, dsrt.RbLocalSynced, dsrt.Usn, dsrt.RbLocalUsn, dsrt.CreatedAt, dsrt.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dsrt._exists = true
	return nil
}

// Update updates a DjmdSongRelatedTrack in the database.
func (c *Client) UpdateDjmdSongRelatedTrack(ctx context.Context, dsrt *DjmdSongRelatedTrack) error {
	db := c.db

	switch {
	case !dsrt._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case dsrt._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE djmdSongRelatedTracks SET ` +
		`RelatedTracksID = $1, ContentID = $2, TrackNo = $3, UUID = $4, rb_data_status = $5, rb_local_data_status = $6, rb_local_deleted = $7, rb_local_synced = $8, usn = $9, rb_local_usn = $10, created_at = $11, updated_at = $12 ` +
		`WHERE ID = $13`
	// run
	logf(sqlstr, dsrt.RelatedTracksID, dsrt.ContentID, dsrt.TrackNo, dsrt.UUID, dsrt.RbDataStatus, dsrt.RbLocalDataStatus, dsrt.RbLocalDeleted, dsrt.RbLocalSynced, dsrt.Usn, dsrt.RbLocalUsn, dsrt.CreatedAt, dsrt.UpdatedAt, dsrt.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dsrt.RelatedTracksID, dsrt.ContentID, dsrt.TrackNo, dsrt.UUID, dsrt.RbDataStatus, dsrt.RbLocalDataStatus, dsrt.RbLocalDeleted, dsrt.RbLocalSynced, dsrt.Usn, dsrt.RbLocalUsn, dsrt.CreatedAt, dsrt.UpdatedAt, dsrt.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the DjmdSongRelatedTrack to the database.
func (c *Client) SaveDjmdSongRelatedTrack(ctx context.Context, dsrt *DjmdSongRelatedTrack) error {
	if dsrt.Exists() {
		return c.UpdateDjmdSongRelatedTrack(ctx, dsrt)
	}
	return c.InsertDjmdSongRelatedTrack(ctx, dsrt)
}

// Upsert performs an upsert for DjmdSongRelatedTrack.
func (c *Client) UpsertDjmdSongRelatedTrack(ctx context.Context, dsrt *DjmdSongRelatedTrack) error {
	db := c.db

	switch {
	case dsrt._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO djmdSongRelatedTracks (` +
		`ID, RelatedTracksID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`)` +
		` ON CONFLICT (ID) DO ` +
		`UPDATE SET ` +
		`RelatedTracksID = EXCLUDED.RelatedTracksID, ContentID = EXCLUDED.ContentID, TrackNo = EXCLUDED.TrackNo, UUID = EXCLUDED.UUID, rb_data_status = EXCLUDED.rb_data_status, rb_local_data_status = EXCLUDED.rb_local_data_status, rb_local_deleted = EXCLUDED.rb_local_deleted, rb_local_synced = EXCLUDED.rb_local_synced, usn = EXCLUDED.usn, rb_local_usn = EXCLUDED.rb_local_usn, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, dsrt.ID, dsrt.RelatedTracksID, dsrt.ContentID, dsrt.TrackNo, dsrt.UUID, dsrt.RbDataStatus, dsrt.RbLocalDataStatus, dsrt.RbLocalDeleted, dsrt.RbLocalSynced, dsrt.Usn, dsrt.RbLocalUsn, dsrt.CreatedAt, dsrt.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dsrt.ID, dsrt.RelatedTracksID, dsrt.ContentID, dsrt.TrackNo, dsrt.UUID, dsrt.RbDataStatus, dsrt.RbLocalDataStatus, dsrt.RbLocalDeleted, dsrt.RbLocalSynced, dsrt.Usn, dsrt.RbLocalUsn, dsrt.CreatedAt, dsrt.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dsrt._exists = true
	return nil
}

// Delete deletes the DjmdSongRelatedTrack from the database.
func (c *Client) DeleteDjmdSongRelatedTrack(ctx context.Context, dsrt *DjmdSongRelatedTrack) error {
	db := c.db

	switch {
	case !dsrt._exists: // doesn't exist
		return nil
	case dsrt._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM djmdSongRelatedTracks ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, dsrt.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dsrt.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	dsrt._deleted = true
	return nil
}

func scanDjmdSongRelatedTrackRows(rows *sql.Rows) ([]*DjmdSongRelatedTrack, error) {
	var res []*DjmdSongRelatedTrack
	for rows.Next() {
		dsrt := DjmdSongRelatedTrack{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dsrt.ID, &dsrt.RelatedTracksID, &dsrt.ContentID, &dsrt.TrackNo, &dsrt.UUID, &dsrt.RbDataStatus, &dsrt.RbLocalDataStatus, &dsrt.RbLocalDeleted, &dsrt.RbLocalSynced, &dsrt.Usn, &dsrt.RbLocalUsn, &dsrt.CreatedAt, &dsrt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dsrt)
	}

	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

func (c *Client) AllDjmdSongRelatedTrack(ctx context.Context) ([]*DjmdSongRelatedTrack, error) {
	db := c.db

	const sqlstr = `SELECT * FROM DjmdSongRelatedTrack`
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()
	res, err := scanDjmdSongRelatedTrackRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongRelatedTracksByContentID retrieves a row from 'djmdSongRelatedTracks' as a DjmdSongRelatedTrack.
//
// Generated from index 'djmd_song_related_tracks__content_i_d'.
func (c *Client) DjmdSongRelatedTracksByContentID(ctx context.Context, contentID sql.NullString) ([]*DjmdSongRelatedTrack, error) {
	// func DjmdSongRelatedTracksByContentID(ctx context.Context, db DB, contentID sql.NullString) ([]*DjmdSongRelatedTrack, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, RelatedTracksID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongRelatedTracks ` +
		`WHERE ContentID = $1`
	// run
	logf(sqlstr, contentID)
	rows, err := db.QueryContext(ctx, sqlstr, contentID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongRelatedTrack
	for rows.Next() {
		dsrt := DjmdSongRelatedTrack{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dsrt.ID, &dsrt.RelatedTracksID, &dsrt.ContentID, &dsrt.TrackNo, &dsrt.UUID, &dsrt.RbDataStatus, &dsrt.RbLocalDataStatus, &dsrt.RbLocalDeleted, &dsrt.RbLocalSynced, &dsrt.Usn, &dsrt.RbLocalUsn, &dsrt.CreatedAt, &dsrt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dsrt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongRelatedTracksByRelatedTracksID retrieves a row from 'djmdSongRelatedTracks' as a DjmdSongRelatedTrack.
//
// Generated from index 'djmd_song_related_tracks__related_tracks_i_d'.
func (c *Client) DjmdSongRelatedTracksByRelatedTracksID(ctx context.Context, relatedTracksID sql.NullString) ([]*DjmdSongRelatedTrack, error) {
	// func DjmdSongRelatedTracksByRelatedTracksID(ctx context.Context, db DB, relatedTracksID sql.NullString) ([]*DjmdSongRelatedTrack, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, RelatedTracksID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongRelatedTracks ` +
		`WHERE RelatedTracksID = $1`
	// run
	logf(sqlstr, relatedTracksID)
	rows, err := db.QueryContext(ctx, sqlstr, relatedTracksID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongRelatedTrack
	for rows.Next() {
		dsrt := DjmdSongRelatedTrack{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dsrt.ID, &dsrt.RelatedTracksID, &dsrt.ContentID, &dsrt.TrackNo, &dsrt.UUID, &dsrt.RbDataStatus, &dsrt.RbLocalDataStatus, &dsrt.RbLocalDeleted, &dsrt.RbLocalSynced, &dsrt.Usn, &dsrt.RbLocalUsn, &dsrt.CreatedAt, &dsrt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dsrt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongRelatedTracksByUUID retrieves a row from 'djmdSongRelatedTracks' as a DjmdSongRelatedTrack.
//
// Generated from index 'djmd_song_related_tracks__u_u_i_d'.
func (c *Client) DjmdSongRelatedTracksByUUID(ctx context.Context, uuid sql.NullString) ([]*DjmdSongRelatedTrack, error) {
	// func DjmdSongRelatedTracksByUUID(ctx context.Context, db DB, uuid sql.NullString) ([]*DjmdSongRelatedTrack, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, RelatedTracksID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongRelatedTracks ` +
		`WHERE UUID = $1`
	// run
	logf(sqlstr, uuid)
	rows, err := db.QueryContext(ctx, sqlstr, uuid)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongRelatedTrack
	for rows.Next() {
		dsrt := DjmdSongRelatedTrack{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dsrt.ID, &dsrt.RelatedTracksID, &dsrt.ContentID, &dsrt.TrackNo, &dsrt.UUID, &dsrt.RbDataStatus, &dsrt.RbLocalDataStatus, &dsrt.RbLocalDeleted, &dsrt.RbLocalSynced, &dsrt.Usn, &dsrt.RbLocalUsn, &dsrt.CreatedAt, &dsrt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dsrt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongRelatedTracksByRbDataStatus retrieves a row from 'djmdSongRelatedTracks' as a DjmdSongRelatedTrack.
//
// Generated from index 'djmd_song_related_tracks_rb_data_status'.
func (c *Client) DjmdSongRelatedTracksByRbDataStatus(ctx context.Context, rbDataStatus sql.NullInt64) ([]*DjmdSongRelatedTrack, error) {
	// func DjmdSongRelatedTracksByRbDataStatus(ctx context.Context, db DB, rbDataStatus sql.NullInt64) ([]*DjmdSongRelatedTrack, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, RelatedTracksID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongRelatedTracks ` +
		`WHERE rb_data_status = $1`
	// run
	logf(sqlstr, rbDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongRelatedTrack
	for rows.Next() {
		dsrt := DjmdSongRelatedTrack{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dsrt.ID, &dsrt.RelatedTracksID, &dsrt.ContentID, &dsrt.TrackNo, &dsrt.UUID, &dsrt.RbDataStatus, &dsrt.RbLocalDataStatus, &dsrt.RbLocalDeleted, &dsrt.RbLocalSynced, &dsrt.Usn, &dsrt.RbLocalUsn, &dsrt.CreatedAt, &dsrt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dsrt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongRelatedTracksByRbLocalDataStatus retrieves a row from 'djmdSongRelatedTracks' as a DjmdSongRelatedTrack.
//
// Generated from index 'djmd_song_related_tracks_rb_local_data_status'.
func (c *Client) DjmdSongRelatedTracksByRbLocalDataStatus(ctx context.Context, rbLocalDataStatus sql.NullInt64) ([]*DjmdSongRelatedTrack, error) {
	// func DjmdSongRelatedTracksByRbLocalDataStatus(ctx context.Context, db DB, rbLocalDataStatus sql.NullInt64) ([]*DjmdSongRelatedTrack, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, RelatedTracksID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongRelatedTracks ` +
		`WHERE rb_local_data_status = $1`
	// run
	logf(sqlstr, rbLocalDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongRelatedTrack
	for rows.Next() {
		dsrt := DjmdSongRelatedTrack{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dsrt.ID, &dsrt.RelatedTracksID, &dsrt.ContentID, &dsrt.TrackNo, &dsrt.UUID, &dsrt.RbDataStatus, &dsrt.RbLocalDataStatus, &dsrt.RbLocalDeleted, &dsrt.RbLocalSynced, &dsrt.Usn, &dsrt.RbLocalUsn, &dsrt.CreatedAt, &dsrt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dsrt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongRelatedTracksByRbLocalDeleted retrieves a row from 'djmdSongRelatedTracks' as a DjmdSongRelatedTrack.
//
// Generated from index 'djmd_song_related_tracks_rb_local_deleted'.
func (c *Client) DjmdSongRelatedTracksByRbLocalDeleted(ctx context.Context, rbLocalDeleted sql.NullInt64) ([]*DjmdSongRelatedTrack, error) {
	// func DjmdSongRelatedTracksByRbLocalDeleted(ctx context.Context, db DB, rbLocalDeleted sql.NullInt64) ([]*DjmdSongRelatedTrack, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, RelatedTracksID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongRelatedTracks ` +
		`WHERE rb_local_deleted = $1`
	// run
	logf(sqlstr, rbLocalDeleted)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDeleted)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongRelatedTrack
	for rows.Next() {
		dsrt := DjmdSongRelatedTrack{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dsrt.ID, &dsrt.RelatedTracksID, &dsrt.ContentID, &dsrt.TrackNo, &dsrt.UUID, &dsrt.RbDataStatus, &dsrt.RbLocalDataStatus, &dsrt.RbLocalDeleted, &dsrt.RbLocalSynced, &dsrt.Usn, &dsrt.RbLocalUsn, &dsrt.CreatedAt, &dsrt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dsrt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongRelatedTracksByRbLocalUsnID retrieves a row from 'djmdSongRelatedTracks' as a DjmdSongRelatedTrack.
//
// Generated from index 'djmd_song_related_tracks_rb_local_usn__i_d'.
func (c *Client) DjmdSongRelatedTracksByRbLocalUsnID(ctx context.Context, rbLocalUsn sql.NullInt64, id sql.NullString) ([]*DjmdSongRelatedTrack, error) {
	// func DjmdSongRelatedTracksByRbLocalUsnID(ctx context.Context, db DB, rbLocalUsn sql.NullInt64, id sql.NullString) ([]*DjmdSongRelatedTrack, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, RelatedTracksID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongRelatedTracks ` +
		`WHERE rb_local_usn = $1 AND ID = $2`
	// run
	logf(sqlstr, rbLocalUsn, id)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalUsn, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongRelatedTrack
	for rows.Next() {
		dsrt := DjmdSongRelatedTrack{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dsrt.ID, &dsrt.RelatedTracksID, &dsrt.ContentID, &dsrt.TrackNo, &dsrt.UUID, &dsrt.RbDataStatus, &dsrt.RbLocalDataStatus, &dsrt.RbLocalDeleted, &dsrt.RbLocalSynced, &dsrt.Usn, &dsrt.RbLocalUsn, &dsrt.CreatedAt, &dsrt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dsrt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongRelatedTrackByID retrieves a row from 'djmdSongRelatedTracks' as a DjmdSongRelatedTrack.
//
// Generated from index 'sqlite_autoindex_djmdSongRelatedTracks_1'.
func (c *Client) DjmdSongRelatedTrackByID(ctx context.Context, id sql.NullString) (*DjmdSongRelatedTrack, error) {
	// func DjmdSongRelatedTrackByID(ctx context.Context, db DB, id sql.NullString) (*DjmdSongRelatedTrack, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, RelatedTracksID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongRelatedTracks ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, id)
	dsrt := DjmdSongRelatedTrack{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&dsrt.ID, &dsrt.RelatedTracksID, &dsrt.ContentID, &dsrt.TrackNo, &dsrt.UUID, &dsrt.RbDataStatus, &dsrt.RbLocalDataStatus, &dsrt.RbLocalDeleted, &dsrt.RbLocalSynced, &dsrt.Usn, &dsrt.RbLocalUsn, &dsrt.CreatedAt, &dsrt.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &dsrt, nil
}
