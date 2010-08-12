package aida
import (
    "xml"
    "strings"
    "fmt"
    "testing"
)

var test_aida_string = `
  <cloud2d name="rz" conversionUpperEdgeY="6.5313711268131796"
    maxEntries="10000" conversionUpperEdgeX="2.1"
    conversionLowerEdgeY="-0.3110176727053895" conversionLowerEdgeX="-0.1" title="rz"
    conversionBinsY="50" conversionBinsX="50" path="/">
    <annotation>
      <item sticky="true" value="rz" key="Title"/>
      <item sticky="true" value="/rz" key="AidaPath"/>
      <item sticky="true"
        value="//var/folders/z0/z001eV9rEzyzsb2JH+XtNE+++TM/-Tmp-/aida4331894366460493866.aida/rz" key="FullPath"/>
    </annotation>
    <entries2d>
      <entry2d valueX="0.0" valueY="0.0"/>
      <entry2d valueX="0.06282151815625668"
        valueY="0.06283185307179585"/>
      <entry2d valueX="0.12558103905862672"
        valueY="0.1256637061435917"/>
      <entry2d valueX="0.1882166266370287"
        valueY="0.18849555921538755"/>
    </entries2d>
  </cloud2d>
`


func TestCloud2D(t *testing.T) {
    var cs *cloud2dStr
    if err := xml.Unmarshal(strings.NewReader(test_aida_string), &cs); err != nil {
        t.Errorf("xml.Unmarshal: %q", err)
    }
    c := cs.Convert()
    if c.Entries2d[0].ValueX != 0.0 {
        t.Errorf("Expected %f, got %q", 0.0, c.Entries2d[0].ValueX)
    }
    fmt.Println(c)
}

