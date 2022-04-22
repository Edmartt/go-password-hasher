package tests

import (
	"reflect"
	"github.com/Edmartt/go-password-hasher/hasher"
	"testing"
)


func TestRandomSalt(t *testing.T){
	text1 := hasher.ConvertToHash("testing")
	text2 := hasher.ConvertToHash("testing")

	if text1 == text2 {
		t.Error()
	}
}


func TestEmptyStringToHash(t *testing.T){
	stringToHash := "" 

	if stringToHash != "" {
		t.Error()
	}
}


func TestStringToHashLenght(t *testing.T){
	text := "12345678"

	if len(text) < 8{
		t.Error()
	}
}


func TestIsString(t *testing.T){
	myInput := "hello world"

	if reflect.TypeOf(myInput).String() != "string"{
		t.Error()
	}
}


func TestIsNotString(t *testing.T){
	myInput := 50

	if reflect.TypeOf(myInput).String() == "string"{
		t.Error()
	}
}


func TestHashEquals(t *testing.T){
	currentHash := hasher.ConvertToHash("Hello")
	isTheSame := hasher.CheckHash(currentHash, "Hello")

	if isTheSame != true{
		t.Error()
	}
}


func TestHashNotEqual(t *testing.T){
	currentHash := hasher.ConvertToHash("Hello")
	isDifferent := hasher.CheckHash(currentHash, "hello")

	if isDifferent != false{
		t.Error()
	}
}
