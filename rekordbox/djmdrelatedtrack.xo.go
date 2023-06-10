package rekordbox

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	nulltype "github.com/mattn/go-nulltype"
)

// DjmdRelatedTrack represents a row from 'djmdRelatedTracks'.
type DjmdRelatedTrack struct {
	ID                nulltype.NullString `json:"id"`                   // ID
	Seq               nulltype.NullInt64  `json:"seq"`                  // Seq
	Name              nulltype.NullString `json:"name"`                 // Name
	Attribute         nulltype.NullInt64  `json:"attribute"`            // Attribute
	ParentID          nulltype.NullString `json:"parent_id"`            // ParentID
	Criteria          nulltype.NullString `json:"criteria"`             // Criteria
	UUID              nulltype.NullString `json:"uuid"`                 // UUID
	RbDataStatus      nulltype.NullInt64  `json:"rb_data_status"`       // rb_data_status
	RbLocalDataStatus nulltype.NullInt64  `json:"rb_local_data_status"` // rb_local_data_status
	RbLocalDeleted    nulltype.NullInt64  `json:"rb_local_deleted"`     // rb_local_deleted
	RbLocalSynced     nulltype.NullInt64  `json:"rb_local_synced"`      // rb_local_synced
	Usn               nulltype.NullInt64  `json:"usn"`                  // usn
	RbLocalUsn        nulltype.NullInt64  `json:"rb_local_usn"`         // rb_local_usn
	CreatedAt         Time                `json:"created_at"`           // created_at
	UpdatedAt         Time                `json:"updated_at"`           // updated_at
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the DjmdRelatedTrack exists in the database.
func (drt *DjmdRelatedTrack) Exists() bool {
	return drt._exists
}

// Deleted returns true when the DjmdRelatedTrack has been marked for deletion from
// the database.
func (drt *DjmdRelatedTrack) Deleted() bool {
	return drt._deleted
}

// Insert inserts the DjmdRelatedTrack to the database.
func (c *Client) InsertDjmdRelatedTrack(ctx context.Context, drt *DjmdRelatedTrack) error {
	db := c.db

	switch {
	case drt._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case drt._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO djmdRelatedTracks (` +
		`ID, Seq, Name, Attribute, ParentID, Criteria, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15` +
		`)`
	// run
	logf(sqlstr, drt.ID, drt.Seq, drt.Name, drt.Attribute, drt.ParentID, drt.Criteria, drt.UUID, drt.RbDataStatus, drt.RbLocalDataStatus, drt.RbLocalDeleted, drt.RbLocalSynced, drt.Usn, drt.RbLocalUsn, drt.CreatedAt, drt.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, drt.ID, drt.Seq, drt.Name, drt.Attribute, drt.ParentID, drt.Criteria, drt.UUID, drt.RbDataStatus, drt.RbLocalDataStatus, drt.RbLocalDeleted, drt.RbLocalSynced, drt.Usn, drt.RbLocalUsn, drt.CreatedAt, drt.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	drt._exists = true
	return nil
}

// Update updates a DjmdRelatedTrack in the database.
func (c *Client) UpdateDjmdRelatedTrack(ctx context.Context, drt *DjmdRelatedTrack) error {
	db := c.db

	switch {
	case !drt._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case drt._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE djmdRelatedTracks SET ` +
		`Seq = $1, Name = $2, Attribute = $3, ParentID = $4, Criteria = $5, UUID = $6, rb_data_status = $7, rb_local_data_status = $8, rb_local_deleted = $9, rb_local_synced = $10, usn = $11, rb_local_usn = $12, created_at = $13, updated_at = $14 ` +
		`WHERE ID = $15`
	// run
	logf(sqlstr, drt.Seq, drt.Name, drt.Attribute, drt.ParentID, drt.Criteria, drt.UUID, drt.RbDataStatus, drt.RbLocalDataStatus, drt.RbLocalDeleted, drt.RbLocalSynced, drt.Usn, drt.RbLocalUsn, drt.CreatedAt, drt.UpdatedAt, drt.ID)
	if _, err := db.ExecContext(ctx, sqlstr, drt.Seq, drt.Name, drt.Attribute, drt.ParentID, drt.Criteria, drt.UUID, drt.RbDataStatus, drt.RbLocalDataStatus, drt.RbLocalDeleted, drt.RbLocalSynced, drt.Usn, drt.RbLocalUsn, drt.CreatedAt, drt.UpdatedAt, drt.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the DjmdRelatedTrack to the database.
func (c *Client) SaveDjmdRelatedTrack(ctx context.Context, drt *DjmdRelatedTrack) error {
	if drt.Exists() {
		return c.UpdateDjmdRelatedTrack(ctx, drt)
	}
	return c.InsertDjmdRelatedTrack(ctx, drt)
}

// Upsert performs an upsert for DjmdRelatedTrack.
func (c *Client) UpsertDjmdRelatedTrack(ctx context.Context, drt *DjmdRelatedTrack) error {
	db := c.db

	switch {
	case drt._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO djmdRelatedTracks (` +
		`ID, Seq, Name, Attribute, ParentID, Criteria, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15` +
		`)` +
		` ON CONFLICT (ID) DO ` +
		`UPDATE SET ` +
		`Seq = EXCLUDED.Seq, Name = EXCLUDED.Name, Attribute = EXCLUDED.Attribute, ParentID = EXCLUDED.ParentID, Criteria = EXCLUDED.Criteria, UUID = EXCLUDED.UUID, rb_data_status = EXCLUDED.rb_data_status, rb_local_data_status = EXCLUDED.rb_local_data_status, rb_local_deleted = EXCLUDED.rb_local_deleted, rb_local_synced = EXCLUDED.rb_local_synced, usn = EXCLUDED.usn, rb_local_usn = EXCLUDED.rb_local_usn, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, drt.ID, drt.Seq, drt.Name, drt.Attribute, drt.ParentID, drt.Criteria, drt.UUID, drt.RbDataStatus, drt.RbLocalDataStatus, drt.RbLocalDeleted, drt.RbLocalSynced, drt.Usn, drt.RbLocalUsn, drt.CreatedAt, drt.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, drt.ID, drt.Seq, drt.Name, drt.Attribute, drt.ParentID, drt.Criteria, drt.UUID, drt.RbDataStatus, drt.RbLocalDataStatus, drt.RbLocalDeleted, drt.RbLocalSynced, drt.Usn, drt.RbLocalUsn, drt.CreatedAt, drt.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	drt._exists = true
	return nil
}

// Delete deletes the DjmdRelatedTrack from the database.
func (c *Client) DeleteDjmdRelatedTrack(ctx context.Context, drt *DjmdRelatedTrack) error {
	db := c.db

	switch {
	case !drt._exists: // doesn't exist
		return nil
	case drt._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM djmdRelatedTracks ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, drt.ID)
	if _, err := db.ExecContext(ctx, sqlstr, drt.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	drt._deleted = true
	return nil
}

func scanDjmdRelatedTrackRows(rows *sql.Rows) ([]*DjmdRelatedTrack, error) {
	var res []*DjmdRelatedTrack
	for rows.Next() {
		drt := DjmdRelatedTrack{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&drt.ID, &drt.Seq, &drt.Name, &drt.Attribute, &drt.ParentID, &drt.Criteria, &drt.UUID, &drt.RbDataStatus, &drt.RbLocalDataStatus, &drt.RbLocalDeleted, &drt.RbLocalSynced, &drt.Usn, &drt.RbLocalUsn, &drt.CreatedAt, &drt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &drt)
	}

	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

func (c *Client) AllDjmdRelatedTrack(ctx context.Context) ([]*DjmdRelatedTrack, error) {
	db := c.db

	const sqlstr = `SELECT * FROM DjmdRelatedTrack`
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()
	res, err := scanDjmdRelatedTrackRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdRelatedTracksByName retrieves a row from 'djmdRelatedTracks' as a DjmdRelatedTrack.
//
// Generated from index 'djmd_related_tracks__name'.
func (c *Client) DjmdRelatedTracksByName(ctx context.Context, name nulltype.NullString) ([]*DjmdRelatedTrack, error) {
	// func DjmdRelatedTracksByName(ctx context.Context, db DB, name nulltype.NullString) ([]*DjmdRelatedTrack, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, Criteria, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdRelatedTracks ` +
		`WHERE Name = $1`
	// run
	logf(sqlstr, name)
	rows, err := db.QueryContext(ctx, sqlstr, name)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdRelatedTrack
	for rows.Next() {
		drt := DjmdRelatedTrack{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&drt.ID, &drt.Seq, &drt.Name, &drt.Attribute, &drt.ParentID, &drt.Criteria, &drt.UUID, &drt.RbDataStatus, &drt.RbLocalDataStatus, &drt.RbLocalDeleted, &drt.RbLocalSynced, &drt.Usn, &drt.RbLocalUsn, &drt.CreatedAt, &drt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &drt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdRelatedTracksByParentID retrieves a row from 'djmdRelatedTracks' as a DjmdRelatedTrack.
//
// Generated from index 'djmd_related_tracks__parent_i_d'.
func (c *Client) DjmdRelatedTracksByParentID(ctx context.Context, parentID nulltype.NullString) ([]*DjmdRelatedTrack, error) {
	// func DjmdRelatedTracksByParentID(ctx context.Context, db DB, parentID nulltype.NullString) ([]*DjmdRelatedTrack, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, Criteria, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdRelatedTracks ` +
		`WHERE ParentID = $1`
	// run
	logf(sqlstr, parentID)
	rows, err := db.QueryContext(ctx, sqlstr, parentID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdRelatedTrack
	for rows.Next() {
		drt := DjmdRelatedTrack{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&drt.ID, &drt.Seq, &drt.Name, &drt.Attribute, &drt.ParentID, &drt.Criteria, &drt.UUID, &drt.RbDataStatus, &drt.RbLocalDataStatus, &drt.RbLocalDeleted, &drt.RbLocalSynced, &drt.Usn, &drt.RbLocalUsn, &drt.CreatedAt, &drt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &drt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdRelatedTracksBySeq retrieves a row from 'djmdRelatedTracks' as a DjmdRelatedTrack.
//
// Generated from index 'djmd_related_tracks__seq'.
func (c *Client) DjmdRelatedTracksBySeq(ctx context.Context, seq nulltype.NullInt64) ([]*DjmdRelatedTrack, error) {
	// func DjmdRelatedTracksBySeq(ctx context.Context, db DB, seq nulltype.NullInt64) ([]*DjmdRelatedTrack, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, Criteria, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdRelatedTracks ` +
		`WHERE Seq = $1`
	// run
	logf(sqlstr, seq)
	rows, err := db.QueryContext(ctx, sqlstr, seq)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdRelatedTrack
	for rows.Next() {
		drt := DjmdRelatedTrack{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&drt.ID, &drt.Seq, &drt.Name, &drt.Attribute, &drt.ParentID, &drt.Criteria, &drt.UUID, &drt.RbDataStatus, &drt.RbLocalDataStatus, &drt.RbLocalDeleted, &drt.RbLocalSynced, &drt.Usn, &drt.RbLocalUsn, &drt.CreatedAt, &drt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &drt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdRelatedTracksByUUID retrieves a row from 'djmdRelatedTracks' as a DjmdRelatedTrack.
//
// Generated from index 'djmd_related_tracks__u_u_i_d'.
func (c *Client) DjmdRelatedTracksByUUID(ctx context.Context, uuid nulltype.NullString) ([]*DjmdRelatedTrack, error) {
	// func DjmdRelatedTracksByUUID(ctx context.Context, db DB, uuid nulltype.NullString) ([]*DjmdRelatedTrack, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, Criteria, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdRelatedTracks ` +
		`WHERE UUID = $1`
	// run
	logf(sqlstr, uuid)
	rows, err := db.QueryContext(ctx, sqlstr, uuid)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdRelatedTrack
	for rows.Next() {
		drt := DjmdRelatedTrack{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&drt.ID, &drt.Seq, &drt.Name, &drt.Attribute, &drt.ParentID, &drt.Criteria, &drt.UUID, &drt.RbDataStatus, &drt.RbLocalDataStatus, &drt.RbLocalDeleted, &drt.RbLocalSynced, &drt.Usn, &drt.RbLocalUsn, &drt.CreatedAt, &drt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &drt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdRelatedTracksByRbDataStatus retrieves a row from 'djmdRelatedTracks' as a DjmdRelatedTrack.
//
// Generated from index 'djmd_related_tracks_rb_data_status'.
func (c *Client) DjmdRelatedTracksByRbDataStatus(ctx context.Context, rbDataStatus nulltype.NullInt64) ([]*DjmdRelatedTrack, error) {
	// func DjmdRelatedTracksByRbDataStatus(ctx context.Context, db DB, rbDataStatus nulltype.NullInt64) ([]*DjmdRelatedTrack, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, Criteria, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdRelatedTracks ` +
		`WHERE rb_data_status = $1`
	// run
	logf(sqlstr, rbDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdRelatedTrack
	for rows.Next() {
		drt := DjmdRelatedTrack{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&drt.ID, &drt.Seq, &drt.Name, &drt.Attribute, &drt.ParentID, &drt.Criteria, &drt.UUID, &drt.RbDataStatus, &drt.RbLocalDataStatus, &drt.RbLocalDeleted, &drt.RbLocalSynced, &drt.Usn, &drt.RbLocalUsn, &drt.CreatedAt, &drt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &drt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdRelatedTracksByRbLocalDataStatus retrieves a row from 'djmdRelatedTracks' as a DjmdRelatedTrack.
//
// Generated from index 'djmd_related_tracks_rb_local_data_status'.
func (c *Client) DjmdRelatedTracksByRbLocalDataStatus(ctx context.Context, rbLocalDataStatus nulltype.NullInt64) ([]*DjmdRelatedTrack, error) {
	// func DjmdRelatedTracksByRbLocalDataStatus(ctx context.Context, db DB, rbLocalDataStatus nulltype.NullInt64) ([]*DjmdRelatedTrack, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, Criteria, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdRelatedTracks ` +
		`WHERE rb_local_data_status = $1`
	// run
	logf(sqlstr, rbLocalDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdRelatedTrack
	for rows.Next() {
		drt := DjmdRelatedTrack{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&drt.ID, &drt.Seq, &drt.Name, &drt.Attribute, &drt.ParentID, &drt.Criteria, &drt.UUID, &drt.RbDataStatus, &drt.RbLocalDataStatus, &drt.RbLocalDeleted, &drt.RbLocalSynced, &drt.Usn, &drt.RbLocalUsn, &drt.CreatedAt, &drt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &drt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdRelatedTracksByRbLocalDeleted retrieves a row from 'djmdRelatedTracks' as a DjmdRelatedTrack.
//
// Generated from index 'djmd_related_tracks_rb_local_deleted'.
func (c *Client) DjmdRelatedTracksByRbLocalDeleted(ctx context.Context, rbLocalDeleted nulltype.NullInt64) ([]*DjmdRelatedTrack, error) {
	// func DjmdRelatedTracksByRbLocalDeleted(ctx context.Context, db DB, rbLocalDeleted nulltype.NullInt64) ([]*DjmdRelatedTrack, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, Criteria, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdRelatedTracks ` +
		`WHERE rb_local_deleted = $1`
	// run
	logf(sqlstr, rbLocalDeleted)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDeleted)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdRelatedTrack
	for rows.Next() {
		drt := DjmdRelatedTrack{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&drt.ID, &drt.Seq, &drt.Name, &drt.Attribute, &drt.ParentID, &drt.Criteria, &drt.UUID, &drt.RbDataStatus, &drt.RbLocalDataStatus, &drt.RbLocalDeleted, &drt.RbLocalSynced, &drt.Usn, &drt.RbLocalUsn, &drt.CreatedAt, &drt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &drt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdRelatedTracksByRbLocalUsnID retrieves a row from 'djmdRelatedTracks' as a DjmdRelatedTrack.
//
// Generated from index 'djmd_related_tracks_rb_local_usn__i_d'.
func (c *Client) DjmdRelatedTracksByRbLocalUsnID(ctx context.Context, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*DjmdRelatedTrack, error) {
	// func DjmdRelatedTracksByRbLocalUsnID(ctx context.Context, db DB, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*DjmdRelatedTrack, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, Criteria, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdRelatedTracks ` +
		`WHERE rb_local_usn = $1 AND ID = $2`
	// run
	logf(sqlstr, rbLocalUsn, id)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalUsn, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdRelatedTrack
	for rows.Next() {
		drt := DjmdRelatedTrack{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&drt.ID, &drt.Seq, &drt.Name, &drt.Attribute, &drt.ParentID, &drt.Criteria, &drt.UUID, &drt.RbDataStatus, &drt.RbLocalDataStatus, &drt.RbLocalDeleted, &drt.RbLocalSynced, &drt.Usn, &drt.RbLocalUsn, &drt.CreatedAt, &drt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &drt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdRelatedTrackByID retrieves a row from 'djmdRelatedTracks' as a DjmdRelatedTrack.
//
// Generated from index 'sqlite_autoindex_djmdRelatedTracks_1'.
func (c *Client) DjmdRelatedTrackByID(ctx context.Context, id nulltype.NullString) (*DjmdRelatedTrack, error) {
	// func DjmdRelatedTrackByID(ctx context.Context, db DB, id nulltype.NullString) (*DjmdRelatedTrack, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, Criteria, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdRelatedTracks ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, id)
	drt := DjmdRelatedTrack{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&drt.ID, &drt.Seq, &drt.Name, &drt.Attribute, &drt.ParentID, &drt.Criteria, &drt.UUID, &drt.RbDataStatus, &drt.RbLocalDataStatus, &drt.RbLocalDeleted, &drt.RbLocalSynced, &drt.Usn, &drt.RbLocalUsn, &drt.CreatedAt, &drt.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &drt, nil
}
