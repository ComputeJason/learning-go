package maps


const (
	ErrWordNotFound DictionaryErr = DictionaryErr("could not find the word you were looking for")
	ErrWordAlreadyExists DictionaryErr = DictionaryErr("cannot add word because it already exists")
	ErrWordNotFoundForOperation DictionaryErr = DictionaryErr("cannot perform operation on word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}



type Dictionary map[string]string


func (d Dictionary) Search(s string) (string, error) {
	
	definition, ok := d[s]

	if !ok {
		return "", ErrWordNotFound
	}
	
	return definition, nil
}

func (d Dictionary) Add (key,val string) error {

	// here because MAP is a reference type (not a struct), its always implicitly passed by reference. The original will always be editted 
	// hence we don't need to use a pointer reciever type 

	_, err := d.Search(key)

	switch err {
	case ErrWordNotFound:
		d[key] = val
	case nil:
		return ErrWordAlreadyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, val string) error {

	_, err := d.Search(key)

	switch err {
	case nil:
		d[key] = val
	case ErrWordNotFound: 
		return ErrWordNotFoundForOperation
	default:
		return err
	}

	return nil

}

func (d Dictionary) Delete(key string) error {

	_, err := d.Search(key)

	switch err {
	case nil:
		delete(d,key)
	case ErrWordNotFound: 
		return ErrWordNotFoundForOperation
	default:
		return err
	}

	return nil
}
