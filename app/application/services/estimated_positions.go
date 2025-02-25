package application

import (
	"encoding/json"
	"fmt"
)

type EstimationPosition struct {
	ID int     `json:"id"`
	X  float32 `json:"x"`
	Y  float32 `json:"y"`
}

func main() {
	// estimated_positionインスタンスを作成
	estimated_position := EstimationPosition{
		ID: 1,
		X:  30,
		Y:  123,
	}

	// estimated_positionインスタンスをjsonに変換
	b, err := json.Marshal(estimated_position)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}
