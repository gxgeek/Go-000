package v1

import (
	"Go-000/Week04/internal/biz/tag_service"
	"Go-000/Week04/internal/pkg/app"
	"Go-000/Week04/internal/pkg/e"
	"Go-000/Week04/internal/pkg/logging"
	"Go-000/Week04/internal/pkg/qrcode"
	"Go-000/Week04/internal/pkg/setting"
	"Go-000/Week04/internal/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func GetTags(c *gin.Context)  {
	appG := app.Gin{C: c}

	queryName := c.Query("name")
	if queryName != "" {
	}
	var state = -1
	if  queryState := c.Query("state"); queryState != ""{
		state = com.StrTo(queryState).MustInt()
	}
	tagService := tag_service.Tag{Name: queryName, State: state,
		PageSize: setting.AppSetting.PageSize, PageNum: util.GetPage(c)}

	lists, err := tagService.GetAll()
	if err != nil {
		appG.FastReturn(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
		return
	}
	count := tagService.Count()
	data := make(map[string]interface{})
	data["lists"] = lists
	data["total"] = count
	appG.FastReturn(http.StatusOK, e.SUCCESS,data)
}

type AddTagForm struct {
	Name      string `form:"name" valid:"Required;MaxSize(100)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" valid:"Range(0,1)"`
}

// @Summary 新增文章标签
// @Accept  json
// @Produce  json
// @Param   name     query    string     true  "Name"
// @Param   state     query   int     true   "State"
// @Param created_by query   string     true   "CreatedBy"
// @Param token query   string     true   "Token"
// @Success 200 {string} json  "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/ [post]
func AddTag(c *gin.Context)  {
	appG := app.Gin{C: c}

	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state","0")).MustInt()
	createdBy := c.Query("created_by")

	var form AddTagForm
	err := app.BindAndVailValue(c, &form)
	if err != nil {
		appG.FastReturnCode(http.StatusOK, e.INVALID_PARAMS)
		return
	}
	tagService := tag_service.Tag{Name: name, State: state,CreatedBy: createdBy}

	if tagService.ExistByName() {
		appG.FastReturnCode(http.StatusOK, e.ERROR_EXIST_TAG)
		return
	}
	tagService.Add()
	appG.FastReturnCode(http.StatusOK, e.SUCCESS)
}

// @Summary 更新文章标签
// @Accept  json
// @Produce  json
// @Param   id     path    int     true  "ID"
// @Param   name     query    string     true  "Name"
// @Param   state     query   int     true   "State"
// @Param modified_by query   string     true   "ModifiedBy"
// @Param token query   string     true   "Token"
// @Success 200 {string} json  "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/{id} [PUT]
func UpdateTags(c *gin.Context)  {
	appG := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	var state  = -1
	if paramState := c.Query("state"); paramState != "" {
		state = com.StrTo(paramState).MustInt()
	}

	valid := validation.Validation{}
	valid.Required(id,"id").Message("id 不能为空")
	//valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if valid.HasErrors() {
		logging.Info(valid.Errors)
		appG.FastReturnCode(http.StatusOK, code)
		return
	}
	tagService := tag_service.Tag{
		ID:         id,
		Name:       name,
		ModifiedBy: modifiedBy,
		State:      state,
	}


	if tagService.ExistByID() {
		appG.FastReturnCode(http.StatusOK, e.ERROR_EXIST_TAG)
		return
	}
	tagService.Edit()
	appG.FastReturnCode(http.StatusOK, e.SUCCESS)

}
func DeleteTag(c *gin.Context)  {
	appG := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Required(id,"id").Message("id 不能为空")
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		logging.Info(valid.Errors)
		appG.FastReturnCode(http.StatusOK, e.INVALID_PARAMS)
		return
	}
	tagService := tag_service.Tag{ID: id}

	if !tagService.ExistByID() {
		appG.FastReturnCode(http.StatusOK, e.ERROR_NOT_EXIST_TAG)
		return
	}
	tagService.Delete()
	appG.FastReturnCode(http.StatusOK, e.SUCCESS)
}


const (
	QRCODE_URL = "https://github.com/EDDYCJY/blog#gin%E7%B3%BB%E5%88%97%E7%9B%AE%E5%BD%95"
)

func GenerateArticlePoster(c *gin.Context)  {
	appG := app.Gin{c}
	qrc := qrcode.NewQrCode(QRCODE_URL, 300, 300, qr.M, qr.Auto)
	path := qrcode.GetQrCodeFullPath()
	_, _, err := qrc.Encode(path)
	if err != nil {
		appG.FastReturn(http.StatusOK, e.ERROR, nil)
		return
	}

	appG.FastReturn(http.StatusOK, e.SUCCESS, nil)

}
