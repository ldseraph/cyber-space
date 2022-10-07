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
				"id": "systemprofiles0",
				"created": "2022-09-03 16:55:51.651",
				"updated": "2022-09-03 16:55:51.651",
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
				"id": "9oAyHCAGIf3xCBD",
				"created": "2022-10-01 03:47:57.116",
				"updated": "2022-10-06 12:09:39.443",
				"name": "cameras",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "ufadguuq",
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
						"id": "tsis4o8y",
						"name": "serialNumber",
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
						"id": "dm9grbxk",
						"name": "status",
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
						"id": "acrgkpud",
						"name": "host",
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
						"id": "v1xczaum",
						"name": "rtspUsername",
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
						"id": "bjtwddto",
						"name": "rtspPwd",
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
						"id": "xguvbx5s",
						"name": "rtspPort",
						"type": "number",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null
						}
					},
					{
						"system": false,
						"id": "738x7avp",
						"name": "channels",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "z96ewo1f8xwblsi",
							"cascadeDelete": false
						}
					},
					{
						"system": false,
						"id": "fknp6o24",
						"name": "groups",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"collectionId": "n7xogpeawxliwqa",
							"cascadeDelete": false
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
