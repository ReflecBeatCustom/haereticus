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

// GetPacks [...]
func (server *HaereticusServer) GetPacks(writer http.ResponseWriter, body []byte) {
	// 1. Parse params
	req := &types.GetPackRequest{}
	err := json.Unmarshal(body, req)
	if err != nil {
		glog.Errorf("Unmarshal body error: %v", err)
		http.Error(writer, http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	// 2. Get fumen result
	packs, err := server.getPacks(req)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	// 3. Generate & send response
	response := types.GetPackResponse{
		Header: &types.Header{
			ID:      req.ID,
			Jsonrpc: req.Jsonrpc,
			Method:  "GetPacks",
		},
		Result: &types.GetPackResult{
			Data:  packs,
			Total: len(packs),
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

func (server *HaereticusServer) getPacks(request *types.GetPackRequest) ([]*types.Packs, error) {
	packs := make([]*types.Packs, 0)
	if server.dbClient == nil {
		return nil, fmt.Errorf("Error cmdb api db is nil")
	}

	title := fmt.Sprintf("%%%s%%", request.Params.Keyword)
	category := request.Params.Category
	offset := request.Params.Start
	limit := request.Params.Page
	if err := server.dbClient.Table("Packs").Where("Title LIKE ? AND Category = ?", title, category).Order("CreateTime DESC").Offset(offset).Limit(limit).Find(&packs).Error; err != nil {
		return nil, err
	}

	return packs, nil
}
