package yandex

// https://yandex.ru/dev/rasp/doc/reference/schedule-point-point.html

type TrainSchedule struct {
	IntervalSegments []interface{} `json:"interval_segments"`
	Pagination       struct {
		Total  int `json:"total"`
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	} `json:"pagination"`
	Segments []struct {
		ExceptDays string `json:"except_days"`
		Arrival    string `json:"arrival"`
		From       struct {
			Code            string `json:"code"`
			Title           string `json:"title"`
			StationType     string `json:"station_type"`
			PopularTitle    string `json:"popular_title"`
			ShortTitle      string `json:"short_title"`
			TransportType   string `json:"transport_type"`
			StationTypeName string `json:"station_type_name"`
			Type            string `json:"type"`
		} `json:"from"`
		Thread struct {
			UID              string `json:"uid"`
			Title            string `json:"title"`
			Number           string `json:"number"`
			ShortTitle       string `json:"short_title"`
			ThreadMethodLink string `json:"thread_method_link"`
			Carrier          struct {
				Code     int         `json:"code"`
				Contacts string      `json:"contacts"`
				URL      string      `json:"url"`
				LogoSvg  interface{} `json:"logo_svg"`
				Title    string      `json:"title"`
				Phone    string      `json:"phone"`
				Codes    struct {
					Icao   interface{} `json:"icao"`
					Sirena interface{} `json:"sirena"`
					Iata   interface{} `json:"iata"`
				} `json:"codes"`
				Address string `json:"address"`
				Logo    string `json:"logo"`
				Email   string `json:"email"`
			} `json:"carrier"`
			TransportType    string      `json:"transport_type"`
			Vehicle          interface{} `json:"vehicle"`
			TransportSubtype struct {
				Color string `json:"color"`
				Code  string `json:"code"`
				Title string `json:"title"`
			} `json:"transport_subtype"`
			ExpressType string `json:"express_type"`
		} `json:"thread"`
		DeparturePlatform string `json:"departure_platform"`
		Departure         string `json:"departure"`
		Stops             string `json:"stops"`
		Days              string `json:"days"`
		To                struct {
			Code            string `json:"code"`
			Title           string `json:"title"`
			StationType     string `json:"station_type"`
			PopularTitle    string `json:"popular_title"`
			ShortTitle      string `json:"short_title"`
			TransportType   string `json:"transport_type"`
			StationTypeName string `json:"station_type_name"`
			Type            string `json:"type"`
		} `json:"to"`
		Duration          float64     `json:"duration"`
		DepartureTerminal interface{} `json:"departure_terminal"`
		ArrivalTerminal   interface{} `json:"arrival_terminal"`
		StartDate         string      `json:"start_date"`
		ArrivalPlatform   string      `json:"arrival_platform"`
	} `json:"segments"`
	Search struct {
		Date string `json:"date"`
		To   struct {
			Code         string `json:"code"`
			Type         string `json:"type"`
			PopularTitle string `json:"popular_title"`
			ShortTitle   string `json:"short_title"`
			Title        string `json:"title"`
		} `json:"to"`
		From struct {
			Code            string `json:"code"`
			Title           string `json:"title"`
			StationType     string `json:"station_type"`
			PopularTitle    string `json:"popular_title"`
			ShortTitle      string `json:"short_title"`
			TransportType   string `json:"transport_type"`
			StationTypeName string `json:"station_type_name"`
			Type            string `json:"type"`
		} `json:"from"`
	} `json:"search"`
}
