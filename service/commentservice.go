package service

import (
	"bucuo/dao"
	"bucuo/model/request"
	"bucuo/model/table"
)

var commentdao = dao.CommentDao{}

type CommentService struct {
}

func (c CommentService) AddComment(req *request.CommentReq, uid uint) string {
	if !commentdao.PostExist(req.Type, req.PostID) {
		return "recoder not found"
	}
	comment := table.Comment{
		Content:   req.Content,
		UserID:    uid,
		OwnerType: req.Type,
		OwnerID:   req.PostID,
	}
	if err := commentdao.AddComment(&comment); err != nil {
		return err.Error()
	}
	return ""
}
func (c CommentService) DeleteCommemnt(id uint, uid uint) string {
	if es := commentdao.DeleteComment(id, uid); es != nil {
		return es.Error()
	} else {
		return ""
	}
}
func (c CommentService) AddReply(req *request.AddReplyReq, uid uint) string {
	if s := commentdao.AddReply(req.CommentID, &table.Reply{Content: req.Content, ReplierID: uid}); s != nil {
		return s.Error()
	}
	return ""
}
func (c CommentService) DeleteReply(id uint, uid uint) string {
	if s := commentdao.DeleteReply(id, uid); s != nil {
		return s.Error()
	} else {
		return ""
	}
}
func (c CommentService) Like(uid uint, commentid uint) string {
	if er := commentdao.Like(uid, commentid); er != nil {
		return er.Error()
	} else {
		return ""
	}
}
func (c CommentService) LikeReply(uid uint, commentid uint) string {
	if er := commentdao.LikeReply(uid, commentid); er != nil {
		return er.Error()
	} else {
		return ""
	}
}
