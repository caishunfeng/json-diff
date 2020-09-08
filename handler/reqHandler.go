package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"json-diff/model"
	"net/http"
	"os"
	"os/exec"
)

type ReqHandler struct {
}

func NewReqHandler() *ReqHandler {
	return &ReqHandler{}
}

func (this *ReqHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := new(model.DiffReq)

	s, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("read request body fail，err:%s", err.Error())))
		return
	}

	err = json.Unmarshal(s, req)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("json unmarshal fail，err:%s", err.Error())))
		return
	}

	liveResponse, err := callAPI(req.EndPoints.Live, req.QueryParams, req.HeaderParams)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("call live api fail，err:%s", err.Error())))
		return
	}
	err = writeJSONStringToFile("./log/live/"+req.TraceId, liveResponse)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("save live api result fail，err:%s", err.Error())))
		return
	}

	localResponse, err := callAPI(req.EndPoints.Local, req.QueryParams, req.HeaderParams)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("call local api fail，err:%s", err.Error())))
		return
	}
	err = writeJSONStringToFile("./log/local/"+req.TraceId, localResponse)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("save local api result fail，err:%s", err.Error())))
		return
	}

	var out []byte
	if isWindows() {
		out, err = exec.Command("cmd", "/c", "node diff.js", req.TraceId).Output()
	} else {
		out, err = exec.Command("sh", "-c", "node diff.js", req.TraceId).Output()
	}

	if err != nil {
		w.Write([]byte(fmt.Sprintf("exec diff fail，msg:%s,err:%s", string(out), err.Error())))
		return
	}

	out, err = ioutil.ReadFile("./log/diff/" + req.TraceId)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("read diff result fail，msg:%s,err:%s", string(out), err.Error())))
		return
	}

	os.Remove("./log/live/" + req.TraceId)
	os.Remove("./log/local/" + req.TraceId)
	os.Remove("./log/diff/" + req.TraceId)

	w.Write(out)
}
