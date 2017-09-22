# go-multi-mock
Multi mock go tester

This package allows you to set mulitple mocks for testing when you make many different api calls.

This test is meant to be used for testing and would be utilized as follows..

```		
      It("Returns not found when wrong person is found", func() {
				response1 := helpers.MockResponse{
					Status: http.StatusOK,
					Body:   mock_json.WrongPersonFoundResponse,
				}
				response2 := helpers.MockResponse{
					Status: http.StatusOK,
					Body:   mock_json.OtherWrongPersonResponse,
				}

				responses := helpers.Responses{&response1, &response2}
				server := StartMocks(responses)
				os.Setenv("YOUR_URL", server.URL)

				class := Class{}
				res, err := class.method("stuff", "number")

				Expect(err).To(HaveOccurred())
			})     
 ```
