package postcontroller

import (
	"net/http"

	"github.com/L-oris/yabb/router/httperror"
)

func (c Controller) new(w http.ResponseWriter, req *http.Request) {
	postForm, err := parsePostForm(w, req, true)
	if err != nil {
		httperror.BadRequest(w, err.Error())
		return
	}

	newPost, err := c.service.Create(postForm.post, postForm.fileBytes)
	if err != nil {
		httperror.InternalServer(w, err.Error())
		return
	}

	w.Header().Set("Location", "/post/"+newPost.ID)
	w.WriteHeader(http.StatusSeeOther)
}

func (c Controller) updateByID(w http.ResponseWriter, req *http.Request) {
	postID, err := getPostIDFromURL(req)
	if err != nil {
		httperror.BadRequest(w, err.Error())
		return
	}

	postForm, err := parsePostForm(w, req, false)
	if err != nil {
		httperror.BadRequest(w, err.Error())
		return
	}

	updatedPost, err := c.service.UpdateByID(postID, postForm.post, postForm.fileBytes)
	if err != nil {
		httperror.InternalServer(w, err.Error())
		return
	}

	w.Header().Set("Location", "/post/"+updatedPost.ID)
	w.WriteHeader(http.StatusSeeOther)
}

func (c Controller) deleteByID(w http.ResponseWriter, req *http.Request) {
	postID, err := getPostIDFromURL(req)
	if err != nil {
		httperror.BadRequest(w, err.Error())
		return
	}

	if err = c.service.DeleteByID(postID); err != nil {
		httperror.InternalServer(w, err.Error())
		return
	}

	w.Header().Set("Location", "/post/all")
	w.WriteHeader(http.StatusSeeOther)
}
