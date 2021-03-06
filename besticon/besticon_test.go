package besticon

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"net/url"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/erkie/besticon/vcr"
)

func TestKicktipp(t *testing.T) {
	actualImages, err, finder := fetchIconsWithVCR("kicktipp.vcr", "http://kicktipp.de", IconFinder{})
	assertEquals(t, nil, err)
	expectedImages := []Icon{
		{URL: "http://info.kicktipp.de/assets/img/jar_cb333387130/assets/img/logos/apple-touch-icon-57x57-precomposed.png", Width: 57, Height: 57, Format: "png", Bytes: 1535, Sha1sum: "79aae9e0df7d52ed50ac47c1dd4bd16e2ddf8b4a"},
		{URL: "http://www.kicktipp.de/apple-touch-icon-precomposed.png", Width: 57, Height: 57, Format: "png", Bytes: 1622, Sha1sum: "fd4306aefd9ed2b4983697ca56301d8810a03987"},
		{URL: "http://www.kicktipp.de/apple-touch-icon.png", Width: 57, Height: 57, Format: "png", Bytes: 1622, Sha1sum: "fd4306aefd9ed2b4983697ca56301d8810a03987"},
		{URL: "http://www.kicktipp.de/favicon.ico", Width: 32, Height: 32, Format: "gif", Bytes: 35275, Sha1sum: "09297d0ffe17149c3d4d4a3a3a8c7e8c51932d58"},
		{URL: "http://info.kicktipp.de/assets/img/jar_cb1652512069/assets/img/logos/favicon.png", Width: 16, Height: 16, Format: "png", Bytes: 1820, Sha1sum: "04b49fac810828f6723cd763600af23f0edbde03"},
	}
	assertEquals(t, expectedImages, actualImages)

	actualImage := finder.IconInSizeRange(SizeRange{20, 80, 500})
	expectedImage := &Icon{URL: "http://info.kicktipp.de/assets/img/jar_cb333387130/assets/img/logos/apple-touch-icon-57x57-precomposed.png", Width: 57, Height: 57, Format: "png", Bytes: 1535, Sha1sum: "79aae9e0df7d52ed50ac47c1dd4bd16e2ddf8b4a"}
	assertEquals(t, expectedImage, actualImage)
}

func TestDaringfireball(t *testing.T) {
	actualImages, err, finder := fetchIconsWithVCR("daringfireball.net.vcr", "http://daringfireball.net", IconFinder{})
	assertEquals(t, nil, err)

	expectedImages := []Icon{
		{URL: "http://daringfireball.net/graphics/apple-touch-icon.png", Width: 314, Height: 314, Format: "png", Bytes: 8642, Error: error(nil), Sha1sum: "f47cf7cf13ec1a74049d99d9f1565354b5b20317"},
		{URL: "http://daringfireball.net/favicon.ico", Width: 32, Height: 32, Format: "ico", Bytes: 6518, Error: error(nil), Sha1sum: "c066c70aa1dd2b4347d3095592aac28b55e85d04"},
		{URL: "http://daringfireball.net/graphics/favicon.ico?v=005", Width: 32, Height: 32, Format: "ico", Bytes: 6518, Error: error(nil), Sha1sum: "c066c70aa1dd2b4347d3095592aac28b55e85d04"},
	}

	assertEquals(t, expectedImages, actualImages)

	actualImage := finder.IconInSizeRange(SizeRange{20, 80, 500})
	expectedImage := &Icon{URL: "http://daringfireball.net/graphics/apple-touch-icon.png", Width: 314, Height: 314, Format: "png", Bytes: 8642, Error: error(nil), Sha1sum: "f47cf7cf13ec1a74049d99d9f1565354b5b20317"}
	assertEquals(t, expectedImage, actualImage)
}

func TestAwsAmazonChangingBaseURL(t *testing.T) {
	actualImages, err, _ := fetchIconsWithVCR("aws.amazon.vcr", "http://aws.amazon.com", IconFinder{})
	assertEquals(t, nil, err)
	expectedImages := []Icon{
		{URL: "http://a0.awsstatic.com/main/images/site/touch-icon-ipad-144-precomposed.png", Width: 144, Height: 144, Format: "png", Bytes: 3944, Sha1sum: "225817df40ff11d083c282d08b49a5ed50fd0f2d"},
		{URL: "http://a0.awsstatic.com/main/images/site/touch-icon-iphone-114-precomposed.png", Width: 114, Height: 114, Format: "png", Bytes: 3081, Sha1sum: "58aabb2a99fcb283710fd200c9feed69c015a29e"},
		{URL: "http://a0.awsstatic.com/main/images/site/favicon.ico", Width: 16, Height: 16, Format: "ico", Bytes: 1150, Sha1sum: "a64f3616e3671b835f8b208b7339518d8b386a08"},
		{URL: "http://aws.amazon.com/favicon.ico", Width: 16, Height: 16, Format: "ico", Bytes: 1150, Sha1sum: "a64f3616e3671b835f8b208b7339518d8b386a08"},
	}

	assertEquals(t, expectedImages, actualImages)
}

