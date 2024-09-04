package basic

import (
	"fmt"

	"github.com/VallabhLakadeTech/golang/grpc/my-protobuf/protogen/basic"
	"google.golang.org/protobuf/encoding/protojson"
)

func PrintUser() {

	user := basic.User{
		Id:       1,
		Username: "Vallabh",
		IsActive: true,
		Password: []byte("somepassword"),
		Emails:   []string{"vallabh.lakade@gmail.com", "vallabh.lakade.1990@gmail.com"},
		Gender:   basic.Gender_GENDER_MALE,
	}
	fmt.Println("User:", user)
}

func ProtoToUserJson() {
	user := basic.User{
		Id:       1,
		Username: "Vallabh",
		IsActive: true,
		Password: []byte("somepassword"),
		Emails:   []string{"vallabh.lakade@gmail.com", "vallabh.lakade.1990@gmail.com"},
		Gender:   basic.Gender_GENDER_MALE,
	}
	userJsonBytes, err := protojson.Marshal(&user)
	if err != nil {
		fmt.Println("Some error in marshalling proto user", err)
		return
	}
	fmt.Println("User JSON: ", string(userJsonBytes))
}

func JSONToProtoUser() {
	json := `{"id":1,"username":"Vallabh","is_active":true,"password":"c29tZXBhc3N3b3Jk","emails":["vallabh.lakade@gmail.com","vallabh.lakade.1990@gmail.com"],"gender":"GENDER_MALE"}`
}
