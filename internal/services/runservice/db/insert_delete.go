// Code generated by go generate; DO NOT EDIT.
package db

import (
	"encoding/json"
	"time"

	idb "agola.io/agola/internal/db"
	"agola.io/agola/internal/errors"
	"agola.io/agola/internal/sql"
	"agola.io/agola/services/runservice/types"

	sq "github.com/Masterminds/squirrel"
)

func (d *DB) InsertOrUpdateSequence(tx *sql.Tx, v *types.Sequence) error {
	var err error
	if v.Revision == 0 {
		err = d.InsertSequence(tx, v)
	} else {
		err = d.UpdateSequence(tx, v)
	}

	return errors.WithStack(err)
}

func (d *DB) InsertSequence(tx *sql.Tx, v *types.Sequence) error {
	if v.Revision != 0 {
		return errors.Errorf("expected revision 0 got %d", v.Revision)
	}

	data, err := d.insertSequenceData(tx, v)
	if err != nil {
		return errors.WithStack(err)
	}

	return d.insertSequenceQ(tx, v, data)
}

func (d *DB) insertSequenceData(tx *sql.Tx, v *types.Sequence) ([]byte, error) {
	v.Revision = 1

	now := time.Now()
	v.SetCreationTime(now)
	v.SetUpdateTime(now)

	data, err := json.Marshal(v)
	if err != nil {
		v.Revision = 0
		return nil, errors.WithStack(err)
	}

	q := sb.Insert("sequence_t").Columns("id", "revision", "data").Values(v.ID, v.Revision, data)
	if _, err := d.exec(tx, q); err != nil {
		v.Revision = 0
		return nil, errors.Wrap(err, "failed to insert sequence_t")
	}

	return data, nil
}

// insertRawSequenceData should be used only for import.
// It won't update object times.
func (d *DB) insertRawSequenceData(tx *sql.Tx, v *types.Sequence) ([]byte, error) {
	v.Revision = 1

	data, err := json.Marshal(v)
	if err != nil {
		v.Revision = 0
		return nil, errors.WithStack(err)
	}

	q := sb.Insert("sequence_t").Columns("id", "revision", "data").Values(v.ID, v.Revision, data)
	if _, err := d.exec(tx, q); err != nil {
		v.Revision = 0
		return nil, errors.Wrap(err, "failed to insert sequence_t")
	}

	return data, nil
}

func (d *DB) UpdateSequence(tx *sql.Tx, v *types.Sequence) error {
	data, err := d.updateSequenceData(tx, v)
	if err != nil {
		return errors.WithStack(err)
	}

	return d.updateSequenceQ(tx, v, data)
}

func (d *DB) updateSequenceData(tx *sql.Tx, v *types.Sequence) ([]byte, error) {
	if v.Revision < 1 {
		return nil, errors.Errorf("expected revision > 0 got %d", v.Revision)
	}

	curRevision := v.Revision
	v.Revision++

	v.SetUpdateTime(time.Now())

	data, err := json.Marshal(v)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	q := sb.Update("sequence_t").SetMap(map[string]interface{}{"id": v.ID, "revision": v.Revision, "data": data}).Where(sq.Eq{"id": v.ID, "revision": curRevision})
	res, err := d.exec(tx, q)
	if err != nil {
		v.Revision = curRevision
		return nil, errors.Wrap(err, "failed to update sequence_t")
	}

	rows, err := res.RowsAffected()
	if err != nil {
		v.Revision = curRevision
		return nil, errors.Wrap(err, "failed to update sequence_t")
	}

	if rows != 1 {
		v.Revision = curRevision
		return nil, idb.ErrConcurrent
	}

	return data, nil
}

func (d *DB) DeleteSequence(tx *sql.Tx, id string) error {
	if err := d.deleteSequenceData(tx, id); err != nil {
		return errors.WithStack(err)
	}

	return d.deleteSequenceQ(tx, id)
}

func (d *DB) deleteSequenceData(tx *sql.Tx, id string) error {
	if _, err := tx.Exec("delete from sequence_t where id = $1", id); err != nil {
		return errors.Wrap(err, "failed to delete sequence_t")
	}

	return nil
}

func (d *DB) InsertOrUpdateChangeGroup(tx *sql.Tx, v *types.ChangeGroup) error {
	var err error
	if v.Revision == 0 {
		err = d.InsertChangeGroup(tx, v)
	} else {
		err = d.UpdateChangeGroup(tx, v)
	}

	return errors.WithStack(err)
}

