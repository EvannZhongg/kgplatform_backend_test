package materials

import (
	"kgplatform-backend/api/materials"
	materialsLogic "kgplatform-backend/internal/logic/materials"
)

type ControllerV1 struct {
	materials *materialsLogic.Materials
}

func NewV1() materials.IMaterialsV1 {
	return &ControllerV1{
		materials: materialsLogic.New(),
	}
}
