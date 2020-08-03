package main

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

func main() {
	currentTime := time.Now()

	var t []string
	for i := 1; i < 32; i += 1 {
		numStr := fmt.Sprintf("%02d", i)
		t = append(t, numStr)
	}

	println()
	fmt.Print(currentTime)
	println()

	tmpl, _ := template.New("").Parse(`
		<table border>
		<tbody>
			<tr>
				<td>
					name
				</td>
				<td colspan="31">
					October
				</td>
			</tr>

			{{ $months := .Months   }}
			{{ $domains := .Domains }}

			<tr>
				<td></td>
				{{ range $month := $months }}
					<td>{{ $month }}</td>
				{{ end }}
			</tr>

			{{ range $domain := $domains }}
				<tr>
					<td>{{ $domain }}</td>
					{{ range $month := $months }}
						<td>X</td>
					{{ end }}
				</tr>
			{{ end }}

		</tbody>
		</table>
	`)

	tmpl.Execute(os.Stdout, struct {
		Months  []string
		Domains []string
	}{
		Months:  []string{"a", "b", "c", "d"},
		Domains: []string{"john", "mary", "bob"},
	})
}
