package vessel

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type VesselRepo interface {
	GetAllVessel() ([]VesselData, error)
}

type DBVesselRepo struct {
}

func (dbVRepo DBVesselRepo) GetAllVessel() ([]VesselData, error) {

	vessels := []VesselData{
		VesselData{ID: 1, Naccs_code: "n1", Vessel_name: "v1", Owner_name: "o1", Modified_person_name: "m1"},
		VesselData{ID: 1, Naccs_code: "n2", Vessel_name: "v3", Owner_name: "o2", Modified_person_name: "m2"},
	}

	return vessels, nil
}

type PostGresDBVesselRepo struct {
}

var dbObj *sql.DB

func (pdb PostGresDBVesselRepo) DBInit(host string, port int, user string, password string, dbname string) {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	dbObj = db

	fmt.Println("Successfully connected!")
	return
}

func (pdb PostGresDBVesselRepo) GetAllVessel() ([]VesselData, error) {

	sqlStatement := `SELECT * FROM vessel LIMIT 10`
	rows, err := dbObj.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var vessels []VesselData
	for rows.Next() {
		var vessel VesselData
		err := rows.Scan(&vessel.ID, &vessel.Naccs_code, &vessel.Vessel_name, &vessel.Owner_name, &vessel.Modified_person_name, &vessel.Notes)
		if err != nil {
			panic(err)
		}
		vessels = append(vessels, vessel)
	}

	return vessels, nil
}