func (d *DB) InsertChangeGroup(tx *sql.Tx, v *types.ChangeGroup) error {
	if v.Revision != 0 {
		return errors.Errorf("expected revision 0 got %d", v.Revision)
	}

	data, err := d.insertChangeGroupData(tx, v)
	if err != nil {
		return errors.WithStack(err)
	}

	return d.insertChangeGroupQ(tx, v, data)
}

func (d *DB) insertChangeGroupData(tx *sql.Tx, v *types.ChangeGroup) ([]byte, error) {
	v.Revision = 1

	now := time.Now()
	v.SetCreationTime(now)
	v.SetUpdateTime(now)

	data, err := json.Marshal(v)
	if err != nil {
		v.Revision = 0
		return nil, errors.WithStack(err)
	}

	q := sb.Insert("changegroup").Columns("id", "revision", "data").Values(v.ID, v.Revision, data)
	if _, err := d.exec(tx, q); err != nil {
		v.Revision = 0
		return nil, errors.Wrap(err, "failed to insert changegroup")
	}

	return data, nil
}

// insertRawChangeGroupData should be used only for import.
// It won't update object times.
func (d *DB) insertRawChangeGroupData(tx *sql.Tx, v *types.ChangeGroup) ([]byte, error) {
	v.Revision = 1

	data, err := json.Marshal(v)
	if err != nil {
		v.Revision = 0
		return nil, errors.WithStack(err)
	}

	q := sb.Insert("changegroup").Columns("id", "revision", "data").Values(v.ID, v.Revision, data)
	if _, err := d.exec(tx, q); err != nil {
		v.Revision = 0
		return nil, errors.Wrap(err, "failed to insert changegroup")
	}

	return data, nil
}

func (d *DB) UpdateChangeGroup(tx *sql.Tx, v *types.ChangeGroup) error {
	data, err := d.updateChangeGroupData(tx, v)
	if err != nil {
		return errors.WithStack(err)
	}

	return d.updateChangeGroupQ(tx, v, data)
}

func (d *DB) updateChangeGroupData(tx *sql.Tx, v *types.ChangeGroup) ([]byte, error) {
	if v.Revision < 1 {
		return nil, errors.Errorf("expected revision > 0 got %d", v.Revision)
	}

	curRevision := v.Revision
	v.Revision++

	v.SetUpdateTime(time.Now())

	data, err := json.Marshal(v)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	q := sb.Update("changegroup").SetMap(map[string]interface{}{"id": v.ID, "revision": v.Revision, "data": data}).Where(sq.Eq{"id": v.ID, "revision": curRevision})
	res, err := d.exec(tx, q)
	if err != nil {
		v.Revision = curRevision
		return nil, errors.Wrap(err, "failed to update changegroup")
	}

	rows, err := res.RowsAffected()
	if err != nil {
		v.Revision = curRevision
		return nil, errors.Wrap(err, "failed to update changegroup")
	}

	if rows != 1 {
		v.Revision = curRevision
		return nil, idb.ErrConcurrent
	}

	return data, nil
}

func (d *DB) DeleteChangeGroup(tx *sql.Tx, id string) error {
	if err := d.deleteChangeGroupData(tx, id); err != nil {
		return errors.WithStack(err)
	}

	return d.deleteChangeGroupQ(tx, id)
}

func (d *DB) deleteChangeGroupData(tx *sql.Tx, id string) error {
	if _, err := tx.Exec("delete from changegroup where id = $1", id); err != nil {
		return errors.Wrap(err, "failed to delete changegroup")
	}

	return nil
}

func (d *DB) InsertOrUpdateRun(tx *sql.Tx, v *types.Run) error {
	var err error
	if v.Revision == 0 {
		err = d.InsertRun(tx, v)
	} else {
		err = d.UpdateRun(tx, v)
	}

	return errors.WithStack(err)
}

func (d *DB) InsertRun(tx *sql.Tx, v *types.Run) error {
	if v.Revision != 0 {
		return errors.Errorf("expected revision 0 got %d", v.Revision)
	}

	data, err := d.insertRunData(tx, v)
	if err != nil {
		return errors.WithStack(err)
	}

	return d.insertRunQ(tx, v, data)
}

