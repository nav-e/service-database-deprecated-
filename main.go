package main
import(
  "github.com/sitaramshelke/service-database/dbhelper"
  _ "fmt"
  "github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
  _ "github.com/sitaramshelke/service-database/pbf_import"
)

func main(){
	
  // dbhelper.SetupPGDatabase()		//setup the database
	
	// pbf_import.ImportFromHttp()		//reads from the pbf file and exports to the database
	
  api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/nodes", dbhelper.GetAllNodeData),
    rest.Get("/node", dbhelper.GetNodeData),
		rest.Get("/greennav/vehicles",dbhelper.GetAllVehicles),
		// rest.Post("/device",AddDevice),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":6833", api.MakeHandler()))
  
}
