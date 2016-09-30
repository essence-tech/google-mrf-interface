package googlemrf

import (
	"github.com/essence-tech/google-mrf-interface/googlemrf"
)

// NewTestClient provides a test version of the MediaClient.
func NewTestClient(
	s bool,
	m *googlemrf.MRF,
	ms []googlemrf.MRF,
	si *googlemrf.Single,
	d *googlemrf.Double,
	ss []googlemrf.Single,
	ds []googlemrf.Double,
) MediaClient {
	return testClient{
		Success: s,
		MRF:     m,
		ManyMRF: ms,
		Single:  si,
		Double:  d,
		Singles: ss,
		Doubles: ds,
	}
}

type testClient struct {
	Success bool
	MRF     *googlemrf.MRF
	ManyMRF []googlemrf.MRF
	Single  *googlemrf.Single
	Double  *googlemrf.Double
	Singles []googlemrf.Single
	Doubles []googlemrf.Double
}

func (t testClient) Close() {}
func (t testClient) ValidateMRF(...string) (bool, error) {
	return t.Success, nil
}
func (t testClient) ValidateAgency(...string) (bool, error) {
	return t.Success, nil
}
func (t testClient) ValidateChannel(...string) (bool, error) {
	return t.Success, nil
}
func (t testClient) ValidateLOB(...string) (bool, error) {
	return t.Success, nil
}
func (t testClient) ValidateMedia(...string) (bool, error) {
	return t.Success, nil
}
func (t testClient) ValidateProduct(...string) (bool, error) {
	return t.Success, nil
}
func (t testClient) ValidateSubMedia(...string) (bool, error) {
	return t.Success, nil
}
func (t testClient) ValidateSubProduct(...string) (bool, error) {
	return t.Success, nil
}
func (t testClient) MRFInfo(string) (*googlemrf.MRF, error) {
	return t.MRF, nil
}
func (t testClient) MRFs() ([]googlemrf.MRF, error) {
	return t.ManyMRF, nil
}
func (t testClient) LOBs() ([]googlemrf.Single, error) {
	return t.Singles, nil
}
func (t testClient) Products() ([]googlemrf.Double, error) {
	return t.Doubles, nil
}
func (t testClient) SubProducts() ([]googlemrf.Double, error) {
	return t.Doubles, nil
}