func (d *DB) insertRunData(tx *sql.Tx, v *types.Run) ([]byte, error) {
	v.Revision = 1

	now := time.Now()
	v.SetCreationTime(now)
	v.SetUpdateTime(now)

	data, err := json.Marshal(v)
	if err != nil {
		v.Revision = 0
		return nil, errors.WithStack(err)
	}

	q := sb.Insert("run").Columns("id", "revision", "data").Values(v.ID, v.Revision, data)
	if _, err := d.exec(tx, q); err != nil {
		v.Revision = 0
		return nil, errors.Wrap(err, "failed to insert run")
	}

	return data, nil
}

// insertRawRunData should be used only for import.
// It won't update object times.
func (d *DB) insertRawRunData(tx *sql.Tx, v *types.Run) ([]byte, error) {
	v.Revision = 1

	data, err := json.Marshal(v)
	if err != nil {
		v.Revision = 0
		return nil, errors.WithStack(err)
	}

	q := sb.Insert("run").Columns("id", "revision", "data").Values(v.ID, v.Revision, data)
	if _, err := d.exec(tx, q); err != nil {
		v.Revision = 0
		return nil, errors.Wrap(err, "failed to insert run")
	}

	return data, nil
}

func (d *DB) UpdateRun(tx *sql.Tx, v *types.Run) error {
	data, err := d.updateRunData(tx, v)
	if err != nil {
		return errors.WithStack(err)
	}

	return d.updateRunQ(tx, v, data)
}

func (d *DB) updateRunData(tx *sql.Tx, v *types.Run) ([]byte, error) {
	if v.Revision < 1 {
		return nil, errors.Errorf("expected revision > 0 got %d", v.Revision)
	}

	curRevision := v.Revision
	v.Revision++

	v.SetUpdateTime(time.Now())

	data, err := json.Marshal(v)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	q := sb.Update("run").SetMap(map[string]interface{}{"id": v.ID, "revision": v.Revision, "data": data}).Where(sq.Eq{"id": v.ID, "revision": curRevision})
	res, err := d.exec(tx, q)
	if err != nil {
		v.Revision = curRevision
		return nil, errors.Wrap(err, "failed to update run")
	}

	rows, err := res.RowsAffected()
	if err != nil {
		v.Revision = curRevision
		return nil, errors.Wrap(err, "failed to update run")
	}

	if rows != 1 {
		v.Revision = curRevision
		return nil, idb.ErrConcurrent
	}

	return data, nil
}

func (d *DB) DeleteRun(tx *sql.Tx, id string) error {
	if err := d.deleteRunData(tx, id); err != nil {
		return errors.WithStack(err)
	}

	return d.deleteRunQ(tx, id)
}

func (d *DB) deleteRunData(tx *sql.Tx, id string) error {
	if _, err := tx.Exec("delete from run where id = $1", id); err != nil {
		return errors.Wrap(err, "failed to delete run")
	}

	return nil
}

func (d *DB) InsertOrUpdateRunConfig(tx *sql.Tx, v *types.RunConfig) error {
	var err error
	if v.Revision == 0 {
		err = d.InsertRunConfig(tx, v)
	} else {
		err = d.UpdateRunConfig(tx, v)
	}

	return errors.WithStack(err)
}

func (d *DB) InsertRunConfig(tx *sql.Tx, v *types.RunConfig) error {
	if v.Revision != 0 {
		return errors.Errorf("expected revision 0 got %d", v.Revision)
	}

	data, err := d.insertRunConfigData(tx, v)
	if err != nil {
		return errors.WithStack(err)
	}

	return d.insertRunConfigQ(tx, v, data)
}

func (d *DB) insertRunConfigData(tx *sql.Tx, v *types.RunConfig) ([]byte, error) {
	v.Revision = 1

	now := time.Now()
	v.SetCreationTime(now)
	v.SetUpdateTime(now)

	data, err := json.Marshal(v)
	if err != nil {
		v.Revision = 0
		return nil, errors.WithStack(err)
	}

	q := sb.Insert("runconfig").Columns("id", "revision", "data").Values(v.ID, v.Revision, data)
	if _, err := d.exec(tx, q); err != nil {
		v.Revision = 0
		return nil, errors.Wrap(err, "failed to insert runconfig")
	}

	return data, nil
}

