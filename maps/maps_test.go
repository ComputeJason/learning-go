package maps

import "testing"

func assertStringAreEqual(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertErrorsAreEqual(t testing.TB, got, want error){
	t.Helper()
	if got.Error() != want.Error(){
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertErrorIsNil(t testing.TB, err error){
	t.Helper()
	if err != nil {
		t.Fatal("expected error to be nil, but its not")
	}
}

func TestSearch(t *testing.T){

	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStringAreEqual(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_, err := dictionary.Search("")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertErrorsAreEqual(t, err, ErrWordNotFound)

	})
}

func TestAdd (t *testing.T){
	t.Run("test add new word", func(t *testing.T) {

		dict := Dictionary{}
		err := dict.Add("test","this is just a test")

		assertErrorIsNil(t,err)

		got, _ := dict.Search("test")
		want := "this is just a test"

		assertStringAreEqual(t, got, want)
	})

	t.Run("test add existing word", func(t *testing.T) {

		dict := Dictionary{"test":"existing test"}
		err := dict.Add("test","this is just a test")

		assertErrorsAreEqual(t,err, ErrWordAlreadyExists)

		got, _ := dict.Search("test")
		want := "existing test"

		assertStringAreEqual(t, got, want)


		// gotcha is that you can READ from a NIL map but WRITE will cause panic 
	})
}

func TestUpdate(t *testing.T){

	t.Run("test update existing key", func(t *testing.T) {

		dict := Dictionary{"test":"this is a test"}

		err := dict.Update("test","this is a updated test")

		assertErrorIsNil(t,err)

		got, _ := dict.Search("test")
		want := "this is a updated test"

		assertStringAreEqual(t, got, want)

	})

	t.Run("test update non-existing key", func(t *testing.T) {
		
		dict := Dictionary{}

		err := dict.Update("test","this is a updated test")

		assertErrorsAreEqual(t,err, ErrWordNotFoundForOperation)

		_ , err = dict.Search("test")

		assertErrorsAreEqual(t, err, ErrWordNotFound)

	})

}

func TestDelete(t *testing.T){
	t.Run("test delete existing word", func(t *testing.T) {
		dict := Dictionary{"test":"this is a test"}

		err := dict.Delete("test")

		assertErrorIsNil(t,err)

		_, err = dict.Search("test")

		assertErrorsAreEqual(t, err, ErrWordNotFound)
	})

	t.Run("test delete existing word", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Delete("test")

		assertErrorsAreEqual(t,err, ErrWordNotFoundForOperation)
	})
}