func TestNetflixWitCookieRedirects(t *testing.T) {
	actualImages, err, _ := fetchIconsWithVCR("netflix.vcr", "http://netflix.com", IconFinder{})
	assertEquals(t, nil, err)
	expectedImages := []Icon{
		{URL: "https://assets.nflxext.com/us/ffe/siteui/common/icons/nficon2016.png", Width: 64, Height: 64, Format: "png", Bytes: 1755, Sha1sum: "867e51e9b4a474c19da52d6454076c007a9d01f2"},
		{URL: "https://assets.nflxext.com/us/ffe/siteui/common/icons/nficon2016.ico", Width: 64, Height: 64, Format: "ico", Bytes: 16958, Sha1sum: "931e18dfc6e7d950dc2f2bbdfe31e1ea720acf7c"},
		{URL: "https://www.netflix.com/favicon.ico", Width: 64, Height: 64, Format: "ico", Bytes: 16958, Sha1sum: "931e18dfc6e7d950dc2f2bbdfe31e1ea720acf7c"},
	}

	assertEquals(t, expectedImages, actualImages)
}

func TestAolWithOnePixelGifs(t *testing.T) {
	actualImages, err, _ := fetchIconsWithVCR("aol.vcr", "http://aol.com", IconFinder{})
	assertEquals(t, nil, err)
	expectedImages := []Icon{
		{URL: "http://www.aol.com/favicon.ico", Width: 32, Height: 32, Format: "ico", Bytes: 7886, Error: error(nil), Sha1sum: "c474f8307362367be553b884878e221f25fcb80b"},
		{URL: "http://www.aol.com/favicon.ico?v=2", Width: 32, Height: 32, Format: "ico", Bytes: 7886, Error: error(nil), Sha1sum: "c474f8307362367be553b884878e221f25fcb80b"},
	}

	assertEquals(t, expectedImages, actualImages)
}

func TestGithubWithIconHrefLinks(t *testing.T) {
	actualImages, err, _ := fetchIconsWithVCR("github.vcr", "http://github.com", IconFinder{})
	assertEquals(t, nil, err)
	expectedImages := []Icon{
		{URL: "https://github.com/apple-touch-icon-144.png", Width: 144, Height: 144, Format: "png", Bytes: 796, Sha1sum: "2626d8f64d5d3a76bd535151dfe84b62d3f3ee63"},
		{URL: "https://github.com/apple-touch-icon.png", Width: 120, Height: 120, Format: "png", Bytes: 676, Sha1sum: "8eb0b1d3f0797c0fe94368f4ad9a2c9513541cd2"},
		{URL: "https://github.com/apple-touch-icon-114.png", Width: 114, Height: 114, Format: "png", Bytes: 648, Sha1sum: "644982478322a731a6bd8fe7fad9afad8f4a3c4b"},
		{URL: "https://github.com/apple-touch-icon-precomposed.png", Width: 57, Height: 57, Format: "png", Bytes: 705, Sha1sum: "f97e9954be83f93ce2a1a85c2d8f93e2235c887f"},
		{URL: "https://assets-cdn.github.com/favicon.ico", Width: 32, Height: 32, Format: "ico", Bytes: 6518, Sha1sum: "4eda7c0f3a36181f483dd0a14efe9f58c8b29814"},
		{URL: "https://github.com/favicon.ico", Width: 32, Height: 32, Format: "ico", Bytes: 6518, Sha1sum: "4eda7c0f3a36181f483dd0a14efe9f58c8b29814"},
	}

	assertEquals(t, expectedImages, actualImages)
}