// insertRawRunConfigData should be used only for import.
// It won't update object times.
func (d *DB) insertRawRunConfigData(tx *sql.Tx, v *types.RunConfig) ([]byte, error) {
	v.Revision = 1

	data, err := json.Marshal(v)
	if err != nil {
		v.Revision = 0
		return nil, errors.WithStack(err)
	}

	q := sb.Insert("runconfig").Columns("id", "revision", "data").Values(v.ID, v.Revision, data)
	if _, err := d.exec(tx, q); err != nil {
		v.Revision = 0
		return nil, errors.Wrap(err, "failed to insert runconfig")
	}

	return data, nil
}

func (d *DB) UpdateRunConfig(tx *sql.Tx, v *types.RunConfig) error {
	data, err := d.updateRunConfigData(tx, v)
	if err != nil {
		return errors.WithStack(err)
	}

	return d.updateRunConfigQ(tx, v, data)
}

func (d *DB) updateRunConfigData(tx *sql.Tx, v *types.RunConfig) ([]byte, error) {
	if v.Revision < 1 {
		return nil, errors.Errorf("expected revision > 0 got %d", v.Revision)
	}

	curRevision := v.Revision
	v.Revision++

	v.SetUpdateTime(time.Now())

	data, err := json.Marshal(v)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	q := sb.Update("runconfig").SetMap(map[string]interface{}{"id": v.ID, "revision": v.Revision, "data": data}).Where(sq.Eq{"id": v.ID, "revision": curRevision})
	res, err := d.exec(tx, q)
	if err != nil {
		v.Revision = curRevision
		return nil, errors.Wrap(err, "failed to update runconfig")
	}

	rows, err := res.RowsAffected()
	if err != nil {
		v.Revision = curRevision
		return nil, errors.Wrap(err, "failed to update runconfig")
	}

	if rows != 1 {
		v.Revision = curRevision
		return nil, idb.ErrConcurrent
	}

	return data, nil
}

func (d *DB) DeleteRunConfig(tx *sql.Tx, id string) error {
	if err := d.deleteRunConfigData(tx, id); err != nil {
		return errors.WithStack(err)
	}

	return d.deleteRunConfigQ(tx, id)
}

func (d *DB) deleteRunConfigData(tx *sql.Tx, id string) error {
	if _, err := tx.Exec("delete from runconfig where id = $1", id); err != nil {
		return errors.Wrap(err, "failed to delete runconfig")
	}

	return nil
}

func (d *DB) InsertOrUpdateRunCounter(tx *sql.Tx, v *types.RunCounter) error {
	var err error
	if v.Revision == 0 {
		err = d.InsertRunCounter(tx, v)
	} else {
		err = d.UpdateRunCounter(tx, v)
	}

	return errors.WithStack(err)
}

func (d *DB) InsertRunCounter(tx *sql.Tx, v *types.RunCounter) error {
	if v.Revision != 0 {
		return errors.Errorf("expected revision 0 got %d", v.Revision)
	}

	data, err := d.insertRunCounterData(tx, v)
	if err != nil {
		return errors.WithStack(err)
	}

	return d.insertRunCounterQ(tx, v, data)
}

func (d *DB) insertRunCounterData(tx *sql.Tx, v *types.RunCounter) ([]byte, error) {
	v.Revision = 1

	now := time.Now()
	v.SetCreationTime(now)
	v.SetUpdateTime(now)

	data, err := json.Marshal(v)
	if err != nil {
		v.Revision = 0
		return nil, errors.WithStack(err)
	}

	q := sb.Insert("runcounter").Columns("id", "revision", "data").Values(v.ID, v.Revision, data)
	if _, err := d.exec(tx, q); err != nil {
		v.Revision = 0
		return nil, errors.Wrap(err, "failed to insert runcounter")
	}

	return data, nil
}

// insertRawRunCounterData should be used only for import.
// It won't update object times.
func (d *DB) insertRawRunCounterData(tx *sql.Tx, v *types.RunCounter) ([]byte, error) {
	v.Revision = 1

	data, err := json.Marshal(v)
	if err != nil {
		v.Revision = 0
		return nil, errors.WithStack(err)
	}

	q := sb.Insert("runcounter").Columns("id", "revision", "data").Values(v.ID, v.Revision, data)
	if _, err := d.exec(tx, q); err != nil {
		v.Revision = 0
		return nil, errors.Wrap(err, "failed to insert runcounter")
	}

	return data, nil
}

