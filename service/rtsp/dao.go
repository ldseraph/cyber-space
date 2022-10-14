package rtsp

import "github.com/pocketbase/pocketbase/daos"

var dao *daos.Dao

// Camera is rtsp camera model
type Camera struct {
	Name string
}

// FindAll find all rtsp camera model
func FindAll() ([]Camera, error) {
	rows := []Camera{}
	collection, err := dao.FindCollectionByNameOrId("rtsp_cameras")
	if err != nil {
		return nil, err
	}
	dao.RecordQuery(collection).All(&rows)
	return rows, nil
}
