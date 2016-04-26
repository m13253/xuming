package main

import (
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func prepareQuotes() {
	for i, s := range quotes {
		encoded := []byte(s)
		decoded := make([]byte, len(encoded))
		for idx := range encoded {
			decoded[idx] = encoded[idx] ^ 1
		}
		quotes[i] = string(decoded)
	}
}

func getQuote() string {
	return quotes[rand.Intn(len(quotes))]
}

// bytes(map(lambda x: x ^ 1, s.encode())).decode()
var quotes = []string {
	"Unn!xntof-!unn!rhlqmd-!rnlduhldr!o`®wd ",
	"蛑烢蘨廏癏怪\uef4d虳鍂邝廏甐怿\uef4d鵢庺膌廏粱䰰⁃",
	"牐幁杊弡牐䱰枤耒弡曅䑌䬖\uef4d弡庋䥨瘛幋杞癮頩罕俖綊罒䱼醵⁃",
	"濏䚼䉷颀牐庭\uef4d䧃笥筻鵞牐庭\uef4d牐庭幁䯛駀狋笽,9ㄣ瑟幋䏺\uef4d幁䯛駀幻䚼䯷⁀幻廻籐廈䎢籕⁃",
	"牐庭龘帉䥦曅䚼䯷\uef4d帱䆡矷袟睻馢䇲幌廇薯裙\uef4d龙癮駀醫䋚皵敞\uef4d蜡牐庭醫䶰⁃",
	"牐䚟侇鴶牐皿溎䝩!0854!丵⁀57!丵曅䥦䬧丵庢\uef4d溎乹牐庭䗝筣揨䅾⁏I`v`hh!fthu`s⁎\uef4d溎乹佸䤎曅!@mni`!&Nd!龘幫笍皳⁃",
	"䝩牐検脓帊䈌\uef4d牐䏺佸幁皳!@mni`!&Nd!龘幫!fthu`s\uef4d鮶䶟蔾䥪廻罕䕰\uef40",
	"H!vntme!mhjd!rqd`j!gdv!vnser/!牐龘筠眤鶞弡庭䀛幁硶曅䶤弝\uef4d䱰癮駀!qsnlnud!ltut`m!toedsru`oehof!cduvddo!nts!uvn!bntoushdr⁃",
	"H!uihoj!H!rqd`j!wdsx!qnns!Dofmhri-!ctu!`oxv`x!H&l!xdu!un!r`x-!uihr!hr!wdsx!hlqnsu`ou/",
	"牐庭曅䇚虞駀䎐爭幌灔旳䋲\uef4d幌灔鈱蘨\uef4d龟溬弝牙曅泿椟⁃",
	"牐庭駀䜛䇲䝱䜛獀䉱仔\uef4d䜛獀䤊牙\uef4d䜛獀䇌䜛獀\uef40牐庭䱰幁䯛邼䥞䎗俖睁䑏曅邝䉨⁃",
	"牐庭幬䌏籐疎癮睈朁俉佻曅䆜逛䋚\uef4d庺弔曅䚱蛿肼䏊幌䁓牐庭⁃幬䚼廻籐癮幌䎮牙邝曅\uef40",
	"弡庭䫓弒䌂幆駀粩煎䔋\uef4d幌駀⁏馀朁袏癮俖蚩⁎䔋⁃",
	"䉛爌弡薯牐䔋\uef4d牐䎮庤䚟欕弡幁䎤⁏疡䎮䤈䐋⁎⁃",
	"弡薯牐畮獀幌畮獀\uef4d牐鮵畮獀⁃牐䱰癏栯鶞弡䐋鮈龘幁惸⁃",
	"弡庭蜟乹懞烈餾痸曅龘幁䤖曅摇鯻⁃弡庭箔檞龙!unn!xntof⁃",
	"牐癮黪溎昿牙廇\uef4d馀俖䥛廇\uef40餾痸曅䒫幁幫䚼䯷牐糠䏺龆\uef5e",
	"弡庭駀枤耒\uef4d濏䚼曅䌏鏰䢪\uef4d胢箕弡庭幌枤耒誙䉱䒫膍䏺䔧⁃牐鶞店鱉歐袏敞\uef40",
	"鮇俖䕕鮇俖䔋\uef5e",
	"弡庭睈幁幫䤼\uef4d䄩幗攍鶐䉱庁帉䝱痸\uef4d弡庭箕䄷店曅餾痸鯱遄鶐俖龙侪⁃弇癮薯眤薯䏺曅薯裙䔋\uef4d肼!unn!rhlqmd-!rnlduhldr!o`®wd\uef40",
	"牐庋䥨癮弝幻幁幫蔾遄鶞弡庭鯳龘幫⁃",
	"牐睈龘幫侄駀䐋鮈弡庭幁惸廻敞曅溎諍⁃",
	"幬䚼睈幁䎤鮜䎪⁏薶䢱䥦䎐鵣⁎⁃牐䱰庁帉鮜帞幌鮵⁃龘癮睁䤼曅⁃",
	"䝩䯢彡幋䱇眤䧃矝弡庭狤耒幋睈䀎䶯\uef4d弡庭駀鵞鵢⁃",
	"伒慷牐庭曅䇲䯛眂帞癮俉膌駀曅⁃䉱胢疷䁘牐庭彛頩灀曅\uef40",
	"弡庭䔋\uef4d幌駀炲佅幫䥦痱薺\uef4d鮵揱䝩䶳溎蓧䯛廇\uef4d䇌狋牐爸䉥幁攫⁃",
	"弡庭䔋\uef4do`®wd\uef40",
	"H&l!`ofsx !牐鶞弡鯳\uef4d弡庭龘硶䬑䔋癮幌頍曅\uef40牐庋䥨毖俖漫廇弡庭幁幊\uef40",
	"Ѯ!ѾѿѽѼЏ!ІЂѵѼѿѴ!ѽѲѼѿѳѴѼЍѴで",
	"ѓѴѺѹѻѹѸ!ѾѿЌЃ-!ѼѱЇѹѿѼѱѺЍѼѱЎ!ѲѿЁѵѿЀЃЍ ",
	"L`swdmntr !牐箕弡暹䶯龝廇⁃牐幁氲幂䚚⁃乂庁帞幌邼鶞弡箕⁃",
	"H&l!rnssx/!H!`l!`!dmdbushb`m!qnvds!dofhodds/",
	"廻䐑\uef4d幌毠弡䜑䝩庁帉弌漯\uef4d店幌邼幌䚟侇鴶皿溎䞸䄺龆曅箌硠⁃",
	"弡幁幫廻曅䐼龑䔋\uef4d伒慷駀蜡醫牐䤊痖\uef4d弇癮帞駀遂阐䉱䏇䎳曅頍橊⁃",
	"爁庤䑏眤牐䱰侴廇幥觗鮖\uef4d䎪⁏銞䉨䚼䯷敞笺庤\uef4d䳃䚡椹槎耾鷊帊⁃⁎",
	"銞䉨䚼䯷敞笺庤\uef4d䳃䚡椹槎耾鷊帊⁃",
	"弡庭溘牐瑟曅龘睭幝餾䔋ででDybhude ",
	"幌眤龘幫紶䌖䑌䰰耖燿廇⁃",
	"龘帉䤼曅袏瘮䑌邝紶䌖駀䥦䋚䯢彡\uef4d䍖廭帞駀䥦䋚䯢彡\uef4d牐䚟䍖廭帞幻弡庭䯢彡䯢彡\uef4d庤䑏龘䰰䱰廻䰰廻紶廇⁃",
}
