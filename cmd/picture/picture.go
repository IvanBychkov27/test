// https://habr.com/ru/post/530134/
// Создание изображений в runtime (динамично)
// с добавлением надписи

package main

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	pageForm = `
<html>
    <title>SimpleCaptcha</title>
 <body>
	<br>
		<img src="data:image/png;base64, {IMAGE}">
    <br>
	<br>
	<form action="/" method="POST" autocomplete="off">
		
		<input autofocus id="value" name="value"/>
		
		<input type="hidden" name="control" value="{CONTROL}"/>
		<input type="submit" value="Send" />
		
		<label for="textcolor">Выберите цвет текста</label>
		<input type="color" id="textcolor" name="textcolor" />

		<label for="foncolor">Выберите цвет фона</label>
		<input type="color" id="foncolor" name="foncolor" />

		<label for="font">Выберите размер шрифта</label>
		<input type="number" step="10" min="10" max="500" id="font" name="font"/>
	</form>
`
	endPage = `
 </body>
</html>
`
)

func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	res := ""
	err := req.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if len(req.Form) != 0 {
		msgColor := req.Form.Get("textcolor")[1:]
		if msgColor != "000000" {
			msgColorDefault = msgColor
		}
		fonColor := req.Form.Get("foncolor")[1:]
		if fonColor != "000000" {
			imgColorDefault = fonColor
		}
		f := req.Form.Get("font")
		if f != "" {
			var fs int
			fs, err = strconv.Atoi(f)
			if err == nil {
				fontSize = fs
			}
		}

		codControl := req.Form.Get("control")
		codUser := req.Form.Get("value")
		codValue := md5Encode(codUser, "SimpleCaptcha")

		if codControl == codValue {
			res += `<h4 style="color:green;"> Код ` + codUser + ` верный!</h4>`
		} else {
			res += `<h4 style="color:red;"> ` + codUser + `  не верный код - проверка не пройдена...</h4>`
		}
	}
	page := pagePicture() // формирование страницы со сгенерированной картинкой с текстом
	_, err = w.Write([]byte(page + res + endPage))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("server is listening... 127.0.0.1:8181")
	http.HandleFunc("/", ServeHTTP)
	err := http.ListenAndServe("localhost:8181", nil)
	chk(err)
}

// Для теста в проекте tch
// http://127.0.0.1:2000/redir?token=2910cbab-615e-4dca-bc3e-a6a74cf5d48f&url=http://ya.ru

// формирование страницы со сгенерированной картинкой с текстом
func pagePicture() string {
	cod := 100 + rand.Intn(899)
	text := strconv.Itoa(cod)
	buf, err := Do(text, fontSize) // генерация картинки с текстом text и размером шрифта fontSize
	if err != nil {
		log.Println("error generation image")
		return "<html><body>error generation image"
	}

	dataBase64 := base64.StdEncoding.EncodeToString(buf.Bytes()) // кодируем картинку из буфера в строку для страницы
	data := strings.Replace(pageForm, "{IMAGE}", dataBase64, -1)

	codControl := md5Encode(strconv.Itoa(cod), "SimpleCaptcha")
	data = strings.Replace(data, "{CONTROL}", codControl, -1)

	return data
}

func md5Encode(data, salt string) string {
	data += salt
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}

// Пакет сборки картинки:
// Параметры по умолчанию
var (
	imgColorDefault = "ffffaa"
	msgColorDefault = "0707FF"
	fontSize        = 80
)

const (
	imgWDefault         = 10
	imgHDefault         = 10
	dpiDefault  float64 = 72
)

// Соберём структуру Текста
type Label struct {
	Text     string
	FontSize int
	Color    string
}

// Соберём структуру Картинки с нужными полями - высота, ширина, цвет и текст
type Img struct {
	Width  int
	Height int
	Color  string
	Label  Label
}

// Do - входная точка
func Do(text string, frotSize int) (*bytes.Buffer, error) {
	label := Label{text, frotSize, msgColorDefault}
	img := Img{0, 0, imgColorDefault, label}
	return img.generate()
}

