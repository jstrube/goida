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


var encoded2 = `
<?xml version="1.0" encoding="ISO-8859-1" ?>
<!DOCTYPE aida SYSTEM "http://aida.freehep.org/schemas/3.3/aida.dtd">
<aida version="3.3">
  <implementation version="1.1" package="FreeHEP"/>
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
      <entry2d valueX="0.2506664671286084"
        valueY="0.2513274122871834"/>
      <entry2d valueX="0.31286893008046174"
        valueY="0.3141592653589793"/>
      <entry2d valueX="0.3747626291714493"
        valueY="0.3769911184307751"/>
      <entry2d valueX="0.43628648279308524"
        valueY="0.43982297150257105"/>
      <entry2d valueX="0.4973797743297096"
        valueY="0.5026548245743668"/>
      <entry2d valueX="0.5579822120784583"
        valueY="0.5654866776461627"/>
      <entry2d valueX="0.6180339887498949"
        valueY="0.6283185307179586"/>
      <entry2d valueX="0.6774758404905828"
        valueY="0.6911503837897545"/>
      <entry2d valueX="0.7362491053693558"
        valueY="0.7539822368615502"/>
      <entry2d valueX="0.7942957812695611"
        valueY="0.816814089933346"/>
      <entry2d valueX="0.8515585831301453"
        valueY="0.8796459430051421"/>
      <entry2d valueX="0.9079809994790936"
        valueY="0.9424777960769379"/>
      <entry2d valueX="0.9635073482034305"
        valueY="1.0053096491487337"/>
      <entry2d valueX="1.0180828315007424"
        valueY="1.0681415022205294"/>
      <entry2d valueX="1.0716535899579933"
        valueY="1.1309733552923253"/>
      <entry2d valueX="1.1241667557042612"
        valueY="1.1938052083641213"/>
      <entry2d valueX="1.1755705045849463"
        valueY="1.2566370614359172"/>
      <entry2d valueX="1.2258141073059532"
        valueY="1.319468914507713"/>
      <entry2d valueX="1.2748479794973795"
        valueY="1.382300767579509"/>
      <entry2d valueX="1.3226237306473037"
        valueY="1.4451326206513047"/>
      <entry2d valueX="1.3690942118573772"
        valueY="1.5079644737231004"/>
      <entry2d valueX="1.4142135623730951"
        valueY="1.5707963267948963"/>
      <entry2d valueX="1.4579372548428229"
        valueY="1.633628179866692"/>
      <entry2d valueX="1.5002221392609192"
        valueY="1.6964600329384882"/>
      <entry2d valueX="1.5410264855515785"
        valueY="1.7592918860102842"/>
      <entry2d valueX="1.5803100247513808"
        valueY="1.8221237390820801"/>
      <entry2d valueX="1.618033988749895"
        valueY="1.8849555921538759"/>
      <entry2d valueX="1.6541611485491234"
        valueY="1.9477874452256716"/>
      <entry2d valueX="1.6886558510040304"
        valueY="2.0106192982974673"/>
      <entry2d valueX="1.7214840540078873"
        valueY="2.0734511513692633"/>
      <entry2d valueX="1.7526133600877272"
        valueY="2.136283004441059"/>
      <entry2d valueX="1.7820130483767358"
        valueY="2.1991148575128547"/>
      <entry2d valueX="1.809654104932039"
        valueY="2.2619467105846507"/>
      <entry2d valueX="1.8355092513679625"
        valueY="2.324778563656447"/>
      <entry2d valueX="1.859552971776503"
        valueY="2.3876104167282426"/>
      <entry2d valueX="1.8817615379084507"
        valueY="2.4504422698000385"/>
      <entry2d valueX="1.9021130325903073"
        valueY="2.5132741228718345"/>
      <entry2d valueX="1.9205873713538861"
        valueY="2.57610597594363"/>
      <entry2d valueX="1.9371663222572624"
        valueY="2.638937829015426"/>
      <entry2d valueX="1.9518335238774949"
        valueY="2.701769682087222"/>
      <entry2d valueX="1.9645745014573774"
        valueY="2.764601535159018"/>
      <entry2d valueX="1.9753766811902753"
        valueY="2.827433388230814"/>
      <entry2d valueX="1.9842294026289555"
        valueY="2.8902652413026093"/>
      <entry2d valueX="1.99112392920616"
        valueY="2.9530970943744053"/>
      <entry2d valueX="1.9960534568565431"
        valueY="3.0159289474462008"/>
      <entry2d valueX="1.9990131207314632"
        valueY="3.0787608005179967"/>
      <entry2d valueX="2.0" valueY="3.1415926535897927"/>
      <entry2d valueX="1.9990131207314632"
        valueY="3.204424506661589"/>
      <entry2d valueX="1.9960534568565433"
        valueY="3.267256359733384"/>
      <entry2d valueX="1.99112392920616"
        valueY="3.330088212805181"/>
      <entry2d valueX="1.9842294026289555"
        valueY="3.3929200658769765"/>
      <entry2d valueX="1.9753766811902753"
        valueY="3.4557519189487724"/>
      <entry2d valueX="1.9645745014573774"
        valueY="3.5185837720205684"/>
      <entry2d valueX="1.9518335238774949"
        valueY="3.5814156250923634"/>
      <entry2d valueX="1.9371663222572622"
        valueY="3.6442474781641603"/>
      <entry2d valueX="1.9205873713538861"
        valueY="3.707079331235956"/>
      <entry2d valueX="1.902113032590307"
        valueY="3.7699111843077517"/>
      <entry2d valueX="1.8817615379084511"
        valueY="3.8327430373795472"/>
      <entry2d valueX="1.8595529717765027"
        valueY="3.895574890451343"/>
      <entry2d valueX="1.8355092513679623"
        valueY="3.9584067435231383"/>
      <entry2d valueX="1.8096541049320392"
        valueY="4.021238596594935"/>
      <entry2d valueX="1.7820130483767354"
        valueY="4.084070449666731"/>
      <entry2d valueX="1.752613360087727"
        valueY="4.146902302738527"/>
      <entry2d valueX="1.7214840540078873"
        valueY="4.209734155810322"/>
      <entry2d valueX="1.6886558510040304"
        valueY="4.272566008882118"/>
      <entry2d valueX="1.6541611485491237"
        valueY="4.335397861953915"/>
      <entry2d valueX="1.618033988749895"
        valueY="4.3982297150257095"/>
      <entry2d valueX="1.5803100247513806"
        valueY="4.461061568097506"/>
      <entry2d valueX="1.5410264855515785"
        valueY="4.523893421169301"/>
      <entry2d valueX="1.5002221392609192"
        valueY="4.586725274241097"/>
      <entry2d valueX="1.4579372548428224"
        valueY="4.649557127312894"/>
      <entry2d valueX="1.4142135623730951"
        valueY="4.71238898038469"/>
      <entry2d valueX="1.369094211857377"
        valueY="4.775220833456485"/>
      <entry2d valueX="1.3226237306473039"
        valueY="4.838052686528281"/>
      <entry2d valueX="1.2748479794973797"
        valueY="4.900884539600077"/>
      <entry2d valueX="1.2258141073059536"
        valueY="4.963716392671872"/>
      <entry2d valueX="1.1755705045849456"
        valueY="5.026548245743669"/>
      <entry2d valueX="1.124166755704261"
        valueY="5.0893800988154645"/>
      <entry2d valueX="1.0716535899579933"
        valueY="5.15221195188726"/>
      <entry2d valueX="1.0180828315007422"
        valueY="5.215043804959056"/>
      <entry2d valueX="0.9635073482034305"
        valueY="5.277875658030852"/>
      <entry2d valueX="0.907980999479093"
        valueY="5.340707511102648"/>
      <entry2d valueX="0.8515585831301451"
        valueY="5.403539364174444"/>
      <entry2d valueX="0.7942957812695614"
        valueY="5.466371217246239"/>
      <entry2d valueX="0.7362491053693555"
        valueY="5.529203070318036"/>
      <entry2d valueX="0.6774758404905826"
        valueY="5.592034923389831"/>
      <entry2d valueX="0.6180339887498942"
        valueY="5.654866776461628"/>
      <entry2d valueX="0.5579822120784583"
        valueY="5.717698629533423"/>
      <entry2d valueX="0.4973797743297096"
        valueY="5.780530482605219"/>
      <entry2d valueX="0.43628648279308463"
        valueY="5.843362335677015"/>
      <entry2d valueX="0.37476262917144915"
        valueY="5.9061941887488105"/>
      <entry2d valueX="0.31286893008046196"
        valueY="5.969026041820606"/>
      <entry2d valueX="0.25066646712860907"
        valueY="6.0318578948924015"/>
      <entry2d valueX="0.18821662663702782"
        valueY="6.094689747964199"/>
      <entry2d valueX="0.12558103905862716"
        valueY="6.157521601035993"/>
      <entry2d valueX="0.06282151815625647"
        valueY="6.22035345410779"/>
    </entries2d>
  </cloud2d>
  <cloud2d name="xy" conversionUpperEdgeY="1.0999999999999999"
    conversionUpperEdgeX="2.1" maxEntries="10000" conversionLowerEdgeY="-1.1"
    conversionLowerEdgeX="-0.1" conversionBinsY="50" title="xy" conversionBinsX="50"
    path="/">
    <annotation>
      <item sticky="true" value="xy" key="Title"/>
      <item sticky="true" value="/xy" key="AidaPath"/>
      <item sticky="true"
        value="//var/folders/z0/z001eV9rEzyzsb2JH+XtNE+++TM/-Tmp-/aida4331894366460493866.aida/xy" key="FullPath"/>
    </annotation>
    <entries2d>
      <entry2d valueX="0.0" valueY="0.0"/>
      <entry2d valueX="0.001973271571728441"
        valueY="0.06279051952931347"/>
      <entry2d valueX="0.007885298685522124"
        valueY="0.1253332335643042"/>
      <entry2d valueX="0.01771274927131139"
        valueY="0.18738131458572468"/>
      <entry2d valueX="0.031416838871368924"
        valueY="0.24868988716485468"/>
      <entry2d valueX="0.04894348370484647"
        valueY="0.3090169943749474"/>
      <entry2d valueX="0.07022351411174865"
        valueY="0.36812455268467803"/>
      <entry2d valueX="0.09517294753398053"
        valueY="0.42577929156507277"/>
      <entry2d valueX="0.12369331995613642"
        valueY="0.4817536741017153"/>
      <entry2d valueX="0.15567207449798492"
        valueY="0.5358267949789964"/>
      <entry2d valueX="0.19098300562505266"
        valueY="0.5877852522924731"/>
      <entry2d valueX="0.22948675722421086"
        valueY="0.6374239897486897"/>
      <entry2d valueX="0.27103137257858845"
        valueY="0.6845471059286886"/>
      <entry2d valueX="0.3154528940713113"
        valueY="0.7289686274214114"/>
      <entry2d valueX="0.36257601025131037"
        valueY="0.7705132427757893"/>
      <entry2d valueX="0.412214747707527"
        valueY="0.8090169943749473"/>
      <entry2d valueX="0.46417320502100345"
        valueY="0.8443279255020151"/>
      <entry2d valueX="0.5182463258982847"
        valueY="0.8763066800438635"/>
      <entry2d valueX="0.5742207084349273"
        valueY="0.9048270524660195"/>
      <entry2d valueX="0.6318754473153222"
        valueY="0.9297764858882513"/>
      <entry2d valueX="0.6909830056250528"
        valueY="0.9510565162951535"/>
      <entry2d valueX="0.7513101128351454"
        valueY="0.9685831611286311"/>
      <entry2d valueX="0.8126186854142755"
        valueY="0.9822872507286886"/>
      <entry2d valueX="0.8746667664356957"
        valueY="0.9921147013144778"/>
      <entry2d valueX="0.9372094804706865"
        valueY="0.9980267284282714"/>
      <entry2d valueX="1.0000000000000002"
        valueY="0.9999999999999999"/>
      <entry2d valueX="1.0627905195293132"
        valueY="0.9980267284282714"/>
      <entry2d valueX="1.1253332335643045"
        valueY="0.9921147013144777"/>
      <entry2d valueX="1.1873813145857248"
        valueY="0.9822872507286885"/>
      <entry2d valueX="1.248689887164855"
        valueY="0.968583161128631"/>
      <entry2d valueX="1.3090169943749477"
        valueY="0.9510565162951534"/>
      <entry2d valueX="1.368124552684678"
        valueY="0.9297764858882512"/>
      <entry2d valueX="1.4257792915650729"
        valueY="0.9048270524660194"/>
      <entry2d valueX="1.4817536741017154"
        valueY="0.8763066800438634"/>
      <entry2d valueX="1.5358267949789965"
        valueY="0.8443279255020151"/>
      <entry2d valueX="1.5877852522924731"
        valueY="0.8090169943749473"/>
      <entry2d valueX="1.6374239897486897"
        valueY="0.770513242775789"/>
      <entry2d valueX="1.6845471059286892"
        valueY="0.728968627421411"/>
      <entry2d valueX="1.7289686274214118"
        valueY="0.6845471059286884"/>
      <entry2d valueX="1.770513242775789"
        valueY="0.6374239897486897"/>
      <entry2d valueX="1.8090169943749477"
        valueY="0.5877852522924727"/>
      <entry2d valueX="1.844327925502015"
        valueY="0.5358267949789964"/>
      <entry2d valueX="1.8763066800438637"
        valueY="0.4817536741017151"/>
      <entry2d valueX="1.9048270524660196"
        valueY="0.4257792915650724"/>
      <entry2d valueX="1.9297764858882516"
        valueY="0.36812455268467764"/>
      <entry2d valueX="1.9510565162951536"
        valueY="0.309016994374947"/>
      <entry2d valueX="1.968583161128631"
        valueY="0.24868988716485468"/>
      <entry2d valueX="1.9822872507286888"
        valueY="0.18738131458572446"/>
      <entry2d valueX="1.9921147013144778"
        valueY="0.12533323356430442"/>
      <entry2d valueX="1.9980267284282716"
        valueY="0.06279051952931347"/>
      <entry2d valueX="2.0" valueY="-4.440892098500626E-16"/>
      <entry2d valueX="1.9980267284282716"
        valueY="-0.0627905195293139"/>
      <entry2d valueX="1.992114701314478"
        valueY="-0.12533323356430398"/>
      <entry2d valueX="1.9822872507286886"
        valueY="-0.18738131458572532"/>
      <entry2d valueX="1.968583161128631"
        valueY="-0.24868988716485513"/>
      <entry2d valueX="1.9510565162951534"
        valueY="-0.30901699437494784"/>
      <entry2d valueX="1.9297764858882513"
        valueY="-0.3681245526846784"/>
      <entry2d valueX="1.9048270524660196"
        valueY="-0.42577929156507277"/>
      <entry2d valueX="1.8763066800438633"
        valueY="-0.4817536741017159"/>
      <entry2d valueX="1.8443279255020149"
        valueY="-0.535826794978997"/>
      <entry2d valueX="1.8090169943749472"
        valueY="-0.5877852522924735"/>
      <entry2d valueX="1.7705132427757893"
        valueY="-0.6374239897486899"/>
      <entry2d valueX="1.7289686274214113"
        valueY="-0.6845471059286888"/>
      <entry2d valueX="1.6845471059286887"
        valueY="-0.7289686274214114"/>
      <entry2d valueX="1.6374239897486895"
        valueY="-0.7705132427757895"/>
      <entry2d valueX="1.5877852522924725"
        valueY="-0.809016994374948"/>
      <entry2d valueX="1.535826794978996"
        valueY="-0.8443279255020154"/>
      <entry2d valueX="1.4817536741017152"
        valueY="-0.8763066800438637"/>
      <entry2d valueX="1.4257792915650729"
        valueY="-0.9048270524660195"/>
      <entry2d valueX="1.3681245526846777"
        valueY="-0.9297764858882516"/>
      <entry2d valueX="1.3090169943749475"
        valueY="-0.9510565162951536"/>
      <entry2d valueX="1.2486898871648544"
        valueY="-0.9685831611286313"/>
      <entry2d valueX="1.1873813145857246"
        valueY="-0.9822872507286888"/>
      <entry2d valueX="1.1253332335643045"
        valueY="-0.9921147013144779"/>
      <entry2d valueX="1.0627905195293121"
        valueY="-0.9980267284282718"/>
      <entry2d valueX="1.0000000000000002" valueY="-1.0"/>
      <entry2d valueX="0.9372094804706862"
        valueY="-0.9980267284282717"/>
      <entry2d valueX="0.8746667664356957"
        valueY="-0.992114701314478"/>
      <entry2d valueX="0.8126186854142756"
        valueY="-0.9822872507286888"/>
      <entry2d valueX="0.7513101128351458"
        valueY="-0.9685831611286314"/>
      <entry2d valueX="0.6909830056250519"
        valueY="-0.9510565162951534"/>
      <entry2d valueX="0.6318754473153216"
        valueY="-0.9297764858882513"/>
      <entry2d valueX="0.5742207084349273"
        valueY="-0.9048270524660196"/>
      <entry2d valueX="0.5182463258982842"
        valueY="-0.8763066800438635"/>
      <entry2d valueX="0.46417320502100323"
        valueY="-0.8443279255020151"/>
      <entry2d valueX="0.4122147477075263"
        valueY="-0.8090169943749471"/>
      <entry2d valueX="0.36257601025131003"
        valueY="-0.7705132427757891"/>
      <entry2d valueX="0.3154528940713113"
        valueY="-0.7289686274214117"/>
      <entry2d valueX="0.2710313725785881"
        valueY="-0.6845471059286884"/>
      <entry2d valueX="0.22948675722421064"
        valueY="-0.6374239897486896"/>
      <entry2d valueX="0.1909830056250521"
        valueY="-0.5877852522924727"/>
      <entry2d valueX="0.1556720744979847"
        valueY="-0.5358267949789964"/>
      <entry2d valueX="0.12369331995613642"
        valueY="-0.4817536741017153"/>
      <entry2d valueX="0.0951729475339802"
        valueY="-0.4257792915650722"/>
      <entry2d valueX="0.07022351411174854"
        valueY="-0.36812455268467786"/>
      <entry2d valueX="0.04894348370484647"
        valueY="-0.3090169943749476"/>
      <entry2d valueX="0.031416838871369035"
        valueY="-0.24868988716485535"/>
      <entry2d valueX="0.017712749271311168"
        valueY="-0.18738131458572382"/>
      <entry2d valueX="0.007885298685522235"
        valueY="-0.12533323356430465"/>
      <entry2d valueX="0.001973271571728441"
        valueY="-0.06279051952931326"/>
    </entries2d>
  </cloud2d>
</aida>
`
