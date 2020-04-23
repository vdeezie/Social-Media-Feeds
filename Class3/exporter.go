package class3

//SocialMedia interface
type SocialMedia interface {
	Feed() []string 
	Fame() int
}

//Facebook struct
type Facebook struct {
	URL 	string `json: "URL"` 		
	Name 	string `json: "Name"`		
	Friends int `json: "Friends"`		

}

//Linkedin struct type
type Linkedin struct {
	URL 		string `json: "URL"`		
	Name		string `json: "Name"` 		
	Connections	int `json: "Connections"` 	

}

//Twitter takes the struct type
type Twitter struct {
	URL 		string `json: "URL"` 		
	Username 	string `json: "Username"` 	
	Followers 	int `json: "Followers"`		

}

// Feed returns the latest Facebook posts
func (f *Facebook) Feed() []string {
	return []string{
		"Facebook feeds",
		"Hey, here's my cool new selfie",
		"What is code?",
	}

}

// Fame tell how famous a user is. The higher
// the number of friends the more famous
func (f *Facebook) Fame() int {
	return f.Friends
}

// Feed returns the latest Linkedin posts
func (l *Linkedin) Feed() []string {
	return []string{
		"LinkedIn feeds",
		"Hey, I just started a new position at Hotels.ng",
	}
}

// Fame tell how famous a user is. The higher
// the number of Connections the more famous
func (l *Linkedin) Fame() int {
	return l.Connections
}

// Feed returns the latest Twitter posts
func (t *Twitter) Feed() []string {
	return []string{
		"Twitter feeds",
		"Coding is basically copying other people's work",
	}
}

// Fame tell how famous a user is. The higher
// the number of followers the more famous
func (t *Twitter) Fame() int {
	return t.Followers
}