package models

type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	Data            map[string]interface{}
	FloatMap        map[string]float32
	IsAuthenticated int
	Username        string
	No_Of_Likes     int
}
