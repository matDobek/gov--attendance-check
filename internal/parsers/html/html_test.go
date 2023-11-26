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
		{"div:2", []element{{tag: "div", nthChild: 2}}},
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
		{"td", []string{"td", "td", "td"}},
		{"tr td", []string{"td", "td", "td"}},
		{"tr:2 td", []string{"td", "td"}},
		{"td:1", []string{"td", "td"}},
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

func TestExtract(t *testing.T) {
	html := `
	<!DOCTYPE html>
	<html>
		<head></head>
		<body>
			<table>
				<tr>
					<td>
						<a>
							<em>
								a1
							</em>
						</a>
						<strong>
							a2
						</strong>
					</td>
				</tr>
				<tr>
					<td>b1</td>
					<td>b2</td>
				</tr>
			</table>
		</body>
	</html>
	`

	tests := [][]any{
		{"tr td", []string{"a1 a2", "b1", "b2"}},
	}

	for _, test := range tests {
		arg := test[0].(string)
		expectedResult := test[1].([]string)

		result, err := Extract(html, arg)
		if err != nil {
			t.Errorf("Error: %v", err)
			continue
		}

		if len(result) != len(expectedResult) {
			t.Errorf("args: %v\n %#v\n %#v\n\n", arg, expectedResult, result)
			continue
		}

		for i, er := range expectedResult {
			if result[i] != er {
				t.Errorf("args: %v\n %#v\n %#v\n\n", arg, er, result[i])
			}
		}
	}
}

func TestExtractAttr(t *testing.T) {
	html := `
	<!DOCTYPE html>
	<html>
		<head></head>
		<body>
			<a href="foo.co">foo</a>
			<a href="bar.co">bar</a>
			<a href="baz.co">baz</a>
		</body>
	</html>
	`

	tests := [][]any{
		{"a", "href", []string{"foo.co", "bar.co", "baz.co"}},
	}

	for _, test := range tests {
		arg0 := test[0].(string)
		arg1 := test[1].(string)
		expectedResult := test[2].([]string)

		result, err := ExtractAttr(html, arg0, arg1)
		if err != nil {
			t.Errorf("Error: %v", err)
			continue
		}

		if len(result) != len(expectedResult) {
			t.Errorf("args: %v %v\n %#v\n %#v\n\n", arg0, arg1, expectedResult, result)
			continue
		}

		for i, er := range expectedResult {
			if result[i] != er {
				t.Errorf("args: %v, %v\n %#v\n %#v\n\n", arg0, arg1, er, result[i])
			}
		}
	}
}
