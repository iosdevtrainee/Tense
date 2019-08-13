package rest

import (
	"../../internal/blog"
	"../../internal/comment"
	"../../internal/thread"
)

type Route interface{}

type PostHandler interface {
	CreatePost(bc blog.Creator) Route
	FetchPostByName(bs blog.Searcher) Route
	FetchPostByID(bf blog.Fetcher) Route
	DeletePost(br blog.Remover) Route
	UpdatePost(bu blog.Updater) Route
}

type CommentHandler interface {
	FetchAllCommentsByBlogID(cf comment.Fetcher) Route
	FetchAllCommentsByUserID(cf comment.Fetcher) Route
	UpdateComment(cu comment.Updater) Route
	DeleteComment(cr comment.Remover) Route
	FetchCommentByID(cf comment.Fetcher) Route
	CreateComment(cc comment.Creator) Route
}

type ThreadHandler interface {
	CreateThread(tc thread.Creator) Route
	FetchThreadByID(tf thread.Fetcher) Route
	DeleteThread(tr thread.Remover) Route
	UpdateThread(tu thread.Updater) Route
}
