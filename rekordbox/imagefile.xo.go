package rekordbox

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	nulltype "github.com/mattn/go-nulltype"
)

// ImageFile represents a row from 'imageFile'.
type ImageFile struct {
	ID                nulltype.NullString `json:"id"`                   // ID
	TableName         nulltype.NullString `json:"table_name"`           // TableName
	TargetUUID        nulltype.NullString `json:"target_uuid"`          // TargetUUID
	TargetID          nulltype.NullString `json:"target_id"`            // TargetID
	Path              nulltype.NullString `json:"path"`                 // Path
	Hash              nulltype.NullString `json:"hash"`                 // Hash
	Size              nulltype.NullInt64  `json:"size"`                 // Size
	RbLocalPath       nulltype.NullString `json:"rb_local_path"`        // rb_local_path
	RbInsyncHash      nulltype.NullString `json:"rb_insync_hash"`       // rb_insync_hash
	RbInsyncLocalUsn  nulltype.NullInt64  `json:"rb_insync_local_usn"`  // rb_insync_local_usn
	RbFileHashDirty   nulltype.NullInt64  `json:"rb_file_hash_dirty"`   // rb_file_hash_dirty
	RbLocalFileStatus nulltype.NullInt64  `json:"rb_local_file_status"` // rb_local_file_status
	RbInProgress      nulltype.NullInt64  `json:"rb_in_progress"`       // rb_in_progress
	RbProcessType     nulltype.NullInt64  `json:"rb_process_type"`      // rb_process_type
	RbTempPath        nulltype.NullString `json:"rb_temp_path"`         // rb_temp_path
	RbPriority        nulltype.NullInt64  `json:"rb_priority"`          // rb_priority
	RbFileSizeDirty   nulltype.NullInt64  `json:"rb_file_size_dirty"`   // rb_file_size_dirty
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

// Exists returns true when the ImageFile exists in the database.
func (ifVal *ImageFile) Exists() bool {
	return ifVal._exists
}

// Deleted returns true when the ImageFile has been marked for deletion from
// the database.
func (ifVal *ImageFile) Deleted() bool {
	return ifVal._deleted
}

// Insert inserts the ImageFile to the database.
func (c *Client) InsertImageFile(ctx context.Context, ifVal *ImageFile) error {
	db := c.db

	switch {
	case ifVal._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case ifVal._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO imageFile (` +
		`ID, TableName, TargetUUID, TargetID, Path, Hash, Size, rb_local_path, rb_insync_hash, rb_insync_local_usn, rb_file_hash_dirty, rb_local_file_status, rb_in_progress, rb_process_type, rb_temp_path, rb_priority, rb_file_size_dirty, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26` +
		`)`
	// run
	logf(sqlstr, ifVal.ID, ifVal.TableName, ifVal.TargetUUID, ifVal.TargetID, ifVal.Path, ifVal.Hash, ifVal.Size, ifVal.RbLocalPath, ifVal.RbInsyncHash, ifVal.RbInsyncLocalUsn, ifVal.RbFileHashDirty, ifVal.RbLocalFileStatus, ifVal.RbInProgress, ifVal.RbProcessType, ifVal.RbTempPath, ifVal.RbPriority, ifVal.RbFileSizeDirty, ifVal.UUID, ifVal.RbDataStatus, ifVal.RbLocalDataStatus, ifVal.RbLocalDeleted, ifVal.RbLocalSynced, ifVal.Usn, ifVal.RbLocalUsn, ifVal.CreatedAt, ifVal.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, ifVal.ID, ifVal.TableName, ifVal.TargetUUID, ifVal.TargetID, ifVal.Path, ifVal.Hash, ifVal.Size, ifVal.RbLocalPath, ifVal.RbInsyncHash, ifVal.RbInsyncLocalUsn, ifVal.RbFileHashDirty, ifVal.RbLocalFileStatus, ifVal.RbInProgress, ifVal.RbProcessType, ifVal.RbTempPath, ifVal.RbPriority, ifVal.RbFileSizeDirty, ifVal.UUID, ifVal.RbDataStatus, ifVal.RbLocalDataStatus, ifVal.RbLocalDeleted, ifVal.RbLocalSynced, ifVal.Usn, ifVal.RbLocalUsn, ifVal.CreatedAt, ifVal.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	ifVal._exists = true
	return nil
}

// Update updates a ImageFile in the database.
func (c *Client) UpdateImageFile(ctx context.Context, ifVal *ImageFile) error {
	db := c.db

	switch {
	case !ifVal._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case ifVal._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE imageFile SET ` +
		`TableName = $1, TargetUUID = $2, TargetID = $3, Path = $4, Hash = $5, Size = $6, rb_local_path = $7, rb_insync_hash = $8, rb_insync_local_usn = $9, rb_file_hash_dirty = $10, rb_local_file_status = $11, rb_in_progress = $12, rb_process_type = $13, rb_temp_path = $14, rb_priority = $15, rb_file_size_dirty = $16, UUID = $17, rb_data_status = $18, rb_local_data_status = $19, rb_local_deleted = $20, rb_local_synced = $21, usn = $22, rb_local_usn = $23, created_at = $24, updated_at = $25 ` +
		`WHERE ID = $26`
	// run
	logf(sqlstr, ifVal.TableName, ifVal.TargetUUID, ifVal.TargetID, ifVal.Path, ifVal.Hash, ifVal.Size, ifVal.RbLocalPath, ifVal.RbInsyncHash, ifVal.RbInsyncLocalUsn, ifVal.RbFileHashDirty, ifVal.RbLocalFileStatus, ifVal.RbInProgress, ifVal.RbProcessType, ifVal.RbTempPath, ifVal.RbPriority, ifVal.RbFileSizeDirty, ifVal.UUID, ifVal.RbDataStatus, ifVal.RbLocalDataStatus, ifVal.RbLocalDeleted, ifVal.RbLocalSynced, ifVal.Usn, ifVal.RbLocalUsn, ifVal.CreatedAt, ifVal.UpdatedAt, ifVal.ID)
	if _, err := db.ExecContext(ctx, sqlstr, ifVal.TableName, ifVal.TargetUUID, ifVal.TargetID, ifVal.Path, ifVal.Hash, ifVal.Size, ifVal.RbLocalPath, ifVal.RbInsyncHash, ifVal.RbInsyncLocalUsn, ifVal.RbFileHashDirty, ifVal.RbLocalFileStatus, ifVal.RbInProgress, ifVal.RbProcessType, ifVal.RbTempPath, ifVal.RbPriority, ifVal.RbFileSizeDirty, ifVal.UUID, ifVal.RbDataStatus, ifVal.RbLocalDataStatus, ifVal.RbLocalDeleted, ifVal.RbLocalSynced, ifVal.Usn, ifVal.RbLocalUsn, ifVal.CreatedAt, ifVal.UpdatedAt, ifVal.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the ImageFile to the database.
func (c *Client) SaveImageFile(ctx context.Context, ifVal *ImageFile) error {
	if ifVal.Exists() {
		return c.UpdateImageFile(ctx, ifVal)
	}
	return c.InsertImageFile(ctx, ifVal)
}

// Upsert performs an upsert for ImageFile.
func (c *Client) UpsertImageFile(ctx context.Context, ifVal *ImageFile) error {
	db := c.db

	switch {
	case ifVal._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO imageFile (` +
		`ID, TableName, TargetUUID, TargetID, Path, Hash, Size, rb_local_path, rb_insync_hash, rb_insync_local_usn, rb_file_hash_dirty, rb_local_file_status, rb_in_progress, rb_process_type, rb_temp_path, rb_priority, rb_file_size_dirty, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26` +
		`)` +
		` ON CONFLICT (ID) DO ` +
		`UPDATE SET ` +
		`TableName = EXCLUDED.TableName, TargetUUID = EXCLUDED.TargetUUID, TargetID = EXCLUDED.TargetID, Path = EXCLUDED.Path, Hash = EXCLUDED.Hash, Size = EXCLUDED.Size, rb_local_path = EXCLUDED.rb_local_path, rb_insync_hash = EXCLUDED.rb_insync_hash, rb_insync_local_usn = EXCLUDED.rb_insync_local_usn, rb_file_hash_dirty = EXCLUDED.rb_file_hash_dirty, rb_local_file_status = EXCLUDED.rb_local_file_status, rb_in_progress = EXCLUDED.rb_in_progress, rb_process_type = EXCLUDED.rb_process_type, rb_temp_path = EXCLUDED.rb_temp_path, rb_priority = EXCLUDED.rb_priority, rb_file_size_dirty = EXCLUDED.rb_file_size_dirty, UUID = EXCLUDED.UUID, rb_data_status = EXCLUDED.rb_data_status, rb_local_data_status = EXCLUDED.rb_local_data_status, rb_local_deleted = EXCLUDED.rb_local_deleted, rb_local_synced = EXCLUDED.rb_local_synced, usn = EXCLUDED.usn, rb_local_usn = EXCLUDED.rb_local_usn, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, ifVal.ID, ifVal.TableName, ifVal.TargetUUID, ifVal.TargetID, ifVal.Path, ifVal.Hash, ifVal.Size, ifVal.RbLocalPath, ifVal.RbInsyncHash, ifVal.RbInsyncLocalUsn, ifVal.RbFileHashDirty, ifVal.RbLocalFileStatus, ifVal.RbInProgress, ifVal.RbProcessType, ifVal.RbTempPath, ifVal.RbPriority, ifVal.RbFileSizeDirty, ifVal.UUID, ifVal.RbDataStatus, ifVal.RbLocalDataStatus, ifVal.RbLocalDeleted, ifVal.RbLocalSynced, ifVal.Usn, ifVal.RbLocalUsn, ifVal.CreatedAt, ifVal.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, ifVal.ID, ifVal.TableName, ifVal.TargetUUID, ifVal.TargetID, ifVal.Path, ifVal.Hash, ifVal.Size, ifVal.RbLocalPath, ifVal.RbInsyncHash, ifVal.RbInsyncLocalUsn, ifVal.RbFileHashDirty, ifVal.RbLocalFileStatus, ifVal.RbInProgress, ifVal.RbProcessType, ifVal.RbTempPath, ifVal.RbPriority, ifVal.RbFileSizeDirty, ifVal.UUID, ifVal.RbDataStatus, ifVal.RbLocalDataStatus, ifVal.RbLocalDeleted, ifVal.RbLocalSynced, ifVal.Usn, ifVal.RbLocalUsn, ifVal.CreatedAt, ifVal.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	ifVal._exists = true
	return nil
}

// Delete deletes the ImageFile from the database.
func (c *Client) DeleteImageFile(ctx context.Context, ifVal *ImageFile) error {
	db := c.db

	switch {
	case !ifVal._exists: // doesn't exist
		return nil
	case ifVal._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM imageFile ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, ifVal.ID)
	if _, err := db.ExecContext(ctx, sqlstr, ifVal.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	ifVal._deleted = true
	return nil
}

func scanImageFileRows(rows *sql.Rows) ([]*ImageFile, error) {
	var res []*ImageFile
	for rows.Next() {
		ifVal := ImageFile{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ifVal.ID, &ifVal.TableName, &ifVal.TargetUUID, &ifVal.TargetID, &ifVal.Path, &ifVal.Hash, &ifVal.Size, &ifVal.RbLocalPath, &ifVal.RbInsyncHash, &ifVal.RbInsyncLocalUsn, &ifVal.RbFileHashDirty, &ifVal.RbLocalFileStatus, &ifVal.RbInProgress, &ifVal.RbProcessType, &ifVal.RbTempPath, &ifVal.RbPriority, &ifVal.RbFileSizeDirty, &ifVal.UUID, &ifVal.RbDataStatus, &ifVal.RbLocalDataStatus, &ifVal.RbLocalDeleted, &ifVal.RbLocalSynced, &ifVal.Usn, &ifVal.RbLocalUsn, &ifVal.CreatedAt, &ifVal.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ifVal)
	}

	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

func (c *Client) AllImageFile(ctx context.Context) ([]*ImageFile, error) {
	db := c.db

	const sqlstr = `SELECT * FROM ImageFile`
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()
	res, err := scanImageFileRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ImageFileByTableNameTargetID retrieves a row from 'imageFile' as a ImageFile.
//
// Generated from index 'image_file__table_name__target_i_d'.
func (c *Client) ImageFileByTableNameTargetID(ctx context.Context, tableName, targetID nulltype.NullString) ([]*ImageFile, error) {
	// func ImageFileByTableNameTargetID(ctx context.Context, db DB, tableName nulltype.NullString, targetID nulltype.NullString) ([]*ImageFile, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, TableName, TargetUUID, TargetID, Path, Hash, Size, rb_local_path, rb_insync_hash, rb_insync_local_usn, rb_file_hash_dirty, rb_local_file_status, rb_in_progress, rb_process_type, rb_temp_path, rb_priority, rb_file_size_dirty, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM imageFile ` +
		`WHERE TableName = $1 AND TargetID = $2`
	// run
	logf(sqlstr, tableName, targetID)
	rows, err := db.QueryContext(ctx, sqlstr, tableName, targetID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*ImageFile
	for rows.Next() {
		ifVal := ImageFile{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ifVal.ID, &ifVal.TableName, &ifVal.TargetUUID, &ifVal.TargetID, &ifVal.Path, &ifVal.Hash, &ifVal.Size, &ifVal.RbLocalPath, &ifVal.RbInsyncHash, &ifVal.RbInsyncLocalUsn, &ifVal.RbFileHashDirty, &ifVal.RbLocalFileStatus, &ifVal.RbInProgress, &ifVal.RbProcessType, &ifVal.RbTempPath, &ifVal.RbPriority, &ifVal.RbFileSizeDirty, &ifVal.UUID, &ifVal.RbDataStatus, &ifVal.RbLocalDataStatus, &ifVal.RbLocalDeleted, &ifVal.RbLocalSynced, &ifVal.Usn, &ifVal.RbLocalUsn, &ifVal.CreatedAt, &ifVal.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ifVal)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ImageFileByTableNameTargetUUIDID retrieves a row from 'imageFile' as a ImageFile.
//
// Generated from index 'image_file__table_name__target_u_u_i_d__i_d'.
func (c *Client) ImageFileByTableNameTargetUUIDID(ctx context.Context, tableName, targetUUID, id nulltype.NullString) ([]*ImageFile, error) {
	// func ImageFileByTableNameTargetUUIDID(ctx context.Context, db DB, tableName nulltype.NullString, targetUUID nulltype.NullString, id nulltype.NullString) ([]*ImageFile, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, TableName, TargetUUID, TargetID, Path, Hash, Size, rb_local_path, rb_insync_hash, rb_insync_local_usn, rb_file_hash_dirty, rb_local_file_status, rb_in_progress, rb_process_type, rb_temp_path, rb_priority, rb_file_size_dirty, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM imageFile ` +
		`WHERE TableName = $1 AND TargetUUID = $2 AND ID = $3`
	// run
	logf(sqlstr, tableName, targetUUID, id)
	rows, err := db.QueryContext(ctx, sqlstr, tableName, targetUUID, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*ImageFile
	for rows.Next() {
		ifVal := ImageFile{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ifVal.ID, &ifVal.TableName, &ifVal.TargetUUID, &ifVal.TargetID, &ifVal.Path, &ifVal.Hash, &ifVal.Size, &ifVal.RbLocalPath, &ifVal.RbInsyncHash, &ifVal.RbInsyncLocalUsn, &ifVal.RbFileHashDirty, &ifVal.RbLocalFileStatus, &ifVal.RbInProgress, &ifVal.RbProcessType, &ifVal.RbTempPath, &ifVal.RbPriority, &ifVal.RbFileSizeDirty, &ifVal.UUID, &ifVal.RbDataStatus, &ifVal.RbLocalDataStatus, &ifVal.RbLocalDeleted, &ifVal.RbLocalSynced, &ifVal.Usn, &ifVal.RbLocalUsn, &ifVal.CreatedAt, &ifVal.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ifVal)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ImageFileByUUID retrieves a row from 'imageFile' as a ImageFile.
//
// Generated from index 'image_file__u_u_i_d'.
func (c *Client) ImageFileByUUID(ctx context.Context, uuid nulltype.NullString) ([]*ImageFile, error) {
	// func ImageFileByUUID(ctx context.Context, db DB, uuid nulltype.NullString) ([]*ImageFile, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, TableName, TargetUUID, TargetID, Path, Hash, Size, rb_local_path, rb_insync_hash, rb_insync_local_usn, rb_file_hash_dirty, rb_local_file_status, rb_in_progress, rb_process_type, rb_temp_path, rb_priority, rb_file_size_dirty, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM imageFile ` +
		`WHERE UUID = $1`
	// run
	logf(sqlstr, uuid)
	rows, err := db.QueryContext(ctx, sqlstr, uuid)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*ImageFile
	for rows.Next() {
		ifVal := ImageFile{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ifVal.ID, &ifVal.TableName, &ifVal.TargetUUID, &ifVal.TargetID, &ifVal.Path, &ifVal.Hash, &ifVal.Size, &ifVal.RbLocalPath, &ifVal.RbInsyncHash, &ifVal.RbInsyncLocalUsn, &ifVal.RbFileHashDirty, &ifVal.RbLocalFileStatus, &ifVal.RbInProgress, &ifVal.RbProcessType, &ifVal.RbTempPath, &ifVal.RbPriority, &ifVal.RbFileSizeDirty, &ifVal.UUID, &ifVal.RbDataStatus, &ifVal.RbLocalDataStatus, &ifVal.RbLocalDeleted, &ifVal.RbLocalSynced, &ifVal.Usn, &ifVal.RbLocalUsn, &ifVal.CreatedAt, &ifVal.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ifVal)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ImageFileByRbDataStatus retrieves a row from 'imageFile' as a ImageFile.
//
// Generated from index 'image_file_rb_data_status'.
func (c *Client) ImageFileByRbDataStatus(ctx context.Context, rbDataStatus nulltype.NullInt64) ([]*ImageFile, error) {
	// func ImageFileByRbDataStatus(ctx context.Context, db DB, rbDataStatus nulltype.NullInt64) ([]*ImageFile, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, TableName, TargetUUID, TargetID, Path, Hash, Size, rb_local_path, rb_insync_hash, rb_insync_local_usn, rb_file_hash_dirty, rb_local_file_status, rb_in_progress, rb_process_type, rb_temp_path, rb_priority, rb_file_size_dirty, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM imageFile ` +
		`WHERE rb_data_status = $1`
	// run
	logf(sqlstr, rbDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*ImageFile
	for rows.Next() {
		ifVal := ImageFile{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ifVal.ID, &ifVal.TableName, &ifVal.TargetUUID, &ifVal.TargetID, &ifVal.Path, &ifVal.Hash, &ifVal.Size, &ifVal.RbLocalPath, &ifVal.RbInsyncHash, &ifVal.RbInsyncLocalUsn, &ifVal.RbFileHashDirty, &ifVal.RbLocalFileStatus, &ifVal.RbInProgress, &ifVal.RbProcessType, &ifVal.RbTempPath, &ifVal.RbPriority, &ifVal.RbFileSizeDirty, &ifVal.UUID, &ifVal.RbDataStatus, &ifVal.RbLocalDataStatus, &ifVal.RbLocalDeleted, &ifVal.RbLocalSynced, &ifVal.Usn, &ifVal.RbLocalUsn, &ifVal.CreatedAt, &ifVal.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ifVal)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ImageFileByRbFileHashDirty retrieves a row from 'imageFile' as a ImageFile.
//
// Generated from index 'image_file_rb_file_hash_dirty'.
func (c *Client) ImageFileByRbFileHashDirty(ctx context.Context, rbFileHashDirty nulltype.NullInt64) ([]*ImageFile, error) {
	// func ImageFileByRbFileHashDirty(ctx context.Context, db DB, rbFileHashDirty nulltype.NullInt64) ([]*ImageFile, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, TableName, TargetUUID, TargetID, Path, Hash, Size, rb_local_path, rb_insync_hash, rb_insync_local_usn, rb_file_hash_dirty, rb_local_file_status, rb_in_progress, rb_process_type, rb_temp_path, rb_priority, rb_file_size_dirty, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM imageFile ` +
		`WHERE rb_file_hash_dirty = $1`
	// run
	logf(sqlstr, rbFileHashDirty)
	rows, err := db.QueryContext(ctx, sqlstr, rbFileHashDirty)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*ImageFile
	for rows.Next() {
		ifVal := ImageFile{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ifVal.ID, &ifVal.TableName, &ifVal.TargetUUID, &ifVal.TargetID, &ifVal.Path, &ifVal.Hash, &ifVal.Size, &ifVal.RbLocalPath, &ifVal.RbInsyncHash, &ifVal.RbInsyncLocalUsn, &ifVal.RbFileHashDirty, &ifVal.RbLocalFileStatus, &ifVal.RbInProgress, &ifVal.RbProcessType, &ifVal.RbTempPath, &ifVal.RbPriority, &ifVal.RbFileSizeDirty, &ifVal.UUID, &ifVal.RbDataStatus, &ifVal.RbLocalDataStatus, &ifVal.RbLocalDeleted, &ifVal.RbLocalSynced, &ifVal.Usn, &ifVal.RbLocalUsn, &ifVal.CreatedAt, &ifVal.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ifVal)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ImageFileByRbFileSizeDirty retrieves a row from 'imageFile' as a ImageFile.
//
// Generated from index 'image_file_rb_file_size_dirty'.
func (c *Client) ImageFileByRbFileSizeDirty(ctx context.Context, rbFileSizeDirty nulltype.NullInt64) ([]*ImageFile, error) {
	// func ImageFileByRbFileSizeDirty(ctx context.Context, db DB, rbFileSizeDirty nulltype.NullInt64) ([]*ImageFile, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, TableName, TargetUUID, TargetID, Path, Hash, Size, rb_local_path, rb_insync_hash, rb_insync_local_usn, rb_file_hash_dirty, rb_local_file_status, rb_in_progress, rb_process_type, rb_temp_path, rb_priority, rb_file_size_dirty, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM imageFile ` +
		`WHERE rb_file_size_dirty = $1`
	// run
	logf(sqlstr, rbFileSizeDirty)
	rows, err := db.QueryContext(ctx, sqlstr, rbFileSizeDirty)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*ImageFile
	for rows.Next() {
		ifVal := ImageFile{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ifVal.ID, &ifVal.TableName, &ifVal.TargetUUID, &ifVal.TargetID, &ifVal.Path, &ifVal.Hash, &ifVal.Size, &ifVal.RbLocalPath, &ifVal.RbInsyncHash, &ifVal.RbInsyncLocalUsn, &ifVal.RbFileHashDirty, &ifVal.RbLocalFileStatus, &ifVal.RbInProgress, &ifVal.RbProcessType, &ifVal.RbTempPath, &ifVal.RbPriority, &ifVal.RbFileSizeDirty, &ifVal.UUID, &ifVal.RbDataStatus, &ifVal.RbLocalDataStatus, &ifVal.RbLocalDeleted, &ifVal.RbLocalSynced, &ifVal.Usn, &ifVal.RbLocalUsn, &ifVal.CreatedAt, &ifVal.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ifVal)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ImageFileByRbLocalDataStatus retrieves a row from 'imageFile' as a ImageFile.
//
// Generated from index 'image_file_rb_local_data_status'.
func (c *Client) ImageFileByRbLocalDataStatus(ctx context.Context, rbLocalDataStatus nulltype.NullInt64) ([]*ImageFile, error) {
	// func ImageFileByRbLocalDataStatus(ctx context.Context, db DB, rbLocalDataStatus nulltype.NullInt64) ([]*ImageFile, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, TableName, TargetUUID, TargetID, Path, Hash, Size, rb_local_path, rb_insync_hash, rb_insync_local_usn, rb_file_hash_dirty, rb_local_file_status, rb_in_progress, rb_process_type, rb_temp_path, rb_priority, rb_file_size_dirty, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM imageFile ` +
		`WHERE rb_local_data_status = $1`
	// run
	logf(sqlstr, rbLocalDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*ImageFile
	for rows.Next() {
		ifVal := ImageFile{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ifVal.ID, &ifVal.TableName, &ifVal.TargetUUID, &ifVal.TargetID, &ifVal.Path, &ifVal.Hash, &ifVal.Size, &ifVal.RbLocalPath, &ifVal.RbInsyncHash, &ifVal.RbInsyncLocalUsn, &ifVal.RbFileHashDirty, &ifVal.RbLocalFileStatus, &ifVal.RbInProgress, &ifVal.RbProcessType, &ifVal.RbTempPath, &ifVal.RbPriority, &ifVal.RbFileSizeDirty, &ifVal.UUID, &ifVal.RbDataStatus, &ifVal.RbLocalDataStatus, &ifVal.RbLocalDeleted, &ifVal.RbLocalSynced, &ifVal.Usn, &ifVal.RbLocalUsn, &ifVal.CreatedAt, &ifVal.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ifVal)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ImageFileByRbLocalDeleted retrieves a row from 'imageFile' as a ImageFile.
//
// Generated from index 'image_file_rb_local_deleted'.
func (c *Client) ImageFileByRbLocalDeleted(ctx context.Context, rbLocalDeleted nulltype.NullInt64) ([]*ImageFile, error) {
	// func ImageFileByRbLocalDeleted(ctx context.Context, db DB, rbLocalDeleted nulltype.NullInt64) ([]*ImageFile, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, TableName, TargetUUID, TargetID, Path, Hash, Size, rb_local_path, rb_insync_hash, rb_insync_local_usn, rb_file_hash_dirty, rb_local_file_status, rb_in_progress, rb_process_type, rb_temp_path, rb_priority, rb_file_size_dirty, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM imageFile ` +
		`WHERE rb_local_deleted = $1`
	// run
	logf(sqlstr, rbLocalDeleted)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDeleted)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*ImageFile
	for rows.Next() {
		ifVal := ImageFile{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ifVal.ID, &ifVal.TableName, &ifVal.TargetUUID, &ifVal.TargetID, &ifVal.Path, &ifVal.Hash, &ifVal.Size, &ifVal.RbLocalPath, &ifVal.RbInsyncHash, &ifVal.RbInsyncLocalUsn, &ifVal.RbFileHashDirty, &ifVal.RbLocalFileStatus, &ifVal.RbInProgress, &ifVal.RbProcessType, &ifVal.RbTempPath, &ifVal.RbPriority, &ifVal.RbFileSizeDirty, &ifVal.UUID, &ifVal.RbDataStatus, &ifVal.RbLocalDataStatus, &ifVal.RbLocalDeleted, &ifVal.RbLocalSynced, &ifVal.Usn, &ifVal.RbLocalUsn, &ifVal.CreatedAt, &ifVal.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ifVal)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ImageFileByRbLocalDeletedRbInProgressRbLocalFileStatusRbProcessTypeRbPriority retrieves a row from 'imageFile' as a ImageFile.
//
// Generated from index 'image_file_rb_local_deleted_rb_in_progress_rb_local_file_status_rb_process_type_rb_priority'.
func (c *Client) ImageFileByRbLocalDeletedRbInProgressRbLocalFileStatusRbProcessTypeRbPriority(ctx context.Context, rbLocalDeleted, rbInProgress, rbLocalFileStatus, rbProcessType, rbPriority nulltype.NullInt64) ([]*ImageFile, error) {
	// func ImageFileByRbLocalDeletedRbInProgressRbLocalFileStatusRbProcessTypeRbPriority(ctx context.Context, db DB, rbLocalDeleted nulltype.NullInt64, rbInProgress nulltype.NullInt64, rbLocalFileStatus nulltype.NullInt64, rbProcessType nulltype.NullInt64, rbPriority nulltype.NullInt64) ([]*ImageFile, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, TableName, TargetUUID, TargetID, Path, Hash, Size, rb_local_path, rb_insync_hash, rb_insync_local_usn, rb_file_hash_dirty, rb_local_file_status, rb_in_progress, rb_process_type, rb_temp_path, rb_priority, rb_file_size_dirty, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM imageFile ` +
		`WHERE rb_local_deleted = $1 AND rb_in_progress = $2 AND rb_local_file_status = $3 AND rb_process_type = $4 AND rb_priority = $5`
	// run
	logf(sqlstr, rbLocalDeleted, rbInProgress, rbLocalFileStatus, rbProcessType, rbPriority)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDeleted, rbInProgress, rbLocalFileStatus, rbProcessType, rbPriority)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*ImageFile
	for rows.Next() {
		ifVal := ImageFile{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ifVal.ID, &ifVal.TableName, &ifVal.TargetUUID, &ifVal.TargetID, &ifVal.Path, &ifVal.Hash, &ifVal.Size, &ifVal.RbLocalPath, &ifVal.RbInsyncHash, &ifVal.RbInsyncLocalUsn, &ifVal.RbFileHashDirty, &ifVal.RbLocalFileStatus, &ifVal.RbInProgress, &ifVal.RbProcessType, &ifVal.RbTempPath, &ifVal.RbPriority, &ifVal.RbFileSizeDirty, &ifVal.UUID, &ifVal.RbDataStatus, &ifVal.RbLocalDataStatus, &ifVal.RbLocalDeleted, &ifVal.RbLocalSynced, &ifVal.Usn, &ifVal.RbLocalUsn, &ifVal.CreatedAt, &ifVal.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ifVal)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ImageFileByRbLocalUsnID retrieves a row from 'imageFile' as a ImageFile.
//
// Generated from index 'image_file_rb_local_usn__i_d'.
func (c *Client) ImageFileByRbLocalUsnID(ctx context.Context, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*ImageFile, error) {
	// func ImageFileByRbLocalUsnID(ctx context.Context, db DB, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*ImageFile, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, TableName, TargetUUID, TargetID, Path, Hash, Size, rb_local_path, rb_insync_hash, rb_insync_local_usn, rb_file_hash_dirty, rb_local_file_status, rb_in_progress, rb_process_type, rb_temp_path, rb_priority, rb_file_size_dirty, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM imageFile ` +
		`WHERE rb_local_usn = $1 AND ID = $2`
	// run
	logf(sqlstr, rbLocalUsn, id)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalUsn, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*ImageFile
	for rows.Next() {
		ifVal := ImageFile{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ifVal.ID, &ifVal.TableName, &ifVal.TargetUUID, &ifVal.TargetID, &ifVal.Path, &ifVal.Hash, &ifVal.Size, &ifVal.RbLocalPath, &ifVal.RbInsyncHash, &ifVal.RbInsyncLocalUsn, &ifVal.RbFileHashDirty, &ifVal.RbLocalFileStatus, &ifVal.RbInProgress, &ifVal.RbProcessType, &ifVal.RbTempPath, &ifVal.RbPriority, &ifVal.RbFileSizeDirty, &ifVal.UUID, &ifVal.RbDataStatus, &ifVal.RbLocalDataStatus, &ifVal.RbLocalDeleted, &ifVal.RbLocalSynced, &ifVal.Usn, &ifVal.RbLocalUsn, &ifVal.CreatedAt, &ifVal.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ifVal)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ImageFileByID retrieves a row from 'imageFile' as a ImageFile.
//
// Generated from index 'sqlite_autoindex_imageFile_1'.
func (c *Client) ImageFileByID(ctx context.Context, id nulltype.NullString) (*ImageFile, error) {
	// func ImageFileByID(ctx context.Context, db DB, id nulltype.NullString) (*ImageFile, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, TableName, TargetUUID, TargetID, Path, Hash, Size, rb_local_path, rb_insync_hash, rb_insync_local_usn, rb_file_hash_dirty, rb_local_file_status, rb_in_progress, rb_process_type, rb_temp_path, rb_priority, rb_file_size_dirty, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM imageFile ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, id)
	ifVal := ImageFile{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&ifVal.ID, &ifVal.TableName, &ifVal.TargetUUID, &ifVal.TargetID, &ifVal.Path, &ifVal.Hash, &ifVal.Size, &ifVal.RbLocalPath, &ifVal.RbInsyncHash, &ifVal.RbInsyncLocalUsn, &ifVal.RbFileHashDirty, &ifVal.RbLocalFileStatus, &ifVal.RbInProgress, &ifVal.RbProcessType, &ifVal.RbTempPath, &ifVal.RbPriority, &ifVal.RbFileSizeDirty, &ifVal.UUID, &ifVal.RbDataStatus, &ifVal.RbLocalDataStatus, &ifVal.RbLocalDeleted, &ifVal.RbLocalSynced, &ifVal.Usn, &ifVal.RbLocalUsn, &ifVal.CreatedAt, &ifVal.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &ifVal, nil
}
