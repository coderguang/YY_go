package httpHandle

import (
	"github.com/coderguang/GameEngine_go/sgfile"
	"github.com/coderguang/GameEngine_go/sgstring"
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

	file,_,err:=sghttp.CheckIsAllowFiles(w,r,[]string{"img","jpg"})
	if err!=nil{
		sglog.Error("file type not allow",err)
		return
	}
	defer file.Close()

	fileBytes,err:=sghttp.CheckFileTypeMatch(w,r,"jpg",file)

	if err!=nil{
		sglog.Error("file type invalid",err)
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
