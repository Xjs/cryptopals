package english

import (
	"testing"
)

func TestIsEnglish(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		text string
		want bool
	}{
		{"wristwatch", "I wish to watch my Irish wristwatch", true},
		{"reference", EnglishText, true},
		{"triple-reference", EnglishText + EnglishText + EnglishText, true},
		{"german", "Deutscher Beispieltext, überhaupt nicht englisch, denn er enthält viel mehr Umlaute und Kapitänsmützen und lange Wörter und sowas.", false}, // fragile
		{"german2", "Ja, ich versuch ihn mal noch ein bißchen ordentlich aufzuschreiben.", false},
		{"japanese", "ごめんなさい！", false},
		{"japanese2", "『ヒックとドラゴン』（原題: How to Train Your Dragon）は、2010年のアメリカの3Dアニメ映画。監督は『リロ・アンド・スティッチ』のディーン・デュボアとクリス・サンダース。イギリスの児童文学作家クレシッダ・コーウェルの同名の児童文学が原作である。北米では約2億1700万ドル以上の興行収入を上げている[1]。また、このヒットを受けて続編の制作が決定した[2]。続編は2014年6月13日に全米公開されている。", false},
		{"garbage", "f240t9ujgn elnfi u2pweiodwd qeu109mq dssd lak;lkasd ckj", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(GetScore(tt.text, ReferenceHistogram))
			if got := IsEnglish(tt.text); got != tt.want {
				t.Errorf("IsEnglish() = %v, want %v (score: %f)", got, tt.want, GetScore(tt.text, ReferenceHistogram))
			}
		})
	}
}
