package model

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/project-flogo/rules/common"
)

func TestOne(t *testing.T) {

	td1 := TupleDescriptor{}
	td1.Name = "a"
	td1.TTLInSeconds = 10
	td1p1 := TuplePropertyDescriptor{}
	td1p1.Name = "p1"
	td1p1.KeyIndex = 3
	td1p1.PropType = data.TypeDouble
	td1p2 := TuplePropertyDescriptor{}
	td1p2.Name = "p2"
	td1p2.KeyIndex = 31
	td1p2.PropType = data.TypeString

	td1.Props = []TuplePropertyDescriptor{td1p1, td1p2}

	str, _ := json.Marshal(&td1)
	fmt.Printf("succes %s\n", str)

	tpdx := TupleDescriptor{}
	tpdx.TTLInSeconds = -1
	json.Unmarshal([]byte(str), &tpdx)

	str1, _ := json.Marshal(&tpdx)
	fmt.Printf("succes %s\n", str1)

}

func TestTwo(t *testing.T) {
	tupleDescAbsFileNm := common.GetAbsPathForResource("src/github.com/project-flogo/rules/examples/rulesapp/rulesapp.json")
	tupleDescriptor := common.FileToString(tupleDescAbsFileNm)

	fmt.Printf("Loaded tuple descriptor: \n%s\n", tupleDescriptor)
	//First register the tuple descriptors
	RegisterTupleDescriptors(tupleDescriptor)

}
