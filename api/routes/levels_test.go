package routes_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/Nearrivers/dnd-grid-server/api/presenter"
	"github.com/Nearrivers/dnd-grid-server/api/routes"
	"github.com/Nearrivers/dnd-grid-server/pkg/models/repository"
	"github.com/gofiber/fiber/v2"
)

var mockLevels = []repository.Levels{
	{
		ID:        1,
		Name:      "test",
		ImagePath: "./test",
		GridWidth: 3,
		GridColor: "#fff",
		GridSpacing: sql.NullInt64{
			Int64: 3,
		},
	},
	{
		ID:        2,
		Name:      "autre",
		ImagePath: "./autre",
		GridWidth: 3,
		GridColor: "#444",
		GridSpacing: sql.NullInt64{
			Int64: 3,
		},
	},
}

type levelServiceMock struct{}

func (l *levelServiceMock) DeleteLevel(id int64) error {
	return nil
}

func (l *levelServiceMock) GetLevel(id int64) (repository.Levels, error) {
	return mockLevels[0], nil
}

func (l *levelServiceMock) GetLevelWithEntities(id int64) ([]repository.GetLevelWithEntitiesRow, error) {
	return []repository.GetLevelWithEntitiesRow{}, nil
}

func (l *levelServiceMock) GetLevels() ([]repository.Levels, error) {
	return mockLevels, nil
}

func (l *levelServiceMock) NewLevel(level repository.NewLevelParams) error {
	return nil
}

func (l *levelServiceMock) UpdateLevel(level repository.UpdateLevelParams) error {
	return nil
}

func setupNewApp() *fiber.App {
	ls := &levelServiceMock{}
	app := fiber.New()
	routes.BookRouter(app, ls)
	return app
}

func TestAddRoute(t *testing.T) {
	cases := []struct {
		description string
		statusCode  int
		l           presenter.Level
	}{
		{
			description: "valid level",
			statusCode:  201,
			l: presenter.Level{
				ID:          0,
				Name:        "name",
				ImagePath:   "./path",
				GridWidth:   3,
				GridColor:   "#fff",
				GridSpacing: 3,
			},
		},
		{
			description: "Missing name field",
			statusCode:  http.StatusBadRequest,
			l: presenter.Level{
				ID:          0,
				ImagePath:   "./path",
				GridWidth:   3,
				GridColor:   "#fff",
				GridSpacing: 3,
			},
		},
		{
			description: "Missing image path field",
			statusCode:  http.StatusBadRequest,
			l: presenter.Level{
				ID:          0,
				Name:        "name",
				GridWidth:   3,
				GridColor:   "#fff",
				GridSpacing: 3,
			},
		},
		{
			description: "Missing grid width field",
			statusCode:  http.StatusBadRequest,
			l: presenter.Level{
				ID:          0,
				Name:        "name",
				ImagePath:   "./path",
				GridColor:   "#fff",
				GridSpacing: 3,
			},
		},
		{
			description: "Missing grid color field",
			statusCode:  http.StatusBadRequest,
			l: presenter.Level{
				ID:          0,
				Name:        "name",
				ImagePath:   "./path",
				GridWidth:   3,
				GridSpacing: 3,
			},
		},
		{
			description: "Missing grid spacing field",
			statusCode:  http.StatusBadRequest,
			l: presenter.Level{
				Name:      "name",
				ID:        0,
				ImagePath: "./path",
				GridWidth: 3,
				GridColor: "#fff",
			},
		},
	}

	for _, tt := range cases {
		app := setupNewApp()
		t.Run(tt.description, func(t *testing.T) {
			ml, err := json.Marshal(tt.l)
			if err != nil {
				t.Fatalf("couldn't marshal level: %v", err)
			}

			r := bytes.NewReader(ml)

			req, err := http.NewRequest(http.MethodPost, "/levels", r)
			if err != nil {
				t.Fatalf("couldn't create request: %v", err)
			}

			req.Header.Add("Content-Type", "application/json")
			res, err := app.Test(req, -1)
			if err != nil {
				t.Fatalf("unexpected error occured: %v", err)
			}

			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("couldn't read response body: %v", err)
			}

			s := string(body)
			if res.StatusCode != tt.statusCode {
				t.Fatalf("wrong status code. got %d, want %d for %s", res.StatusCode, tt.statusCode, s)
			}
		})
	}
}

