package httpHandle

import (
	"github.com/coderguang/GameEngine_go/sgfile"
	"github.com/coderguang/GameEngine_go/sgstring"
	"io/ioutil"
	"github.com/coderguang/GameEngine_go/sgnet/sghttp"
	"github.com/coderguang/GameEngine_go/sglog"
	"net/http"
)


func ReceiveClientData(w http.ResponseWriter, r *http.Request, flag chan bool) {

	defer func(){
		flag<- true
	}()

	sglog.Info("receive client send data");

	err:=sghttp.CheckFileMaxSize(w,r,10240)
	if err!=nil{
		sglog.Error("file size too big,",err)
		return 
	}

	file, _, err := r.FormFile("sg")	
	if err != nil {	
		sglog.Error("get form data error,",err)
		return 
	}

	contentType,err:=sghttp.GetFileDetectContentType(file)
	if err!=nil{
		sglog.Error("get form file detect type error,",contentType,err)
	}

	sglog.Info("contentType:",contentType)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		sglog.Error("read all file error:",err)
		return 
	}

	filePath:="./data/"
	filename:=sgstring.RandNumStringRunes(5)

	if _,_,err:=sgfile.WriteFile(filePath,filename,fileBytes);err!=nil{
		sglog.Error("write file error");
		return;
	}

	sglog.Info("receive and write file ok");
}
