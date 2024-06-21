package http

import (
	"fmt"
	"forum/internal/models"
	"net/http"
	"path"
	"strconv"
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
		err := r.ParseForm()
		if err != nil {
			t.badRequest(w)
			return
		}

		checkedCategories := r.Form["cat"]
		page := r.FormValue("page")
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			pageInt = 1
		}
		if pageInt < 1 {
			t.notFound(w)
			return
		}
		checkedList, idList, err := t.GetCategoriesForTemplate(checkedCategories)
		if err != nil {
			t.internalServerError(w, err)
			return
		}

		user, _ := r.Context().Value("user").(*models.User)
		countPostsOnPage := t.configs.PostsOnPage
		posts, countPosts, err := t.service.GetPostsForHome(pageInt, countPostsOnPage, idList, nil)
		if err != nil {
			t.internalServerError(w, err)
			return
		}
		countPage := countPosts / countPostsOnPage
		if countPosts%countPostsOnPage > 0 {
			countPage++
		}

		data := &TemplateData{
			Data: struct {
				Header     string
				Posts      *[]models.Post
				Categories *[]CheckedCategory
				Page       interface{}
			}{
				Header:     "Latest Posts",
				Posts:      posts,
				Categories: checkedList,
				Page: struct {
					Count []int
					Num   int
				}{
					Count: make([]int, countPage+1),
					Num:   pageInt,
				},
			},
			User: user,
		}

		t.render(w, http.StatusOK, "home.html", data)
	} else {
		t.methodNotAllowed(w)
	}
}

func (t *Transport) postView(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		baseID := path.Base(r.URL.Path)
		id, err := strconv.Atoi(baseID)
		if err != nil || id < 1 {
			t.badRequest(w)
			return
		}

		user, _ := r.Context().Value("user").(*models.User)
		post, err := t.service.GetPostByID(id, user)
		if err != nil {
			t.notFound(w)
			return
		}
		data := &TemplateData{
			Data: post,
			User: user,
		}

		t.render(w, http.StatusOK, "post-view.html", data)
	} else if r.Method == http.MethodPost {
		baseID := path.Base(r.URL.Path)
		id, err := strconv.Atoi(baseID)
		if err != nil || id < 1 {
			t.badRequest(w)
			return
		}

		user, _ := r.Context().Value("user").(*models.User)
		if user == nil {
			http.Redirect(w, r, fmt.Sprintf("/post/view/%d", id), http.StatusSeeOther)
			return
		}

		err = r.ParseForm()
		if err != nil {
			t.badRequest(w)
			return
		}
		if r.PostForm.Has("post-reactions") {
			reaction := r.PostForm.Get("post-reactions")
			posttId := r.PostForm.Get("post-id")
			id, err := strconv.Atoi(posttId)
			if err != nil {
				t.badRequest(w)
				return
			}
			reactionInt, err := strconv.Atoi(reaction)
			if err != nil {
				t.badRequest(w)
				return
			}

			err = t.service.ReactionsToPost(&models.Post{Post_ID: int64(id)}, user, reactionInt)
			if err != nil {
				t.internalServerError(w, err)
				return
			}
		}
		if r.PostForm.Has("comment-reactions") {
			reaction := r.PostForm.Get("comment-reactions")
			commentId := r.PostForm.Get("comment-id")
			id, err := strconv.Atoi(commentId)
			if err != nil {
				t.badRequest(w)
				return
			}
			reactionInt, err := strconv.Atoi(reaction)
			if err != nil {
				t.badRequest(w)
				return
			}

			err = t.service.ReactionsToComment(id, user, reactionInt)
			if err != nil {
				t.internalServerError(w, err)
				return
			}
		}
		if r.PostForm.Has("create-comment") {
			comment := r.PostForm.Get("comment")
			thisPost, err := t.service.GetPostByID(id, user)
			if err != nil {
				t.badRequest(w)
				return
			}
			err = t.service.CraeteCommentary(thisPost, &models.Comment{
				User:                *user,
				Commentraie_Content: comment,
			})
			if err != nil {
				t.internalServerError(w, err)
				return
			}
		}

		http.Redirect(w, r, fmt.Sprintf("/post/view/%d", id), http.StatusSeeOther)
	} else {
		t.methodNotAllowed(w)
	}
}

