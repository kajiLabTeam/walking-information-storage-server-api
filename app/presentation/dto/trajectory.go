package dto_presenttion

type GetTrajectoriesResponse struct {
	Floor        Floor        `json:"floor"`
	Trajectories []Trajectory `json:"trajectories"`
}

type Position struct {
	ID       string  `json:"id"`
	X        float32 `json:"x"`
	Y        float32 `json:"y"`
	WalkedAt string  `json:"walked_at"`
}

type Trajectory struct {
	ID        string     `json:"id"`
	Estimated []Position `json:"estimated"`
	Correct   []Position `json:"correct"`
}

type Floor struct {
	ID          string `json:"id"`
	MapImageURL string `json:"map_image_url"`
}
