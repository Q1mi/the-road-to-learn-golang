package palindrome

import "testing"

func TestJudge(t *testing.T) {
	type test struct {
		str  string
		want bool
	}

	var tests = map[string]test{
		"simple":  test{"沙河有沙又有河", false},
		"with":    test{"Madam,I’mAdam", true},
		"chinese": test{"油灯少灯油", true},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Judge(tc.str)
			if got != tc.want {
				t.Errorf("excepted:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

func BenchmarkJudge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Judge("沙河有沙又有河沙")
	}
}

func BenchmarkJudgeParallel(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Judge("沙河有沙又有河沙")
		}
	})
}
