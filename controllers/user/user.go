package user

import (
	//"github.com/astaxie/beego"
	"github.com/thanzen/eq/controllers/base"
	"github.com/thanzen/eq/models/user"
)

type UserController struct {
	base.BaseRouter
}

func (this *UserController) getUser(u *user.User) bool {
	username := this.GetString(":username")
	u.Username = username
	err := this.UserService.Read(u, "Username")
	if err != nil {
		this.Abort("404")
		return true
	}
	this.Data["TheUser"] = &u

	return false
}

func (this *UserController) Index() {
	//this.TplNames = "user/home.html"
	//this.TplNames = "index.html"
	var user user.User
	this.Data["TheUser"] = &user
	if this.getUser(&user) {
		return
	}
    this.Data["json"] = &user
    this.ServeJson()

	//limit := 5

	//var posts []*models.Post
	//var comments []*models.Comment

	//user.RecentPosts().Limit(limit).RelatedSel().All(&posts)
	//user.RecentComments().Limit(limit).RelatedSel().All(&comments)

	//var ftopics []*models.FollowTopic
	//var topics []*models.Topic
	//nums, _ := models.FollowTopics().Filter("User", &user.Id).Limit(8).OrderBy("-Created").RelatedSel("Topic").All(&ftopics, "Topic")
	//if nums > 0 {
	//	topics = make([]*models.Topic, 0, nums)
	//	for _, ft := range ftopics {
	//		topics = append(topics, ft.Topic)
	//	}
	//}
	//this.Data["TheUserTopics"] = topics
	//this.Data["TheUserTopicsMore"] = nums >= 8

	//this.Data["TheUserPosts"] = posts
	//this.Data["TheUserComments"] = comments	//limit := 5

	//var posts []*models.Post
	//var comments []*models.Comment

	//user.RecentPosts().Limit(limit).RelatedSel().All(&posts)
	//user.RecentComments().Limit(limit).RelatedSel().All(&comments)

	//var ftopics []*models.FollowTopic
	//var topics []*models.Topic
	//nums, _ := models.FollowTopics().Filter("User", &user.Id).Limit(8).OrderBy("-Created").RelatedSel("Topic").All(&ftopics, "Topic")
	//if nums > 0 {
	//	topics = make([]*models.Topic, 0, nums)
	//	for _, ft := range ftopics {
	//		topics = append(topics, ft.Topic)
	//	}
	//}
	//this.Data["TheUserTopics"] = topics
	//this.Data["TheUserTopicsMore"] = nums >= 8

	//this.Data["TheUserPosts"] = posts
	//this.Data["TheUserComments"] = comments	//limit := 5

	//var posts []*models.Post
	//var comments []*models.Comment

	//user.RecentPosts().Limit(limit).RelatedSel().All(&posts)
	//user.RecentComments().Limit(limit).RelatedSel().All(&comments)

	//var ftopics []*models.FollowTopic
	//var topics []*models.Topic
	//nums, _ := models.FollowTopics().Filter("User", &user.Id).Limit(8).OrderBy("-Created").RelatedSel("Topic").All(&ftopics, "Topic")
	//if nums > 0 {
	//	topics = make([]*models.Topic, 0, nums)
	//	for _, ft := range ftopics {
	//		topics = append(topics, ft.Topic)
	//	}
	//}
	//this.Data["TheUserTopics"] = topics
	//this.Data["TheUserTopicsMore"] = nums >= 8

	//this.Data["TheUserPosts"] = posts
	//this.Data["TheUserComments"] = comments	//limit := 5

	//var posts []*models.Post
	//var comments []*models.Comment

	//user.RecentPosts().Limit(limit).RelatedSel().All(&posts)
	//user.RecentComments().Limit(limit).RelatedSel().All(&comments)

	//var ftopics []*models.FollowTopic
	//var topics []*models.Topic
	//nums, _ := models.FollowTopics().Filter("User", &user.Id).Limit(8).OrderBy("-Created").RelatedSel("Topic").All(&ftopics, "Topic")
	//if nums > 0 {
	//	topics = make([]*models.Topic, 0, nums)
	//	for _, ft := range ftopics {
	//		topics = append(topics, ft.Topic)
	//	}
	//}
	//this.Data["TheUserTopics"] = topics
	//this.Data["TheUserTopicsMore"] = nums >= 8

	//this.Data["TheUserPosts"] = posts
	//this.Data["TheUserComments"] = comments
}
