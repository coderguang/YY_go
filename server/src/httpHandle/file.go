package httpHandle

import (
	"io"
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


func ReceiveMultiClientData(w http.ResponseWriter,r *http.Request,flag chan bool){
	defer func(){
		flag<- true
	}()

	read_form,err:=r.MultipartReader()

	if err!=nil{
		sglog.Error("r.MultipartReader err",err);
		return
	}	

	for{
		part, err_part := read_form.NextPart()
    	if err_part == io.EOF {
        	break
		}
		name:=part.FormName();
		if name==""{
			continue
		}
		filename:=part.FileName()
		sglog.Debug("get formname:",name,",file:",filename)

		if filename==""{

		}else{
			file,_,err:=r.FormFile(name)
			if err!=nil{
				sglog.Error("read file error,file:",file,err)
				continue;
			}
			fileBytes,err:=ioutil.ReadAll(file)
			if err!=nil{
				sglog.Error("read data from file error,file:",file,err)
				continue
			}
			filePath:="./data"
			if _,_,err:=sgfile.WriteFile(filePath,filename,fileBytes);err!=nil{
				sglog.Error("write to file file error,file:",file,err)
				continue
			}
		}
	}


	sglog.Info("receive ReceiveMultiClientData send data complete");
}