func (d *DB) UpdateRunCounter(tx *sql.Tx, v *types.RunCounter) error {
	data, err := d.updateRunCounterData(tx, v)
	if err != nil {
		return errors.WithStack(err)
	}

	return d.updateRunCounterQ(tx, v, data)
}

func (d *DB) updateRunCounterData(tx *sql.Tx, v *types.RunCounter) ([]byte, error) {
	if v.Revision < 1 {
		return nil, errors.Errorf("expected revision > 0 got %d", v.Revision)
	}

	curRevision := v.Revision
	v.Revision++

	v.SetUpdateTime(time.Now())

	data, err := json.Marshal(v)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	q := sb.Update("runcounter").SetMap(map[string]interface{}{"id": v.ID, "revision": v.Revision, "data": data}).Where(sq.Eq{"id": v.ID, "revision": curRevision})
	res, err := d.exec(tx, q)
	if err != nil {
		v.Revision = curRevision
		return nil, errors.Wrap(err, "failed to update runcounter")
	}

	rows, err := res.RowsAffected()
	if err != nil {
		v.Revision = curRevision
		return nil, errors.Wrap(err, "failed to update runcounter")
	}

	if rows != 1 {
		v.Revision = curRevision
		return nil, idb.ErrConcurrent
	}

	return data, nil
}

func (d *DB) DeleteRunCounter(tx *sql.Tx, id string) error {
	if err := d.deleteRunCounterData(tx, id); err != nil {
		return errors.WithStack(err)
	}

	return d.deleteRunCounterQ(tx, id)
}

func (d *DB) deleteRunCounterData(tx *sql.Tx, id string) error {
	if _, err := tx.Exec("delete from runcounter where id = $1", id); err != nil {
		return errors.Wrap(err, "failed to delete runcounter")
	}

	return nil
}

func (d *DB) InsertOrUpdateRunEvent(tx *sql.Tx, v *types.RunEvent) error {
	var err error
	if v.Revision == 0 {
		err = d.InsertRunEvent(tx, v)
	} else {
		err = d.UpdateRunEvent(tx, v)
	}

	return errors.WithStack(err)
}

func (d *DB) InsertRunEvent(tx *sql.Tx, v *types.RunEvent) error {
	if v.Revision != 0 {
		return errors.Errorf("expected revision 0 got %d", v.Revision)
	}

	data, err := d.insertRunEventData(tx, v)
	if err != nil {
		return errors.WithStack(err)
	}

	return d.insertRunEventQ(tx, v, data)
}

func (d *DB) insertRunEventData(tx *sql.Tx, v *types.RunEvent) ([]byte, error) {
	v.Revision = 1

	now := time.Now()
	v.SetCreationTime(now)
	v.SetUpdateTime(now)

	data, err := json.Marshal(v)
	if err != nil {
		v.Revision = 0
		return nil, errors.WithStack(err)
	}

	q := sb.Insert("runevent").Columns("id", "revision", "data").Values(v.ID, v.Revision, data)
	if _, err := d.exec(tx, q); err != nil {
		v.Revision = 0
		return nil, errors.Wrap(err, "failed to insert runevent")
	}

	return data, nil
}

// insertRawRunEventData should be used only for import.
// It won't update object times.
func (d *DB) insertRawRunEventData(tx *sql.Tx, v *types.RunEvent) ([]byte, error) {
	v.Revision = 1

	data, err := json.Marshal(v)
	if err != nil {
		v.Revision = 0
		return nil, errors.WithStack(err)
	}

	q := sb.Insert("runevent").Columns("id", "revision", "data").Values(v.ID, v.Revision, data)
	if _, err := d.exec(tx, q); err != nil {
		v.Revision = 0
		return nil, errors.Wrap(err, "failed to insert runevent")
	}

	return data, nil
}

func (d *DB) UpdateRunEvent(tx *sql.Tx, v *types.RunEvent) error {
	data, err := d.updateRunEventData(tx, v)
	if err != nil {
		return errors.WithStack(err)
	}

	return d.updateRunEventQ(tx, v, data)
}

