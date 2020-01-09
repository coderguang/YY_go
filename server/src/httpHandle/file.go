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


func TryReceiveOne(w http.ResponseWriter,r *http.Request,flag chan bool){
	defer func(){
		flag<- true
	}()

	
	_,err:=r.MultipartReader()

	if err!=nil{
		sglog.Error("r.MultipartReader err",err);
		return
	}	

	filename:="pic"
	file,_,err:=r.FormFile(filename)
	if err!=nil{
		sglog.Error("read file error,file:",file,err)
		return;
	}
	fileBytes,err:=ioutil.ReadAll(file)
	if err!=nil{
		sglog.Error("read data from file error,file:",file,err)
		return
	}
	filePath:="./data/"
	if _,_,err:=sgfile.WriteFile(filePath,filename,fileBytes);err!=nil{
		sglog.Error("write to file file error,file:",file,err)
		return
	}
	sglog.Info("receive ",filename,",complete by try")
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
			fileBytes:=make([]byte,10000)
			num,err:=part.Read(fileBytes)
			if err!=nil&&err!=io.EOF{
				sglog.Error("read part file error,file:",filename,err)
				continue
			}
			sglog.Info("read part file success,num:",num)
			part.Close()

			filePath:="./data/"
			wrieteFileName,err:=sgfile.GetFileName(filename)
			if err!=nil{
				sglog.Error("unknow filename,err:",filename,err)
				wrieteFileName="unknow_"+sgstring.RandNumStringRunes(5)
			}
			if _,_,err:=sgfile.WriteFile(filePath,wrieteFileName,fileBytes[:num]);err!=nil{
				sglog.Error("write to file file error,file:",wrieteFileName,err)
				continue
			}
		}
	}


	sglog.Info("receive ReceiveMultiClientData send data complete");
}
