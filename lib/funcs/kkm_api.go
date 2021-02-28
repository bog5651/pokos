package funcs

import (
	"database/sql"
	"fmt"
	"github.com/guark/guark/app"
	"github.com/mitchellh/mapstructure"
	"pokos/lib/database"
	"pokos/lib/types"
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

func UpdateKKM(c app.Context) (interface{}, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	raw := c.Get("kkm")
	var kkm types.KKM
	err = mapstructure.Decode(raw, &kkm)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`
		UPDATE kkm_registers 
		SET client_id = ?, cash_desk_id = ?, serial_number = ?, register_date = ?, ofd = ?, isExcise = ?, system_no = ?, type = ?, fn = ?, address = ?, end_date_fn = ?, end_date_ofd = ?, inspection_day_count = ?, comment = ?
		WHERE id = ?`,
		kkm.ClientId,
		kkm.ModelId,
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

	return GetKKMByID(db, kkm.ID)
}

func DeleteKKM(c app.Context) (interface{}, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	id := c.Get("id").(float64)
	_, err = db.Exec(`DELETE FROM kkm_registers WHERE id = ?`, int64(id))
	if err != nil {
		return nil, err
	}

	return "OK", nil
}

func GetKKMByID(db *sql.DB, id int64) (types.KKM, error) {
	var KKM types.KKM

	err := db.QueryRow(
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
				 LEFT JOIN cash_desk AS cd ON cd.id = k.cash_desk_id
		WHERE k.id = ? LIMIT 1`,
		id,
	).Scan(
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

	return KKM, err
}

func CreateKKM(c app.Context) (interface{}, error) {
	db, err := database.GetInstance()
	if err != nil {
		return nil, err
	}

	raw := c.Get("kkm")
	var kkm types.KKM
	err = mapstructure.Decode(raw, &kkm)
	if err != nil {
		return nil, err
	}
	c.App.Log.Error(fmt.Sprintf("got %+v", kkm))

	insertInfo, err := db.Exec(`
		INSERT INTO kkm_registers (client_id, cash_desk_id, serial_number, register_date, ofd, isExcise, system_no, type, fn, address, end_date_fn, end_date_ofd, inspection_day_count, comment) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		kkm.ClientId,
		kkm.ModelId,
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
	)
	if err != nil {
		return nil, err
	}

	id, err := insertInfo.LastInsertId()
	c.App.Log.Error(fmt.Sprintf("inserted id %d", id))
	if err != nil {
		return nil, err
	}

	return GetKKMByID(db, id)
}
