package main
   import(
	
	
    
    "github.com/2021-ZeroGravity-backend/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func main() {
	
	model.DB,_ = gorm.Open("mysql", "Debian-sys-maint/O6sbcOfIxG2o33qA@(127.0.0.1:3306)/ZeroGravity?parseTime=True")
	
}
