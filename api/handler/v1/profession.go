package v1

import "github.com/gin-gonic/gin"

// Create Profession godoc
// @ID create-profession
// @Router /v1/profession [POST]
// @Summary create profession
// @Description Create Profession
// @Tags profession
// @Accept json
// @Produce json
// @Param profession body models.CreateProfession true "profession"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateProfession(c *gin.Context) {

}

// Get Profession godoc
// @ID get-profession
// @Router /v1/profession/{profession_id} [GET]
// @Summary get profession
// @Description Get Profession
// @Tags profession
// @Accept json
// @Produce json
// @Param profession_id path string true "profession_id"
// @Success 200 {object} models.ResponseModel{data=models.Profession} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetProfession(c *gin.Context) {

}

// Get All Profession godoc
// @ID get-all-profession
// @Router /v1/profession [GET]
// @Summary get all profession
// @Description Get All Profession
// @Tags profession
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} models.ResponseModel{data=models.GetAllProfessionResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllProfessions(c *gin.Context) {

}
