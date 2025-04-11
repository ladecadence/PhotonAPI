package main

import (
	"crypto/sha256"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"

	// "os"
	"net/http"

	"github.com/ladecadence/PhotonAPI/pkg/color"
	"github.com/ladecadence/PhotonAPI/pkg/config"
	"github.com/ladecadence/PhotonAPI/pkg/database"
	"github.com/ladecadence/PhotonAPI/pkg/models"
	"github.com/ladecadence/PhotonAPI/pkg/routes"
)

//go:embed assets/wall0003.jpg
var img []byte

func initData(db database.Database) {
	user := models.User{
		Name:     "testuser",
		Password: fmt.Sprintf("%x", sha256.Sum256([]byte("testpassword"))),
		Email:    "test@email.com",
		Role:     models.UserRoleUser,
	}
	err := db.UpsertUser(user)
	if err != nil {
		panic("Error Inserting initial data: " + err.Error())
	}
	user = models.User{
		Name:     "testadmin",
		Password: fmt.Sprintf("%x", sha256.Sum256([]byte("testAdminpassword"))),
		Email:    "admin@email.com",
		Role:     models.UserRoleAdmin,
	}
	err = db.UpsertUser(user)
	if err != nil {
		panic("Error Inserting initial data: " + err.Error())
	}

	wall := models.Wall{
		Uid:         "1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05",
		Name:        "WallTBK",
		Description: "Best test wall",
		Adjustable:  true,
		DegMin:      5.0,
		DegMax:      40.0,
		Image:       img,
		ImgW:        961,
		ImgH:        1280,
		Holds:       "[{'id':0,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':257.0,'y':109.0},{'id':1,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':368.0,'y':94.0},{'id':2,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':474.0,'y':96.0},{'id':3,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':586.0,'y':99.0},{'id':4,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':698.0,'y':98.0},{'id':5,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':696.0,'y':208.0},{'id':6,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':580.0,'y':233.0},{'id':7,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':475.0,'y':230.0},{'id':8,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':361.0,'y':225.0},{'id':9,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':251.0,'y':221.0},{'id':10,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':247.0,'y':352.0},{'id':11,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':364.0,'y':345.0},{'id':12,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':481.0,'y':361.0},{'id':13,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':586.0,'y':351.0},{'id':14,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':704.0,'y':357.0},{'id':15,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':698.0,'y':486.0},{'id':16,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':585.0,'y':466.0},{'id':17,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':473.0,'y':468.0},{'id':18,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':372.0,'y':471.0},{'id':19,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':259.0,'y':472.0},{'id':20,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':262.0,'y':583.0},{'id':21,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':375.0,'y':593.0},{'id':22,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':485.0,'y':581.0},{'id':23,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':593.0,'y':586.0},{'id':24,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':691.0,'y':584.0},{'id':25,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':691.0,'y':701.0},{'id':26,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':574.0,'y':700.0},{'id':27,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':474.0,'y':702.0},{'id':28,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':369.0,'y':699.0},{'id':29,'size':40,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':263.0,'y':714.0},{'id':30,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':264.0,'y':810.0},{'id':31,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':368.0,'y':815.0},{'id':32,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':473.0,'y':804.0},{'id':33,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':585.0,'y':820.0},{'id':34,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':701.0,'y':821.0},{'id':35,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':689.0,'y':936.0},{'id':36,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':580.0,'y':935.0},{'id':37,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':479.0,'y':926.0},{'id':38,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':370.0,'y':926.0},{'id':39,'size':30,'type':0,'wallid':'1fddf17c-3ddf-4dc7-a3d0-e3ac3d9f8b05','x':264.0,'y':926.0}]",
	}
	err = db.UpsertWall(wall)
	if err != nil {
		panic("Error Inserting initial data: " + err.Error())
	}
}

func main() {
	// flags
	testDataFlag := flag.Bool("testdata", false, "load test data into database")
	configFileFlag := flag.String("conf", "config.toml", "config file path")
	flag.Parse()

	// read configuration
	config := config.Config{ConfFile: *configFileFlag}
	config.GetConfig()

	// open and init database
	database := &database.SQLite{}
	_, err := database.Open(config.Database)
	if err != nil {
		panic("Error opening DB: " + err.Error())
	}
	err = database.Init()
	if err != nil {
		panic("Error initializing DB: " + err.Error())
	}

	// load testa data in database if -testdata arg
	if *testDataFlag {
		initData(database)
	}

	// walls, err := database.GetWalls()
	// for _, w := range walls {
	// 	fmt.Printf("%v :  %v\n", w.Uid, w.Name)
	// 	err = os.WriteFile(w.Uid+".jpg", w.Image, 0644)
	// }

	// create server mux and routes
	mux := http.NewServeMux()
	routes.RegisterRoutes(database, config, mux)

	const logo = `
	
░▒▓███████▓▒░░▒▓█▓▒░░▒▓█▓▒░░▒▓██████▓▒░▒▓████████▓▒░▒▓██████▓▒░░▒▓███████▓▒░ ░▒▓██████▓▒░░▒▓███████▓▒░░▒▓█▓▒░ 
░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░  ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░ 
░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░  ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░ 
░▒▓███████▓▒░░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░  ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓████████▓▒░▒▓███████▓▒░░▒▓█▓▒░ 
░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░  ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░ 
░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▒░  ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░ 
░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░░▒▓██████▓▒░  ░▒▓█▓▒░   ░▒▓██████▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░ 
                                                                                                                                                                                                               
	`
	fmt.Fprintf(os.Stderr, "%s", color.Cyan+logo+color.Reset)
	log.Printf("📡 "+color.Green+"PhotonAPI version "+color.Purple+"%s"+color.Green+" listening on port "+color.Yellow+"%d"+color.Reset, config.Version, config.Port)

	// launch server
	addr := fmt.Sprintf("%s:%d", config.Addr, config.Port)
	log.Fatal(http.ListenAndServe(addr, mux))

}