// generate - соберёт картинку по нужным размерам, цветом и текстом
func (i Img) generate() (*bytes.Buffer, error) {
	// Если есть размеры и нет требований по Тексту - соберём Текст по умолчанию.
	if i.Width == 0 || i.Height == 0 {
		i.Width, i.Height = imgWDefault, imgHDefault
	}
	if i.Label.Text == "" {
		i.Label.Text = fmt.Sprintf("%d x %d", i.Width, i.Height)
	}
	// Если нет требований по размеру шрифта - подберём его исходя из размеров картинки.
	if i.Label.FontSize > i.Height {
		i.Height = i.Label.FontSize
	}
	if i.Label.FontSize == 0 {
		i.Label.FontSize = i.Width / 5
		if i.Height < i.Width {
			i.Label.FontSize = i.Height / 5
		}
	}
	lenText := int(float64(len(i.Label.Text)*i.Label.FontSize) * 0.6)
	if lenText > i.Width {
		i.Width = lenText
	}

	// Переведём цвет из строки в color.RGBA.
	clr, err := ToRGBA(i.Color)
	if err != nil {
		return nil, err
	}
	// Создадим in-memory картинку с нужными размерами.
	m := image.NewRGBA(image.Rect(0, 0, i.Width, i.Height))
	// Отрисуем картинку: - по размерам (Bounds); - и с цветом (Uniform - обёртка над color.Color c Image функциями)
	// - исходя из точки (Point), как базовой картинки; - заполним цветом нашу Uniform (draw.Src)
	draw.Draw(m, m.Bounds(), image.NewUniform(clr), image.Point{}, draw.Src)
	// Добавим текст в картинку
	if err = i.drawLabel(m); err != nil {
		return nil, err
	}
	var im image.Image = m
	// Выделим память под нашы данные (байты картинки)
	buffer := &bytes.Buffer{}
	// Закодируем картинку в buffer
	err = png.Encode(buffer, im)

	return buffer, err
}

// drawLabel - добавит текст на картинку
func (i *Img) drawLabel(m *image.RGBA) error {
	// Разберём цвет текста из строки в RGBA.
	clr, err := ToRGBA(i.Label.Color)
	chk(err)
	// Получим шрифт (должен работать и с латиницей и с кириллицей).
	//fontBytes, err := ioutil.ReadFile(string(goregular.TTF))
	//if err != nil {
	//	return err
	//}
	fnt, err := truetype.Parse(goregular.TTF)
	chk(err)
	// Подготовим Drawer для отрисовки текста на картинке.
	d := &font.Drawer{
		Dst: m,
		Src: image.NewUniform(clr),
		Face: truetype.NewFace(fnt, &truetype.Options{
			Size:    float64(i.Label.FontSize),
			DPI:     dpiDefault,
			Hinting: font.HintingNone,
		}),
	}
	// Зададим базовую линию.
	d.Dot = fixed.Point26_6{
		X: (fixed.I(i.Width) - d.MeasureString(i.Label.Text)) / 2,
		Y: fixed.I((i.Height+i.Label.FontSize)/2 - i.Label.FontSize/10),
	}
	// Непосредственно отрисовка текста в нашу RGBA картинку.
	d.DrawString(i.Label.Text)

	return nil
}

// Пакет для работы с цветами
type rgb struct {
	red   uint8
	green uint8
	blue  uint8
}

func ToRGBA(h string) (color.RGBA, error) {
	rgb, err := hex2RGB(h)
	if err != nil {
		return color.RGBA{}, err
	}
	return color.RGBA{R: rgb.red, G: rgb.green, B: rgb.blue, A: 255}, nil
}
func hex2RGB(hex string) (rgb, error) {
	values, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return rgb{}, err
	}
	return rgb{
		red:   uint8(values >> 16),
		green: uint8((values >> 8) & 0xFF),
		blue:  uint8(values & 0xFF),
	}, nil
}

func chk(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
