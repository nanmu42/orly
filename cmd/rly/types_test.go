/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by an MIT License.
 * You may find a license copy in project root.
 *
 */

package main

import (
	"image/color"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseGenerationQuery(t *testing.T) {
	const queryString = `g_loc=LR&g_text=abcdef&color=d2691e&img_id=66858&author=nanmu42&top_text=galaxy&title=HelloWorld`
	values, err := url.ParseQuery(queryString)
	require.NoError(t, err, "ParseQuery")
	cq, err := ParseCoverQuery(values)
	require.NoError(t, err, "ParseCoverQuery")

	assert.Equal(t, "HelloWorld", cq.Title)
	assert.Equal(t, "galaxy", cq.TopText)
	assert.Equal(t, "nanmu42", cq.Author)
	assert.Equal(t, int64(66858), cq.ImageID)
	assert.Equal(t, color.RGBA{210, 105, 30, 255}, cq.PrimaryColor)
	assert.Equal(t, "abcdef", cq.GuideText)
	assert.Equal(t, "LR", cq.GuideTextPlacement)
}
