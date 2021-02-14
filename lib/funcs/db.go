package funcs

import (
	"database/sql"
	"errors"
	"fmt"
	"pokos/lib/database"
	"pokos/lib/types"
	"pokos/lib/utils"

	"github.com/guark/guark/app"
)

func GetKKM(c app.Context) (interface{}, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	KKMs := make([]types.KKM, 0)

	rows, err := db.Query(
		`SELECT k.id,
			   c.name,
			   cd.name,
			   k.client_id,
			   k.cash_desk_id,
			   k.serial_number,
			   k.register_date,
			   k.ofd,
			   k.isExcise,
			   k.system_no,
			   k.type,
			   k.fn,
			   k.address,
			   k.end_date_fn,
			   k.end_date_ofd,
			   k.inspection_day_count,
			   k.comment
		FROM kkm_registers k
				 LEFT JOIN clients AS c ON c.id = k.client_id
				 LEFT JOIN cash_desk AS cd ON cd.id = k.cash_desk_id`,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var KKM types.KKM
		err = rows.Scan(
			&KKM.ID,
			&KKM.ClientName,
			&KKM.ModelName,
			&KKM.ClientId,
			&KKM.ModelId,
			&KKM.SerialNumber,
			&KKM.RegisterDate,
			&KKM.OFD,
			&KKM.IsExcise,
			&KKM.SystemNO,
			&KKM.Type,
			&KKM.FN,
			&KKM.Address,
			&KKM.EndDateFN,
			&KKM.EndDateOFD,
			&KKM.InspectionDayCount,
			&KKM.Comment,
		)
		if err != nil {
			c.App.Log.Warn("Error get KKM row: %s", err)
		}

		KKMs = append(KKMs, KKM)
	}

	return KKMs, nil
}

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

func EditKKM(c app.Context) (interface{}, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	kkm := c.Get("kkm").(types.KKM)
	_, err = db.Exec(`
		UPDATE kkm_registers 
		SET client_id = ?, cash_desk_id = ?, serial_number = ?, register_date = ?, ofd = ?, isExcise = ?, system_no = ?, type = ?, fn = ?, address = ?, end_date_fn = ?, end_date_ofd = ?, inspection_day_count = ?, comment = ?
		WHERE id = ?`,
		kkm.ClientId,
		kkm.SerialNumber,
		kkm.RegisterDate,
		kkm.OFD,
		kkm.IsExcise,
		kkm.SystemNO,
		kkm.Type,
		kkm.FN,
		kkm.Address,
		kkm.EndDateFN,
		kkm.EndDateOFD,
		kkm.InspectionDayCount,
		kkm.Comment,
		kkm.ID,
	)
	if err != nil {
		return nil, err
	}

	return "OK", nil
}

func EditClient(c app.Context) (interface{}, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	client := c.Get("client").(types.Client)
	_, err = db.Exec(`
		UPDATE clients SET name =? WHERE id = ?`,
		client.Name,
		client.ID,
	)
	if err != nil {
		return nil, err
	}

	return "OK", nil
}

func EditModelKKM(c app.Context) (interface{}, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	model := c.Get("modelKkm").(types.ModelKKM)
	_, err = db.Exec(`
		UPDATE cash_desk SET name =? WHERE id = ?`,
		model.Name,
		model.ID,
	)
	if err != nil {
		return nil, err
	}

	return "OK", nil
}

func DeleteKKM(c app.Context) (interface{}, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	id := c.Get("id").(int64)
	_, err = db.Exec(`DELETE FROM kkm_registers WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}

	return "OK", nil
}

func DeleteClient(c app.Context) (interface{}, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	id := c.Get("id").(int64)

	rows, err := db.Query(`SELECT id FROM kkm_registers WHERE client_id = ?`, id)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return nil, err
	}
	if !errors.Is(sql.ErrNoRows, err){
		kkmListIds := make([]int64, 0)
		for rows.Next() {
			var id int64
			err := rows.Scan(&id)
			if err != nil {
				return nil, fmt.Errorf("Удаление запрещено. Есть Регистрации ККМ. Ошибка получения списка ID")
			}
			kkmListIds = append(kkmListIds, id)
		}
		return nil, fmt.Errorf("Удаление запрещено. Есть Регистрации ККМ cо следующими ID - %s", utils.IntJoin(kkmListIds, ", "))
	}

	_, err = db.Exec(`DELETE FROM clients WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}

	return "OK", nil
}

func DeleteModelKKM(c app.Context) (interface{}, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	id := c.Get("id").(int64)

	rows, err := db.Query(`SELECT id FROM kkm_registers WHERE cash_desk_id = ?`, id)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return nil, err
	}
	if !errors.Is(sql.ErrNoRows, err){
		kkmListIds := make([]int64, 0)
		for rows.Next() {
			var id int64
			err := rows.Scan(&id)
			if err != nil {
				return nil, fmt.Errorf("Удаление запрещено. Есть Регистрации ККМ. Ошибка получения списка ID")
			}
			kkmListIds = append(kkmListIds, id)
		}
		return nil, fmt.Errorf("Удаление запрещено. Есть Регистрации ККМ cо следующими ID - %s", utils.IntJoin(kkmListIds, ", "))
	}

	_, err = db.Exec(`DELETE FROM cash_desk WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}

	return "OK", nil
}
