package vessel

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type VesselRepo interface {
	GetAllVessel() ([]VesselData, error)
	GetVesselByNaccsCode(vesselID string) (VesselData, error)
	AddVessel(vsl *VesselData)
	UpdateVessel(vsl *VesselData) (res bool)
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

func (dbVRepo DBVesselRepo) GetVesselByNaccsCode(vesselID string) (VesselData, error) {

	vessels := []VesselData{
		VesselData{ID: 1, Naccs_code: "n1", Vessel_name: "v1", Owner_name: "o1", Modified_person_name: "m1"},
		VesselData{ID: 1, Naccs_code: "n2", Vessel_name: "v3", Owner_name: "o2", Modified_person_name: "m2"},
	}

	return vessels[0], nil
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

func (pdb PostGresDBVesselRepo) CreateTable() {

	// Create the table if it doesn't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS vessel (
		id SERIAL,
		naccs_code text PRIMARY KEY,
		vessel_name TEXT,
		owner_name TEXT,
		modified_person_name TEXT,
		notes TEXT
	  );`
	_, err := dbObj.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB created!")
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

func (pdb PostGresDBVesselRepo) GetVesselByNaccsCode(naccsCode string) (VesselData, error) {
	fmt.Printf("\nGiven naccs code : %s", naccsCode)
	sqlStatement := `SELECT naccs_code, vessel_name FROM vessel WHERE naccs_code=$1;`
	//sqlStatement := `SELECT naccs_code, vessel_name FROM vessel where naccs_code = $1`
	var vessel_name string
	var naccs_code string
	var VesselData VesselData

	row := dbObj.QueryRow(sqlStatement, naccsCode)
	switch err := row.Scan(&naccs_code, &vessel_name); err {
	case sql.ErrNoRows:
		fmt.Println("No row returned")
	case nil:
		fmt.Println("Matching data returned")
		VesselData.Naccs_code = naccs_code
		VesselData.Vessel_name = vessel_name
	default:
		fmt.Printf("\nPanic no rows returned with %s", sqlStatement)
	}

	return VesselData, nil
}

func (pdb PostGresDBVesselRepo) AddVessel(vsl *VesselData) {
	fmt.Printf("\n New vessel name : %s", vsl.Vessel_name)
	sqlStatement := `insert into vessel (vessel_name, naccs_code, owner_name, modified_person_name, notes) VALUES ($1, $2, $3, $4, $5) returning id;`

	id := 0
	err := dbObj.QueryRow(sqlStatement, vsl.Vessel_name, vsl.Naccs_code, vsl.Owner_name, vsl.Modified_person_name, vsl.Notes).Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("New record ID is:", id)
}

func (pdb PostGresDBVesselRepo) UpdateVessel(vsl *VesselData) (res bool) {
	fmt.Printf("\n Update vessel name : %s", vsl.Vessel_name)

	vslData, err := pdb.GetVesselByNaccsCode(vsl.Naccs_code)
	if err != nil {
		fmt.Printf("\n Error in search naccs code : %s", vsl.Naccs_code)
		panic(err)
	}
	if vslData.Naccs_code == "" {
		fmt.Printf("\n No vessel found for Given naccs code : %s", vslData.Naccs_code)
		return false
	}

	sqlStatement := `update vessel set vessel_name=$1, owner_name=$2 where naccs_code=$3;`

	_, err1 := dbObj.Exec(sqlStatement, vsl.Vessel_name, vsl.Owner_name, vsl.Naccs_code)
	if err1 != nil {
		panic(err)
	}

	fmt.Println("Vessel record is update for naccs code :", vsl.Naccs_code)
	return true
}
