package http

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type CheckedCategory struct {
	Category_id   int
	Category_name string
	IsChecked     bool
}

func (t *Transport) render(w http.ResponseWriter, status int, page string, data *TemplateData) {
	ts, ok := t.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		t.ErrLog.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	pageName := strings.Split(page, ".")[0]
	data.PageName = pageName

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		t.ErrLog.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}

func (t *Transport) notFound(w http.ResponseWriter) {
	t.render(w, http.StatusNotFound, "error.html", &TemplateData{Data: "Page Not Found"})
}

func (t *Transport) internalServerError(w http.ResponseWriter, err error) {
	t.ErrLog.Println(err)
	t.render(w, http.StatusInternalServerError, "error.html", &TemplateData{Data: http.StatusText(http.StatusInternalServerError)})
}

func (t *Transport) GetCategoriesForTemplate(categoriesList []string) (checkedList *[]CheckedCategory, idList []int, err error) {
	categories, err := t.service.GetCategiries()
	if err != nil {
		return nil, nil, err
	}

	idList = []int{}
	for _, c := range categoriesList {
		num, err := strconv.Atoi(c)
		if err != nil {
			return nil, nil, err
		}
		idList = append(idList, num)
	}

	checkedList = &[]CheckedCategory{}
	for _, c := range *categories {
		checked := func() bool {
			for _, num := range idList {
				if num == c.Category_id {
					return true
				}
			}
			return false
		}()
		*checkedList = append(*checkedList, CheckedCategory{
			Category_id:   c.Category_id,
			Category_name: c.Category_name,
			IsChecked:     checked,
		})
	}
	return checkedList, idList, nil
}
