package store

import (
	"context"
	"database/sql"
	"fmt"
	"todolist/pkg/structs"

	"github.com/jmoiron/sqlx"
)

type sqlStore struct {
	db *sqlx.DB
}

type sqlStoreTxn struct {
	txn *sqlx.Tx
}

func NewSQLStore(db *sqlx.DB) Store {
	return &sqlStore{
		db: db,
	}
}

func (s *sqlStore) Update(action func(tx Txn) error) error {
	dbtx, err := s.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			_ = dbtx.Rollback()
			panic(r)
		}
	}()

	tx := &sqlStoreTxn{
		txn: dbtx,
	}

	err = action(tx)
	if err != nil {
		_ = dbtx.Rollback()
		return err
	}

	return dbtx.Commit()
}

func readRecord(rows *sql.Rows, record *structs.TodoItem) error {
	return rows.Scan(
		&record.Id,
		&record.Item,
		&record.Item_order,
	)
}

func (tx *sqlStoreTxn) DbTx() interface{} {
	return tx.txn
}

func (tx *sqlStoreTxn) Add(ctx context.Context, record *structs.TodoItem) error {
	var maxOrder int
	err := tx.txn.QueryRowContext(ctx,
		`SELECT COALESCE(MAX(ITEM_ORDER),0) FROM TODOLIST`).Scan(&maxOrder)
	if err != nil {
		return err
	}
	maxOrder++
	_, err = tx.txn.ExecContext(ctx,
		tx.txn.Rebind("INSERT INTO TODOLIST(ID, ITEM,ITEM_ORDER) VALUES(?, ?, ?)"),
		record.Id,
		record.Item,
		maxOrder,
	)
	return err
}

func (tx *sqlStoreTxn) Delete(ctx context.Context, id string) error {
	var itemOrder int
	err := tx.txn.QueryRowContext(ctx, `SELECT ITEM_ORDER FROM TODOLIST WHERE ID=?`, id).Scan(&itemOrder)
	if err != nil {
		return err
	}

	result, err := tx.txn.ExecContext(ctx, tx.txn.Rebind("DELETE FROM TODOLIST WHERE ID=?"), id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("unknown id")
	}

	_, err = tx.txn.ExecContext(ctx,
		`UPDATE TODOLIST
		SET ITEM_ORDER = ITEM_ORDER-1
		WHERE ITEM_ORDER>?`,
		itemOrder)
	if err != nil {
		return err
	}

	return nil
}

func (tx *sqlStoreTxn) Update(ctx context.Context, record *structs.TodoItem) error {
	result, err := tx.txn.ExecContext(ctx,
		tx.txn.Rebind(`UPDATE TODOLIST SET
			ITEM=?
			WHERE ID=?`),
		record.Item,
		record.Id,
	)

	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("unknown id")
	}
	return nil
}

func (tx *sqlStoreTxn) Get(ctx context.Context, id string, item *structs.TodoItem) error {
	queryStmt := "SELECT ID, ITEM, ITEM_ORDER FROM TODOLIST WHERE ID=?"

	rows, err := tx.txn.QueryContext(ctx, tx.txn.Rebind(queryStmt), id)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return fmt.Errorf("unknown id")
	}

	if err := readRecord(rows, item); err != nil {
		return err
	}

	return nil
}

func (tx *sqlStoreTxn) List(ctx context.Context, items *structs.TodoItemList) error {
	queryStmt := "SELECT ID, ITEM, ITEM_ORDER FROM TODOLIST"

	rows, err := tx.txn.QueryContext(ctx, tx.txn.Rebind(queryStmt))
	if err != nil {
		return err
	}
	defer rows.Close()

	items.Items = make([]structs.TodoItem, 0)
	var record structs.TodoItem
	for rows.Next() {
		if err := readRecord(rows, &record); err != nil {
			return err
		}
		items.Items = append(items.Items, record)
		items.Count++
	}
	return nil
}

func (tx *sqlStoreTxn) ReorderItem(ctx context.Context, id string, newOrder int) error {
	var currentOrder int
	err := tx.txn.QueryRowContext(ctx, `SELECT ITEM_ORDER FROM TODOLIST WHERE ID=?`, id).Scan(&currentOrder)
	if err != nil {
		return err
	}

	if currentOrder > newOrder {
		_, err := tx.txn.ExecContext(ctx,
			`UPDATE TODOLIST
			SET ITEM_ORDER = ITEM_ORDER+1
			WHERE ITEM_ORDER>=? AND ITEM_ORDER <?`,
			newOrder, currentOrder)
		if err != nil {
			return err
		}
	} else if newOrder > currentOrder {
		_, err := tx.txn.ExecContext(ctx,
			`UPDATE TODOLIST
			SET ITEM_ORDER = ITEM_ORDER-1
			WHERE ITEM_ORDER>? AND ITEM_ORDER <=?`,
			currentOrder, newOrder)
		if err != nil {
			return err
		}
	}

	_, err = tx.txn.ExecContext(ctx,
		`UPDATE TODOLIST
		SET ITEM_ORDER=?
		WHERE ID=?`,
		newOrder,
		id)
	if err != nil {
		return err
	}
	return nil
}
