package html

import (
	"testing"

	"github.com/matDobek/gov--attendance-check/internal/testing/assert"
)

func TestToElements(t *testing.T) {
	t.Run("trims whitespace", func(t *testing.T) {
		t.Parallel()

		got := toElements("\n\t   div   \t\n")
		want := []element{
			{tag: "div"},
		}

		assert.Equal(t, got, want)
	})

	t.Run("handles nested elements", func(t *testing.T) {
		t.Parallel()

		got := toElements("table tbody")
		want := []element{
			{tag: "table"}, {tag: "tbody"},
		}

		assert.Equal(t, got, want)
	})

	t.Run("handles css classes", func(t *testing.T) {
		t.Parallel()

		got := toElements("table.foo.bar")
		want := []element{
			{tag: "table", class: []string{"foo", "bar"}},
		}

		assert.Equal(t, got, want)
	})

	t.Run("handles css ids", func(t *testing.T) {
		t.Parallel()

		got := toElements("table#foo#bar")
		want := []element{
			{tag: "table", id: []string{"foo", "bar"}},
		}

		assert.Equal(t, got, want)
	})

	t.Run("handles all elements at once", func(t *testing.T) {
		t.Parallel()

		got := toElements("table#foo.foz thead#bar.baz")
		want := []element{
			{tag: "table", id: []string{"foo"}, class: []string{"foz"}},
			{tag: "thead", id: []string{"bar"}, class: []string{"baz"}},
		}

		assert.Equal(t, got, want)
	})

	t.Run("allow names to have special characters (`-`  `_`  `0-9`)", func(t *testing.T) {
		t.Parallel()

		got := toElements("table.align-left__2px")
		want := []element{
			{tag: "table", class: []string{"align-left__2px"}},
		}

		assert.Equal(t, got, want)
	})

	t.Run("allow names to have special characters (`-`  `_`  `0-9`)", func(t *testing.T) {
		t.Parallel()

		got := toElements("div:2")
		want := []element{
			{tag: "div", nthChild: 2},
		}

		assert.Equal(t, got, want)
	})

}

func TestSearch(t *testing.T) {
	t.Parallel()

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
		ex, err := Search(html, test[0].(string))
		assert.Error(t, err)

		got := []string{}
		for _, e := range ex {
			got = append(got, e.Data)
		}
		want := test[1].([]string)

		assert.Equal(t, got, want)
	}
}

func TestExtract(t *testing.T) {
	t.Parallel()

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
		got, err := Extract(html, test[0].(string))
		assert.Error(t, err)

		want := test[1].([]string)

		assert.Equal(t, got, want)
	}
}

func TestExtractAttr(t *testing.T) {
	t.Parallel()

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
		got, err := ExtractAttr(html, test[0].(string), test[1].(string))
		assert.Error(t, err)

		want := test[2].([]string)

		assert.Equal(t, got, want)
	}
}
