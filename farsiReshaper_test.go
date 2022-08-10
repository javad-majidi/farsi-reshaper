package FarsiReshaper

import "testing"

type sampleFarsiToBeReshaped struct {
	got, wanted string
}

var sampleFarsiToBeTested = []sampleFarsiToBeReshaped{
	sampleFarsiToBeReshaped{got: Reshape("سلام"), wanted: "ﻡﻼﺳ"},
	sampleFarsiToBeReshaped{got: Reshape("چطوری شما؟"), wanted: "؟ﺎﻤﺷ ﯼﺭﻮﻄﭼ"},
	sampleFarsiToBeReshaped{got: Reshape("گچ کاری پژمان"), wanted: "ﻥﺎﻣﮋﭘ ﯼﺭﺎﮐ ﭻﮔ"},
	sampleFarsiToBeReshaped{got: Reshape("در این پروژه از bobardo.com استفاده شده است"), wanted: "ﺖﺳﺍ ﻩﺪﺷ ﻩﺩﺎﻔﺘﺳﺍ bobardo.com ﺯﺍ ﻩﮊﻭﺮﭘ ﻦﯾﺍ ﺭﺩ"},
	sampleFarsiToBeReshaped{got: Reshape("This is sample english for sanity check"), wanted: "This is sample english for sanity check"},
}

func TestReshape(t *testing.T) {
	for _, element := range sampleFarsiToBeTested {
		if element.got != element.wanted {
			t.Errorf("Output %q not equal to expected %q", element.got, element.wanted)
		}
	}
}