func TestWithIconDataImages(t *testing.T) {
	actualImages, err, _ := fetchIconsWithVCR("dpa.vcr", "https://dpa.com", IconFinder{})
	assertEquals(t, nil, err)
	expectedImages := []Icon{
		{URL: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyhpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTExIDc5LjE1ODMyNSwgMjAxNS8wOS8xMC0wMToxMDoyMCAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENDIDIwMTUgKE1hY2ludG9zaCkiIHhtcE1NOkluc3RhbmNlSUQ9InhtcC5paWQ6NTNDODdFRjkxQTg1MTFFNkFDRjBGQjJBNUY1NTA2RkEiIHhtcE1NOkRvY3VtZW50SUQ9InhtcC5kaWQ6NTNDODdFRkExQTg1MTFFNkFDRjBGQjJBNUY1NTA2RkEiPiA8eG1wTU06RGVyaXZlZEZyb20gc3RSZWY6aW5zdGFuY2VJRD0ieG1wLmlpZDo1M0M4N0VGNzFBODUxMUU2QUNGMEZCMkE1RjU1MDZGQSIgc3RSZWY6ZG9jdW1lbnRJRD0ieG1wLmRpZDo1M0M4N0VGODFBODUxMUU2QUNGMEZCMkE1RjU1MDZGQSIvPiA8L3JkZjpEZXNjcmlwdGlvbj4gPC9yZGY6UkRGPiA8L3g6eG1wbWV0YT4gPD94cGFja2V0IGVuZD0iciI/PrTtwmQAAAKJSURBVHja7JpLSFRRHIdnbo1llEgEBkoQgZBGFm6iB0nSA0toUWizbBEEKkTbaJdFmxbhwkWLNpELN2oGUdJCAheREhFEBVlqoIhCDx/h9B04yTDqzD3nTm7md+DjiPP/z+/cj/s4V4ynUqlYIY8gVuBDAiRAAiRAAiRAAiRAAiRAAiRAAiRAAiRAAiRAAiRAAiSggEY88xf3B6+VMB2AUvgOb1qP3lvMdzA5W23O9v+VQ8YmphrYCVM24/eqAiguZ7oNTVCUVjNrPjaf0fwrD4syi7kFSdickdMB7eT8jJixjekmXIGStI/m4CHcIGNqWQANB5meQlmW730Np2icjrCwKqZnUJ6lbBhO/lugp+DnUJ2lbNRmfAhoMKd6T46DN6MWHkU85ftyHHzMXhZd1Mc9Msw9rTvHwZuxC3qpLzYN16EiZMZpms54OmiD3SFrT0CjR8ZFOByythKuGgGXHEOSngKS65DjnGEE7HFsqvI4NTcw7XVs2+8hoNqxfp/PPmCLR0/CY8+RWIe1JQL7+HEZn11Xxd12zj7rXcYXDwGfHOu/BvbO7DL6PO8BTxzrez0ynI8lsJufPyEbvtmNhM+4Awsha8fhgUdGJ0yGrDU7wrsBp+c7fmgJ0WB2gRcyt5IOl8FHuzNLhVhYE/U/PDJmmJphPkfpElymfjSwjZ32GbrWdToCR6gbirJFpd+cPedhbI2St3CMusEIGQNMdfA+yy7wLHWPV7wM8bgyd9FzcCjtZch84QANS3l8STHvAA02ZwdMwEt4ka8cMjYy1cPxtJehV9BPxvKlGNe/yekPIhIgARIgARIgARIgARIgARIgARIgARIgARIgARIgARIgAYUy/gowAD0rp5mdS0siAAAAAElFTkSuQmCC", Width: 64, Height: 64, Format: "png", Bytes: 1563, Sha1sum: "b006bb6a138582b13da64f5c97107394388ea5d1"},
	}

	assertEquals(t, expectedImages, actualImages)
}

func TestWithIconOverride(t *testing.T) {
	override := IconFinder{}
	override.OverrideFinder = func(inputUrl string) []string {
		if strings.Contains(inputUrl, "dpa.com") {
			parsed, err := url.Parse(inputUrl)
			if err != nil {
				return nil
			}
			if parsed.Host == "dpa.com" {
				return []string{"https://static.feedercdn.com/logo/dpa.png"}
			}
		}
		return nil
	}
	actualImages, err, _ := fetchIconsWithVCR("dpa_override.vcr", "https://dpa.com", override)
	assertEquals(t, nil, err)
	expectedImages := []Icon{
		{URL: "https://static.feedercdn.com/logo/dpa.png", Width: 343, Height: 343, Format: "png", Bytes: 8914, Sha1sum: "2f95d6155514e93febd6d637fdf524820adbb24b"},
	}

	assertEquals(t, expectedImages, actualImages)
}

func TestEat24WithBaseTag(t *testing.T) {
	actualImages, err, finder := fetchIconsWithVCR("eat24.vcr", "http://eat24.com", IconFinder{})
	assertEquals(t, nil, err)
	expectedImages := []Icon{
		{URL: "http://eat24hours.com/favicon.ico", Width: 16, Height: 16, Format: "ico", Bytes: 1406, Sha1sum: "f8914a1135e718b11cc93b7a362655ca358c16fb"},
	}
	assertEquals(t, expectedImages, actualImages)

	actualImage := finder.IconInSizeRange(SizeRange{20, 50, 500})
	assertEquals(t, (*Icon)(nil), actualImage)
}

func TestAlibabaWithBaseTagWithoutScheme(t *testing.T) {
	actualImages, err, _ := fetchIconsWithVCR("alibaba.vcr", "http://alibaba.com", IconFinder{})
	assertEquals(t, nil, err)
	expectedImages := []Icon{
		{URL: "http://is.alicdn.com/simg/single/icon/favicon.ico", Width: 16, Height: 16, Format: "ico", Bytes: 1406, Sha1sum: "4ffbef9b6044c62cd6c8b1ee0913ba93e6e80072"},
		{URL: "http://www.alibaba.com/favicon.ico", Width: 16, Height: 16, Format: "ico", Bytes: 1406, Sha1sum: "4ffbef9b6044c62cd6c8b1ee0913ba93e6e80072"},
	}

	assertEquals(t, expectedImages, actualImages)
}

func TestARDWithSortBySize(t *testing.T) {
	actualImages, err, _ := fetchIconsWithVCR("ard.vcr", "http://ard.de", IconFinder{})
	assertEquals(t, nil, err)
	expectedImages := []Icon{
		{URL: "http://www.ard.de/ARD-144.png", Width: 144, Height: 144, Format: "png", Bytes: 29228, Sha1sum: "a6be15435a80e9de7978d203a3f2723940ab6bda"},
		{URL: "http://www.ard.de/apple-touch-icon-precomposed.png", Width: 144, Height: 144, Format: "png", Bytes: 29228, Sha1sum: "a6be15435a80e9de7978d203a3f2723940ab6bda"},
		{URL: "http://www.ard.de/apple-touch-icon.png", Width: 144, Height: 144, Format: "png", Bytes: 29228, Sha1sum: "a6be15435a80e9de7978d203a3f2723940ab6bda"},
		{URL: "http://www.ard.de/favicon.ico", Width: 144, Height: 144, Format: "ico", Bytes: 116094, Sha1sum: "e5bd22dda5647420c5d2053ee9fd21b543dc09a8"},
	}

	assertEquals(t, expectedImages, actualImages)
}

func TestMortenmøllerWithIDNAHost(t *testing.T) {
	actualImages, err, _ := fetchIconsWithVCR("mortenmoller.vcr", "https://mortenmøller.dk", IconFinder{})
	assertEquals(t, nil, err)
	assertEquals(t, 13, len(actualImages))
}

func TestYoutubeWithDomainRewrite(t *testing.T) {
	// This test can only work because with HostOnlyDomains accordingly
	_, err, finder := fetchIconsWithVCR("youtube.vcr", "http://youtube.com/does-not-exist", IconFinder{})
	ico := finder.IconInSizeRange(SizeRange{0, 80, 200})
	assertEquals(t, &Icon{URL: "https://s.ytimg.com/yts/img/favicon_96-vfldSA3ca.png", Width: 96, Height: 96, Format: "png", Bytes: 1510, Sha1sum: "7149bef987538d34e2ab6e069d08211d0a6e407d"}, ico)
	assertEquals(t, nil, err)
}

func TestParsingInexistentSite(t *testing.T) {
	actualImages, err, _ := fetchIconsWithVCR("not_existent.vcr", "http://wikipedia.org/does-not-exist", IconFinder{})

	assertEquals(t, errors.New("besticon: not found"), err)
	assertEquals(t, 0, len(actualImages))
}

func TestParsingEmptyResponse(t *testing.T) {
	actualImages, err, _ := fetchIconsWithVCR("empty_body.vcr", "http://foobar.com", IconFinder{})

	assertEquals(t, 0, len(actualImages))
	assertEquals(t, errors.New("besticon: empty response"), err)
}

func mustFindIconLinks(html []byte) []string {
	doc, e := docFromHTML(html)
	check(e)
	links := extractIconTags(doc)
	sort.Strings(links)
	return links
}

func TestMainColorForIconsWithBrokenImageData(t *testing.T) {
	icn := Icon{Format: "png", ImageData: []byte("broken-image-data")}
	colr := MainColorForIcons([]Icon{icn})
	assertEquals(t, (*color.RGBA)(nil), colr)
}

func TestFindBestIconNoIcons(t *testing.T) {
	icons, _, _ := fetchIconsWithVCR("example.com.vcr", "http://example.com", IconFinder{})
	assertEquals(t, 0, len(icons))
}

func TestLinkExtraction(t *testing.T) {
	links := mustFindIconLinks(mustReadFile("testdata/daringfireball.html"))
	assertEquals(t, []string{
		"/graphics/apple-touch-icon.png",
		"/graphics/favicon.ico?v=005",
	}, links)

	links = mustFindIconLinks(mustReadFile("testdata/newyorker.html"))
	assertEquals(t, []string{
		"/wp-content/assets/dist/img/icon/apple-touch-icon-114x114-precomposed.png",
		"/wp-content/assets/dist/img/icon/apple-touch-icon-144x144-precomposed.png",
		"/wp-content/assets/dist/img/icon/apple-touch-icon-57x57-precomposed.png",
		"/wp-content/assets/dist/img/icon/apple-touch-icon-precomposed.png",
		"/wp-content/assets/dist/img/icon/apple-touch-icon.png",
		"/wp-content/assets/dist/img/icon/favicon.ico",
	}, links)
}

func TestImageSizeDetection(t *testing.T) {
	assertEquals(t, 1, getImageWidthForFile("testdata/pixel.png"))
	assertEquals(t, 1, getImageWidthForFile("testdata/pixel.gif"))
	assertEquals(t, 48, getImageWidthForFile("testdata/favicon.ico"))
}

func TestParseSizeRange(t *testing.T) {
	// This single num behaviour ensures backwards compatibility for
	// people who pant (at least) pixel perfect icons.
	sizeRange, err := ParseSizeRange("120")
	assertEquals(t, &SizeRange{120, 120, MaxIconSize}, sizeRange)

	sizeRange, err = ParseSizeRange("0..120..256")
	assertEquals(t, &SizeRange{0, 120, 256}, sizeRange)

	sizeRange, err = ParseSizeRange("120..120..120")
	assertEquals(t, &SizeRange{120, 120, 120}, sizeRange)

	_, err = ParseSizeRange("")
	assertEquals(t, errBadSize, err)

	_, err = ParseSizeRange(" ")
	assertEquals(t, errBadSize, err)

	// Max < Perfect not allowed
	_, err = ParseSizeRange("16..120..80")
	assertEquals(t, errBadSize, err)

	// Perfect < Min  not allowed
	_, err = ParseSizeRange("120..16..80")
	assertEquals(t, errBadSize, err)

	// Min too small
	_, err = ParseSizeRange("-1..2..3")
	assertEquals(t, errBadSize, err)

	// Max too big
	_, err = ParseSizeRange("1..2..501")
	assertEquals(t, errBadSize, err)
}

func TestParseSize(t *testing.T) {
	size, ok := parseSize("120")
	assertEquals(t, ok, true)
	assertEquals(t, 120, size)

	_, ok = parseSize("")
	assertEquals(t, ok, false)

	_, ok = parseSize("-10")
	assertEquals(t, ok, false)
}

const testdataDir = "testdata/"

func fetchIconsWithVCR(vcrFile string, url string, finder IconFinder) ([]Icon, error, *IconFinder) {
	path := testdataDir + vcrFile
	c, f, err := vcr.Client(path)
	if err != nil {
		return nil, err, nil
	}
	defer f.Close()

	setHTTPClient(c)
	finder.HostOnlyDomains = []string{"youtube.com"}
	icons, e := finder.FetchIcons(url)
	return icons, e, &finder
}

func getImageWidthForFile(filename string) int {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	icfg, _, err := image.DecodeConfig(f)
	check(err)
	return icfg.Width
}

func mustReadFile(filename string) []byte {
	bytes, e := ioutil.ReadFile(filename)
	check(e)
	return bytes
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func assertEquals(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		fail(t, fmt.Sprintf("Not equal: %#v (expected)\n"+
			"        != %#v (actual)", expected, actual))
	}
}

func fail(t *testing.T, failureMessage string) {
	t.Errorf("\t%s\n"+
		"\r\t",
		failureMessage)
}

func init() {
	keepImageBytes = false
}
