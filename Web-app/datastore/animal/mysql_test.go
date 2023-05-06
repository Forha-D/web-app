package animal

import (
	"database/sql"
	"os"
	"reflect"
	"testing"
	"github.com/Forha-D/web-app/Web-app/driver"
	"github.com/Forha-D/web-app/Web-app/entities"
)

func initializemySQL(t *testing.T) *sql.DB {

	config := driver.MySQLConfig {
		Host :os.Getenv ("SQL_HOST"),
		User :os.Getenv ("SQL_USER"),
		Password :os.Getenv ("SQL_PASSWORD"),
		Port :os.Getenv ("SQL_PORT"),
		Db :os.Getenv ("SQL_DB"),
	}

	var err error

	db,err := driver.ConnectToMySQL(config)

	if err != nil {
		t.Errorf ("Could not connect to sql, err:%v", err)
	}

	return db


}

func testAnimalStorer_Create(t *testing.T, db AnimalStorer){

testcases := [] struct  {
	req           entities.Animal
	response      entities.Animal
}{
	{entities.Animal{ Name : "Hen", Age: 1}, entities.Animal {4, "Hen", 1}},
     {entities.Animal{ Name : "Duck", Age: 3}, entities.Animal {6, "Hen", 3}},
}

for i,v := range testcases {
	 resp,_ := db.Create(v.req)

            if !reflect.DeepEqual(resp, v.response) {
            t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.response)
}

}

}

func testAnimalstorer_Get( t *testing.T, db AnimalStorer){
testcases := [] struct {
	id int
	resp []entities.Animal
}{

	{ 0, []entities.Animal{{1, "Hippo", 10}, {2, "Ele", 20}}},
	{1, []entities.Animal{{1, "Hippo", 10}}},

    if !reflect.DeepEqual(resp, v.resp){
                t.Errorf ("[TEST%d]Failed. Got %v\tExpected:%v\n", i+1, resp, v.resp)
	}
}

}




func TestDatastore (t *testing.T) {
	db := initializemySQL(t)

	a := New(db)

	testAnimalstorer_Get(t, a)
	testAnimalStorer_Create(t, a)
}