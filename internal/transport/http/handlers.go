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
			// TODO add errors top
			// t.Error = errors.New()
			http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
		}

		checkedCategories := r.Form["cat"]
		checkedList, idList, err := t.GetCategoriesForTemplate(checkedCategories)
		if err != nil {
			t.internalServerError(w, err)
			return
		}

		countPosts, err := t.service.GetCountOfPosts()
		if err != nil {
			t.internalServerError(w, err)
			return
		}
		countPostsOnPage := t.configs.PostsOnPage
		countPage := countPosts / countPostsOnPage
		if countPosts%countPostsOnPage > 0 {
			countPage++
		}
		user, _ := r.Context().Value("user").(*models.User)
		// TODO why need user
		posts, err := t.service.GetPostsForHome(1, countPostsOnPage, idList, nil)
		if err != nil {
			t.internalServerError(w, err)
			return
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
					Num:   0,
				},
			},
			User: user,
		}

		t.render(w, http.StatusOK, "home.html", data)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (t *Transport) homePages(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := r.ParseForm()
		if err != nil {
			// TODO add errors top
			// t.Error = errors.New()
			http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
		}
		checkedCategories := r.Form["cat"]
		checkedList, idList, err := t.GetCategoriesForTemplate(checkedCategories)
		if err != nil {
			t.internalServerError(w, err)
			return
		}

		// TODO add GetCountOfPosts after filters
		countPosts, err := t.service.GetCountOfPosts()
		if err != nil {
			t.internalServerError(w, err)
			return
		}
		countPostsOnPage := t.configs.PostsOnPage
		numPageString := path.Base(r.URL.Path)
		numPage, err := strconv.Atoi(numPageString)
		if err != nil || numPage < 1 || (numPage-1)*countPostsOnPage > countPosts {
			t.notFound(w)
			return
		}
		countPage := countPosts / countPostsOnPage
		if countPosts%countPostsOnPage > 0 {
			countPage++
		}
		user, _ := r.Context().Value("user").(*models.User)
		posts, err := t.service.GetPostsForHome(numPage, countPostsOnPage, idList, nil)
		if err != nil {
			t.internalServerError(w, err)
			return
		}

		data := &TemplateData{
			Data: struct {
				Header     string
				Posts      *[]models.Post
				Categories *[]CheckedCategory
				Page       interface{}
			}{
				Header:     "",
				Posts:      posts,
				Categories: checkedList,
				Page: struct {
					Count []int
					Num   int
				}{
					Count: make([]int, countPage+1),
					Num:   numPage,
				},
			},
			User: user,
		}

		t.render(w, http.StatusOK, "home.html", data)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (t *Transport) postView(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		baseID := path.Base(r.URL.Path)
		id, err := strconv.Atoi(baseID)
		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}

		user, _ := r.Context().Value("user").(*models.User)
		post, err := t.service.GetPostByID(id, user)
		if err != nil {
			t.internalServerError(w, err)
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
			http.NotFound(w, r)
			return
		}

		user, _ := r.Context().Value("user").(*models.User)
		if user == nil {
			http.Redirect(w, r, fmt.Sprintf("/post/view/%d", id), http.StatusSeeOther)
			return
		}

		err = r.ParseForm()
		if err != nil {
			// TODO add errors top
			// t.Error = errors.New()
			http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
		}
		if r.PostForm.Has("post-reactions") {
			reaction := r.PostForm.Get("post-reactions")
			posttId := r.PostForm.Get("post-id")
			id, err := strconv.Atoi(posttId)
			if err != nil {
				t.internalServerError(w, err)
				return
			}
			reactionInt, err := strconv.Atoi(reaction)
			if err != nil {
				t.internalServerError(w, err)
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
				t.internalServerError(w, err)
				return
			}
			reactionInt, err := strconv.Atoi(reaction)
			if err != nil {
				t.internalServerError(w, err)
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
				t.internalServerError(w, err)
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
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
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
			// TODO add errors top
			// t.Error = errors.New()
			http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
		}

		categoriesList := r.PostForm["categories"]
		categoriesId := []int{}
		for _, c := range categoriesList {
			num, err := strconv.Atoi(c)
			if err != nil {
				t.internalServerError(w, err)
				return
			}
			categoriesId = append(categoriesId, num)
		}
		title := r.PostForm.Get("title")
		content := r.PostForm.Get("content")
		u := r.Context().Value("user")
		user, ok := u.(*models.User)
		if !ok {
			// internal
		}

		newPost := &models.Post{
			User:         *user,
			Post_Title:   title,
			Post_Content: content,
		}

		err = t.service.CreatePost(newPost, categoriesId)
		if err != nil {
			t.internalServerError(w, err)
			return
		}
		id := newPost.Post_ID
		http.Redirect(w, r, fmt.Sprintf("/post/view/%d", id), http.StatusSeeOther)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
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
		err := r.ParseForm()
		if err != nil {
			// TODO add errors top
			// t.Error = errors.New()
			http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
		}
		email := r.PostForm.Get("email")
		pass := r.PostForm.Get("pass")
		cook, err := r.Cookie("auth")
		var newUuid string
		if err != nil {
			newUuid, err = t.service.CreateNewCookie()
			if err != nil {
				// internal
			}
			http.SetCookie(w, &http.Cookie{Name: "auth", Value: newUuid, Path: "/"})
			cook = &http.Cookie{Name: "auth", Value: newUuid, Path: "/"}
		}

		// user, err := t.service.LoginByEmailAndPass(email, pass)
		_, err = t.service.LoginByEmailAndPass(email, pass, cook.Value)
		if err != nil {
			// TODO add errors top
			// t.Error = errors.New()
			http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (t *Transport) logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		c, _ := r.Cookie("auth")
		_, err := t.service.DeregisterByCookieValue(c.Value)
		if err != nil {
			fmt.Println(err)
		}

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
			// TODO add errors top
			// t.Error = errors.New()
			http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
		}
		name := r.PostForm.Get("name")
		email := r.PostForm.Get("email")
		pass := r.PostForm.Get("pass")
		passConfirm := r.PostForm.Get("pass-confirm")
		if pass != passConfirm {
			// TODO add errors top
			// t.Error = errors.New()
			http.Redirect(w, r, fmt.Sprintf("/user/signup"), http.StatusSeeOther)
			return
		}
		_, err = t.service.CreateNewUser(&models.User{User_email: email, User_nickname: name}, pass)
		if err != nil {
			// TODO add errors top
			// t.Error = errors.New()
		}

		http.Redirect(w, r, "/user/signup", http.StatusSeeOther)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
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
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
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
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
