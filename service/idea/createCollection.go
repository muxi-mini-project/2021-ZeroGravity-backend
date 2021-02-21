package idea

import (
	"github.com/2021-ZeroGravity-backend/model"
)

//CreateCollection is used to create collection
func CreateCollection(id, uid int) error {

	var collection *model.CollectionModel

	collection = &model.CollectionModel{

		CollectorId: uid,
		IdeaId:      id,
	}

	return collection.Create()

}
