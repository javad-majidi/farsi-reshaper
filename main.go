package FarsiReshaper

import (
	"github.com/robertkrimen/otto"
)

func Reshape(input string) string {
	vm := otto.New()
	_, err := vm.Run(`
//default Values:
var defLang = "fa"; //set default language:	["ar":arabic, "fa":Persian, "en":english]
var e_numbers = 1; //enable/disable arabic numbers: [0, 1]
var f_numbers = 1; //enable/disable persian numbers: [0, 1]
var e_harakat = 1; //enable/disable arabic harakat: [0, 1]
var dir = "rtl";

//initialize global vars:
var y = "";
var g;
var old = "";
var tstr = "";
var csr1 = (csr2 = 0); //cursor location
var laIndex = 168; //position of laa characters in the unicode string

var left = "ڤـئظشسيیبلپتنمكکگطضصثقفغعهخچحج"; //defining letters that can connect from the left
var right = "ڤـئؤرلالآىیآةوزژظشسيپبللأاأتنمكکگطضصثقفغعهخحچجدذلإإ"; //defining letters that can connect from the right
var arnumbs = "٠١٢٣٤٥٦٧٨٩";
var fanumbs = "۰۱۲۳۴۵۶۷۸۹";
var ennumbs = "0123456789";
var harakat = "ًٌٍَُِّْ"; //defining the harakat
var symbols = "ـ.،؟ @#$%^&*-+|\/=~,:"; //defining other symbols
var unicode =
    "ﺁﺁﺂﺂ" +
    "ﺃﺃﺄﺄ" +
    "ﺇﺇﺈﺈ" +
    "ﺍﺍﺎﺎ" +
    "ﺏﺑﺒﺐ" +
    "ﺕﺗﺘﺖ" +
    "ﺙﺛﺜﺚ" +
    "ﺝﺟﺠﺞ" +
    "ﺡﺣﺤﺢ" +
    "ﺥﺧﺨﺦ" +
    "ﺩﺩﺪﺪ" +
    "ﺫﺫﺬﺬ" +
    "ﺭﺭﺮﺮ" +
    "ﺯﺯﺰﺰ" +
    "ﺱﺳﺴﺲ" +
    "ﺵﺷﺸﺶ" +
    "ﺹﺻﺼﺺ" +
    "ﺽﺿﻀﺾ" +
    "ﻁﻃﻄﻂ" +
    "ﻅﻇﻈﻆ" +
    "ﻉﻋﻌﻊ" +
    "ﻍﻏﻐﻎ" +
    "ﻑﻓﻔﻒ" +
    "ﻕﻗﻘﻖ" +
    "ﻙﻛﻜﻚ" +
    "ﻝﻟﻠﻞ" +
    "ﻡﻣﻤﻢ" +
    "ﻥﻧﻨﻦ" +
    "ﻩﻫﻬﻪ" +
    "ﻭﻭﻮﻮ" +
    "ﻱﻳﻴﻲ" +
    "ﺓﺓﺔﺔ" +
    "ﺅﺅﺆﺆ" +
    "ﺉﺋﺌﺊ" +
    "ﻯﻯﻰﻰ" +
    "گﮔﮕﮓ" +
    "چﭼﭽﭻ" +
    "پﭘﭙﭗ" +
    "ژﮊﮋﮋ" +
    "ﯼﯾﯿﯽ" +
    "کﮐﮑﮏ" +
    "ﭪﭬﭭﭫ" +
    "ﻵﻵﻶﻶ" +
    "ﻷﻷﻸﻸ" +
    "ﻹﻹﻺﻺ" +
    "ﻻﻻﻼﻼ"; //defining arabic and persian unicode chars(individual, start, middle, end)
var arabic =
    "آ" +
    "أ" +
    "إ" +
    "ا" +
    "ب" +
    "ت" +
    "ث" +
    "ج" +
    "ح" +
    "خ" +
    "د" +
    "ذ" +
    "ر" +
    "ز" +
    "س" +
    "ش" +
    "ص" +
    "ض" +
    "ط" +
    "ظ" +
    "ع" +
    "غ" +
    "ف" +
    "ق" +
    "ك" +
    "ل" +
    "م" +
    "ن" +
    "ه" +
    "و" +
    "ي" +
    "ة" +
    "ؤ" +
    "ئ" +
    "ى" +
    "گ" +
    "چ" +
    "پ" +
    "ژ" +
    "ی" +
    "ک" +
    "ڤ";
var notEng = arabic + harakat + "ء،؟"; //defining all arabic letters + harakat + arabic symbols
var brackets = "(){}[]";

var msie = (opera = 0); //checking which browser is in use

var output_text = "";

function ProcessInput(text) {
    //the processing function
    old = "";
    tstr = "";
    y = text;
    x = y.split("");
    len = x.length;

    for (
        g = 0;
        g < len;
        g++ //process each letter, submit it to tests and then add it to the output string
    ) {
        //ignoring the harakat
        b = a = 1;
        while (harakat.indexOf(x[g - b]) >= 0) b = b + 1;
        while (harakat.indexOf(x[g + a]) >= 0) a = a + 1;
        if (g == 0) {
            //determine the position of each letter
            if (right.indexOf(x[a]) >= 0) pos = 1;
            else pos = 0;
        } else if (g == len - 1) {
            if (left.indexOf(x[len - b - 1]) >= 0) pos = 3;
            else pos = 0;
        } else {
            if (left.indexOf(x[g - b]) < 0) {
                if (right.indexOf(x[g + a]) < 0) pos = 0;
                else pos = 1;
            } else if (left.indexOf(x[g - b]) >= 0) {
                if (right.indexOf(x[g + a]) >= 0) pos = 2;
                else pos = 3;
            }
        }
        if (x[g] == "\n") {
            //if new line occurs, save old data in a temp, process new data, then regroup
            old = old + output_text + "\n";
            output_text = "";
        } else if (x[g] == "\r") {
            //if this char is carriage return, skip it.
        } else if (x[g] == "ء") addChar("ﺀ");
        else if (brackets.indexOf(x[g]) >= 0) {
            //if this char is a bracket, reverse it
            asd = brackets.indexOf(x[g]);
            if (asd % 2 == 0) addChar(brackets.charAt(asd + 1));
            else addChar(brackets.charAt(asd - 1));
        } else if (arabic.indexOf(x[g]) >= 0) {
            //if the char is an Arabic letter.. convert it to Unicode
            if (x[g] == "ل") {
                //if this letter is (laam)
                //check if its actually a (laa) combination
                ar_pos = arabic.indexOf(x[g + 1]);
                //alert(ar_pos)
                if (ar_pos >= 0 && ar_pos < 4) {
                    addChar(unicode.charAt(ar_pos * 4 + pos + laIndex));
                    g = g + 1;
                } //if its just (laam)
                else addChar(unicode.charAt(arabic.indexOf(x[g]) * 4 + pos));
            } //if its any arabic letter other than (laam)
            else addChar(unicode.charAt(arabic.indexOf(x[g]) * 4 + pos));
        } else if (symbols.indexOf(x[g]) >= 0)
            //if the char is a symbol, add it
            addChar(x[g]);
        else if (harakat.indexOf(x[g]) >= 0) {
            //if the char is a haraka, and harakat are enabled, add it
            if (e_harakat == 1) addChar(x[g]);
            else {
            }
        } else if (unicode.indexOf(x[g]) >= 0) {
            //if the char is an arabic reversed letter, reverse it back!
            uni_pos = unicode.indexOf(x[g]);
            la_pos = unicode.indexOf(x[g]);
            if (la_pos >= laIndex)
                //if its a laa combination
                for (
                    temp = 4;
                    temp < 20;
                    temp += 4 //find which laa
                ) {
                    if (la_pos < temp + laIndex) {
                        addChar(arabic.charAt(temp / 4 - 1));
                        addChar("ل");
                        temp = 30;
                    }
                }
            //if its any other letter
            else
                for (temp = 4; temp < 180; temp += 4) {
                    if (uni_pos < temp) {
                        addChar(arabic.charAt(temp / 4 - 1));
                        temp = 200;
                    }
                }
        } //if the char is none of the above, then treat it as english text (don't reverse) (english chars + numbers + symbols (as is))
        else {
            h = g;
            while (
                notEng.indexOf(x[h]) < 0 &&
                unicode.indexOf(x[h]) < 0 &&
                brackets.indexOf(x[h]) < 0 &&
                x[h] != undefined
                ) {
                //if this is an english sentence, or numbers, put it all in one string
                if (ennumbs.indexOf(x[h]) >= 0) {
                    mynumb = ennumbs.indexOf(x[h]);

                    if (e_numbers == 1) {
                        x[h] = arnumbs.charAt(mynumb);
                    } else if (f_numbers == 1) {
                        // AMIB
                        x[h] = fanumbs.charAt(mynumb);
                    }
                } else if (arnumbs.indexOf(x[h]) >= 0) {
                    mynumb = arnumbs.indexOf(x[h]);

                    if (e_numbers == 0) {
                        x[h] = ennumbs.charAt(mynumb);
                    }
                } else if (fanumbs.indexOf(x[h]) >= 0) {
                    // AMIB
                    mynumb = arnumbs.indexOf(x[h]);

                    if (f_numbers == 0) {
                        x[h] = ennumbs.charAt(mynumb);
                    }
                }
                tstr = tstr + x[h];
                h = h + 1;
                if (msie == 1 || opera == 1) {
                    //solving a ie/opera difference in javascript
                    temp = h + 1;
                } else {
                    temp = h;
                }
                if (x[temp] == "\n") break;
            }
            xstr = tstr.split("");
            r = xstr.length - 1;
            if (r == 1 && xstr[1] == " ")
                //make sure spaces between arabic and english text display properly
                tstr = " " + xstr[0];
            else {
                while (xstr[r] == " ") {
                    tstr = " " + tstr.substring(0, tstr.length - 1);
                    r = r - 1;
                }
            }
            output_text = tstr + output_text; //put together the arabic text + the new english text
            tstr = "";
            g = h - 1; //set the loop pointer to the first char after the english text.
        }
    }
    return old + output_text; //put together the old text and the last sentence
}

function addChar(chr) {
    //add arabic chars (change to Unicode)
    output_text = chr + output_text;
}
// console.log("The value of abc is ");
var generatedText = ProcessInput("` + input + `");
`)
	if err != nil {
		return ""
	}

	reshapedFarsi := ""
	if value, err := vm.Get("generatedText"); err == nil {
		reshapedFarsi = value.String()
	}
	return reshapedFarsi
}
