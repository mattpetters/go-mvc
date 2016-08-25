package viewmodels

type Lights struct {
  Title string
  Active string
  Lights []Light
}

type Light struct {
  Room string
  Description string
  IsOn bool
}

func GetLights() Lights {
	result := Lights{
		Title:  "Vedado Control Center - Lights",
		Active: "lights",
	}
  light1 := Light {
    Room: "Living Room",
    Description: "White light by TV",
    IsOn: false,
  }
  light2 := Light {
    Room: "Guest bathroom",
    Description: "In the bathroom",
    IsOn: true,
  }
  light3 := Light {
    Room: "Kitchen",
    Description: "Above the stove",
    IsOn: true,
  }

  result.Lights = []Light{
  light1,
  light2,
  light3,
  }

	return result
}
