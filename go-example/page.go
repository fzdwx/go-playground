package page

import "io/ioutil"

type Page struct {
	Title string
	Body  []byte
}

// save writes the page to the file.
func (this Page) save() error {
	return ioutil.WriteFile(this.Title, this.Body, 0666)
}

// load reads the page from the file.
func load(p *Page) (err error) {
	body, err := ioutil.ReadFile(p.Title)
	if err != nil {
		return err
	}
	p.Body = body
	return nil
}
