package model

import(
	"strconv"
	"github.com/gin-gonic/gin"
	"log"
)


type User struct {
	Id             string    `json:"id"`
	Account        string    `json:"account"`
	Nickname       string    `json:"nickname"`
	Password       string    `json:"password"`
	Avatar         string    `json:"avatar"`
	Energy         int       `json:"energy"`
}


func NewUser(user User) string {
	
	if err := database.DB.Table("tbl_user").Create(&user); err != nil {
		log.Println("NewUser" + err.Error())
	     return
	}
	return user.Nickname
}






//登录验证
func VerifyPassword(id string, password string) bool {
	var user = User
	if err := database.DB.Table("tbl_user").Where("account=?", id).First(&user).Error; err != nil {
		log.Println("VerifyPasswordError" + err.Error())
		return false
	}
	if user.Password == password {
		return true
	}
	log.Println(user.Password)
	log.Println(password)
	return false
}

const (
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin    = "请重新登陆"
)


//解析token
func VerifyToken(strToken string) (*JwtClaims, error) {
	//	strToken := c.Param("token")    //c.Param()可以获取单个参数,路径的也行

	//解析
	token, err := jwt.ParseWithClaims(strToken, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	if err != nil {
		return nil, errors.New(ErrorReason_ServerBusy + ",或token解析失败")
	}
	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReason_ReLogin)
	}

	return claims, nil //JwtClaims结构体类型
}




type Idea struct{
	IdeaId         int       `json:"idea_id"`
    Content        string    `json:"content"`
	ReleaseDate	   string    `json:"releaseDate"`
	PublisherId    string 	 `json:"publisher_id"`
	likessum	   int 	     `json:"likes_sum"`
	CommentSum	   int       `json:"comment_sum"`
	CollectionSum  int 		 `json:"collection_sum"`
}
 


func NewIdea (idea Idea) int {
	
	if err := database.DB.Table("tbl_idea").Create(&idea).Error; err != nil {
		log.Println("NewIdea " + err.Error())
		return 0
	}
	return idea.IdeaId
}


func DeleteIdea (idea Idea) string {
	const result = "删除成功"
	if err := database.DB.Table("tbl_idea").Delete((&idea,1).Error; err != nil {
		log.Println("DeleteIdea " + err.Error())
		return 
	}
	return   result

}



type Comment struct{
	Id             int       `json:"id"`
	CommenterId    string    `json:"commenter_id"`
	CommentedId    string    `json:"commented_id"`
	likessum       string    `json:"likes_sum"`
	ReleaseDate    string    `json:"release_date"`
	Content        string    `json:"content"`

}


func NewComment(comment  Comment) int {
	
	if err := database.DB.Table("tbl_comments").Create(&comment); err != nil {
		log.Println("NewComment" + err.Error())
		return 0
	}
	return Comment.Id
}






type Collection struct{
	Id             int       `json:"id"`
	CollectorId    int       `json:"collector_id"`
	IdeaId         int       `json:"idea_id"`

}

func NewCollection(collection Collection) int {
	
	if err := database.DB.Table("tbl_favorite_records").Create(&collection); err != nil {
		log.Println("NewCollection" + err.Error())
		return 0
	}
	return collection.Id
}






type IdeaLikeRecord  struct{
	Id             int       `json:"id"`
	IdeaId         int       `json:"idea_id"`
	LikersId       string    `json:"likers_id"`
	BelikedId      string    `json:"beliked_id"`
    
}

func NewIdeaLike(ideaLikeRecord IdeaLikeRecord) int {
	
	if err := database.DB.Table("tbl_like_record_idea").Create(&ideaLikeRecord); err != nil {
		log.Println("NewIdeaLike" + err.Error())
		return 0
	}
	return ideaLikeRecord.Id
}




type CommentLikeRecord  struct{
	Id             int       `json:"id"`
	CommentId      int       `json:"comment_id"`
	LikersId       string    `json:"likers_id"`
	BelikedId      string    `json:"beliked_id"`
}


func NewCommentLike(commentLikeRecord CommentLikeRecord) int {
	
	if err := database.DB.Table("tbl_like_record_comment").Create(&commentLikeRecord); err != nil {
		log.Println("NewCommentLike" + err.Error())
		return 0
	}
	return commentLikeRecord.Id
}








JwtCliams...
type JwtClaims struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
}

var Secret = "sault" //加盐

func GetToken(claims *JwtClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		log.Println(err)
		return ""
	}
	return signedToken
}

const (
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin    = "请重新登陆"
)

func VerifyToken(strToken string) (*JwtClaims, error) {
	//	strToken := c.Param("token")    //c.Param()可以获取单个参数,路径的也行

	//解析
	token, err := jwt.ParseWithClaims(strToken, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	if err != nil {
		return nil, errors.New(ErrorReason_ServerBusy + ",或token解析失败")
	}
	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReason_ReLogin)
	}

	return claims, nil //JwtClaims结构体类型
}