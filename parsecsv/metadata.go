package parsecsv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/amarnath262/testcsv/action/flow/test"
	"github.com/amarnath262/testcsv/core/activity"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	fields := make([]interface{}, 3)
	fields[0] = "test"
	fields[1] = "test2"
	fields[2] = "test3"
	tc.SetInput("fieldNames", fields)
	tc.SetInput("delimiter", ",")
	tc.SetInput("csv", "hello,my,name\ntest,test,test")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	//check result attr
	b, _ := json.Marshal(tc.GetOutput("output"))
	fmt.Println(string(b))
}

func TestFile(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	fields := make([]interface{}, 2)
	fields[0] = "test"
	fields[1] = "test2"
	tc.SetInput("fieldNames", fields)
	tc.SetInput("file", "./test.csv")
	tc.SetInput("delimiter", ",")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	//check result attr
	b, _ := json.Marshal(tc.GetOutput("output"))
	fmt.Println(string(b))
}

func TestWithMissmatchFields(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	fields := make([]interface{}, 2)
	fields[0] = "test"
	fields[1] = "test2"
	tc.SetInput("fieldNames", fields)
	tc.SetInput("delimiter", ",")
	tc.SetInput("csv", "hello,my,name\ntest,test,test")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	//check result attr
	b, _ := json.Marshal(tc.GetOutput("output"))
	fmt.Println(string(b))
}

func TestWithQuotedstrings(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	fields := make([]interface{}, 3)
	fields[0] = "test"
	fields[1] = "test2"
	fields[2] = "test3"
	tc.SetInput("fieldNames", fields)
	tc.SetInput("delimiter", ",")
	tc.SetInput("csv", "\"hello\",\"my name\",name\ntest,\"test, data\",test")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	//check result attr
	b, _ := json.Marshal(tc.GetOutput("output"))
	fmt.Println(string(b))
}