func (t *Transport) postCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		categories, err := t.service.GetCategiries()
		if err != nil {
			t.internalServerError(w, err)
			return
		}

		user, _ := r.Context().Value("user").(*models.User)
		data := &TemplateData{
			Data: struct {
				Categories *[]models.Category
			}{
				Categories: categories,
			},
			User: user,
		}

		t.render(w, http.StatusOK, "post-create.html", data)
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			t.badRequest(w)
			return
		}

		categoriesList := r.PostForm["categories"]
		categoriesId := []int{}
		for _, c := range categoriesList {
			num, err := strconv.Atoi(c)
			if err != nil {
				t.badRequest(w)
				return
			}
			categoriesId = append(categoriesId, num)
		}
		title := r.PostForm.Get("title")
		content := r.PostForm.Get("content")
		u := r.Context().Value("user")
		user, ok := u.(*models.User)
		if !ok {
			t.internalServerError(w, err)
			return
		}

		newPost := &models.Post{
			User:         *user,
			Post_Title:   title,
			Post_Content: content,
		}

		err = t.service.CreatePost(newPost, categoriesId)
		if err != nil {
			t.badRequest(w)
			return
		}

		id := newPost.Post_ID
		http.Redirect(w, r, fmt.Sprintf("/post/view/%d", id), http.StatusSeeOther)
	} else {
		t.methodNotAllowed(w)
	}
}

func (t *Transport) login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		user, _ := r.Context().Value("user").(*models.User)
		data := &TemplateData{
			User: user,
		}

		t.render(w, http.StatusOK, "login.html", data)
	} else if r.Method == http.MethodPost {
		data := &TemplateData{}
		err := r.ParseForm()
		if err != nil {
			t.badRequest(w)
			return
		}
		email := r.PostForm.Get("email")
		pass := r.PostForm.Get("pass")
		authCookie, err := r.Cookie("auth")
		if err != nil {
			authCookie, err = t.CreateCookie()
			if err != nil {
				t.internalServerError(w, err)
				return
			}
			http.SetCookie(w, authCookie)
		}

		_, err = t.service.LoginByEmailAndPass(email, pass, authCookie.Value)
		if err != nil {
			data.Data = err
			t.render(w, http.StatusUnprocessableEntity, "login.html", data)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		t.methodNotAllowed(w)
	}
}

func (t *Transport) logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		c, err := r.Cookie("auth")
		if err != nil {
			t.badRequest(w)
			return
		}
		_, err = t.service.DeregisterByCookieValue(c.Value)
		if err != nil {
			t.internalServerError(w, err)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		t.methodNotAllowed(w)
	}
}

func (t *Transport) signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data := &TemplateData{}

		t.render(w, http.StatusOK, "signup.html", data)
	} else if r.Method == http.MethodPost {
		data := &TemplateData{}

		err := r.ParseForm()
		if err != nil {
			t.badRequest(w)
			return
		}
		name := r.PostForm.Get("name")
		email := r.PostForm.Get("email")
		pass := r.PostForm.Get("pass")
		passConfirm := r.PostForm.Get("pass-confirm")
		if pass != passConfirm {
			data.Data = err
			t.render(w, http.StatusUnprocessableEntity, "login.html", data)
			return
		}
		_, err = t.service.CreateNewUser(&models.User{User_email: email, User_nickname: name}, pass)
		if err != nil {
			data.Data = err
			t.render(w, http.StatusUnprocessableEntity, "signup.html", data)
			return
		}

		http.Redirect(w, r, "/user/signup", http.StatusSeeOther)
	} else {
		t.methodNotAllowed(w)
	}
}

func (t *Transport) myPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		user, _ := r.Context().Value("user").(*models.User)
		posts, err := t.service.GetOnlyMyPosts(user)
		if err != nil {
			t.internalServerError(w, err)
			return
		}

		data := &TemplateData{
			Data: struct {
				Header     string
				Posts      *[]models.Post
				Categories *[]models.Category
				Page       interface{}
			}{
				Header: "My posts",
				Posts:  posts,
			},
			User: user,
		}
		t.render(w, http.StatusOK, "home.html", data)
	} else {
		t.methodNotAllowed(w)
	}
}

func (t *Transport) liked(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		user, _ := r.Context().Value("user").(*models.User)
		posts, err := t.service.GetMyPostReactions(user)
		if err != nil {
			t.internalServerError(w, err)
			return
		}

		data := &TemplateData{
			Data: struct {
				Header     string
				Posts      *[]models.Post
				Categories *[]models.Category
				Page       interface{}
			}{
				Header: "Liked posts",
				Posts:  posts,
			},
			User: user,
		}
		t.render(w, http.StatusOK, "home.html", data)
	} else {
		t.methodNotAllowed(w)
	}
}
