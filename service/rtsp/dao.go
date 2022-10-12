package rtsp

import "github.com/pocketbase/pocketbase/daos"

var dao *daos.Dao

type RtspCameras struct {
	Name string
}

func FindAll() ([]RtspCameras, error) {
	rows := []RtspCameras{}
	collection, err := dao.FindCollectionByNameOrId("rtsp_cameras")
	if err != nil {
		return nil, err
	}
	dao.RecordQuery(collection).All(&rows)
	return rows, nil
}
