package testing

import (
	"Signature/operations"
	"math/big"
	"testing"
)

func Test(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()
	if gPoint.X.String() != "55066263022277343669578718895168534326250603453777594175500187360389116729240" || gPoint.Y.String() != "32670510020758816978083085130507043184471273380659243275938904335757337482424" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test1(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()
	point := operations.Multiply(gPoint, big.NewInt(2))
	if point.X.String() != "89565891926547004231252920425935692360644145829622209833684329913297188986597" || point.Y.String() != "12158399299693830322967808612713398636155367887041628176798871954788371653930" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test2(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()
	point := operations.Multiply(gPoint, big.NewInt(3))
	if point.X.String() != "112711660439710606056748659173929673102114977341539408544630613555209775888121" || point.Y.String() != "25583027980570883691656905877401976406448868254816295069919888960541586679410" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test3(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()
	point := operations.Multiply(gPoint, big.NewInt(4))
	if point.X.String() != "103388573995635080359749164254216598308788835304023601477803095234286494993683" || point.Y.String() != "37057141145242123013015316630864329550140216928701153669873286428255828810018" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test4(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()
	point := operations.Multiply(gPoint, big.NewInt(5))
	if point.X.String() != "21505829891763648114329055987619236494102133314575206970830385799158076338148" || point.Y.String() != "98003708678762621233683240503080860129026887322874138805529884920309963580118" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test5(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()
	point := operations.Multiply(gPoint, big.NewInt(6))
	if point.X.String() != "115780575977492633039504758427830329241728645270042306223540962614150928364886" || point.Y.String() != "78735063515800386211891312544505775871260717697865196436804966483607426560663" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test6(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()
	point := operations.Multiply(gPoint, big.NewInt(7))
	if point.X.String() != "41948375291644419605210209193538855353224492619856392092318293986323063962044" || point.Y.String() != "48361766907851246668144012348516735800090617714386977531302791340517493990618" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test7(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()
	point := operations.Multiply(gPoint, big.NewInt(8))
	if point.X.String() != "21262057306151627953595685090280431278183829487175876377991189246716355947009" || point.Y.String() != "41749993296225487051377864631615517161996906063147759678534462689479575333124" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test8(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()
	point := operations.Multiply(gPoint, big.NewInt(9))
	if point.X.String() != "78173298682877769088723994436027545680738210601369041078747105985693655485630" || point.Y.String() != "92362876758821804597230797234617159328445543067760556585160674174871431781431" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test9(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()
	point := operations.Multiply(gPoint, big.NewInt(10))
	if point.X.String() != "72488970228380509287422715226575535698893157273063074627791787432852706183111" || point.Y.String() != "62070622898698443831883535403436258712770888294397026493185421712108624767191" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test10(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()

	nStr := "115792089237316195423570985008687907852837564279074904382605163141518161494336"
	n, successN := new(big.Int).SetString(nStr, 10)

	if !successN {
		panic("Error setting y value")
	}

	point := operations.Multiply(gPoint, n)
	if point.X.String() != "55066263022277343669578718895168534326250603453777594175500187360389116729240" || point.Y.String() != "83121579216557378445487899878180864668798711284981320763518679672151497189239" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test11(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()

	nStr := "115792089237316195423570985008687907852837564279074904382605163141518161494335"
	n, successN := new(big.Int).SetString(nStr, 10)

	if !successN {
		panic("Error setting y value")
	}

	point := operations.Multiply(gPoint, n)
	if point.X.String() != "89565891926547004231252920425935692360644145829622209833684329913297188986597" || point.Y.String() != "103633689937622365100603176395974509217114616778598935862658712053120463017733" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test12(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()

	nStr := "115792089237316195423570985008687907852837564279074904382605163141518161494334"
	n, successN := new(big.Int).SetString(nStr, 10)

	if !successN {
		panic("Error setting y value")
	}

	point := operations.Multiply(gPoint, n)
	if point.X.String() != "112711660439710606056748659173929673102114977341539408544630613555209775888121" || point.Y.String() != "90209061256745311731914079131285931446821116410824268969537695047367247992253" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test13(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()

	nStr := "115792089237316195423570985008687907852837564279074904382605163141518161494333"
	n, successN := new(big.Int).SetString(nStr, 10)

	if !successN {
		panic("Error setting y value")
	}

	point := operations.Multiply(gPoint, n)
	if point.X.String() != "103388573995635080359749164254216598308788835304023601477803095234286494993683" || point.Y.String() != "78734948092074072410555668377823578303129767736939410369584297579653005861645" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test14(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()

	nStr := "115792089237316195423570985008687907852837564279074904382605163141518161494332"
	n, successN := new(big.Int).SetString(nStr, 10)

	if !successN {
		panic("Error setting y value")
	}

	point := operations.Multiply(gPoint, n)
	if point.X.String() != "21505829891763648114329055987619236494102133314575206970830385799158076338148" || point.Y.String() != "17788380558553574189887744505607047724243097342766425233927699087598871091545" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test15(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()

	nStr := "115792089237316195423570985008687907852837564279074904382605163141518161494331"
	n, successN := new(big.Int).SetString(nStr, 10)

	if !successN {
		panic("Error setting y value")
	}

	point := operations.Multiply(gPoint, n)
	if point.X.String() != "115780575977492633039504758427830329241728645270042306223540962614150928364886" || point.Y.String() != "37057025721515809211679672464182131982009266967775367602652617524301408111000" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test16(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()

	nStr := "115792089237316195423570985008687907852837564279074904382605163141518161494330"
	n, successN := new(big.Int).SetString(nStr, 10)

	if !successN {
		panic("Error setting y value")
	}

	point := operations.Multiply(gPoint, n)
	if point.X.String() != "41948375291644419605210209193538855353224492619856392092318293986323063962044" || point.Y.String() != "67430322329464948755426972660171172053179366951253586508154792667391340681045" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test17(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()

	nStr := "115792089237316195423570985008687907852837564279074904382605163141518161494329"
	n, successN := new(big.Int).SetString(nStr, 10)

	if !successN {
		panic("Error setting y value")
	}

	point := operations.Multiply(gPoint, n)
	if point.X.String() != "21262057306151627953595685090280431278183829487175876377991189246716355947009" || point.Y.String() != "74042095941090708372193120377072390691273078602492804360923121318429259338539" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test18(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()

	nStr := "1157920892373161954235709850086879078528375642790749043826"
	n, successN := new(big.Int).SetString(nStr, 10)

	if !successN {
		panic("Error setting y value")
	}

	point := operations.Multiply(gPoint, n)
	if point.X.String() != "3244136006132833744462889224152708771538725381841114352938415213976847208540" || point.Y.String() != "40609115319813868216280712574872299580414903578367214853494075651471732711693" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test19(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()

	nStr := "50086879078528375642790749043826"
	n, successN := new(big.Int).SetString(nStr, 10)

	if !successN {
		panic("Error setting y value")
	}

	point := operations.Multiply(gPoint, n)
	if point.X.String() != "94269476549127915542463146211203539333198952564967457583147901943972078843758" || point.Y.String() != "30385150139282804446670070983661983443735183985003177252872694918605756888634" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}

func Test20(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Test didn't call panic")
		}
	}()

	gPoint, _ := operations.CreateGPoint()

	nStr := "115792089237316195423570985008687907852837564279074904382605163141518161494337"
	n, _ := new(big.Int).SetString(nStr, 10)

	point := operations.Multiply(gPoint, n)

	if point != nil {
		t.Error("Test didn't call panic")
	}
}

func Test21(t *testing.T) {
	gPoint, _ := operations.CreateGPoint()

	nStr := "1"
	n, _ := new(big.Int).SetString(nStr, 10)

	point := operations.Multiply(gPoint, n)

	if point.X.String() != "55066263022277343669578718895168534326250603453777594175500187360389116729240" || point.Y.String() != "32670510020758816978083085130507043184471273380659243275938904335757337482424" {
		t.Errorf("Invalid coord:\nX:%s\nY:%s", gPoint.X.String(), gPoint.Y.String())
	}
}
