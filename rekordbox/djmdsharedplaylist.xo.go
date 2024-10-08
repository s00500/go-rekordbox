package rekordbox

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	nulltype "github.com/mattn/go-nulltype"
)

// DjmdSharedPlaylist represents a row from 'djmdSharedPlaylist'.
type DjmdSharedPlaylist struct {
	ID            nulltype.NullString `json:"id"`             // ID
	DataSelection nulltype.NullInt64  `json:"data_selection"` // data_selection
	EditedAt      *Time               `json:"edited_at"`      // edited_at
	Int1          nulltype.NullInt64  `json:"int1"`           // int_1
	Int2          nulltype.NullInt64  `json:"int2"`           // int_2
	Str1          nulltype.NullString `json:"str1"`           // str_1
	Str2          nulltype.NullString `json:"str2"`           // str_2
	Text1         nulltype.NullString `json:"text1"`          // text_1
	Text2         nulltype.NullString `json:"text2"`          // text_2
	CreatedAt     Time                `json:"created_at"`     // created_at
	UpdatedAt     Time                `json:"updated_at"`     // updated_at
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the DjmdSharedPlaylist exists in the database.
func (dsp *DjmdSharedPlaylist) Exists() bool {
	return dsp._exists
}

// Deleted returns true when the DjmdSharedPlaylist has been marked for deletion from
// the database.
func (dsp *DjmdSharedPlaylist) Deleted() bool {
	return dsp._deleted
}

// Insert inserts the DjmdSharedPlaylist to the database.
func (c *Client) InsertDjmdSharedPlaylist(ctx context.Context, dsp *DjmdSharedPlaylist) error {
	db := c.db

	switch {
	case dsp._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case dsp._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO djmdSharedPlaylist (` +
		`ID, data_selection, edited_at, int_1, int_2, str_1, str_2, text_1, text_2, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`)`
	// run
	logf(sqlstr, dsp.ID, dsp.DataSelection, dsp.EditedAt, dsp.Int1, dsp.Int2, dsp.Str1, dsp.Str2, dsp.Text1, dsp.Text2, dsp.CreatedAt, dsp.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dsp.ID, dsp.DataSelection, dsp.EditedAt, dsp.Int1, dsp.Int2, dsp.Str1, dsp.Str2, dsp.Text1, dsp.Text2, dsp.CreatedAt, dsp.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dsp._exists = true
	return nil
}

// Update updates a DjmdSharedPlaylist in the database.
func (c *Client) UpdateDjmdSharedPlaylist(ctx context.Context, dsp *DjmdSharedPlaylist) error {
	db := c.db

	switch {
	case !dsp._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case dsp._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE djmdSharedPlaylist SET ` +
		`data_selection = $1, edited_at = $2, int_1 = $3, int_2 = $4, str_1 = $5, str_2 = $6, text_1 = $7, text_2 = $8, created_at = $9, updated_at = $10 ` +
		`WHERE ID = $11`
	// run
	logf(sqlstr, dsp.DataSelection, dsp.EditedAt, dsp.Int1, dsp.Int2, dsp.Str1, dsp.Str2, dsp.Text1, dsp.Text2, dsp.CreatedAt, dsp.UpdatedAt, dsp.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dsp.DataSelection, dsp.EditedAt, dsp.Int1, dsp.Int2, dsp.Str1, dsp.Str2, dsp.Text1, dsp.Text2, dsp.CreatedAt, dsp.UpdatedAt, dsp.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the DjmdSharedPlaylist to the database.
func (c *Client) SaveDjmdSharedPlaylist(ctx context.Context, dsp *DjmdSharedPlaylist) error {
	if dsp.Exists() {
		return c.UpdateDjmdSharedPlaylist(ctx, dsp)
	}
	return c.InsertDjmdSharedPlaylist(ctx, dsp)
}

// Upsert performs an upsert for DjmdSharedPlaylist.
func (c *Client) UpsertDjmdSharedPlaylist(ctx context.Context, dsp *DjmdSharedPlaylist) error {
	db := c.db

	switch {
	case dsp._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO djmdSharedPlaylist (` +
		`ID, data_selection, edited_at, int_1, int_2, str_1, str_2, text_1, text_2, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`)` +
		` ON CONFLICT (ID) DO ` +
		`UPDATE SET ` +
		`data_selection = EXCLUDED.data_selection, edited_at = EXCLUDED.edited_at, int_1 = EXCLUDED.int_1, int_2 = EXCLUDED.int_2, str_1 = EXCLUDED.str_1, str_2 = EXCLUDED.str_2, text_1 = EXCLUDED.text_1, text_2 = EXCLUDED.text_2, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, dsp.ID, dsp.DataSelection, dsp.EditedAt, dsp.Int1, dsp.Int2, dsp.Str1, dsp.Str2, dsp.Text1, dsp.Text2, dsp.CreatedAt, dsp.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dsp.ID, dsp.DataSelection, dsp.EditedAt, dsp.Int1, dsp.Int2, dsp.Str1, dsp.Str2, dsp.Text1, dsp.Text2, dsp.CreatedAt, dsp.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dsp._exists = true
	return nil
}

// Delete deletes the DjmdSharedPlaylist from the database.
func (c *Client) DeleteDjmdSharedPlaylist(ctx context.Context, dsp *DjmdSharedPlaylist) error {
	db := c.db

	switch {
	case !dsp._exists: // doesn't exist
		return nil
	case dsp._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM djmdSharedPlaylist ` +
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

func scanDjmdSharedPlaylistRows(rows *sql.Rows) ([]*DjmdSharedPlaylist, error) {
	var res []*DjmdSharedPlaylist
	for rows.Next() {
		dsp := DjmdSharedPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dsp.ID, &dsp.DataSelection, &dsp.EditedAt, &dsp.Int1, &dsp.Int2, &dsp.Str1, &dsp.Str2, &dsp.Text1, &dsp.Text2, &dsp.CreatedAt, &dsp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dsp)
	}

	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

func (c *Client) AllDjmdSharedPlaylist(ctx context.Context) ([]*DjmdSharedPlaylist, error) {
	db := c.db

	const sqlstr = `SELECT * FROM DjmdSharedPlaylist`
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()
	res, err := scanDjmdSharedPlaylistRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSharedPlaylistByID retrieves a row from 'djmdSharedPlaylist' as a DjmdSharedPlaylist.
//
// Generated from index 'sqlite_autoindex_djmdSharedPlaylist_1'.
func (c *Client) DjmdSharedPlaylistByID(ctx context.Context, id nulltype.NullString) (*DjmdSharedPlaylist, error) {
	// func DjmdSharedPlaylistByID(ctx context.Context, db DB, id nulltype.NullString) (*DjmdSharedPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, data_selection, edited_at, int_1, int_2, str_1, str_2, text_1, text_2, created_at, updated_at ` +
		`FROM djmdSharedPlaylist ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, id)
	dsp := DjmdSharedPlaylist{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&dsp.ID, &dsp.DataSelection, &dsp.EditedAt, &dsp.Int1, &dsp.Int2, &dsp.Str1, &dsp.Str2, &dsp.Text1, &dsp.Text2, &dsp.CreatedAt, &dsp.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &dsp, nil
}
