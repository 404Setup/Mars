package gnew

import (
	"crypto/md5"
	"errors"
	"net/url"
	"time"

	hash2 "github.com/3JoB/ulib/hash"
	"github.com/3JoB/ulib/litefmt"
	"github.com/3JoB/unsafeConvert"
	"github.com/savsgio/atreugo/v11"
	"gorm.io/gorm"

	"Mars/database/controller"
	"Mars/lib/uuid"
	"Mars/server/helper"
	"Mars/server/schemas"
	"Mars/shared/configure"
	schemas2 "Mars/shared/schemas"
	"Mars/shared/utils/hash"
	"Mars/shared/utils/json"
)

func ExternalDownload(c *atreugo.RequestCtx) error {
	defer func() {
		if p := recover(); p != nil {
			_ = helper.HandleInternalError(c, p)
			return
		}
	}()

	var jar *schemas2.NewDownloadSchema
	_ = json.JSON.Unmarshal(c.Request.Body(), &jar)
	if jar == nil {
		return c.JSONResponse(schemas.NewError("missing parameter"), 400)
	}

	if jar.Project == "" || jar.Version == "" || jar.Build == 0 {
		return c.JSONResponse(schemas.NewError("project or version or build is nil"), 400)
	}

	bud, err := controller.FindBuildByProjectAndVersionAndNumber(jar.Project, jar.Version, jar.Build)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSONResponse(schemas.NewError("project or version or buildId is not found"), 400)
		}
		return c.JSONResponse(schemas.NewError(err.Error()), 500)
	}

	budDL := bud.UnmarshalDownloads()
	if budDL == nil {
		budDL = map[string]schemas2.ApplicationVersionsSchema{}
	}

	if len(jar.File) != 0 {
		var e error
		for a, b := range jar.File {
			_, err = url.Parse(b.Url)
			if err != nil {
				e = errors.New(litefmt.Sprint(">> ", b.Url, " << not an effective URL"))
				break
			}
			if configure.Get().ActiveSniffing && (b.Sha256 == "" && b.Name == "") {
				name, hashe, err := hash.GetFilenameAndHash(b.Url, false)
				if err != nil {
					e = errors.New(litefmt.Sprint(">> ", err.Error()))
					break
				}
				if hashe == "" {
					e = errors.New(litefmt.Sprint(">> ", "Cannot calculate hash in streams on ", b.Url))
					break
				}
				r := schemas2.ApplicationVersionsSchema{
					Name:   name,
					Sha256: hashe,
					Url:    b.Url,
				}
				buildDownloadUrl(a, r, budDL)
				continue
			}
			buildDownloadUrl(a, b, budDL)
		}
		if e != nil {
			return c.JSONResponse(schemas.NewError(e.Error()), 400)
		}
	}
	if err = bud.MarshalDownloads(budDL); err != nil {
		return c.JSONResponse(schemas.NewError(err.Error()), 500)
	}

	if err = controller.CreateDownload(jar.Project, jar.Version, jar.Build, bud); err != nil {
		return c.JSONResponse(schemas.NewError(err.Error()), 500)
	}

	return c.JSONResponse(schemas.NewResult("success"), 200)
}

func buildDownloadUrl(id string, d schemas2.ApplicationVersionsSchema, b map[string]schemas2.ApplicationVersionsSchema) {
	if d.Name != "" {
		if d.Sha256 != "" {
			b[id] = schemas2.ApplicationVersionsSchema{
				Name:   d.Name,
				Sha256: d.Sha256,
				Url:    d.Url,
			}
			return
		}
		b[id] = schemas2.ApplicationVersionsSchema{
			Name:   d.Name,
			Sha256: "",
			Url:    d.Url,
		}
		return
	}

	b[id] = schemas2.ApplicationVersionsSchema{
		Name:   uuid.GenerateUUIDv9(hash2.CreateHMAC(unsafeConvert.BytePointer(d.Url), unsafeConvert.BytePointer(time.Now().String()), md5.New).Hex()).String(),
		Sha256: "",
		Url:    d.Url,
	}
}