func (d *DB) updateRunEventData(tx *sql.Tx, v *types.RunEvent) ([]byte, error) {
	if v.Revision < 1 {
		return nil, errors.Errorf("expected revision > 0 got %d", v.Revision)
	}

	curRevision := v.Revision
	v.Revision++

	v.SetUpdateTime(time.Now())

	data, err := json.Marshal(v)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	q := sb.Update("runevent").SetMap(map[string]interface{}{"id": v.ID, "revision": v.Revision, "data": data}).Where(sq.Eq{"id": v.ID, "revision": curRevision})
	res, err := d.exec(tx, q)
	if err != nil {
		v.Revision = curRevision
		return nil, errors.Wrap(err, "failed to update runevent")
	}

	rows, err := res.RowsAffected()
	if err != nil {
		v.Revision = curRevision
		return nil, errors.Wrap(err, "failed to update runevent")
	}

	if rows != 1 {
		v.Revision = curRevision
		return nil, idb.ErrConcurrent
	}

	return data, nil
}

func (d *DB) DeleteRunEvent(tx *sql.Tx, id string) error {
	if err := d.deleteRunEventData(tx, id); err != nil {
		return errors.WithStack(err)
	}

	return d.deleteRunEventQ(tx, id)
}

func (d *DB) deleteRunEventData(tx *sql.Tx, id string) error {
	if _, err := tx.Exec("delete from runevent where id = $1", id); err != nil {
		return errors.Wrap(err, "failed to delete runevent")
	}

	return nil
}

func (d *DB) InsertOrUpdateExecutor(tx *sql.Tx, v *types.Executor) error {
	var err error
	if v.Revision == 0 {
		err = d.InsertExecutor(tx, v)
	} else {
		err = d.UpdateExecutor(tx, v)
	}

	return errors.WithStack(err)
}

func (d *DB) InsertExecutor(tx *sql.Tx, v *types.Executor) error {
	if v.Revision != 0 {
		return errors.Errorf("expected revision 0 got %d", v.Revision)
	}

	data, err := d.insertExecutorData(tx, v)
	if err != nil {
		return errors.WithStack(err)
	}

	return d.insertExecutorQ(tx, v, data)
}

func (d *DB) insertExecutorData(tx *sql.Tx, v *types.Executor) ([]byte, error) {
	v.Revision = 1

	now := time.Now()
	v.SetCreationTime(now)
	v.SetUpdateTime(now)

	data, err := json.Marshal(v)
	if err != nil {
		v.Revision = 0
		return nil, errors.WithStack(err)
	}

	q := sb.Insert("executor").Columns("id", "revision", "data").Values(v.ID, v.Revision, data)
	if _, err := d.exec(tx, q); err != nil {
		v.Revision = 0
		return nil, errors.Wrap(err, "failed to insert executor")
	}

	return data, nil
}

// insertRawExecutorData should be used only for import.
// It won't update object times.
func (d *DB) insertRawExecutorData(tx *sql.Tx, v *types.Executor) ([]byte, error) {
	v.Revision = 1

	data, err := json.Marshal(v)
	if err != nil {
		v.Revision = 0
		return nil, errors.WithStack(err)
	}

	q := sb.Insert("executor").Columns("id", "revision", "data").Values(v.ID, v.Revision, data)
	if _, err := d.exec(tx, q); err != nil {
		v.Revision = 0
		return nil, errors.Wrap(err, "failed to insert executor")
	}

	return data, nil
}

func (d *DB) UpdateExecutor(tx *sql.Tx, v *types.Executor) error {
	data, err := d.updateExecutorData(tx, v)
	if err != nil {
		return errors.WithStack(err)
	}

	return d.updateExecutorQ(tx, v, data)
}

func (d *DB) updateExecutorData(tx *sql.Tx, v *types.Executor) ([]byte, error) {
	if v.Revision < 1 {
		return nil, errors.Errorf("expected revision > 0 got %d", v.Revision)
	}

	curRevision := v.Revision
	v.Revision++

	v.SetUpdateTime(time.Now())

	data, err := json.Marshal(v)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	q := sb.Update("executor").SetMap(map[string]interface{}{"id": v.ID, "revision": v.Revision, "data": data}).Where(sq.Eq{"id": v.ID, "revision": curRevision})
	res, err := d.exec(tx, q)
	if err != nil {
		v.Revision = curRevision
		return nil, errors.Wrap(err, "failed to update executor")
	}

	rows, err := res.RowsAffected()
	if err != nil {
		v.Revision = curRevision
		return nil, errors.Wrap(err, "failed to update executor")
	}

	if rows != 1 {
		v.Revision = curRevision
		return nil, idb.ErrConcurrent
	}

	return data, nil
}

