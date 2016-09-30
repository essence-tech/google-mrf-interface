package googlemrf

import (
	"errors"
	"io"

	"github.com/essence-tech/google-mrf-interface/googlemrf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// ErrDoesNotExist returned if an entity does not exist.
var ErrDoesNotExist = errors.New("Does not exist")

// MediaClient the gateway to Google media schema info.
type MediaClient interface {
	Close()
	ValidateMRF(...string) (bool, error)
	ValidateAgency(...string) (bool, error)
	ValidateChannel(...string) (bool, error)
	ValidateLOB(...string) (bool, error)
	ValidateMedia(...string) (bool, error)
	ValidateProduct(...string) (bool, error)
	ValidateSubMedia(...string) (bool, error)
	ValidateSubProduct(...string) (bool, error)
	MRFInfo(string) (*googlemrf.MRF, error)
	MRFs() ([]googlemrf.MRF, error)
	LOBs() ([]googlemrf.Single, error)
	Products() ([]googlemrf.Double, error)
	SubProducts() ([]googlemrf.Double, error)
}

type validationFunc func(context.Context, *googlemrf.Query, ...grpc.CallOption) (*googlemrf.Single, error)

// mediaClient the gateway to Google media schema info.
type mediaClient struct {
	client googlemrf.GoogleMRFClient
	conn   *grpc.ClientConn
}

// New returns a new MediaClient with default options.
func New(host string) MediaClient {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	_client := googlemrf.NewGoogleMRFClient(conn)
	return &mediaClient{_client, conn}
}

// Close closes the connection to the rcp server.
func (m *mediaClient) Close() {
	m.conn.Close()
}

// ValidateMRF checks an MRF ID for valididty.
func (m mediaClient) ValidateMRF(inputs ...string) (bool, error) {
	mrf, err := m.client.ValidateMRF(context.Background(), inputsToQuery(inputs...))
	if err != nil {
		return false, err
	}
	if mrf.ID != "" {
		return true, nil
	}
	return false, nil
}

// ValidateAgency checks an agency name for validity.
func (m mediaClient) ValidateAgency(inputs ...string) (bool, error) {
	return validate(m.client.ValidateAgency, inputs...)
}

// ValidateChannel checks a channel for validity.
func (m mediaClient) ValidateChannel(inputs ...string) (bool, error) {
	return validate(m.client.ValidateChannel, inputs...)
}

// ValidateLOB checks a line of business for validity.
func (m mediaClient) ValidateLOB(inputs ...string) (bool, error) {
	return validate(m.client.ValidateLOB, inputs...)
}

// ValidateMedia checks a media for validitiy.
func (m mediaClient) ValidateMedia(inputs ...string) (bool, error) {
	return validate(m.client.ValidateMedia, inputs...)
}

// ValidateProduct checks a product for validity.
func (m mediaClient) ValidateProduct(inputs ...string) (bool, error) {
	return validate(m.client.ValidateProduct, inputs...)
}

// ValidateSubMedia checks a submedia for validitiy.
func (m mediaClient) ValidateSubMedia(inputs ...string) (bool, error) {
	return validate(m.client.ValidateSubMedia, inputs...)
}

// ValidateSubProduct checks a subproduct for validity.
func (m mediaClient) ValidateSubProduct(inputs ...string) (bool, error) {
	return validate(m.client.ValidateSubProduct, inputs...)
}

// MRFInfo provides more info about a MRF code.
func (m mediaClient) MRFInfo(code string) (*googlemrf.MRF, error) {
	mrf, err := m.client.ValidateMRF(context.Background(), &googlemrf.Query{Name: code})
	if err != nil {
		return nil, err
	}
	if mrf.ID != "" {
		return mrf, nil
	}
	return nil, nil
}

// MRFs returns all MRF codes.
func (m mediaClient) MRFs() ([]googlemrf.MRF, error) {
	stream, err := m.client.MRFs(context.Background(), &googlemrf.Empty{})
	if err != nil {
		return nil, err
	}

	mrfs := []googlemrf.MRF{}
	for {
		mrf, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		mrfs = append(mrfs, *mrf)
	}
	return mrfs, nil
}

// LOBs returns all LOBs.
func (m mediaClient) LOBs() ([]googlemrf.Single, error) {
	stream, err := m.client.LOBs(context.Background(), &googlemrf.Empty{})
	if err != nil {
		return nil, err
	}

	lobs := []googlemrf.Single{}
	for {
		lob, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		lobs = append(lobs, *lob)
	}
	return lobs, nil
}

// Products returns all products.
func (m mediaClient) Products() ([]googlemrf.Double, error) {
	stream, err := m.client.Products(context.Background(), &googlemrf.Empty{})
	if err != nil {
		return nil, err
	}

	pros := []googlemrf.Double{}
	for {
		p, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		pros = append(pros, *p)
	}
	return pros, nil
}

// SubProducts returns all subproducts.
func (m mediaClient) SubProducts() ([]googlemrf.Double, error) {
	stream, err := m.client.SubProducts(context.Background(), &googlemrf.Empty{})
	if err != nil {
		return nil, err
	}

	sps := []googlemrf.Double{}
	for {
		sp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		sps = append(sps, *sp)
	}
	return sps, nil
}

func validate(f validationFunc, inputs ...string) (bool, error) {
	s, err := f(context.Background(), inputsToQuery(inputs...))
	if err != nil {
		return false, err
	}
	if s.Name != "" {
		return true, nil
	}
	return false, nil
}

func inputsToQuery(inputs ...string) *googlemrf.Query {
	if len(inputs) == 1 {
		return &googlemrf.Query{
			Name: inputs[0],
		}
	}
	if len(inputs) > 1 {
		return &googlemrf.Query{
			Name:   inputs[1],
			Parent: inputs[0],
		}
	}
	return &googlemrf.Query{}
}
