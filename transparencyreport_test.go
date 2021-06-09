package transparencyreport

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUnmarshalJSON(t *testing.T) {
	data := ``
	m := []int{}
	err := unmarshalJSON([]byte(data), &m)
	require.NotNil(t, err)
	require.Equal(t, "Unknown prefix", err.Error())

	data = `)]}'[[1, 2, 3]]`
	err = unmarshalJSON([]byte(data), &m)
	require.Nil(t, err)
	require.Equal(t, []int{1, 2, 3}, m)
}

func TestCertUnmarshalJSON(t *testing.T) {
	data := `)]}'

	[["https.ct.chr",["81:14:BA:EC:B4:5B:05:58:81:54:EB:09:BF:BE:22:D5","C\u003dRU, O\u003dSOLAR SEKYURITI, OOO, OU\u003dInstantSSL, L\u003dMoscow, ST\u003dMoscow, STREET\u003dTagilskaya str 4/5, postalCode\u003d107143, CN\u003dqualysapi.va.rt-solar.ru","C\u003dGB, O\u003dCOMODO CA Limited, L\u003dSalford, ST\u003dGreater Manchester, CN\u003dCOMODO RSA Organization Validation Secure Server CA",1540771200000,1572393599000,null,null,["qualysapi.va.rt-solar.ru","www.qualysapi.va.rt-solar.ru"]
	]
	,[["cloudflare_nimbus2019",null,42894737]
	,["google_rocketeer",null,480359570]
	]
	]
	]`

	cert := Cert{}
	err := unmarshalJSON([]byte(data), &cert)
	require.Nil(t, err)
	require.Equal(t, "81:14:BA:EC:B4:5B:05:58:81:54:EB:09:BF:BE:22:D5", cert.SerialNumber)
}

func TestCertsResponseUnmarshalJSON(t *testing.T) {
	data := `)]}'

	[["https.ct.cdsr",[[null,"appscreener.rt-solar.ru","Let's Encrypt Authority X3",1601441778000,1609217778000,"J2RTFJ1CV6XFaVJXgAhwdWhCSytSPx5Hm+dv6C3RgPo\u003d",2,null,1]
	,[null,"*.rt-solar.ru","GlobalSign RSA OV SSL CA 2018",1594302104000,1629177413000,"NCyijfiyBLECXXcur5AE11utRH8zPT+Vx0ieDEwNQ/w\u003d",6,null,2]
	,[null,"rt-solar.ru","Let's Encrypt Authority X3",1533220085000,1540996085000,"NFN8Jp5TUElZ9Dw7fx/+7q0mN4xNFKEO9mOoBOcNZUo\u003d",5,null,1]
	,[null,"appscreener.rt-solar.ru","Let's Encrypt Authority X3",1601441778000,1609217778000,"TsPfP1pD4qS2lY2onS3gkgPjWEQ6tISp1Vpbl63sR14\u003d",2,null,1]
	,[null,"*.rt-solar.ru","GlobalSign RSA OV SSL CA 2018",1563340613000,1594963013000,"kp5Gw9wSmrI4+x0MivWed5OXDccXgoQRMKXkzCq+O7g\u003d",3,null,2]
	,[null,"*.rt-solar.ru","GlobalSign RSA OV SSL CA 2018",1594302104000,1629177413000,"lYz840xuzQHc4dyqBrXQry+JubxooVTAQYBmGqDVFHM\u003d",1,null,2]
	,[null,"*.rt-solar.ru","COMODO RSA Domain Validation Secure Server CA",1532390400000,1564012799000,"l4Ybq0eqxHiCPiR+2DOBST5jRL0hXu8scP1wEoqLv4E\u003d",2,null,2]
	,[null,"*.rt-solar.ru","COMODO RSA Domain Validation Secure Server CA",1532390400000,1564012799000,"rGwt1+IVhB3Z1IGL2M1x/bgMrIbtxF5YOhoLlN/3Po4\u003d",3,null,2]
	,[null,"*.rt-solar.ru","GlobalSign RSA OV SSL CA 2018",1563340613000,1594963013000,"7rgLYo25KcWNjC+22NGkAbMd74HRmbYwf6ADZX9bdLk\u003d",6,null,2]
	,[null,"rt-solar.ru","Let's Encrypt Authority X3",1533220085000,1540996085000,"79lm2oVushsShS9ZtJlnvBlRj8hvM+TeR9eHPo+T7dU\u003d",3,null,1]
	]
	,[["2318898313753020486",null,"C\u003dBE, O\u003dGlobalSign nv-sa, CN\u003dGlobalSign RSA OV SSL CA 2018",17]
	,["2758712965900849059",null,"C\u003dGB, O\u003dCOMODO CA Limited, L\u003dSalford, ST\u003dGreater Manchester, CN\u003dCOMODO RSA Domain Validation Secure Server CA",8]
	,["4428624498008853827",null,"C\u003dUS, O\u003dLet's Encrypt, CN\u003dLet's Encrypt Authority X3",186]
	,["-5555328073833692941",null,"C\u003dUS, O\u003dDigiCert Inc, OU\u003dwww.digicert.com, CN\u003dGeoTrust RSA CA 2018",4]
	,["306338310514697924",null,"C\u003dRU, O\u003dRU-Center (ЗАО Региональный Сетевой Информационный Центр), L\u003dMoscow, ST\u003dMoscow, CN\u003dRU-CENTER High Assurance Services CA 2",4]
	,["4761097160778966641",null,"C\u003dGB, O\u003dSectigo Limited, L\u003dSalford, ST\u003dGreater Manchester, CN\u003dSectigo RSA Domain Validation Secure Server CA",20]
	,["-5541429204046699144",null,"C\u003dGB, O\u003dCOMODO CA Limited, L\u003dSalford, ST\u003dGreater Manchester, CN\u003dCOMODO RSA Organization Validation Secure Server CA",16]
	,["-8394976244263496152",null,"C\u003dBE, O\u003dGlobalSign nv-sa, CN\u003dGlobalSign Organization Validation CA - SHA256 - G2",4]
	]
	,[null,"cnQtc29sYXIucnU6dHJ1ZTp0cnVlOjpDQW9RQVE9PQ",null,1,26]
	]
	]`

	response := certsResponse{}
	err := unmarshalJSON([]byte(data), &response)
	// hashes, pageID, err := parseResponse([]byte(data))
	require.Nil(t, err)
	require.Equal(t, "cnQtc29sYXIucnU6dHJ1ZTp0cnVlOjpDQW9RQVE9PQ", response.NextPageID)
	require.Equal(t, 10, len(response.Certs))
	require.Equal(t, "J2RTFJ1CV6XFaVJXgAhwdWhCSytSPx5Hm+dv6C3RgPo\u003d", response.Certs[0].Hash)
	require.Equal(t, "79lm2oVushsShS9ZtJlnvBlRj8hvM+TeR9eHPo+T7dU\u003d", response.Certs[9].Hash)

	// Thu Aug 02 2018 14:28:05 GMT+0000
	date := time.Date(2018, time.August, 02, 14, 28, 05, 00, time.UTC)
	if !response.Certs[9].NotBefore.UTC().Equal(date) {
		t.Errorf("%s != %s", response.Certs[9].NotBefore.UTC(), date)
	}

	// Wed Oct 31 2018 14:28:05 GMT+0000
	date = time.Date(2018, time.October, 31, 14, 28, 05, 00, time.UTC)
	if !response.Certs[9].NotAfter.UTC().Equal(date) {
		t.Errorf("%s != %s", response.Certs[9].NotAfter.UTC(), date)
	}
}
