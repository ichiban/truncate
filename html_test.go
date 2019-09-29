package truncate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TruncateHTML(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		assert := assert.New(t)

		s, err := HTML(`<p>test</p>`, 3)
		assert.NoError(err)
		assert.Equal("<p>te…</p>", s)
	})

	t.Run("nested", func(t *testing.T) {
		assert := assert.New(t)

		s, err := HTML(`<div><div><p>test</p></div></div>`, 3)
		assert.NoError(err)
		assert.Equal("<div><div><p>te…</p></div></div>", s)
	})

	t.Run("with trailing tags", func(t *testing.T) {
		assert := assert.New(t)

		s, err := HTML(`<p>test<br/></p><br/>`, 3)
		assert.NoError(err)
		assert.Equal("<p>te…</p>", s)
	})

	t.Run("fairly complex", func(t *testing.T) {
		assert := assert.New(t)

		s, err := HTML(`<h1>test</h1><div>a<strong>b</strong></div><blockquote>e</blockquote><div><br>f<br><br>g<br><br>h</div>`, 6)
		assert.NoError(err)
		assert.Equal("<h1>test</h1><div>a<strong>b</strong></div>", s)
	})
}
