package main

import (
	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/responses"
	"mime/multipart"
	"net/http"
	"net/textproto"
)

type FileUploaderRequest struct {
	Simple  string                `formData:"simple" description:"Simple scalar value in body."`
	Query   int                   `query:"in_query" description:"Simple scalar value in query."`
	Upload1 *multipart.FileHeader `formData:"upload1" description:"Upload with *multipart.FileHeader."`
	Upload2 multipart.File        `formData:"upload2" description:"Upload with multipart.File."`
}

type FileUploaderResponse struct {
	Filename    string               `json:"filename"`
	Header      textproto.MIMEHeader `json:"header"`
	Size        int64                `json:"size"`
	Upload1Peek string               `json:"peek1"`
	Upload2Peek string               `json:"peek2"`
	Simple      string               `json:"simple"`
	Query       int                  `json:"inQuery"`
}

func WrapFileUploader(rh *huma.Resource) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RunFileUploader(rh)
	})
}

func FileUploader() func(ctx huma.Context, in FileUploaderRequest) {
	return func(ctx huma.Context, in FileUploaderRequest) {
		out := FileUploaderResponse{}

		out.Query = in.Query
		out.Simple = in.Simple
		out.Filename = in.Upload1.Filename
		out.Header = in.Upload1.Header

		out.Size = in.Upload1.Size
		f, err := in.Upload1.Open()
		if err != nil {
			ctx.WriteError(http.StatusBadRequest, "got an error when open upload1", err)
			return
		}

		defer func() {
			clErr := f.Close()
			if clErr != nil && err == nil {
				err = clErr
			}

			clErr = in.Upload2.Close()
			if clErr != nil && err == nil {
				err = clErr
			}
		}()

		p := make([]byte, 100)
		_, err = f.Read(p)
		if err != nil {
			ctx.WriteError(http.StatusBadRequest, "got an error when open upload1", err)
			return
		}

		out.Upload1Peek = string(p)

		p = make([]byte, 100)
		_, err = in.Upload2.Read(p)
		if err != nil {
			ctx.WriteError(http.StatusBadRequest, "got an error when open upload1", err)
			return
		}

		out.Upload2Peek = string(p)

		ctx.WriteModel(http.StatusOK, out)
	}
}

func RunFileUploader(r *huma.Resource) {
	r.Post("RunFileUploader", "File Upload With 'multipart/form-data'",
		responses.BadRequest(),
		responses.OK().Model(FileUploaderResponse{}),
	).Run(FileUploader)
}
