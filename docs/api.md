# API設計

## 最終目標
フロアに溜まった推定軌跡及び正解軌跡を取得するAPI

<img width="677" alt="image.png (37.8 kB)" src="https://img.esa.io/uploads/production/attachments/13979/2024/11/22/168726/72f077cd-764c-469b-8640-875950691fd1.png">

# API設計 
## フロアに溜まった推定軌跡及び正解軌跡を取得するAPI
例.14号館に溜まった推定座標A,B,Cを取得

### reqest
- building_id (建物ID) ・・・14号館
- floor_id(フロアID)・・・5階

### 欲しい情報 (JSON)
- estimation_potisions(推定軌跡)
- correct_potisions(正解軌跡)


## 正解軌跡が存在するとき
### GET `api/buildings/${building_id}/floors/${floor_id}/trajectories?is_correct_included=true&count=100`
#### response
```json
{
	"floor_map_image": {
		"url": "署名付きURL"
	},
	"estimation_trajectories": [
		{
			"estimation_trajectories_id": 1,
			"estimation_position": [
				{
					"id": 1,
					"x": 36.1834166,
					"y": 138.110467,
					"walked_at": 1560000000
				},
				{
					"id": 2,
					"x": 35.1834166,
					"y": 137.110467,
					"walked_at": 1560000000
				}
			]
		},
		{
			"estimation_trajectories_id": 2,
			"estimation_position": [
				{
					"id": 3,
					"x": 35.1834166,
					"y": 137.110467,
					"walked_at": 1560000000
				},
				{
					"id": 4,
					"x": 35.1834166,
					"y": 137.110467,
					"walked_at": 1560000000
				}
			]
		}
	],
	"correct_trajectories": [
		{
			"correct_trajectories_id": 1,
			"correct_position": [
				{
					"id": 1,
					"x": 30.1834166,
					"y": 137.110467
                    "walked_at": 1560000000
				},
				{
					"id": 2,
					"x": 31.1834166,
					"y": 137.110467
                    "walked_at": 1560000000
				}
			]
		},
		{
			"correct_trajectories_id": 2,
			"correct_position": [
				{
					"id": 3,
					"x": 32.18341661,
					"y": 137.110467
                    "walked_at": 1560000000
				},
				{
					"id": 4,
					"x": 33.183416,
					"y": 137.110467
                    "walked_at": 1560000000
				}
			]
		}
	]
}

```

### GET `api/buildings/${building_id}/floors/${floor_id}/trajectories?is_correct_included=false&count=100`
### - 正解軌跡が存在しない場合
#### response
```json
{
	"floor_map_image": {
		"url": "署名付きURL"
	},
	"estimation_trajectories": [
		{
			"estimation_trajectories_id": 1,
			"estimation_position": [
				{
					"id": 1,
					"x": 36.1834166,
					"y": 138.110467,
					"walked_at": 1560000000
				},
				{
					"id": 2,
					"x": 35.1834166,
					"y": 137.110467,
					"walked_at": 1560000000
				}
			]
		},
		{
			"estimation_trajectories_id": 2,
			"estimation_position": [
				{
					"id": 3,
					"x": 35.1834166,
					"y": 137.110467,
					"walked_at": 1560000000
				},
				{
					"id": 4,
					"x": 35.1834166,
					"y": 137.110467,
					"walked_at": 1560000000
				}
			]
		}
	],
	"correct_trajectories": []
}

```

"correct_trajectories": []
正解軌跡を空[]で返す

## メモ
<img width="1057" alt="ER図_addメモ.png (99.8 kB)" src="https://img.esa.io/uploads/production/attachments/13979/2025/02/10/168726/b056b031-731d-4a5b-a3fb-0799e13a8d83.png">


### `pedestrians`
- 軌跡蓄積システムを使う歩行者
### `trajectories`
- **歩行軌跡**を管理
- 複数の正解座標・推定座標を管理
### `waking_samples`
歩行サンプル
### `estimation_position`
-  推定座標
### `correct_potions`
- 正解座標

## `buildings`
- 建物（14号館、一号館など）のこと
- 複数のフロアから構成される

### `floors`
- 建物の一階層のこと
