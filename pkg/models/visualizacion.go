package models

import (
	"database/sql"
	"fmt"
)

type Visualizacion struct {
	Id_movie int `json:"id_movie"`
	Contador int `json:"visualizaciones"`
}

func GetVisualizaciones(db *sql.DB) ([]*Visualizacion, error) {

	var visualizaciones []*Visualizacion

	query := `SELECT * FROM visualizaciones 
					ORDER BY visualizaciones DESC
					LIMIT 3 
					`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error en query rows: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var row Visualizacion
		if err := rows.Scan(&row.Id_movie, &row.Contador); err != nil {
			return nil, fmt.Errorf("error haciendo scan a BD: %v", err)
		}
		fmt.Printf("Row: %+v\n", row)
		visualizaciones = append(visualizaciones, &row)
	}

	return visualizaciones, nil
}