func TestGetLevels(t *testing.T) {
	app := setupNewApp()

	req, err := http.NewRequest(http.MethodGet, "/levels", nil)
	if err != nil {
		t.Fatalf("couldn't create request : %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("unexpected error occured: %v", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("couldn't read response body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("wrong status code, got %d, want %d", res.StatusCode, http.StatusOK)
	}

	var want struct {
		Status bool              `json:"status"`
		Data   []presenter.Level `json:"data"`
		Error  error             `json:"error"`
	}

	err = json.Unmarshal(body, &want)
	if err != nil {
		t.Fatalf("couldn't unmarshal response body: %v", err)
	}

	if want.Error != nil {
		t.Fatalf("got an error response but didn't expect one: %v", want.Error)
	}
}

// func createTestImage() *image.RGBA {
// 	width := 20
// 	height := 10

// 	upLeft := image.Point{0, 0}
// 	lowRight := image.Point{width, height}

// 	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

// 	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
// 	cyan := color.RGBA{100, 200, 200, 0xff}

// 	// Set color for each pixel.
// 	for x := 0; x < width; x++ {
// 		for y := 0; y < height; y++ {
// 			switch {
// 			case x < width/2 && y < height/2: // upper left quadrant
// 				img.Set(x, y, cyan)
// 			case x >= width/2 && y >= height/2: // lower right quadrant
// 				img.Set(x, y, color.White)
// 			default:
// 				// Use zero value.
// 			}
// 		}
// 	}

// 	return img
// }

// func TestUploadImage(t *testing.T) {
// 	cases := []struct {
// 		description   string
// 		imgMime       string
// 		imgEncodeFunc func(w io.Writer, m image.Image) error
// 	}{
// 		{
// 			description:   "upload png",
// 			imgMime:       "png",
// 			imgEncodeFunc: png.Encode,
// 		},
// 		{
// 			description: "upload jpeg",
// 			imgMime:     "jpeg",
// 			imgEncodeFunc: func(w io.Writer, m image.Image) error {
// 				return jpeg.Encode(w, m, &jpeg.Options{})
// 			},
// 		},
// 	}

// 	for _, tt := range cases {
// 		app := setupNewApp()

// 		t.Run(tt.description, func(t *testing.T) {
// 			pr, pw := io.Pipe()
// 			defer pr.Close()

// 			mwriter := multipart.NewWriter(pw)

// 			go func() {
// 				part, err := mwriter.CreateFormFile("image", "testimg."+tt.imgMime)
// 				if err != nil {
// 					pw.CloseWithError(err)
// 					return
// 				}

// 				// https://yourbasic.org/golang/create-image/
// 				img := createTestImage()

// 				var b bytes.Buffer
// 				err = tt.imgEncodeFunc(&b, img)
// 				if err != nil {
// 					pw.CloseWithError(err)
// 					return
// 				}

// 				if _, err := io.Copy(part, &b); err != nil {
// 					pw.CloseWithError(err)
// 					return
// 				}

// 				pw.Close()
// 			}()

// 			defer mwriter.Close()

// 			req, err := http.NewRequest(http.MethodPost, "/levels/image", pr)
// 			if err != nil {
// 				t.Fatalf("couldn't create request: %v", err)
// 			}

// 			req.Header.Add("Content-Type", mwriter.FormDataContentType())
// 			req.Header.Add("Content-Disposition", `form-date; name="image"`)
// 			res, err := app.Test(req, -1)

// 			if err != nil {
// 				t.Fatalf("got an error but didn't expect one: %v", err)
// 			}

// 			want := http.StatusCreated
// 			if res.StatusCode != want {
// 				t.Fatalf("wrong status code. got %d, want %d", res.StatusCode, want)
// 			}
// 		})
// 	}
// }
