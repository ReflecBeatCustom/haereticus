package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	// import by gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/ReflecBeatCustom/haereticus/pkg/types"
)

// GetFumens [...]
func (server *HaereticusServer) GetFumens(writer http.ResponseWriter, body []byte) {
	// 1. Parse params
	req := &types.GetFumenRequest{}
	err := json.Unmarshal(body, req)
	if err != nil {
		glog.Errorf("Unmarshal body error: %v", err)
		http.Error(writer, http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	// 2. Get fumen result
	fumens, err := server.getFumens(req)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	// 3. Generate & send response
	response := types.GetFumenResponse{
		Header: &types.Header{
			ID:      req.ID,
			Jsonrpc: req.Jsonrpc,
			Method:  "GetFumens",
		},
		Result: &types.GetFumenResult{
			Data:  fumens,
			Total: len(fumens),
		},
	}
	responseBytes, err := json.Marshal(response)
	if err != nil {
		glog.Errorf("Marshal response error: %v", err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	responseOK(writer, responseBytes)
}

func (server *HaereticusServer) getFumens(request *types.GetFumenRequest) ([]*types.Songs, error) {
	fumens := make([]*types.Songs, 0)
	if server.dbClient == nil {
		return nil, fmt.Errorf("Error cmdb api db is nil")
	}

	title := fmt.Sprintf("%%%s%%", request.Params.Keyword)
	artist := "%"
	creator := "%"
	category := request.Params.Category
	offset := request.Params.Start
	limit := request.Params.Page
	if request.Params.Artist != "" {
		artist = request.Params.Artist
	}
	if request.Params.Creator != "" {
		creator = request.Params.Creator
	}
	if err := server.dbClient.Table("Songs").Where("Title LIKE ? AND Artist LIKE ? AND ChartAuthor LIKE ? AND Category = ?", title, artist, creator, category).Order("CreateTime DESC").Offset(offset).Limit(limit).Find(&fumens).Error; err != nil {
		return nil, err
	}

	return fumens, nil
}
