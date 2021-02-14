package funcs

import (
	"pokos/lib/database"
	"pokos/lib/types"

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

func GetModelsKKM(c app.Context) (interface{}, error) {db, err := database.GetInstance()
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
