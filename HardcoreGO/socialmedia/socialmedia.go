package socialmedia

import "time"

//go:generate stringer -type=MoodState
type MoodState		int

// Here We define all the possible mood states using an
//  iota enumerator
const (
	MoodStateNeutral MoodState = iota
	MoodStateHappy
	MoodStateSad
	MoodStateAngry
	MoodStateHopeful
	MoodStateThrilled
	MoodStateBored
	MoodStateShy
	MoodStateOnCloudNine
	MoodStateComical
)

// This is a type we embed into types we want to keep
//   a check on for auditing purposes
type AuditableContent struct {
	TimeCreated		time.Time
	TimeModified 	time.Time
	CreatedBy 		string
	ModifiedBy 		string
}

// This is the type that represents a Social Media Post
type Post struct{
	AuditableContent	// Embeded type
	Caption			 	string
	MessageBody 		string
	URL 				string
	ImageURI			string
	ThumbnailURI		string
	Keywords			[]string
	Likers 				[]string
	AuthorMood			MoodState
}

// Map that hols the various Mood States with keys to server as
//  aliases to their respective mod state
var Moods map[string]MoodState

// The init() function is responsible for initializing the mood state
func init(){
	Moods = map[string]MoodState{
		"neutral"	: MoodStateNeutral,
		"happy"		: MoodStateHappy,
		"sad"		: MoodStateSad,
		"angry"		: MoodStateAngry,
		"hopeful"	: MoodStateHopeful,
		"thrilled"	: MoodStateThrilled,
		"bored"		: MoodStateBored,
		"shy"		: MoodStateShy,
		"comical"	: MoodStateComical,
		"cloudnine" : MoodStateOnCloudNine,
	}
}

// This is the function responsible for creating a new Social Media Post
func NewPost(
	username string,
	mood MoodState,
	caption string,
	messageBody string,
	url string,
	imageURI string,
	thumbnailURI string,
	keywords []string,
) *Post{
	auditableContent := AuditableContent{
		CreatedBy: username,
		TimeCreated: time.Now(),
	}

	return &Post{
		Caption:  		caption,
		MessageBody:	messageBody,
		URL:			url,
		ImageURI: 		imageURI,
		ThumbnailURI: 	thumbnailURI,
		AuthorMood: 	mood,
		Keywords:		keywords,
		AuditableContent: auditableContent,
	}
}