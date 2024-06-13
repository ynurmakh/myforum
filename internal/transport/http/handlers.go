package http

import (
	"bytes"
	"fmt"
	"forum/internal/models"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
)

type TemplateData struct {
	Data     interface{}
	User     *models.User
	PageName string
}

func (t *Transport) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		t.notFound(w)
		return
	}
	if r.Method == http.MethodGet {
		categories, err := t.service.GetCategiries()

		posts, err := t.service.GetPostsForHome(1, 20, []int{})
		if err != nil {
			fmt.Println("posts not found")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		data := &TemplateData{
			Data: struct {
				Posts      *[]models.Post
				Categories *[]models.Categories
			}{
				Posts:      posts,
				Categories: categories,
			},
			User: t.User,
		}
		t.render(w, http.StatusOK, "home.html", data)
	} else if r.Method == http.MethodPost {
		categories, err := t.service.GetCategiries()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
		}

		categoriesList := r.PostForm["categories"]
		categoriesId := []int{}
		for _, c := range categoriesList {
			num, err := strconv.Atoi(c)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			categoriesId = append(categoriesId, num)
		}

		posts, err := t.service.GetPostsForHome(1, 20, categoriesId)
		if err != nil {
			fmt.Println("posts not found")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		data := &TemplateData{
			Data: struct {
				Posts      *[]models.Post
				Categories *[]models.Categories
			}{
				Posts:      posts,
				Categories: categories,
			},
			User: t.User,
		}
		t.render(w, http.StatusOK, "home.html", data)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (t *Transport) postView(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	baseID := path.Base(r.URL.Path)
	id, err := strconv.Atoi(baseID)
	if err != nil || id < 1 {
		fmt.Println("atoi err")
		http.NotFound(w, r)
		return
	}

	post, err := t.service.GetPostByID(id)
	if err != nil {
		fmt.Println("post not found")
	}
	if err != nil {
		fmt.Println("user not found")
	}
	// mock category
	post.Post_Categories = append(post.Post_Categories, models.Categories{
		Category_id:   0,
		Category_name: "Trash",
	})
	data := &TemplateData{
		Data: post,
		User: t.User,
	}
	t.render(w, http.StatusOK, "post-view.html", data)
}

func (t *Transport) postCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		categories, err := t.service.GetCategiries()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		data := &TemplateData{
			Data: struct {
				Categories *[]models.Categories
			}{
				Categories: categories,
			},
			User: t.User,
		}

		t.render(w, http.StatusOK, "post-create.html", data)
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		categoriesList := r.PostForm["categories"]
		categoriesId := []int{}
		for _, c := range categoriesList {
			num, err := strconv.Atoi(c)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			categoriesId = append(categoriesId, num)
		}
		title := r.PostForm.Get("title")
		content := r.PostForm.Get("content")
		newPost := &models.Post{
			User:         *t.User,
			Post_Title:   title,
			Post_Content: content,
		}

		err = t.service.CreatePost(newPost, categoriesId)
		id := newPost.Post_ID
		http.Redirect(w, r, fmt.Sprintf("/post/view/%d", id), http.StatusSeeOther)
		// http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (t *Transport) login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// user, err := t.service.GetUser(t.UserId)
		// if err != nil {
		// 	fmt.Println("user not found")
		// }
		data := &TemplateData{
			// User: user,
		}
		t.render(w, http.StatusOK, "login.html", data)
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// email := r.PostForm.Get("email")
		// pass := r.PostForm.Get("pass")
		// id, err := t.service.Login(email, pass)
		// fmt.Println("login user ID:", id)
		if err != nil {
			fmt.Println(err)
			return
		}
		// t.UserId = id
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (t *Transport) logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		t.UserId = 0
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (t *Transport) signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data := &TemplateData{}
		t.render(w, http.StatusOK, "signup.html", data)
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// name := r.PostForm.Get("name")
		// email := r.PostForm.Get("email")
		// pass := r.PostForm.Get("pass")
		// passConfirm := r.PostForm.Get("pass-confirm")
		// id, err := t.service.CreateUser(name, email, pass)
		// fmt.Println("create user:", id, name, email, pass, passConfirm)
		if err != nil {
			fmt.Println(err)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (t *Transport) render(w http.ResponseWriter, status int, page string, data *TemplateData) {
	ts, ok := t.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		log.Println(err)
		return
	}

	pageName := strings.Split(page, ".")[0]
	data.PageName = pageName

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}

func (t *Transport) notFound(w http.ResponseWriter) {
	t.render(w, http.StatusNotFound, "notfound.html", &TemplateData{Data: "Page Not Found"})
}
