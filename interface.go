package exporter


type SocialMedia interface{
	Feed() []string
	Fame() int
}


