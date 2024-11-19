package repository

import (
	"fmt"
	"walking-information-storage-server-api/model"
)

func GetPedestrian() {

	var padeestrian model.Pedestrian

	id := "01F8VYXK67BGC1F9RP1E4S9YAK"
	err := db.QueryRow("SELECT * FROM pedestrians WHERE id = $1", id).Scan(&padeestrian.ID, &padeestrian.CreatedAt, &padeestrian.UpdatedAt, &padeestrian.DeletedAt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(padeestrian)
}
