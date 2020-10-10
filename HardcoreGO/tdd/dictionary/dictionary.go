package dictionary

import (
	"errors"
)

type Dictionary map[string]string

func (d Dictionary)Search(word string) (string,error){
	value, ok := d[word]

	if !ok {
		return "", errors.New("word not found in dictionary")
	}

	return value, nil
}

func (d Dictionary)Add(word string, defination string) error{

	_, err := d.Search(word)

	if err == nil{
		return errors.New("Word already Exist")
	}

	d[word] = defination
	return nil
}