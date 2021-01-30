package main
   import (
	
	

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func main() {
	 model.DB, err = gorm.Open("mysql", "Debian-sys-maint/O6sbcOfIxG2o33qA@(127.0.0.1:3306)/ZeroGravity?parseTime=True")
	 if err != nil {
		panic(err)
	}
}
