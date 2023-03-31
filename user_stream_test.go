package binance_connector

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type userStreamTestSuite struct {
	baseTestSuite
}

func TestUserStream(t *testing.T) {
	suite.Run(t, new(userStreamTestSuite))
}

func (s *userStreamTestSuite) TestStartUserStream() {
	data := []byte(`{
        "listenKey": "pqia91ma19a5s61cv6a81va65sdf19v8a65a1a5s61cv6a81va65sdf19v8a65a1"
    }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		s.assertRequestEqual(newRequest(), r)
	})

	listenKey, err := s.client.NewCreateListenKeyService().Do(newContext())
	s.r().NoError(err)
	s.r().Equal("pqia91ma19a5s61cv6a81va65sdf19v8a65a1a5s61cv6a81va65sdf19v8a65a1", listenKey)
}

func (s *userStreamTestSuite) TestKeepaliveUserStream() {
	data := []byte(`{}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	listenKey := "dummykey"
	s.assertReq(func(r *request) {
		s.assertRequestEqual(newRequest().setParam("listenKey", listenKey), r)
	})

	err := s.client.NewPingUserStream().ListenKey(listenKey).Do(newContext())
	s.r().NoError(err)
}

func (s *userStreamTestSuite) TestCloseUserStream() {
	data := []byte(`{}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	listenKey := "dummykey"
	s.assertReq(func(r *request) {
		s.assertRequestEqual(newRequest().setParam("listenKey", listenKey), r)
	})

	err := s.client.NewCloseUserStream().ListenKey(listenKey).Do(newContext())
	s.r().NoError(err)
}