func (d *DB) DeleteExecutor(tx *sql.Tx, id string) error {
	if err := d.deleteExecutorData(tx, id); err != nil {
		return errors.WithStack(err)
	}

	return d.deleteExecutorQ(tx, id)
}

func (d *DB) deleteExecutorData(tx *sql.Tx, id string) error {
	if _, err := tx.Exec("delete from executor where id = $1", id); err != nil {
		return errors.Wrap(err, "failed to delete executor")
	}

	return nil
}

func (d *DB) InsertOrUpdateExecutorTask(tx *sql.Tx, v *types.ExecutorTask) error {
	var err error
	if v.Revision == 0 {
		err = d.InsertExecutorTask(tx, v)
	} else {
		err = d.UpdateExecutorTask(tx, v)
	}

	return errors.WithStack(err)
}

func (d *DB) InsertExecutorTask(tx *sql.Tx, v *types.ExecutorTask) error {
	if v.Revision != 0 {
		return errors.Errorf("expected revision 0 got %d", v.Revision)
	}

	data, err := d.insertExecutorTaskData(tx, v)
	if err != nil {
		return errors.WithStack(err)
	}

	return d.insertExecutorTaskQ(tx, v, data)
}

func (d *DB) insertExecutorTaskData(tx *sql.Tx, v *types.ExecutorTask) ([]byte, error) {
	v.Revision = 1

	now := time.Now()
	v.SetCreationTime(now)
	v.SetUpdateTime(now)

	data, err := json.Marshal(v)
	if err != nil {
		v.Revision = 0
		return nil, errors.WithStack(err)
	}

	q := sb.Insert("executortask").Columns("id", "revision", "data").Values(v.ID, v.Revision, data)
	if _, err := d.exec(tx, q); err != nil {
		v.Revision = 0
		return nil, errors.Wrap(err, "failed to insert executortask")
	}

	return data, nil
}

// insertRawExecutorTaskData should be used only for import.
// It won't update object times.
func (d *DB) insertRawExecutorTaskData(tx *sql.Tx, v *types.ExecutorTask) ([]byte, error) {
	v.Revision = 1

	data, err := json.Marshal(v)
	if err != nil {
		v.Revision = 0
		return nil, errors.WithStack(err)
	}

	q := sb.Insert("executortask").Columns("id", "revision", "data").Values(v.ID, v.Revision, data)
	if _, err := d.exec(tx, q); err != nil {
		v.Revision = 0
		return nil, errors.Wrap(err, "failed to insert executortask")
	}

	return data, nil
}

func (d *DB) UpdateExecutorTask(tx *sql.Tx, v *types.ExecutorTask) error {
	data, err := d.updateExecutorTaskData(tx, v)
	if err != nil {
		return errors.WithStack(err)
	}

	return d.updateExecutorTaskQ(tx, v, data)
}

func (d *DB) updateExecutorTaskData(tx *sql.Tx, v *types.ExecutorTask) ([]byte, error) {
	if v.Revision < 1 {
		return nil, errors.Errorf("expected revision > 0 got %d", v.Revision)
	}

	curRevision := v.Revision
	v.Revision++

	v.SetUpdateTime(time.Now())

	data, err := json.Marshal(v)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	q := sb.Update("executortask").SetMap(map[string]interface{}{"id": v.ID, "revision": v.Revision, "data": data}).Where(sq.Eq{"id": v.ID, "revision": curRevision})
	res, err := d.exec(tx, q)
	if err != nil {
		v.Revision = curRevision
		return nil, errors.Wrap(err, "failed to update executortask")
	}

	rows, err := res.RowsAffected()
	if err != nil {
		v.Revision = curRevision
		return nil, errors.Wrap(err, "failed to update executortask")
	}

	if rows != 1 {
		v.Revision = curRevision
		return nil, idb.ErrConcurrent
	}

	return data, nil
}

func (d *DB) DeleteExecutorTask(tx *sql.Tx, id string) error {
	if err := d.deleteExecutorTaskData(tx, id); err != nil {
		return errors.WithStack(err)
	}

	return d.deleteExecutorTaskQ(tx, id)
}

func (d *DB) deleteExecutorTaskData(tx *sql.Tx, id string) error {
	if _, err := tx.Exec("delete from executortask where id = $1", id); err != nil {
		return errors.Wrap(err, "failed to delete executortask")
	}

	return nil
}
