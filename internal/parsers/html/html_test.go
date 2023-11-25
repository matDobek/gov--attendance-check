package html

import (
	"reflect"
	"testing"
)

func TestToElements(t *testing.T) {
	tests := [][]interface{}{
		{"   div   ", []element{{tag: "div"}}},
		{"table tbody", []element{{tag: "table"}, {tag: "tbody"}}},
		{"table#Foo", []element{{tag: "table", id: []string{"Foo"}}}},
		{"table#Foo#Bar", []element{{tag: "table", id: []string{"Foo", "Bar"}}}},
		{"table.Foo", []element{{tag: "table", class: []string{"Foo"}}}},
		{"table#Foo.Bar", []element{{tag: "table", id: []string{"Foo"}, class: []string{"Bar"}}}},
		{"table.align-left__2px", []element{{tag: "table", class: []string{"align-left__2px"}}}},
		{"table#Foo thead#Fooz.Bar.Baz", []element{
			{tag: "table", id: []string{"Foo"}},
			{tag: "thead", id: []string{"Fooz"}, class: []string{"Bar", "Baz"}},
		}},
	}

	for _, test := range tests {
		arg := test[0].(string)
		result := toElements(arg)
		expectedResult := test[1].([]element)

		if !reflect.DeepEqual(result, expectedResult) {
			t.Errorf("args: %v\n %#v\n %#v\n\n", arg, expectedResult, result)
		}
	}

}

func TestSearch(t *testing.T) {
	html := `
	<!DOCTYPE html>
	<html>
		<head></head>

		<body id="main" class="container container-main">
			<table>
				<thead>
				</thead>

				<tbody>
					<tr>
						<td>a1</td>
						<td>a2</td>
					</tr>
					<tr>
						<td>b1</td>
						<td>b2</td>
					</tr>
				</tbody>
			</table>
		</body>
	</html>
	`

	tests := [][]any{
		{"body", []string{"body"}},
		{"body#main.container.container-main", []string{"body"}},
		{"non-existingtag#main.container", []string{}},
		{"body#non-existing", []string{}},
		{"body.non-existing", []string{}},
		{"body tbody", []string{"tbody"}},
		{"td", []string{"td", "td", "td", "td"}},
		{"tr td", []string{"td", "td", "td", "td"}},
	}

	for _, test := range tests {
		arg := test[0].(string)
		expectedResult := test[1].([]string)

		result, err := Search(html, arg)
		if err != nil {
			t.Errorf("Error: %v", err)
			continue
		}

		if len(result) != len(expectedResult) {
			t.Errorf("args: %v\n %#v\n %#v\n\n", arg, expectedResult, result)
			continue
		}

		for i, er := range expectedResult {
			if result[i].Data != er {
				t.Errorf("args: %v\n %#v\n %#v\n\n", arg, er, result[i])
			}
		}
	}
}
