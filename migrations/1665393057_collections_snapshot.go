package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

// Auto generated migration with the most recent collections configuration.
func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "z96ewo1f8xwblsi",
				"created": "2022-10-06 06:16:29.943",
				"updated": "2022-10-06 06:22:02.894",
				"name": "channels",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "xsorlwhu",
						"name": "url",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"listRule": "",
				"viewRule": "",
				"createRule": "",
				"updateRule": "",
				"deleteRule": ""
			},
			{
				"id": "n7xogpeawxliwqa",
				"created": "2022-10-06 06:20:46.408",
				"updated": "2022-10-07 04:14:47.938",
				"name": "groups",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "gohcoafn",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "xhrh0wkc",
						"name": "subGroups",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 50,
							"collectionId": "n7xogpeawxliwqa",
							"cascadeDelete": false
						}
					},
					{
						"system": false,
						"id": "lfrr0ul9",
						"name": "isRoot",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					}
				],
				"listRule": "",
				"viewRule": "",
				"createRule": "",
				"updateRule": "",
				"deleteRule": ""
			},
			{
				"id": "systemprofiles0",
				"created": "2022-10-08 03:47:58.400",
				"updated": "2022-10-08 03:47:58.402",
				"name": "profiles",
				"system": true,
				"schema": [
					{
						"system": true,
						"id": "pbfielduser",
						"name": "userId",
						"type": "user",
						"required": true,
						"unique": true,
						"options": {
							"maxSelect": 1,
							"cascadeDelete": true
						}
					},
					{
						"system": false,
						"id": "pbfieldname",
						"name": "name",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "pbfieldavatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpg",
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif"
							],
							"thumbs": null
						}
					}
				],
				"listRule": "userId = @request.user.id",
				"viewRule": "userId = @request.user.id",
				"createRule": "userId = @request.user.id",
				"updateRule": "userId = @request.user.id",
				"deleteRule": null
			},
			{
				"id": "ls3av7rrzo7xwy2",
				"created": "2022-10-10 08:09:30.156",
				"updated": "2022-10-10 08:41:31.434",
				"name": "rtsp_cameras",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "h9nmwgcy",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": true,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "26goun7y",
						"name": "host",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "e1qekasc",
						"name": "port",
						"type": "number",
						"required": true,
						"unique": false,
						"options": {
							"min": 0,
							"max": 65535
						}
					},
					{
						"system": false,
						"id": "25yp8uh3",
						"name": "user",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "7zuzkhbf",
						"name": "password",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "961yfef7",
						"name": "description",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "5calrpnl",
						"name": "status",
						"type": "select",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"values": [
								"online",
								"offline"
							]
						}
					}
				],
				"listRule": "",
				"viewRule": "",
				"createRule": "",
				"updateRule": "",
				"deleteRule": ""
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		// no revert since the configuration on the environment, on which
		// the migration was executed, could have changed via the UI/API
		return nil
	})
}
