package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"

	"github.com/haereticus/pkg/types"
)

// GetFumen ...
func (server *HaereticusServer) GetFumens(writer http.ResponseWriter, body []byte) {
	// 1. Parse params
	var req types.GetFumenRequest
	err := json.Unmarshal(body, &req)
	if err != nil {
		glog.Errorf("Unmarshal body error: %v", err)
		http.Error(writer, http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	// 2. Get fumen result
	fumens, err := server.getFumens(req.Params.Keyword, req.Params.Page, req.Params.Start)
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
		Fumens: fumens,
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

func (server *HaereticusServer) getFumens(keyword string, page, startIndex int) ([]*types.Fumen, error) {
	fumens := make([]*types.Fumen, 0)

	var sql string
	if keyword == "" {
		sql = fmt.Sprintf("SELECT fumen_id, artist_name, music_name, bpm_max, bpm_min, basic_level, medium_level, hard_level FROM fumen LIMIT %d OFFSET %d ", page, startIndex)
	} else {
		sql = fmt.Sprintf("SELECT fumen_id, artist_name, music_name, bpm_max, bpm_min, basic_level, medium_level, hard_level FROM fumen WHERE artist_name LIKE '%%%s%%' OR music_name LIKE '%%%s%%' LIMIT %d OFFSET %d", keyword, keyword, page, startIndex)
	}
	rows, err := server.dbClient.Query(sql)
	if err != nil {
		glog.Errorf("Get fumen metadata from db err: %v", err)
		return fumens, err
	}
	defer rows.Close()

	for rows.Next() {
		fumen := &types.Fumen{}
		err := rows.Scan(&fumen.FumenID,
			&fumen.ArtistName,
			&fumen.MusicName,
			&fumen.BPMMax,
			&fumen.BPMMin,
			&fumen.BasicLevel,
			&fumen.MediumLevel,
			&fumen.HardLevel)
		if err != nil {
			glog.Errorf("Get fumen metadata from db err: %v", err)
			return fumens, err
		}
		fumens = append(fumens, fumen)
	}

	return fumens, nil
}
