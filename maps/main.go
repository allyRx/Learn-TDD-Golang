package main


type Dictionnaire map[string]string
type DictionnaireErr string

const (
	ErrNotFounds = DictionnaireErr("couldn't find the  word you were looking for")
 	ErrExist = DictionnaireErr("word already in dictionary")
)


func (e DictionnaireErr) Error() string {
	return string(e)
}
func (d Dictionnaire) Search(word string) (string, error) {

	definition, ok := d[word]

	if !ok {
		return "", ErrNotFounds
	}

	return definition , nil
}


func (d Dictionnaire) Add(word , definition string) (error){
		_ , err := d.Search(word)

		switch err {
			case ErrNotFounds :
				d[word] = definition
			case nil : 
				return ErrExist
			default:
				return err		

		}
	return nil	
}

func (d Dictionnaire) Update(word , newDefinition string) (error){
	_ , err := d.Search(word)
	switch err {
		case ErrNotFounds :
			d.Add(word , newDefinition)
		case nil :
			d[word] = newDefinition	
		default :
			return err	
	}
	
	return nil
}

func (d Dictionnaire) Delete(word string) {
	delete(d , word)
}