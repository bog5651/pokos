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

func GetModelsKKM(c app.Context) (interface{}, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	models := make([]types.ModelKKM, 0)

	rows, err := db.Query(`SELECT id, name FROM cash_desk;`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var model types.ModelKKM
		err = rows.Scan(
			&model.ID,
			&model.Name,
		)
		if err != nil {
			c.App.Log.Warn("Error get model row: %s", err)
			continue
		}

		models = append(models, model)
	}
	return models, nil
}

func UpdateModelKKM(c app.Context) (interface{}, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	raw := c.Get("modelKkm")
	var model types.ModelKKM
	err = mapstructure.Decode(raw, &model)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`
		UPDATE cash_desk SET name =? WHERE id = ?`,
		model.Name,
		model.ID,
	)
	if err != nil {
		return nil, err
	}

	return GetModelsKKMByID(db, model.ID)
}

func DeleteModelKKM(c app.Context) (interface{}, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	id := c.Get("id").(float64)

	rows, err := db.Query(`SELECT id FROM kkm_registers WHERE cash_desk_id = ?`, int64(id))
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
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

	_, err = db.Exec(`DELETE FROM cash_desk WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}

	return "OK", nil
}

func GetModelsKKMByID(db *sql.DB, id int64) (types.ModelKKM, error) {
	var model types.ModelKKM

	err := db.QueryRow(`SELECT id, name FROM cash_desk WHERE id = ? LIMIT 1`, id).Scan(
		&model.ID,
		&model.Name,
	)
	return model, err
}

func CreateModelKKM(c app.Context) (interface{}, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	name := c.Get("name").(string)
	insertInfo, err := db.Exec(`
		INSERT INTO cash_desk (name) VALUES (?)`,
		name,
	)
	if err != nil {
		return nil, err
	}

	id, err := insertInfo.LastInsertId()
	if err != nil {
		return nil, err
	}

	return GetModelsKKMByID(db, id)
}
