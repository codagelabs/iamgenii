package handlers

import (
	"io"
	"os"

	"github.com/iamgenii/utils/http_utils"
	"github.com/iamgenii/validator"

	"net/http"
	//"github.com/iamgenii/svc/images/pkg/v1/services"
)

// ImageHandlers for handlers
type ImageHandlers struct {
	//imageSvc      services.ImageServices
	httpReader       http_utils.HTTPReader
	httpWriter       http_utils.HTTPWriter
	requestValidator validator.RequestValidator
}

func NewImageHandlers(httpReader http_utils.HTTPReader, httpWriter http_utils.HTTPWriter, requestValidator validator.RequestValidator) *ImageHandlers {
	return &ImageHandlers{
		//imageSvc:      imageSvc,
		httpReader:       httpReader,
		httpWriter:       httpWriter,
		requestValidator: requestValidator,
	}
}

// CreateImageHandleFunc handler Function
func (imageHandler ImageHandlers) UploadImage(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	//POST takes the uploaded file(s) and saves it to disk.
	case "POST":
		//parse the multipart form in the request
		err := req.ParseMultipartForm(100000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//get a ref to the parsed multipart form
		m := req.MultipartForm

		//get the *fileheaders
		files := m.File["image_upload"]
		for i, _ := range files {
			//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//create destination file making sure the path is writeable.
			dst, err := os.Create("/var/www/images/" + files[i].Filename)
			defer dst.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		}
		//display success message.
		imageHandler.httpWriter.WriteCustomHTTPError(w, http.StatusCreated, "Upload successful.")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
