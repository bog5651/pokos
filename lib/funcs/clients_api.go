package funcs

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/guark/guark/app"
	"github.com/mitchellh/mapstructure"
	"pokos/lib/database"
	"pokos/lib/types"
	"pokos/lib/utils"
)

func GetClients(c app.Context) (interface{}, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	clients := make([]types.Client, 0)

	rows, err := db.Query(`SELECT id, name FROM clients;`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var client types.Client
		err = rows.Scan(
			&client.ID,
			&client.Name,
		)
		if err != nil {
			c.App.Log.Warn("Error get client row: %s", err)
			continue
		}

		clients = append(clients, client)
	}
	return clients, nil
}

func UpdateClient(c app.Context) (interface{}, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	raw := c.Get("client")
	var client types.Client
	err = mapstructure.Decode(raw, &client)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		UPDATE clients SET name = ? WHERE id = ?`,
		client.Name,
		client.ID,
	)
	if err != nil {
		return nil, err
	}

	return GetClientsByID(db, client.ID)
}

func DeleteClient(c app.Context) (interface{}, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	id := c.Get("id").(float64)

	rows, err := db.Query(`SELECT id FROM kkm_registers WHERE client_id = ?`, int64(id))
	if err != nil {
		if !errors.Is(sql.ErrNoRows, err) {
			return nil, err
		}
	} else {
		kkmListIds := make([]int64, 0)
		for rows.Next() {
			var id int64
			err := rows.Scan(&id)
			if err != nil {
				return nil, fmt.Errorf("Удаление запрещено. Есть Регистрации ККМ. Ошибка получения списка ID")
			}
			kkmListIds = append(kkmListIds, id)
		}
		if len(kkmListIds) > 0 {
			return nil, fmt.Errorf("Удаление запрещено. Есть Регистрации ККМ cо следующими ID - %s", utils.IntJoin(kkmListIds, ", "))
		}
	}

	_, err = db.Exec(`DELETE FROM clients WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}

	return "OK", nil
}

func GetClientsByID(db *sql.DB, id int64) (types.Client, error) {
	var client types.Client

	err := db.QueryRow(`SELECT id, name FROM clients WHERE id = ? LIMIT 1;`, id).Scan(
		&client.ID,
		&client.Name,
	)
	return client, err
}

func CreateClient(c app.Context) (interface{}, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	client := c.Get("name").(string)
	insertInfo, err := db.Exec(`
		INSERT INTO clients (name) VALUES (?)`,
		client,
	)
	if err != nil {
		return nil, err
	}

	id, err := insertInfo.LastInsertId()
	if err != nil {
		return nil, err
	}

	return GetClientsByID(db, id)
}
