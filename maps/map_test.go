package main

import "testing"

func TestMap(t *testing.T) {
	dictionary := Dictionnaire{"test": "this is just a test", "true": "it's a true"}

	t.Run("Know word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknow word", func(t *testing.T) {
		_, got := dictionary.Search("uknow")
		want := "couldn't find the  word you were looking for"
		if got == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t , got , ErrNotFounds)
		assertStrings(t, got.Error(), want)
	})

}

//fonction pour verifier le want et le got
func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q ", got, want)
	}
}

//fonction pour afficher l'erreur
func assertError(t  testing.TB, got , want error ){
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

//fonction pour ajouter
func TestAdd(t *testing.T){
	

	t.Run("new word", func(t *testing.T){
		dictionnaire := Dictionnaire{}
		word := "test"
		definition := "this is just a test"
		err := dictionnaire.Add(word , definition)

		assertError(t , err , nil)
		assertDefinition(t , dictionnaire ,word , definition)
		
	})


	t.Run("existing word" , func(t *testing.T){
		dictionnaire := Dictionnaire{"test" : "this is just a test"}
		word := "test"
		definition := "this is just a test"
		
		err := dictionnaire.Add(word , definition)
		assertError(t ,  err , ErrExist)
		assertDefinition(t , dictionnaire , word , definition)
	})
}

//fonction pour chercher le mot et la definition
func assertDefinition(t testing.TB , dictionary Dictionnaire ,word , definition string){
	t.Helper()

	got , err := dictionary.Search(word)

	if err != nil{
		t.Fatal("should find added word: ", err)
	}

	assertStrings(t , got , definition)
}

//pour la mise a jour
func TestUpdate(t *testing.T){

	t.Run("existing word" , func(t *testing.T) {
		word := "test"
		definition := "this is just exemple"
		dictionnaire := Dictionnaire{word : definition}

		newDefinition := "new definition"

		dictionnaire.Update(word , newDefinition)

		assertDefinition(t , dictionnaire ,word , newDefinition)

	})

	t.Run("new Word" , func(t *testing.T) {
		word := "test"
		definition := "this is just exemple"
		dictionnaire := Dictionnaire{}
		dictionnaire.Add(word, definition)

		 assertDefinition(t, dictionnaire, word, definition)
	})
	
}

func TestDelete(t *testing.T){
	word := "test"
	definition := "It is just test"
	dictionnaire := Dictionnaire{word : definition}

	dictionnaire.Delete(word)

	_ , err := dictionnaire.Search(word)
	assertError(t , err , ErrNotFounds)

}